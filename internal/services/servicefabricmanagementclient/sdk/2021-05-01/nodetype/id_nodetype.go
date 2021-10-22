package nodetype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type NodeTypeId struct {
	SubscriptionId     string
	ResourceGroup      string
	ManagedClusterName string
	Name               string
}

func NewNodeTypeID(subscriptionId, resourceGroup, managedClusterName, name string) NodeTypeId {
	return NodeTypeId{
		SubscriptionId:     subscriptionId,
		ResourceGroup:      resourceGroup,
		ManagedClusterName: managedClusterName,
		Name:               name,
	}
}

func (id NodeTypeId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
		fmt.Sprintf("Managed Cluster Name %q", id.ManagedClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Node Type", segmentsStr)
}

func (id NodeTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/nodeTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ManagedClusterName, id.Name)
}

// ParseNodeTypeID parses a NodeType ID into an NodeTypeId struct
func ParseNodeTypeID(input string) (*NodeTypeId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := NodeTypeId{
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
	if resourceId.Name, err = id.PopSegment("nodeTypes"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// ParseNodeTypeIDInsensitively parses an NodeType ID into an NodeTypeId struct, insensitively
// This should only be used to parse an ID for rewriting to a consistent casing,
// the ParseNodeTypeID method should be used instead for validation etc.
func ParseNodeTypeIDInsensitively(input string) (*NodeTypeId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := NodeTypeId{
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

	// find the correct casing for the 'nodeTypes' segment
	nodeTypesKey := "nodeTypes"
	for key := range id.Path {
		if strings.EqualFold(key, nodeTypesKey) {
			nodeTypesKey = key
			break
		}
	}
	if resourceId.Name, err = id.PopSegment(nodeTypesKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
