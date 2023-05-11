package awsgreengrassv2


// Contains information about a container in which AWS Lambda functions run on AWS IoT Greengrass core devices.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaContainerParamsProperty := &LambdaContainerParamsProperty{
//   	Devices: []interface{}{
//   		&LambdaDeviceMountProperty{
//   			AddGroupOwner: jsii.Boolean(false),
//   			Path: jsii.String("path"),
//   			Permission: jsii.String("permission"),
//   		},
//   	},
//   	MemorySizeInKb: jsii.Number(123),
//   	MountRoSysfs: jsii.Boolean(false),
//   	Volumes: []interface{}{
//   		&LambdaVolumeMountProperty{
//   			AddGroupOwner: jsii.Boolean(false),
//   			DestinationPath: jsii.String("destinationPath"),
//   			Permission: jsii.String("permission"),
//   			SourcePath: jsii.String("sourcePath"),
//   		},
//   	},
//   }
//
type CfnComponentVersion_LambdaContainerParamsProperty struct {
	// The list of system devices that the container can access.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// The memory size of the container, expressed in kilobytes.
	//
	// Default: `16384` (16 MB).
	MemorySizeInKb *float64 `field:"optional" json:"memorySizeInKb" yaml:"memorySizeInKb"`
	// Whether or not the container can read information from the device's `/sys` folder.
	//
	// Default: `false`.
	MountRoSysfs interface{} `field:"optional" json:"mountRoSysfs" yaml:"mountRoSysfs"`
	// The list of volumes that the container can access.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}
