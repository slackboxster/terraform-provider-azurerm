package managedclusterversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = EnvironmentId{}

func TestEnvironmentIDFormatter(t *testing.T) {
	actual := NewEnvironmentID("{subscriptionId}", "{location}", "{environment}").ID()
	expected := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParseEnvironmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnvironmentId
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
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}",
			Expected: &EnvironmentId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{environment}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/PROVIDERS/MICROSOFT.SERVICEFABRIC/LOCATIONS/{LOCATION}/ENVIRONMENTS/{ENVIRONMENT}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnvironmentID(v.Input)
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

func TestParseEnvironmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnvironmentId
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
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}",
			Expected: &EnvironmentId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{environment}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}",
			Expected: &EnvironmentId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{environment}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LOCATIONS/{location}/ENVIRONMENTS/{environment}",
			Expected: &EnvironmentId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{environment}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/LoCaTiOnS/{location}/EnViRoNmEnTs/{environment}",
			Expected: &EnvironmentId{
				SubscriptionId: "{subscriptionId}",
				LocationName:   "{location}",
				Name:           "{environment}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnvironmentIDInsensitively(v.Input)
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
