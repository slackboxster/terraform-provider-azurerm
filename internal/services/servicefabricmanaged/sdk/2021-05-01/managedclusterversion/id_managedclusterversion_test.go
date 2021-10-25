package managedclusterversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = ManagedClusterVersionId{}

func TestManagedClusterVersionIDFormatter(t *testing.T) {
	actual := NewManagedClusterVersionID("{subscriptionId}", "{location}", "{clusterVersion}").ID()
	expected := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/{clusterVersion}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParseManagedClusterVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedClusterVersionId
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
			// missing Name
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/{clusterVersion}",
			Expected: &ManagedClusterVersionId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{clusterVersion}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/PROVIDERS/MICROSOFT.SERVICEFABRIC/LOCATIONS/{LOCATION}/MANAGEDCLUSTERVERSIONS/{CLUSTERVERSION}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedClusterVersionID(v.Input)
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
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}

func TestParseManagedClusterVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedClusterVersionId
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
			// missing Name
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/{clusterVersion}",
			Expected: &ManagedClusterVersionId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{clusterVersion}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedclusterversions/{clusterVersion}",
			Expected: &ManagedClusterVersionId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{clusterVersion}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LOCATIONS/{location}/MANAGEDCLUSTERVERSIONS/{clusterVersion}",
			Expected: &ManagedClusterVersionId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{clusterVersion}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LoCaTiOnS/{location}/MaNaGeDcLuStErVeRsIoNs/{clusterVersion}",
			Expected: &ManagedClusterVersionId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{clusterVersion}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedClusterVersionIDInsensitively(v.Input)
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
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}
