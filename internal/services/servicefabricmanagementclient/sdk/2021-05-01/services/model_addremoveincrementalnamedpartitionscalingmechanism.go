package services

import (
	"encoding/json"
	"fmt"
)

type AddRemoveIncrementalNamedPartitionScalingMechanism struct {
	MaxPartitionCount int64 `json:"maxPartitionCount"`
	MinPartitionCount int64 `json:"minPartitionCount"`
	ScaleIncrement    int64 `json:"scaleIncrement"`
}

func (c *AddRemoveIncrementalNamedPartitionScalingMechanism) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		MaxPartitionCount int64           `json:"maxPartitionCount"`
		MinPartitionCount int64           `json:"minPartitionCount"`
		ScaleIncrement    int64           `json:"scaleIncrement"`
		Kind              json.RawMessage `json:"kind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.MaxPartitionCount = intermediate.MaxPartitionCount
	c.MinPartitionCount = intermediate.MinPartitionCount
	c.ScaleIncrement = intermediate.ScaleIncrement

	kind, err := unmarshalServiceScalingMechanismKind(intermediate.Kind)
	if err != nil {
		return fmt.Errorf("unmarshaling kind: %+v", err)
	}
	c.Kind = kind

	return nil
}

var _ json.Marshaler = AddRemoveIncrementalNamedPartitionScalingMechanism{}

func (o AddRemoveIncrementalNamedPartitionScalingMechanism) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"kind":              "AddRemoveIncrementalNamedPartition",
		"maxPartitionCount": o.MaxPartitionCount,
		"minPartitionCount": o.MinPartitionCount,
		"scaleIncrement":    o.ScaleIncrement,
	})
}
