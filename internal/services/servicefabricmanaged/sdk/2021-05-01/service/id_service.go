package service

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ServiceId struct {
	SubscriptionId     string
	ResourceGroup      string
	ManagedClusterName string
	ApplicationName    string
	Name               string
}

func NewServiceID(subscriptionId, resourceGroup, managedClusterName, applicationName, name string) ServiceId {
	return ServiceId{
		SubscriptionId:     subscriptionId,
		ResourceGroup:      resourceGroup,
		ManagedClusterName: managedClusterName,
		ApplicationName:    applicationName,
		Name:               name,
	}
}

func (id ServiceId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Application Name %q", id.ApplicationName),
		fmt.Sprintf("Managed Cluster Name %q", id.ManagedClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Service", segmentsStr)
}

func (id ServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applications/%s/services/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ManagedClusterName, id.ApplicationName, id.Name)
}

// ParseServiceID parses a Service ID into an ServiceId struct
func ParseServiceID(input string) (*ServiceId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ServiceId{
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
	if resourceId.ApplicationName, err = id.PopSegment("applications"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("services"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseServiceIDInsensitively parses an Service ID into an ServiceId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseServiceID method should be used instead for validation etc.
func ParseServiceIDInsensitively(input string) (*ServiceId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ServiceId{
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
	if resourceId.ApplicationName, err = id.PopSegment(applicationsKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'services' segment
	servicesKey := "services"
	for key := range id.Path {
		if strings.EqualFold(key, servicesKey) {
			servicesKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(servicesKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
