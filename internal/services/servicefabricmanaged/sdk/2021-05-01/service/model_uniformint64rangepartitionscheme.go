package service

import (
	"encoding/json"
	"fmt"
)

type UniformInt64RangePartitionScheme struct {
	Count   int64 `json:"count"`
	HighKey int64 `json:"highKey"`
	LowKey  int64 `json:"lowKey"`
}

func (c *UniformInt64RangePartitionScheme) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		Count           int64           `json:"count"`
		HighKey         int64           `json:"highKey"`
		LowKey          int64           `json:"lowKey"`
		PartitionScheme json.RawMessage `json:"partitionScheme"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.Count = intermediate.Count
	c.HighKey = intermediate.HighKey
	c.LowKey = intermediate.LowKey

	partitionScheme, err := unmarshalPartitionScheme(intermediate.PartitionScheme)
	if err != nil {
		return fmt.Errorf("unmarshaling partitionScheme: %+v", err)
	}
	c.PartitionScheme = partitionScheme

	return nil
}

var _ json.Marshaler = UniformInt64RangePartitionScheme{}

func (o UniformInt64RangePartitionScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"count":           o.Count,
		"highKey":         o.HighKey,
		"lowKey":          o.LowKey,
		"partitionScheme": "UniformInt64Range",
	})
}
