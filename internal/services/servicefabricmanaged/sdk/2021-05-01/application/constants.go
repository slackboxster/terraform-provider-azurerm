package application

type FailureAction string

const (
	FailureActionManual   FailureAction = "Manual"
	FailureActionRollback FailureAction = "Rollback"
)

type RollingUpgradeMode string

const (
	RollingUpgradeModeMonitored       RollingUpgradeMode = "Monitored"
	RollingUpgradeModeUnmonitoredAuto RollingUpgradeMode = "UnmonitoredAuto"
)
