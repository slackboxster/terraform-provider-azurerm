package servicefabricmanaged_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/utils"

	"github.com/hashicorp/go-azure-helpers/response"

	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/sdk/2021-05-01/managedcluster"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type ClusterResource struct{}

func TestAccServiceFabricManagedCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_service_fabric_managed_cluster", "test")
	r := ClusterResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep("password"),
	})
}

func (r ClusterResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	resourceID, err := managedcluster.ParseManagedClusterID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("while parsing resource ID: %+v", err)
	}

	client := clients.ServiceFabricManaged.ManagedClusterClient
	resp, err := client.Get(ctx, *resourceID)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("while checking for cluster's %q existence: %+v", resourceID.String(), err)
	}
	return utils.Bool(resp.HttpResponse.StatusCode == 200), nil
}

func (r ClusterResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-sfmc-%[1]d"
  location = "%[2]s"
}

resource "azurerm_service_fabric_managed_cluster" "test" {
  name                = "testacc-sfmc-%[3]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard"
  username            = "testUser"
  password            = "NotV3ryS3cur3P@$$w0rd"
  dns_service         = true

  client_connection_port = 12345
  http_gateway_port      = 23456

  lb_rules {
    backend_port       = 8000
    frontend_port      = 443
    probe_protocol     = "http"
    protocol           = "tcp"
    probe_request_path = "/"
  }

  node_type {
    data_disk_size = 128
    name = "test1"
    primary = true
    application_port_range = "7000-9000"
    ephemeral_port_range = "10000-20000"

    vm_size = "Standard_DS2_v2"
    vm_image_publisher = "MicrosoftWindowsServer"
    vm_image_sku = "2016-Datacenter"
    vm_image_offer = "WindowsServer"
    vm_image_version = "latest"
    vm_instance_count = 5
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}
