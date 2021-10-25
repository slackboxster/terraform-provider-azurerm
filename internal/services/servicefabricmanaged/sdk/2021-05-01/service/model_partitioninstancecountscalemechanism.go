package service

import (
	"encoding/json"
	"fmt"
)

type PartitionInstanceCountScaleMechanism struct {
	MaxInstanceCount int64 `json:"maxInstanceCount"`
	MinInstanceCount int64 `json:"minInstanceCount"`
	ScaleIncrement   int64 `json:"scaleIncrement"`
}

func (c *PartitionInstanceCountScaleMechanism) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		MaxInstanceCount int64           `json:"maxInstanceCount"`
		MinInstanceCount int64           `json:"minInstanceCount"`
		ScaleIncrement   int64           `json:"scaleIncrement"`
		Kind             json.RawMessage `json:"kind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.MaxInstanceCount = intermediate.MaxInstanceCount
	c.MinInstanceCount = intermediate.MinInstanceCount
	c.ScaleIncrement = intermediate.ScaleIncrement

	kind, err := unmarshalServiceScalingMechanismKind(intermediate.Kind)
	if err != nil {
		return fmt.Errorf("unmarshaling kind: %+v", err)
	}
	c.Kind = kind

	return nil
}

var _ json.Marshaler = PartitionInstanceCountScaleMechanism{}

func (o PartitionInstanceCountScaleMechanism) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"kind":             "ScalePartitionInstanceCount",
		"maxInstanceCount": o.MaxInstanceCount,
		"minInstanceCount": o.MinInstanceCount,
		"scaleIncrement":   o.ScaleIncrement,
	})
}
