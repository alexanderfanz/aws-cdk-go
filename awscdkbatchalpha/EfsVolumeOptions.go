package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
)

// Options for configuring an EfsVolume.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import efs "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFileSystem iFileSystem
//
//
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   		Volumes: []ecsVolume{
//   			batch.*ecsVolume_Efs(&EfsVolumeOptions{
//   				Name: jsii.String("myVolume"),
//   				FileSystem: myFileSystem,
//   				ContainerPath: jsii.String("/Volumes/myVolume"),
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type EfsVolumeOptions struct {
	// the path on the container where this volume is mounted.
	// Experimental.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// the name of this volume.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// if set, the container will have readonly access to the volume.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
	// The EFS File System that supports this volume.
	// Experimental.
	FileSystem awsefs.IFileSystem `field:"required" json:"fileSystem" yaml:"fileSystem"`
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, `rootDirectory` must either be omitted or set to `/`
	// which enforces the path set on the EFS access point.
	// If an access point is used, `enableTransitEncryption` must be `true`.
	// See: https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html
	//
	// Experimental.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Enables encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	// See: https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html
	//
	// Experimental.
	EnableTransitEncryption *bool `field:"optional" json:"enableTransitEncryption" yaml:"enableTransitEncryption"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume is used instead.
	// Specifying `/` has the same effect as omitting this parameter.
	// The maximum length is 4,096 characters.
	// Experimental.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// The value must be between 0 and 65,535.
	// See: https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html
	//
	// Experimental.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
	// Whether or not to use the AWS Batch job IAM role defined in a job definition when mounting the Amazon EFS file system.
	//
	// If specified, `enableTransitEncryption` must be `true`.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html#efs-volume-accesspoints
	//
	// Experimental.
	UseJobRole *bool `field:"optional" json:"useJobRole" yaml:"useJobRole"`
}

