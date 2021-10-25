package service

import (
	"encoding/json"
	"fmt"
)

type ServicePlacementPolicy interface {
}

func unmarshalServicePlacementPolicy(body []byte) (ServicePlacementPolicy, error) {
	type intermediateType struct {
		Type string `json:"type"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(body, &intermediate); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	switch intermediate.Type {

	case "InvalidDomain":
		{
			var out ServicePlacementInvalidDomainPolicy
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "ServicePlacementInvalidDomainPolicy", err)
			}
			return &out, nil
		}

	case "NonPartiallyPlaceService":
		{
			var out ServicePlacementNonPartiallyPlaceServicePolicy
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "ServicePlacementNonPartiallyPlaceServicePolicy", err)
			}
			return &out, nil
		}

	case "PreferredPrimaryDomain":
		{
			var out ServicePlacementPreferPrimaryDomainPolicy
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "ServicePlacementPreferPrimaryDomainPolicy", err)
			}
			return &out, nil
		}

	case "RequiredDomainDistribution":
		{
			var out ServicePlacementRequireDomainDistributionPolicy
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "ServicePlacementRequireDomainDistributionPolicy", err)
			}
			return &out, nil
		}

	case "RequiredDomain":
		{
			var out ServicePlacementRequiredDomainPolicy
			if err := json.Unmarshal(body, &out); err != nil {
				return nil, fmt.Errorf("unmarshalling into %q: %+v", "ServicePlacementRequiredDomainPolicy", err)
			}
			return &out, nil
		}

	}

	return nil, fmt.Errorf("unknown value for Type: %q", intermediate.Type)
}
