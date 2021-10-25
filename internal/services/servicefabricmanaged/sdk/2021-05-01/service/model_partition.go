package service

import (
	"encoding/json"
	"fmt"
)

type Partition interface {
}

func unmarshalPartition(body []byte) (Partition, error) {
	type intermediateType struct {
		Type string `json:"partitionScheme"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(body, &intermediate); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	switch intermediate.Type {

	case "Named":
		{
			var out NamedPartitionScheme
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "NamedPartitionScheme", err)
			}
			return &out, nil
		}

	case "Singleton":
		{
			var out SingletonPartitionScheme
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "SingletonPartitionScheme", err)
			}
			return &out, nil
		}

	case "UniformInt64Range":
		{
			var out UniformInt64RangePartitionScheme
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "UniformInt64RangePartitionScheme", err)
			}
			return &out, nil
		}

	}

	return nil, fmt.Errorf("unknown value for PartitionScheme: %q", intermediate.Type)
}
