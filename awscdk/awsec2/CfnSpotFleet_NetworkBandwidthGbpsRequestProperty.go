package awsec2


// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
//
// Default: No minimum or maximum limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkBandwidthGbpsRequestProperty := &NetworkBandwidthGbpsRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
type CfnSpotFleet_NetworkBandwidthGbpsRequestProperty struct {
	// The maximum amount of network bandwidth, in Gbps.
	//
	// To specify no maximum limit, omit this parameter.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of network bandwidth, in Gbps.
	//
	// To specify no minimum limit, omit this parameter.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

