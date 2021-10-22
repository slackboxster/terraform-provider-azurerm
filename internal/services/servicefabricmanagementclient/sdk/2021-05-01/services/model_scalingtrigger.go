package services

import (
	"encoding/json"
	"fmt"
)

type ScalingTrigger interface {
}

func unmarshalScalingTrigger(body []byte) (ScalingTrigger, error) {
	type intermediateType struct {
		Type string `json:"kind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(body, &intermediate); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	switch intermediate.Type {

	case "AveragePartitionLoadTrigger":
		{
			var out AveragePartitionLoadScalingTrigger
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "AveragePartitionLoadScalingTrigger", err)
			}
			return &out, nil
		}

	case "AverageServiceLoadTrigger":
		{
			var out AverageServiceLoadScalingTrigger
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "AverageServiceLoadScalingTrigger", err)
			}
			return &out, nil
		}

	}

	return nil, fmt.Errorf("unknown value for Kind: %q", intermediate.Type)
}
