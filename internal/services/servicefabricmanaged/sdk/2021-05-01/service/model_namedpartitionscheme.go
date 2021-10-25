package service

import (
	"encoding/json"
	"fmt"
)

type NamedPartitionScheme struct {
	Names []string `json:"names"`
}

func (c *NamedPartitionScheme) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		Names           []string        `json:"names"`
		PartitionScheme json.RawMessage `json:"partitionScheme"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.Names = intermediate.Names

	partitionScheme, err := unmarshalPartitionScheme(intermediate.PartitionScheme)
	if err != nil {
		return fmt.Errorf("unmarshaling partitionScheme: %+v", err)
	}
	c.PartitionScheme = partitionScheme

	return nil
}

var _ json.Marshaler = NamedPartitionScheme{}

func (o NamedPartitionScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"names":           o.Names,
		"partitionScheme": "Named",
	})
}
