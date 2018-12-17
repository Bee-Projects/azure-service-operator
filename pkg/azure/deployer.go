package azure

import (
	"context"
	"encoding/json"
	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("deployer")

type deploymentStatus string

const (
	DeploymentStatusNotFound  deploymentStatus = "NOT_FOUND"
	DeploymentStatusRunning   deploymentStatus = "RUNNING"
	DeploymentStatusSucceeded deploymentStatus = "SUCCEEDED"
	DeploymentStatusFailed    deploymentStatus = "FAILED"
	DeploymentStatusUnknown   deploymentStatus = "UNKNOWN"
)

// Deployer is an interface to be implemented by any component capable of
// deploying resource to Azure using an ARM template
//type ArmDeployer interface {
//	Deploy(
//		deploymentName string,
//		resourceGroupName string,
//		location string,
//		template []byte,
//		goParams interface{},
//		armParams map[string]interface{},
//		tags map[string]string,
//	) (map[string]interface{}, error)
//	Delete(deploymentName string, resourceGroupName string) error
//}

// deployer is an ARM-based implementation of the Deployer interface
type ArmDeployer struct {
	groupsClient      resourcesSDK.GroupsClient
	deploymentsClient resourcesSDK.DeploymentsClient
}

func NewDeployer(
	groupsClient resourcesSDK.GroupsClient,
	deploymentsClient resourcesSDK.DeploymentsClient,
) *ArmDeployer {
	return &ArmDeployer{
		groupsClient:      groupsClient,
		deploymentsClient: deploymentsClient,
	}
}

func (d *ArmDeployer) GetDeploymentAndStatus(
	deploymentName string,
	resourceGroupName string,
) (*resourcesSDK.DeploymentExtended, deploymentStatus, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	deployment, err := d.deploymentsClient.Get(
		ctx,
		resourceGroupName,
		deploymentName,
	)
	if err != nil {
		detailedErr, ok := err.(autorest.DetailedError)
		if !ok || detailedErr.StatusCode != http.StatusNotFound {
			return nil, "", err
		}
		return nil, DeploymentStatusNotFound, nil
	}
	switch *deployment.Properties.ProvisioningState {
	case "Running":
		return &deployment, DeploymentStatusRunning, nil
	case "Succeeded":
		return &deployment, DeploymentStatusSucceeded, nil
	case "Failed":
		return &deployment, DeploymentStatusFailed, nil
	default:
		return &deployment, DeploymentStatusUnknown, nil
	}
}

func (d *ArmDeployer) Deploy(
	deploymentName string,
	resourceGroupName string,
	templateFileName string,
	armParams map[string]interface{}) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := Asset(templateFileName)
	if err != nil {
		log.Error(err, "")
	}

	log.Info("Template data", templateFileName, string(data))
	var armTemplateMap map[string]interface{}
	err = json.Unmarshal(data, &armTemplateMap)
	if err != nil {
		return err
	}

	result, err := d.deploymentsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		resourcesSDK.Deployment{
			Properties: &resourcesSDK.DeploymentProperties{
				Template: armTemplateMap,
				Parameters: &armParams,
				Mode: resourcesSDK.Incremental,
			},
		})
	if err != nil {
		log.Error(err, "")
	}

	log.Info("Deployment Result", "value", result)
	return nil
}

func (d *ArmDeployer) Delete(
	deploymentName string,
	resourceGroupName string,
) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := d.deploymentsClient.Delete(
		ctx,
		resourceGroupName,
		deploymentName,
	)
	if err != nil {
		return err
	}
	if err := result.WaitForCompletion(
		ctx,
		d.deploymentsClient.Client,
	); err != nil {
		return err
	}
	return nil
}