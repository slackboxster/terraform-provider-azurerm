package applicationtypeversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type VersionId struct {
	SubscriptionId      string
	ResourceGroup       string
	ManagedClusterName  string
	ApplicationTypeName string
	Name                string
}

func NewVersionID(subscriptionId, resourceGroup, managedClusterName, applicationTypeName, name string) VersionId {
	return VersionId{
		SubscriptionId:      subscriptionId,
		ResourceGroup:       resourceGroup,
		ManagedClusterName:  managedClusterName,
		ApplicationTypeName: applicationTypeName,
		Name:                name,
	}
}

func (id VersionId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Application Type Name %q", id.ApplicationTypeName),
		fmt.Sprintf("Managed Cluster Name %q", id.ManagedClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Version", segmentsStr)
}

func (id VersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applicationTypes/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ManagedClusterName, id.ApplicationTypeName, id.Name)
}

// ParseVersionID parses a Version ID into an VersionId struct
func ParseVersionID(input string) (*VersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := VersionId{
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
	if resourceId.ApplicationTypeName, err = id.PopSegment("applicationTypes"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("versions"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseVersionIDInsensitively parses an Version ID into an VersionId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseVersionID method should be used instead for validation etc.
func ParseVersionIDInsensitively(input string) (*VersionId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := VersionId{
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
	if resourceId.ApplicationTypeName, err = id.PopSegment(applicationTypesKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'versions' segment
	versionsKey := "versions"
	for key := range id.Path {
		if strings.EqualFold(key, versionsKey) {
			versionsKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(versionsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
