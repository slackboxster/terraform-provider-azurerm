package service

import (
	"encoding/json"
	"fmt"
)

type ScalingMechanism interface {
}

func unmarshalScalingMechanism(body []byte) (ScalingMechanism, error) {
	type intermediateType struct {
		Type string `json:"kind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(body, &intermediate); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	switch intermediate.Type {

	case "AddRemoveIncrementalNamedPartition":
		{
			var out AddRemoveIncrementalNamedPartitionScalingMechanism
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "AddRemoveIncrementalNamedPartitionScalingMechanism", err)
			}
			return &out, nil
		}

	case "ScalePartitionInstanceCount":
		{
			var out PartitionInstanceCountScaleMechanism
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "PartitionInstanceCountScaleMechanism", err)
			}
			return &out, nil
		}

	}

	return nil, fmt.Errorf("unknown value for Kind: %q", intermediate.Type)
}
