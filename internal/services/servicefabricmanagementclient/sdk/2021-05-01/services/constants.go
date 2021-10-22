package services

type MoveCost string

const (
	MoveCostHigh   MoveCost = "High"
	MoveCostLow    MoveCost = "Low"
	MoveCostMedium MoveCost = "Medium"
	MoveCostZero   MoveCost = "Zero"
)

type PartitionScheme string

const (
	PartitionSchemeNamed                  PartitionScheme = "Named"
	PartitionSchemeSingleton              PartitionScheme = "Singleton"
	PartitionSchemeUniformIntSixFourRange PartitionScheme = "UniformInt64Range"
)

type ServiceCorrelationScheme string

const (
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = "AlignedAffinity"
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

type ServiceKind string

const (
	ServiceKindStateful  ServiceKind = "Stateful"
	ServiceKindStateless ServiceKind = "Stateless"
)

type ServiceLoadMetricWeight string

const (
	ServiceLoadMetricWeightHigh   ServiceLoadMetricWeight = "High"
	ServiceLoadMetricWeightLow    ServiceLoadMetricWeight = "Low"
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = "Medium"
	ServiceLoadMetricWeightZero   ServiceLoadMetricWeight = "Zero"
)

type ServicePackageActivationMode string

const (
	ServicePackageActivationModeExclusiveProcess ServicePackageActivationMode = "ExclusiveProcess"
	ServicePackageActivationModeSharedProcess    ServicePackageActivationMode = "SharedProcess"
)

type ServicePlacementPolicyType string

const (
	ServicePlacementPolicyTypeInvalidDomain              ServicePlacementPolicyType = "InvalidDomain"
	ServicePlacementPolicyTypeNonPartiallyPlaceService   ServicePlacementPolicyType = "NonPartiallyPlaceService"
	ServicePlacementPolicyTypePreferredPrimaryDomain     ServicePlacementPolicyType = "PreferredPrimaryDomain"
	ServicePlacementPolicyTypeRequiredDomain             ServicePlacementPolicyType = "RequiredDomain"
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

type ServiceScalingMechanismKind string

const (
	ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition ServiceScalingMechanismKind = "AddRemoveIncrementalNamedPartition"
	ServiceScalingMechanismKindScalePartitionInstanceCount        ServiceScalingMechanismKind = "ScalePartitionInstanceCount"
)

type ServiceScalingTriggerKind string

const (
	ServiceScalingTriggerKindAveragePartitionLoad ServiceScalingTriggerKind = "AveragePartitionLoad"
	ServiceScalingTriggerKindAverageServiceLoad   ServiceScalingTriggerKind = "AverageServiceLoad"
)
