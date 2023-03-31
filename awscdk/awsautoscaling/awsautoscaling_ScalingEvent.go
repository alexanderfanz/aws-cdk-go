package awsautoscaling


// Fleet scaling events.
// Experimental.
type ScalingEvent string

const (
	// Notify when an instance was launched.
	// Experimental.
	ScalingEvent_INSTANCE_LAUNCH ScalingEvent = "INSTANCE_LAUNCH"
	// Notify when an instance was terminated.
	// Experimental.
	ScalingEvent_INSTANCE_TERMINATE ScalingEvent = "INSTANCE_TERMINATE"
	// Notify when an instance failed to terminate.
	// Experimental.
	ScalingEvent_INSTANCE_TERMINATE_ERROR ScalingEvent = "INSTANCE_TERMINATE_ERROR"
	// Notify when an instance failed to launch.
	// Experimental.
	ScalingEvent_INSTANCE_LAUNCH_ERROR ScalingEvent = "INSTANCE_LAUNCH_ERROR"
	// Send a test notification to the topic.
	// Experimental.
	ScalingEvent_TEST_NOTIFICATION ScalingEvent = "TEST_NOTIFICATION"
)

