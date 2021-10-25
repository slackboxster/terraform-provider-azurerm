package services

import (
	"encoding/json"
	"fmt"
)

type AveragePartitionLoadScalingTrigger struct {
	LowerLoadThreshold float64 `json:"lowerLoadThreshold"`
	MetricName         string  `json:"metricName"`
	ScaleInterval      string  `json:"scaleInterval"`
	UpperLoadThreshold float64 `json:"upperLoadThreshold"`
}

func (c *AveragePartitionLoadScalingTrigger) UnmarshalJSON(input []byte) error {
	type intermediateType struct {
		LowerLoadThreshold float64         `json:"lowerLoadThreshold"`
		MetricName         string          `json:"metricName"`
		ScaleInterval      string          `json:"scaleInterval"`
		UpperLoadThreshold float64         `json:"upperLoadThreshold"`
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

	kind, err := unmarshalServiceScalingTriggerKind(intermediate.Kind)
	if err != nil {
		return fmt.Errorf("unmarshaling kind: %+v", err)
	}
	c.Kind = kind

	return nil
}

var _ json.Marshaler = AveragePartitionLoadScalingTrigger{}

func (o AveragePartitionLoadScalingTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"kind":               "AveragePartitionLoadTrigger",
		"lowerLoadThreshold": o.LowerLoadThreshold,
		"metricName":         o.MetricName,
		"scaleInterval":      o.ScaleInterval,
		"upperLoadThreshold": o.UpperLoadThreshold,
	})
}
