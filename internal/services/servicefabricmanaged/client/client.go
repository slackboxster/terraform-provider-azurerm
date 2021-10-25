package client

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/sdk/2021-05-01/managedcluster"
)

type Client struct {
	ManagedClusterClient *managedcluster.ManagedClusterClient
	tokenFunc            func(endpoint string) (autorest.Authorizer, error)
	configureClientFunc  func(c *autorest.Client, authorizer autorest.Authorizer)
}

func NewClient(o *common.ClientOptions) *Client {
	managedCluster := managedcluster.NewManagedClusterClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&managedCluster.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ManagedClusterClient: &managedCluster,
		tokenFunc:            o.TokenFunc,
		configureClientFunc:  o.ConfigureClient,
	}
}
