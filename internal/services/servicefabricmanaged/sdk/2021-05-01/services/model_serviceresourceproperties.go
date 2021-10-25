package services

import (
	"encoding/json"
	"fmt"
)

type ServiceResourceProperties interface {
}

func unmarshalServiceResourceProperties(body []byte) (ServiceResourceProperties, error) {
	type intermediateType struct {
		Type string `json:"serviceKind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(body, &intermediate); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	switch intermediate.Type {

	case "Stateful":
		{
			var out StatefulServiceProperties
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "StatefulServiceProperties", err)
			}
			return &out, nil
		}

	case "Stateless":
		{
			var out StatelessServiceProperties
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "StatelessServiceProperties", err)
			}
			return &out, nil
		}

	}

	return nil, fmt.Errorf("unknown value for ServiceKind: %q", intermediate.Type)
}
