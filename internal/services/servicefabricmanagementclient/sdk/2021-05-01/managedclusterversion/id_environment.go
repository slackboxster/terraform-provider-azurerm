package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type EnvironmentId struct {
	SubscriptionId string
	LocationName   string
	Name           string
}

func NewEnvironmentID(subscriptionId, locationName, name string) EnvironmentId {
	return EnvironmentId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		Name:           name,
	}
}

func (id EnvironmentId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Location Name %q", id.LocationName),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Environment", segmentsStr)
}

func (id EnvironmentId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/environments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.Name)
}

// ParseEnvironmentID parses a Environment ID into an EnvironmentId struct
func ParseEnvironmentID(input string) (*EnvironmentId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := EnvironmentId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.LocationName, err = id.PopSegment("locations"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("environments"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseEnvironmentIDInsensitively parses an Environment ID into an EnvironmentId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseEnvironmentID method should be used instead for validation etc.
func ParseEnvironmentIDInsensitively(input string) (*EnvironmentId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := EnvironmentId{
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
	if resourceId.Name, err = id.PopSegment(environmentsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
