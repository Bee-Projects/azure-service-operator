package storageaccount

import (
	"context"
	"time"

	operatorv1alpha1 "github.com/bee-projects/azure-service-operator/pkg/apis/operator/v1alpha1"
	"github.com/bee-projects/azure-service-operator/pkg/azure"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
)

var log = logf.Log.WithName("controller_storageaccount")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new StorageAccount Controller and adds it to the Manager. The Manager will set fields on the Controller
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

	storageClient := storage.NewAccountsClient(config.SubscriptionID)
	storageClient.Authorizer = authorizer
	return &ReconcileStorageAccount{client: mgr.GetClient(), scheme: mgr.GetScheme(), deployer: armDeployer, storageClient: storageClient}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("storageaccount-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource StorageAccount
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.StorageAccount{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner StorageAccount
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.StorageAccount{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileStorageAccount{}

// ReconcileStorageAccount reconciles a StorageAccount object
type ReconcileStorageAccount struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	deployer *azure.ArmDeployer
	storageClient storage.AccountsClient
}

// Reconcile reads that state of the cluster for a StorageAccount object and makes changes based on the state read
// and what is in the StorageAccount.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileStorageAccount) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling StorageAccount")

	// Fetch the StorageAccount instance
	instance := &operatorv1alpha1.StorageAccount{}
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
		reqLogger.Info("Deleting Storage", "value", instance)

		// Delete the resources
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, err = r.storageClient.Delete(ctx, instance.Spec.ResourceGroupName, instance.Name)
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
			"pkg/azure/templates/storageaccount.json",
			map[string]interface{}{
				"storageAccountName" : map[string]interface{}{
					"value": instance.Name,
				},
				"location": map[string]interface{}{
					"value": instance.Spec.Location,
				},
				"accountType": map[string]interface{}{
					"value": instance.Spec.AccountType,
				},
				"kind": map[string]interface{}{
					"value": instance.Spec.Kind,
				},
				"accessTier": map[string]interface{}{
					"value": instance.Spec.AccessTier,
				},
				"supportsHttpsTrafficOnly": map[string]interface{}{
					"value": instance.Spec.SupportsHttpsTrafficOnly,
				},
			},
		)
	}
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *operatorv1alpha1.StorageAccount) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
