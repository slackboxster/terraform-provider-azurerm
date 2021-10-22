package service

import (
	"encoding/json"
	"fmt"
)

type AverageServiceLoadScalingTrigger struct {
	LowerLoadThreshold float64 `json:"lowerLoadThreshold"`
	MetricName         string  `json:"metricName"`
	ScaleInterval      string  `json:"scaleInterval"`
	UpperLoadThreshold float64 `json:"upperLoadThreshold"`
	UseOnlyPrimaryLoad bool    `json:"useOnlyPrimaryLoad"`
}

func (c *AverageServiceLoadScalingTrigger) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		LowerLoadThreshold float64         `json:"lowerLoadThreshold"`
		MetricName         string          `json:"metricName"`
		ScaleInterval      string          `json:"scaleInterval"`
		UpperLoadThreshold float64         `json:"upperLoadThreshold"`
		UseOnlyPrimaryLoad bool            `json:"useOnlyPrimaryLoad"`
		Kind               json.RawMessage `json:"kind"`
	}
	var intermediate intermediateType
	if err := json.Unmarshal(input, &intermediate); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	c.LowerLoadThreshold = intermediate.LowerLoadThreshold
	c.MetricName = intermediate.MetricName
	c.ScaleInterval = intermediate.ScaleInterval
	c.UpperLoadThreshold = intermediate.UpperLoadThreshold
	c.UseOnlyPrimaryLoad = intermediate.UseOnlyPrimaryLoad

	kind, err := unmarshalServiceScalingTriggerKind(intermediate.Kind)
	if err != nil {
		return fmt.Errorf("unmarshaling kind: %+v", err)
	}
	c.Kind = kind

	return nil
}

var _ json.Marshaler = AverageServiceLoadScalingTrigger{}

func (o AverageServiceLoadScalingTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"kind":               "AverageServiceLoadTrigger",
		"lowerLoadThreshold": o.LowerLoadThreshold,
		"metricName":         o.MetricName,
		"scaleInterval":      o.ScaleInterval,
		"upperLoadThreshold": o.UpperLoadThreshold,
		"useOnlyPrimaryLoad": o.UseOnlyPrimaryLoad,
	})
}
