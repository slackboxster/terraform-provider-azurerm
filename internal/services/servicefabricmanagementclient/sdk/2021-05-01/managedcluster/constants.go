package managedcluster

type Access string

const (
	AccessAllow Access = "allow"
	AccessDeny  Access = "deny"
)

type AddonFeatures string

const (
	AddonFeaturesBackupRestoreService   AddonFeatures = "BackupRestoreService"
	AddonFeaturesDnsService             AddonFeatures = "DnsService"
	AddonFeaturesResourceMonitorService AddonFeatures = "ResourceMonitorService"
)

type ClusterState string

const (
	ClusterStateBaselineUpgrade ClusterState = "BaselineUpgrade"
	ClusterStateDeploying       ClusterState = "Deploying"
	ClusterStateReady           ClusterState = "Ready"
	ClusterStateUpgradeFailed   ClusterState = "UpgradeFailed"
	ClusterStateUpgrading       ClusterState = "Upgrading"
	ClusterStateWaitingForNodes ClusterState = "WaitingForNodes"
)

type ClusterUpgradeCadence string

const (
	ClusterUpgradeCadenceWaveOne  ClusterUpgradeCadence = "Wave1"
	ClusterUpgradeCadenceWaveTwo  ClusterUpgradeCadence = "Wave2"
	ClusterUpgradeCadenceWaveZero ClusterUpgradeCadence = "Wave0"
)

type ClusterUpgradeMode string

const (
	ClusterUpgradeModeAutomatic ClusterUpgradeMode = "Automatic"
	ClusterUpgradeModeManual    ClusterUpgradeMode = "Manual"
)

type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

type ManagedResourceProvisioningState string

const (
	ManagedResourceProvisioningStateCanceled  ManagedResourceProvisioningState = "Canceled"
	ManagedResourceProvisioningStateCreated   ManagedResourceProvisioningState = "Created"
	ManagedResourceProvisioningStateCreating  ManagedResourceProvisioningState = "Creating"
	ManagedResourceProvisioningStateDeleted   ManagedResourceProvisioningState = "Deleted"
	ManagedResourceProvisioningStateDeleting  ManagedResourceProvisioningState = "Deleting"
	ManagedResourceProvisioningStateFailed    ManagedResourceProvisioningState = "Failed"
	ManagedResourceProvisioningStateNone      ManagedResourceProvisioningState = "None"
	ManagedResourceProvisioningStateOther     ManagedResourceProvisioningState = "Other"
	ManagedResourceProvisioningStateSucceeded ManagedResourceProvisioningState = "Succeeded"
	ManagedResourceProvisioningStateUpdating  ManagedResourceProvisioningState = "Updating"
)

type NsgProtocol string

const (
	NsgProtocolAh    NsgProtocol = "ah"
	NsgProtocolEsp   NsgProtocol = "esp"
	NsgProtocolHttp  NsgProtocol = "http"
	NsgProtocolHttps NsgProtocol = "https"
	NsgProtocolIcmp  NsgProtocol = "icmp"
	NsgProtocolTcp   NsgProtocol = "tcp"
	NsgProtocolUdp   NsgProtocol = "udp"
)

type ProbeProtocol string

const (
	ProbeProtocolHttp  ProbeProtocol = "http"
	ProbeProtocolHttps ProbeProtocol = "https"
	ProbeProtocolTcp   ProbeProtocol = "tcp"
)

type Protocol string

const (
	ProtocolTcp Protocol = "tcp"
	ProtocolUdp Protocol = "udp"
)

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameStandard SkuName = "Standard"
)
