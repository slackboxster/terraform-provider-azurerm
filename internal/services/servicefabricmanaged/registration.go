package servicefabricmanaged

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type Registration struct{}

var _ sdk.TypedServiceRegistration = Registration{}
var _ sdk.UntypedServiceRegistration = Registration{}

func (r Registration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{
		ClusterResource{},
	}
}

func (r Registration) Resources() []sdk.Resource {
	return []sdk.Resource{}
}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Service Fabric Managed Clusters"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Service Fabric Managed Clusters",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return nil
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return nil
}