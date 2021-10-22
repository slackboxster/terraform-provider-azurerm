package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type EnvironmentManagedClusterVersionId struct {
	SubscriptionId            string
	LocationName              string
	EnvironmentName           string
	ManagedClusterVersionName string
}

func NewEnvironmentManagedClusterVersionID(subscriptionId, locationName, environmentName, managedClusterVersionName string) EnvironmentManagedClusterVersionId {
	return EnvironmentManagedClusterVersionId{
		SubscriptionId:            subscriptionId,
		LocationName:              locationName,
		EnvironmentName:           environmentName,
		ManagedClusterVersionName: managedClusterVersionName,
	}
}

func (id EnvironmentManagedClusterVersionId) String() string {
	segments := []string{
		fmt.Sprintf("Managed Cluster Version Name %q", id.ManagedClusterVersionName),
		fmt.Sprintf("Environment Name %q", id.EnvironmentName),
		fmt.Sprintf("Location Name %q", id.LocationName),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Environment Managed Cluster Version", segmentsStr)
}

func (id EnvironmentManagedClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/environments/%s/managedClusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.EnvironmentName, id.ManagedClusterVersionName)
}

// ParseEnvironmentManagedClusterVersionID parses a EnvironmentManagedClusterVersion ID into an EnvironmentManagedClusterVersionId struct
func ParseEnvironmentManagedClusterVersionID(input string) (*EnvironmentManagedClusterVersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := EnvironmentManagedClusterVersionId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.LocationName, err = id.PopSegment("locations"); err != nil {
		return nil, err
	}
	if resourceId.EnvironmentName, err = id.PopSegment("environments"); err != nil {
		return nil, err
	}
	if resourceId.ManagedClusterVersionName, err = id.PopSegment("managedClusterVersions"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseEnvironmentManagedClusterVersionIDInsensitively parses an EnvironmentManagedClusterVersion ID into an EnvironmentManagedClusterVersionId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseEnvironmentManagedClusterVersionID method should be used instead for validation etc.
func ParseEnvironmentManagedClusterVersionIDInsensitively(input string) (*EnvironmentManagedClusterVersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := EnvironmentManagedClusterVersionId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	// find the correct casing for the 'locations' segment
	locationsKey := "locations"
	for key := range id.Path {
		if strings.EqualFold(key, locationsKey) {
			locationsKey = key
			break
		}
	}
	if resourceId.LocationName, err = id.PopSegment(locationsKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'environments' segment
	environmentsKey := "environments"
	for key := range id.Path {
		if strings.EqualFold(key, environmentsKey) {
			environmentsKey = key
			break
		}
	}
	if resourceId.EnvironmentName, err = id.PopSegment(environmentsKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'managedClusterVersions' segment
	managedClusterVersionsKey := "managedClusterVersions"
	for key := range id.Path {
		if strings.EqualFold(key, managedClusterVersionsKey) {
			managedClusterVersionsKey = key
			break
		}
	}
	if resourceId.ManagedClusterVersionName, err = id.PopSegment(managedClusterVersionsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
