package nodetype

type DiskType string

const (
	DiskTypePremiumLRS     DiskType = "Premium_LRS"
	DiskTypeStandardLRS    DiskType = "Standard_LRS"
	DiskTypeStandardSSDLRS DiskType = "StandardSSD_LRS"
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
