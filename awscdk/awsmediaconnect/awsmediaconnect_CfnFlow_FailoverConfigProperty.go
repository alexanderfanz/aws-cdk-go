package awsmediaconnect


// The settings for source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConfigProperty := &failoverConfigProperty{
//   	recoveryWindow: jsii.Number(123),
//   	state: jsii.String("state"),
//   }
//
type CfnFlow_FailoverConfigProperty struct {
	// The size of the buffer (delay) that the service maintains.
	//
	// A larger buffer means a longer delay in transmitting the stream, but more room for error correction. A smaller buffer means a shorter delay, but less room for error correction.
	RecoveryWindow *float64 `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// The state of source failover on the flow.
	//
	// If the state is disabled, the flow can have only one source. If the state is enabled, the flow can have one or two sources.
	State *string `field:"optional" json:"state" yaml:"state"`
}
