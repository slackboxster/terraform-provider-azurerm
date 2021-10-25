package applicationtypeversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ApplicationTypeId struct {
	SubscriptionId     string
	ResourceGroup      string
	ManagedClusterName string
	Name               string
}

func NewApplicationTypeID(subscriptionId, resourceGroup, managedClusterName, name string) ApplicationTypeId {
	return ApplicationTypeId{
		SubscriptionId:     subscriptionId,
		ResourceGroup:      resourceGroup,
		ManagedClusterName: managedClusterName,
		Name:               name,
	}
}

func (id ApplicationTypeId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Managed Cluster Name %q", id.ManagedClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Application Type", segmentsStr)
}

func (id ApplicationTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applicationTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ManagedClusterName, id.Name)
}

// ParseApplicationTypeID parses a ApplicationType ID into an ApplicationTypeId struct
func ParseApplicationTypeID(input string) (*ApplicationTypeId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ApplicationTypeId{
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
	if resourceId.Name, err = id.PopSegment("applicationTypes"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseApplicationTypeIDInsensitively parses an ApplicationType ID into an ApplicationTypeId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseApplicationTypeID method should be used instead for validation etc.
func ParseApplicationTypeIDInsensitively(input string) (*ApplicationTypeId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ApplicationTypeId{
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

	// find the correct casing for the 'applicationTypes' segment
	applicationTypesKey := "applicationTypes"
	for key := range id.Path {
		if strings.EqualFold(key, applicationTypesKey) {
			applicationTypesKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(applicationTypesKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
