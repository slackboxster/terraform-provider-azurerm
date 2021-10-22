package service

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/formatting"
	"github.com/hashicorp/terraform-provider-azurerm/internal/identity"
)

type ServicePlacementNonPartiallyPlaceServicePolicy struct {

}



func (c *ServicePlacementNonPartiallyPlaceServicePolicy) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
	Type json.RawMessage `json:"type"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}


	type, err := unmarshalServicePlacementPolicyType(intermediate.Type)
	if err != nil {
		return fmt.Errorf("unmarshaling type: %+v", err)
	}
	c.Type = type


	return nil
}


var _ json.Marshaler = ServicePlacementNonPartiallyPlaceServicePolicy{}

func (o ServicePlacementNonPartiallyPlaceServicePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
"type": "NonPartiallyPlaceService",
	})
}

