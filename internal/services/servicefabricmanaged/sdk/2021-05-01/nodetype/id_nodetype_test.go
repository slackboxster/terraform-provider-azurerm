package nodetype

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = NodeTypeId{}

func TestNodeTypeIDFormatter(t *testing.T) {
	actual := NewNodeTypeID("{subscriptionId}", "{resourceGroupName}", "{clusterName}", "{nodeTypeName}").ID()
	expected := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParseNodeTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeTypeId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing ManagedClusterName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/",
			Error: true,
		},

		{
			// missing value for ManagedClusterName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/",
			Error: true,
		},

		{
			// missing Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}",
			Expected: &NodeTypeId{
				SubscriptionId:     "{subscriptionId}",
				ResourceGroup:      "{resourceGroupName}",
				ManagedClusterName: "{clusterName}",
				Name:               "{nodeTypeName}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/RESOURCEGROUPS/{RESOURCEGROUPNAME}/PROVIDERS/MICROSOFT.SERVICEFABRIC/MANAGEDCLUSTERS/{CLUSTERNAME}/NODETYPES/{NODETYPENAME}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeTypeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.ManagedClusterName != v.Expected.ManagedClusterName {
			t.Fatalf("Expected %q but got %q for ManagedClusterName", v.Expected.ManagedClusterName, actual.ManagedClusterName)
		}
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}

func TestParseNodeTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeTypeId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing ManagedClusterName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/",
			Error: true,
		},

		{
			// missing value for ManagedClusterName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/",
			Error: true,
		},

		{
			// missing Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}",
			Expected: &NodeTypeId{
				SubscriptionId:     "{subscriptionId}",
				ResourceGroup:      "{resourceGroupName}",
				ManagedClusterName: "{clusterName}",
				Name:               "{nodeTypeName}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedclusters/{clusterName}/nodetypes/{nodeTypeName}",
			Expected: &NodeTypeId{
				SubscriptionId:     "{subscriptionId}",
				ResourceGroup:      "{resourceGroupName}",
				ManagedClusterName: "{clusterName}",
				Name:               "{nodeTypeName}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/MANAGEDCLUSTERS/{clusterName}/NODETYPES/{nodeTypeName}",
			Expected: &NodeTypeId{
				SubscriptionId:     "{subscriptionId}",
				ResourceGroup:      "{resourceGroupName}",
				ManagedClusterName: "{clusterName}",
				Name:               "{nodeTypeName}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/MaNaGeDcLuStErS/{clusterName}/NoDeTyPeS/{nodeTypeName}",
			Expected: &NodeTypeId{
				SubscriptionId:     "{subscriptionId}",
				ResourceGroup:      "{resourceGroupName}",
				ManagedClusterName: "{clusterName}",
				Name:               "{nodeTypeName}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeTypeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.ManagedClusterName != v.Expected.ManagedClusterName {
			t.Fatalf("Expected %q but got %q for ManagedClusterName", v.Expected.ManagedClusterName, actual.ManagedClusterName)
		}
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}
