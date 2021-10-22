package service

import (
	"encoding/json"
	"fmt"
)

type StatelessServiceProperties struct {
	CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
	DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
	InstanceCount                int64                         `json:"instanceCount"`
	MinInstanceCount             *int64                        `json:"minInstanceCount,omitempty"`
	MinInstancePercentage        *int64                        `json:"minInstancePercentage,omitempty"`
	PartitionDescription         Partition                     `json:"partitionDescription"`
	PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
	ProvisioningState            *string                       `json:"provisioningState,omitempty"`
	ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
	ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
	ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
	ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
	ServiceTypeName              string                        `json:"serviceTypeName"`
}

func (c *StatelessServiceProperties) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
		DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
		InstanceCount                int64                         `json:"instanceCount"`
		MinInstanceCount             *int64                        `json:"minInstanceCount,omitempty"`
		MinInstancePercentage        *int64                        `json:"minInstancePercentage,omitempty"`
		PartitionDescription         Partition                     `json:"partitionDescription"`
		PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
		ProvisioningState            *string                       `json:"provisioningState,omitempty"`
		ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
		ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
		ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
		ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
		ServiceTypeName              string                        `json:"serviceTypeName"`
		ServiceKind                  json.RawMessage               `json:"serviceKind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.CorrelationScheme = intermediate.CorrelationScheme
	c.DefaultMoveCost = intermediate.DefaultMoveCost
	c.InstanceCount = intermediate.InstanceCount
	c.MinInstanceCount = intermediate.MinInstanceCount
	c.MinInstancePercentage = intermediate.MinInstancePercentage
	c.PartitionDescription = intermediate.PartitionDescription
	c.PlacementConstraints = intermediate.PlacementConstraints
	c.ProvisioningState = intermediate.ProvisioningState
	c.ScalingPolicies = intermediate.ScalingPolicies
	c.ServiceLoadMetrics = intermediate.ServiceLoadMetrics
	c.ServicePackageActivationMode = intermediate.ServicePackageActivationMode
	c.ServicePlacementPolicies = intermediate.ServicePlacementPolicies
	c.ServiceTypeName = intermediate.ServiceTypeName

	serviceKind, err := unmarshalServiceKind(intermediate.ServiceKind)
	if err != nil {
		return fmt.Errorf("unmarshaling serviceKind: %+v", err)
	}
	c.ServiceKind = serviceKind

	return nil
}

var _ json.Marshaler = StatelessServiceProperties{}

func (o StatelessServiceProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"correlationScheme":            o.CorrelationScheme,
		"defaultMoveCost":              o.DefaultMoveCost,
		"instanceCount":                o.InstanceCount,
		"minInstanceCount":             o.MinInstanceCount,
		"minInstancePercentage":        o.MinInstancePercentage,
		"partitionDescription":         o.PartitionDescription,
		"placementConstraints":         o.PlacementConstraints,
		"provisioningState":            o.ProvisioningState,
		"scalingPolicies":              o.ScalingPolicies,
		"serviceKind":                  "Stateless",
		"serviceLoadMetrics":           o.ServiceLoadMetrics,
		"servicePackageActivationMode": o.ServicePackageActivationMode,
		"servicePlacementPolicies":     o.ServicePlacementPolicies,
		"serviceTypeName":              o.ServiceTypeName,
	})
}
