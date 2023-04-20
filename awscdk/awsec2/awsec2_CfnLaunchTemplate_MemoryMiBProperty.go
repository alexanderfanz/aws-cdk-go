package awsec2


// The minimum and maximum amount of memory, in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryMiBProperty := &MemoryMiBProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
type CfnLaunchTemplate_MemoryMiBProperty struct {
	// The maximum amount of memory, in MiB.
	//
	// To specify no maximum limit, omit this parameter.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory, in MiB.
	//
	// To specify no minimum limit, specify `0` .
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

