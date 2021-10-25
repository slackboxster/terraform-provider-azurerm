package service

import (
	"encoding/json"
	"fmt"
)

type StatefulServiceProperties struct {
	CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
	DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
	HasPersistedState            *bool                         `json:"hasPersistedState,omitempty"`
	MinReplicaSetSize            *int64                        `json:"minReplicaSetSize,omitempty"`
	PartitionDescription         Partition                     `json:"partitionDescription"`
	PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
	ProvisioningState            *string                       `json:"provisioningState,omitempty"`
	QuorumLossWaitDuration       *string                       `json:"quorumLossWaitDuration,omitempty"`
	ReplicaRestartWaitDuration   *string                       `json:"replicaRestartWaitDuration,omitempty"`
	ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
	ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
	ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
	ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
	ServicePlacementTimeLimit    *string                       `json:"servicePlacementTimeLimit,omitempty"`
	ServiceTypeName              string                        `json:"serviceTypeName"`
	StandByReplicaKeepDuration   *string                       `json:"standByReplicaKeepDuration,omitempty"`
	TargetReplicaSetSize         *int64                        `json:"targetReplicaSetSize,omitempty"`
}

func (c *StatefulServiceProperties) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
		DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
		HasPersistedState            *bool                         `json:"hasPersistedState,omitempty"`
		MinReplicaSetSize            *int64                        `json:"minReplicaSetSize,omitempty"`
		PartitionDescription         Partition                     `json:"partitionDescription"`
		PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
		ProvisioningState            *string                       `json:"provisioningState,omitempty"`
		QuorumLossWaitDuration       *string                       `json:"quorumLossWaitDuration,omitempty"`
		ReplicaRestartWaitDuration   *string                       `json:"replicaRestartWaitDuration,omitempty"`
		ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
		ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
		ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
		ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
		ServicePlacementTimeLimit    *string                       `json:"servicePlacementTimeLimit,omitempty"`
		ServiceTypeName              string                        `json:"serviceTypeName"`
		StandByReplicaKeepDuration   *string                       `json:"standByReplicaKeepDuration,omitempty"`
		TargetReplicaSetSize         *int64                        `json:"targetReplicaSetSize,omitempty"`
		ServiceKind                  json.RawMessage               `json:"serviceKind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.CorrelationScheme = intermediate.CorrelationScheme
	c.DefaultMoveCost = intermediate.DefaultMoveCost
	c.HasPersistedState = intermediate.HasPersistedState
	c.MinReplicaSetSize = intermediate.MinReplicaSetSize
	c.PartitionDescription = intermediate.PartitionDescription
	c.PlacementConstraints = intermediate.PlacementConstraints
	c.ProvisioningState = intermediate.ProvisioningState
	c.QuorumLossWaitDuration = intermediate.QuorumLossWaitDuration
	c.ReplicaRestartWaitDuration = intermediate.ReplicaRestartWaitDuration
	c.ScalingPolicies = intermediate.ScalingPolicies
	c.ServiceLoadMetrics = intermediate.ServiceLoadMetrics
	c.ServicePackageActivationMode = intermediate.ServicePackageActivationMode
	c.ServicePlacementPolicies = intermediate.ServicePlacementPolicies
	c.ServicePlacementTimeLimit = intermediate.ServicePlacementTimeLimit
	c.ServiceTypeName = intermediate.ServiceTypeName
	c.StandByReplicaKeepDuration = intermediate.StandByReplicaKeepDuration
	c.TargetReplicaSetSize = intermediate.TargetReplicaSetSize

	serviceKind, err := unmarshalServiceKind(intermediate.ServiceKind)
	if err != nil {
		return fmt.Errorf("unmarshaling serviceKind: %+v", err)
	}
	c.ServiceKind = serviceKind

	return nil
}

var _ json.Marshaler = StatefulServiceProperties{}

func (o StatefulServiceProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"correlationScheme":            o.CorrelationScheme,
		"defaultMoveCost":              o.DefaultMoveCost,
		"hasPersistedState":            o.HasPersistedState,
		"minReplicaSetSize":            o.MinReplicaSetSize,
		"partitionDescription":         o.PartitionDescription,
		"placementConstraints":         o.PlacementConstraints,
		"provisioningState":            o.ProvisioningState,
		"quorumLossWaitDuration":       o.QuorumLossWaitDuration,
		"replicaRestartWaitDuration":   o.ReplicaRestartWaitDuration,
		"scalingPolicies":              o.ScalingPolicies,
		"serviceKind":                  "Stateful",
		"serviceLoadMetrics":           o.ServiceLoadMetrics,
		"servicePackageActivationMode": o.ServicePackageActivationMode,
		"servicePlacementPolicies":     o.ServicePlacementPolicies,
		"servicePlacementTimeLimit":    o.ServicePlacementTimeLimit,
		"serviceTypeName":              o.ServiceTypeName,
		"standByReplicaKeepDuration":   o.StandByReplicaKeepDuration,
		"targetReplicaSetSize":         o.TargetReplicaSetSize,
	})
}
