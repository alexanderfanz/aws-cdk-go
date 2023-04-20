package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for the FSx file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//   var securityGroup securityGroup
//   var vpc vpc
//
//   fileSystemProps := &FileSystemProps{
//   	StorageCapacityGiB: jsii.Number(123),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	BackupId: jsii.String("backupId"),
//   	KmsKey: key,
//   	RemovalPolicy: monocdk.RemovalPolicy_DESTROY,
//   	SecurityGroup: securityGroup,
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
// Experimental.
type FileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	// Experimental.
	StorageCapacityGiB *float64 `field:"required" json:"storageCapacityGiB" yaml:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Experimental.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

