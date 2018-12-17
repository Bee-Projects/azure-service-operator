package azure

import (
	"fmt"
	"github.com/Azure/open-service-broker-azure/pkg/version"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

// GetBearerTokenAuthorizer returns a *autorest.BearerAuthorizer used for
// authenticating outbound requests to the Azure APIs
func GetBearerTokenAuthorizer(
	azureEnvironment azure.Environment,
	tenantID string,
	clientID string,
	clientSecret string,
) (*autorest.BearerAuthorizer, error) {
	// Get a token used for authorizing requests to Azure
	oauthConfig, err := adal.NewOAuthConfig(
		azureEnvironment.ActiveDirectoryEndpoint,
		tenantID,
	)
	if err != nil {
		return nil, fmt.Errorf("error building oauth config: %s", err)
	}
	spt, err := adal.NewServicePrincipalToken(
		*oauthConfig,
		clientID,
		clientSecret,
		azureEnvironment.ResourceManagerEndpoint,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting service principal token: %s", err)
	}
	return autorest.NewBearerAuthorizer(spt), nil
}

func GetUserAgent(client autorest.Client) string {
	return fmt.Sprintf(
		"%s; azure-service-operator/%s",
		client.UserAgent,
		version.GetVersion(),
	)
}

func GetDeploymentName(name string, resourceGroupName string) string {
	return name + "-" + resourceGroupName
}