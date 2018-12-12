package arm

type deploymentStatus string

const (
	deploymentStatusNotFound  deploymentStatus = "NOT_FOUND"
	deploymentStatusRunning   deploymentStatus = "RUNNING"
	deploymentStatusSucceeded deploymentStatus = "SUCCEEDED"
	deploymentStatusFailed    deploymentStatus = "FAILED"
	deploymentStatusUnknown   deploymentStatus = "UNKNOWN"
)
