package awsec2


// Base block device options for an EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsDeviceOptionsBase := &EbsDeviceOptionsBase{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	VolumeType: awscdk.Aws_ec2.EbsDeviceVolumeType_STANDARD,
//   }
//
type EbsDeviceOptionsBase struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for `volumeType`: `EbsDeviceVolumeType.IO1`
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}

