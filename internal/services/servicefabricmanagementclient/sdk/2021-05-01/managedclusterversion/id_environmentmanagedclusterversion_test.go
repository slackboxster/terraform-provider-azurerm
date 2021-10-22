package managedclusterversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = EnvironmentManagedClusterVersionId{}

func TestEnvironmentManagedClusterVersionIDFormatter(t *testing.T) {
	actual := NewEnvironmentManagedClusterVersionID("{subscriptionId}", "{location}", "{environment}", "{clusterVersion}").ID()
	expected := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/{clusterVersion}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParseEnvironmentManagedClusterVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnvironmentManagedClusterVersionId
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
			// missing LocationName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/",
			Error: true,
		},

		{
			// missing value for LocationName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/",
			Error: true,
		},

		{
			// missing EnvironmentName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/",
			Error: true,
		},

		{
			// missing value for EnvironmentName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/",
			Error: true,
		},

		{
			// missing ManagedClusterVersionName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/",
			Error: true,
		},

		{
			// missing value for ManagedClusterVersionName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/{clusterVersion}",
			Expected: &EnvironmentManagedClusterVersionId{
				SubscriptionId:            "{subscriptionId}",
				LocationName:              "{location}",
				EnvironmentName:           "{environment}",
				ManagedClusterVersionName: "{clusterVersion}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/PROVIDERS/MICROSOFT.SERVICEFABRIC/LOCATIONS/{LOCATION}/ENVIRONMENTS/{ENVIRONMENT}/MANAGEDCLUSTERVERSIONS/{CLUSTERVERSION}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnvironmentManagedClusterVersionID(v.Input)
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
		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}
		if actual.EnvironmentName != v.Expected.EnvironmentName {
			t.Fatalf("Expected %q but got %q for EnvironmentName", v.Expected.EnvironmentName, actual.EnvironmentName)
		}
		if actual.ManagedClusterVersionName != v.Expected.ManagedClusterVersionName {
			t.Fatalf("Expected %q but got %q for ManagedClusterVersionName", v.Expected.ManagedClusterVersionName, actual.ManagedClusterVersionName)
		}
	}
}

func TestParseEnvironmentManagedClusterVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnvironmentManagedClusterVersionId
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
			// missing LocationName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/",
			Error: true,
		},

		{
			// missing value for LocationName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/",
			Error: true,
		},

		{
			// missing EnvironmentName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/",
			Error: true,
		},

		{
			// missing value for EnvironmentName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/",
			Error: true,
		},

		{
			// missing ManagedClusterVersionName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/",
			Error: true,
		},

		{
			// missing value for ManagedClusterVersionName
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/{clusterVersion}",
			Expected: &EnvironmentManagedClusterVersionId{
				SubscriptionId:            "{subscriptionId}",
				LocationName:              "{location}",
				EnvironmentName:           "{environment}",
				ManagedClusterVersionName: "{clusterVersion}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedclusterversions/{clusterVersion}",
			Expected: &EnvironmentManagedClusterVersionId{
				SubscriptionId:            "{subscriptionId}",
				LocationName:              "{location}",
				EnvironmentName:           "{environment}",
				ManagedClusterVersionName: "{clusterVersion}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LOCATIONS/{location}/ENVIRONMENTS/{environment}/MANAGEDCLUSTERVERSIONS/{clusterVersion}",
			Expected: &EnvironmentManagedClusterVersionId{
				SubscriptionId:            "{subscriptionId}",
				LocationName:              "{location}",
				EnvironmentName:           "{environment}",
				ManagedClusterVersionName: "{clusterVersion}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LoCaTiOnS/{location}/EnViRoNmEnTs/{environment}/MaNaGeDcLuStErVeRsIoNs/{clusterVersion}",
			Expected: &EnvironmentManagedClusterVersionId{
				SubscriptionId:            "{subscriptionId}",
				LocationName:              "{location}",
				EnvironmentName:           "{environment}",
				ManagedClusterVersionName: "{clusterVersion}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnvironmentManagedClusterVersionIDInsensitively(v.Input)
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
		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}
		if actual.EnvironmentName != v.Expected.EnvironmentName {
			t.Fatalf("Expected %q but got %q for EnvironmentName", v.Expected.EnvironmentName, actual.EnvironmentName)
		}
		if actual.ManagedClusterVersionName != v.Expected.ManagedClusterVersionName {
			t.Fatalf("Expected %q but got %q for ManagedClusterVersionName", v.Expected.ManagedClusterVersionName, actual.ManagedClusterVersionName)
		}
	}
}
