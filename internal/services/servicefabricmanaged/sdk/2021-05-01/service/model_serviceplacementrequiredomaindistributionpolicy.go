package service

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/formatting"
	"github.com/hashicorp/terraform-provider-azurerm/internal/identity"
)

type ServicePlacementRequireDomainDistributionPolicy struct {
	DomainName string `json:"domainName"`
}



func (c *ServicePlacementRequireDomainDistributionPolicy) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
	DomainName string `json:"domainName"`
	Type json.RawMessage `json:"type"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.DomainName = intermediate.DomainName

	type, err := unmarshalServicePlacementPolicyType(intermediate.Type)
	if err != nil {
		return fmt.Errorf("unmarshaling type: %+v", err)
	}
	c.Type = type


	return nil
}


var _ json.Marshaler = ServicePlacementRequireDomainDistributionPolicy{}

func (o ServicePlacementRequireDomainDistributionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
"domainName": o.DomainName,
"type": "RequiredDomainDistribution",
	})
}

