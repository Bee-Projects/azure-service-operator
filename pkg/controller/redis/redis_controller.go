package redis

import (
	"context"
	"time"

	operatorv1alpha1 "github.com/bee-projects/azure-service-operator/pkg/apis/operator/v1alpha1"
	"github.com/bee-projects/azure-service-operator/pkg/azure"
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

	rediscache "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
)

var log = logf.Log.WithName("controller_redis")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Redis Controller and adds it to the Manager. The Manager will set fields on the Controller
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

	redisClient := rediscache.NewClient(config.SubscriptionID)
	redisClient.Authorizer = authorizer

	return &ReconcileRedis{client: mgr.GetClient(), scheme: mgr.GetScheme(), deployer: armDeployer, redisClient: redisClient}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("redis-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Redis
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.Redis{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Redis
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.Redis{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileRedis{}

// ReconcileRedis reconciles a Redis object
type ReconcileRedis struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	deployer *azure.ArmDeployer
	redisClient rediscache.Client
}

// Reconcile reads that state of the cluster for a Redis object and makes changes based on the state read
// and what is in the Redis.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileRedis) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Redis")

	// Fetch the Redis instance
	instance := &operatorv1alpha1.Redis{}
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
		reqLogger.Info("Deleting Redis", "value", instance)

		// Delete the resources
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, err = r.redisClient.Delete(ctx, instance.Spec.ResourceGroupName, instance.Name)
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

	} else {
		_, status, _ := r.deployer.GetDeploymentAndStatus(azure.GetDeploymentName(instance.Name, instance.Spec.ResourceGroupName),
			instance.Spec.ResourceGroupName,
		)

		if status == azure.DeploymentStatusNotFound {
			r.deployer.Deploy(
				azure.GetDeploymentName(instance.Name, instance.Spec.ResourceGroupName),
				instance.Spec.ResourceGroupName,
				"pkg/azure/templates/redis.json",
				map[string]interface{}{
					"name" : map[string]interface{}{
						"value": instance.Name,
					},
					"location": map[string]interface{}{
						"value": instance.Spec.Location,
					},
					"skuName": map[string]interface{}{
						"value": instance.Spec.SkuName,
					},
					"skuFamily": map[string]interface{}{
						"value": instance.Spec.SkuFamily,
					},
					"capacity": map[string]interface{}{
						"value": instance.Spec.Capacity,
					},
					"shardCount": map[string]interface{}{
						"value": instance.Spec.ShardCount,
					},
					"enableNonSslPort": map[string]interface{}{
						"value": instance.Spec.EnableNonSSLPort,
					},
				},
			)
		}
	}
	return reconcile.Result{}, nil
}