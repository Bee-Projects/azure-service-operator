package acr

import (
	"context"
	"github.com/bee-projects/azure-service-operator/pkg/azure"
	"time"

	// "github.com/Azure/go-autorest/autorest/azure"
	// "github.com/Azure/go-autorest/autorest/azure/auth"
	operatorv1alpha1 "github.com/bee-projects/azure-service-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	cr "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2018-09-01/containerregistry"
)

var log = logf.Log.WithName("controller_acr")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Acr Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager, config azure.Config) error {
	return add(mgr, newReconciler(mgr, config))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, config azure.Config) reconcile.Reconciler {
	azureSubscriptionID := config.SubscriptionID

	authorizer, err := azure.GetBearerTokenAuthorizer(
		config.Environment,
		config.TenantID,
		config.ClientID,
		config.ClientSecret,
	)
	if err != nil {
		return nil
	}

	resourceGroupsClient := resourcesSDK.NewGroupsClientWithBaseURI(
		config.Environment.ResourceManagerEndpoint,
		azureSubscriptionID,
	)
	resourceGroupsClient.Authorizer = authorizer
	resourceGroupsClient.UserAgent = azure.GetUserAgent(resourceGroupsClient.Client)
	resourceDeploymentsClient := resourcesSDK.NewDeploymentsClientWithBaseURI(
		config.Environment.ResourceManagerEndpoint,
		azureSubscriptionID,
	)
	resourceDeploymentsClient.Authorizer = authorizer
	resourceDeploymentsClient.UserAgent =
		azure.GetUserAgent(resourceDeploymentsClient.Client)
	resourceDeploymentsClient.PollingDuration = time.Minute * 45
	armDeployer := azure.NewDeployer(
		resourceGroupsClient,
		resourceDeploymentsClient,
	)

	acrClient := cr.NewRegistriesClient(config.SubscriptionID)
	acrClient.Authorizer = authorizer
	return &ReconcileAcr{client: mgr.GetClient(), scheme: mgr.GetScheme(), deployer: armDeployer, acrClient: acrClient}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("acr-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Acr
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.Acr{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Acr
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.Acr{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileAcr{}

// ReconcileAcr reconciles a Acr object
type ReconcileAcr struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	deployer *azure.ArmDeployer
	acrClient cr.RegistriesClient
}

// Reconcile reads that state of the cluster for a Acr object and makes changes based on the state read
// and what is in the Acr.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileAcr) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Acr", "Request", request)

	// Fetch the Acr instance
	instance := &operatorv1alpha1.Acr{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// If the resource is deleted
	if instance.ObjectMeta.DeletionTimestamp != nil && instance.ObjectMeta.Finalizers != nil {
		reqLogger.Info("Deleting ACR", "value", instance)

		// Delete the resources
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, err = r.acrClient.Delete(ctx, instance.Spec.ResourceGroupName, instance.Name)
		if err != nil {
			log.Error(err, "")
			return reconcile.Result{}, err
		}

		// Delete the deployment
		err := r.deployer.Delete(azure.GetDeploymentName(instance.Name, instance.Spec.ResourceGroupName),
			instance.Spec.ResourceGroupName,)
		if err != nil {
			log.Error(err, "")
			return reconcile.Result{}, err
		} else {
			instance.SetFinalizers(nil)
			r.client.Update(context.TODO(), instance)
			return reconcile.Result{}, nil
		}

	}

	_, status, err := r.deployer.GetDeploymentAndStatus(azure.GetDeploymentName(instance.Name, instance.Spec.ResourceGroupName),
		instance.Spec.ResourceGroupName,
		)

	if status == azure.DeploymentStatusNotFound {
		r.deployer.Deploy(
			azure.GetDeploymentName(instance.Name, instance.Spec.ResourceGroupName),
			instance.Spec.ResourceGroupName,
			"pkg/azure/templates/acr.json",
			map[string]interface{}{
				"registryName" : map[string]interface{}{
					"value": instance.Name,
				},
				"registryLocation": map[string]interface{}{
					"value": instance.Spec.Location,
				},
				"registrySku": map[string]interface{}{
					"value": instance.Spec.Sku,
				},
				"adminUserEnabled": map[string]interface{}{
					"value": instance.Spec.AdminEnabled,
				},
			},
		)
	}

	return reconcile.Result{}, nil
}

