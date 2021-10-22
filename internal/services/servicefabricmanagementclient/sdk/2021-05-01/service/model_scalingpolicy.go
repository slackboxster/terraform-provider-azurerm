package service

type ScalingPolicy struct {
	ScalingMechanism ScalingMechanism `json:"scalingMechanism"`
	ScalingTrigger   ScalingTrigger   `json:"scalingTrigger"`
}
