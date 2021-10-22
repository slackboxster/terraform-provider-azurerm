package service

import (
	"encoding/json"
	"fmt"
)

type SingletonPartitionScheme struct {
}

func (c *SingletonPartitionScheme) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		PartitionScheme json.RawMessage `json:"partitionScheme"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	partitionScheme, err := unmarshalPartitionScheme(intermediate.PartitionScheme)
	if err != nil {
		return fmt.Errorf("unmarshaling partitionScheme: %+v", err)
	}
	c.PartitionScheme = partitionScheme

	return nil
}

var _ json.Marshaler = SingletonPartitionScheme{}

func (o SingletonPartitionScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"partitionScheme": "Singleton",
	})
}
