package awsimagebuilder


// Defines block device mappings for the instance used to configure your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceBlockDeviceMappingProperty := &InstanceBlockDeviceMappingProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	Ebs: &EbsInstanceBlockDeviceSpecificationProperty{
//   		DeleteOnTermination: jsii.Boolean(false),
//   		Encrypted: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SnapshotId: jsii.String("snapshotId"),
//   		Throughput: jsii.Number(123),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	NoDevice: jsii.String("noDevice"),
//   	VirtualName: jsii.String("virtualName"),
//   }
//
type CfnContainerRecipe_InstanceBlockDeviceMappingProperty struct {
	// The device to which these mappings apply.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Use to remove a mapping from the base image.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Use to manage instance ephemeral devices.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}
