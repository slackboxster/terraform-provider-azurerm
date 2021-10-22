package application

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ApplicationId struct {
	SubscriptionId     string
	ResourceGroup      string
	ManagedClusterName string
	Name               string
}

func NewApplicationID(subscriptionId, resourceGroup, managedClusterName, name string) ApplicationId {
	return ApplicationId{
		SubscriptionId:     subscriptionId,
		ResourceGroup:      resourceGroup,
		ManagedClusterName: managedClusterName,
		Name:               name,
	}
}

func (id ApplicationId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Managed Cluster Name %q", id.ManagedClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Application", segmentsStr)
}

func (id ApplicationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applications/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ManagedClusterName, id.Name)
}

// ParseApplicationID parses a Application ID into an ApplicationId struct
func ParseApplicationID(input string) (*ApplicationId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ApplicationId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.ManagedClusterName, err = id.PopSegment("managedClusters"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("applications"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseApplicationIDInsensitively parses an Application ID into an ApplicationId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseApplicationID method should be used instead for validation etc.
func ParseApplicationIDInsensitively(input string) (*ApplicationId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ApplicationId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	// find the correct casing for the 'managedClusters' segment
	managedClustersKey := "managedClusters"
	for key := range id.Path {
		if strings.EqualFold(key, managedClustersKey) {
			managedClustersKey = key
			break
		}
	}
	if resourceId.ManagedClusterName, err = id.PopSegment(managedClustersKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'applications' segment
	applicationsKey := "applications"
	for key := range id.Path {
		if strings.EqualFold(key, applicationsKey) {
			applicationsKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(applicationsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
