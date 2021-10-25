package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ManagedClusterVersionId struct {
	SubscriptionId string
	LocationName   string
	Name           string
}

func NewManagedClusterVersionID(subscriptionId, locationName, name string) ManagedClusterVersionId {
	return ManagedClusterVersionId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		Name:           name,
	}
}

func (id ManagedClusterVersionId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Location Name %q", id.LocationName),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Managed Cluster Version", segmentsStr)
}

func (id ManagedClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/managedClusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.Name)
}

// ParseManagedClusterVersionID parses a ManagedClusterVersion ID into an ManagedClusterVersionId struct
func ParseManagedClusterVersionID(input string) (*ManagedClusterVersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ManagedClusterVersionId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.LocationName, err = id.PopSegment("locations"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("managedClusterVersions"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseManagedClusterVersionIDInsensitively parses an ManagedClusterVersion ID into an ManagedClusterVersionId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseManagedClusterVersionID method should be used instead for validation etc.
func ParseManagedClusterVersionIDInsensitively(input string) (*ManagedClusterVersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ManagedClusterVersionId{
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

	// find the correct casing for the 'managedClusterVersions' segment
	managedClusterVersionsKey := "managedClusterVersions"
	for key := range id.Path {
		if strings.EqualFold(key, managedClusterVersionsKey) {
			managedClusterVersionsKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(managedClusterVersionsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
