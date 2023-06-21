package awsglobalaccelerator


// Construct properties of the Accelerator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorProps := &AcceleratorProps{
//   	AcceleratorName: jsii.String("acceleratorName"),
//   	Enabled: jsii.Boolean(false),
//   }
//
type AcceleratorProps struct {
	// The name of the accelerator.
	AcceleratorName *string `field:"optional" json:"acceleratorName" yaml:"acceleratorName"`
	// Indicates whether the accelerator is enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

