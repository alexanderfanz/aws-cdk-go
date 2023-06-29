package awsefs


// Properties for defining a `CfnFileSystem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystemPolicy interface{}
//
//   cfnFileSystemProps := &CfnFileSystemProps{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   	BackupPolicy: &BackupPolicyProperty{
//   		Status: jsii.String("status"),
//   	},
//   	BypassPolicyLockoutSafetyCheck: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	FileSystemPolicy: fileSystemPolicy,
//   	FileSystemTags: []elasticFileSystemTagProperty{
//   		&elasticFileSystemTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LifecyclePolicies: []interface{}{
//   		&LifecyclePolicyProperty{
//   			TransitionToIa: jsii.String("transitionToIa"),
//   			TransitionToPrimaryStorageClass: jsii.String("transitionToPrimaryStorageClass"),
//   		},
//   	},
//   	PerformanceMode: jsii.String("performanceMode"),
//   	ProvisionedThroughputInMibps: jsii.Number(123),
//   	ThroughputMode: jsii.String("throughputMode"),
//   }
//
type CfnFileSystemProps struct {
	// Used to create a file system that uses One Zone storage classes.
	//
	// It specifies the AWS Availability Zone in which to create the file system. Use the format `us-east-1a` to specify the Availability Zone. For more information about One Zone storage classes, see [Using EFS storage classes](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide* .
	//
	// > One Zone storage classes are not available in all Availability Zones in AWS Regions where Amazon EFS is available.
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// Use the `BackupPolicy` to turn automatic backups on or off for the file system.
	BackupPolicy interface{} `field:"optional" json:"backupPolicy" yaml:"backupPolicy"`
	// (Optional) A boolean that specifies whether or not to bypass the `FileSystemPolicy` lockout safety check.
	//
	// The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future `PutFileSystemPolicy` requests on this file system. Set `BypassPolicyLockoutSafetyCheck` to `True` only when you intend to prevent the IAM principal that is making the request from making subsequent `PutFileSystemPolicy` requests on this file system. The default value is `False` .
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// A Boolean value that, if true, creates an encrypted file system.
	//
	// When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing AWS KMS key . If you don't specify a KMS key , then the default KMS key for Amazon EFS , `/aws/elasticfilesystem` , is used to protect the encrypted file system.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The `FileSystemPolicy` for the EFS file system.
	//
	// A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using IAM to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide* .
	FileSystemPolicy interface{} `field:"optional" json:"fileSystemPolicy" yaml:"fileSystemPolicy"`
	// Use to create one or more tags associated with the file system.
	//
	// Each tag is a user-defined key-value pair. Name your file system on creation by including a `"Key":"Name","Value":"{value}"` key-value pair. Each key must be unique. For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference Guide* .
	FileSystemTags *[]*CfnFileSystem_ElasticFileSystemTagProperty `field:"optional" json:"fileSystemTags" yaml:"fileSystemTags"`
	// The ID of the AWS KMS key to be used to protect the encrypted file system.
	//
	// This parameter is only required if you want to use a nondefault KMS key . If this parameter is not specified, the default KMS key for Amazon EFS is used. This ID can be in one of the following formats:
	//
	// - Key ID - A unique identifier of the key, for example `1234abcd-12ab-34cd-56ef-1234567890ab` .
	// - ARN - An Amazon Resource Name (ARN) for the key, for example `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` .
	// - Key alias - A previously created display name for a key, for example `alias/projectKey1` .
	// - Key alias ARN - An ARN for a key alias, for example `arn:aws:kms:us-west-2:444455556666:alias/projectKey1` .
	//
	// If `KmsKeyId` is specified, the `Encrypted` parameter must be set to true.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of `LifecyclePolicy` objects that define the file system's `LifecycleConfiguration` object.
	//
	// A `LifecycleConfiguration` object informs EFS lifecycle management and intelligent tiering of the following:
	//
	// - When to move files in the file system from primary storage to the IA storage class.
	// - When to move files that are in IA storage to primary storage.
	//
	// > Amazon EFS requires that each `LifecyclePolicy` object have only a single transition. This means that in a request body, `LifecyclePolicies` needs to be structured as an array of `LifecyclePolicy` objects, one object for each transition, `TransitionToIA` , `TransitionToPrimaryStorageClass` . See the example requests in the following section for more information.
	LifecyclePolicies interface{} `field:"optional" json:"lifecyclePolicies" yaml:"lifecyclePolicies"`
	// The performance mode of the file system.
	//
	// We recommend `generalPurpose` performance mode for most file systems. File systems using the `maxIO` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created.
	//
	// > The `maxIO` mode is not supported on file systems using One Zone storage classes.
	//
	// Default is `generalPurpose` .
	PerformanceMode *string `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating.
	//
	// Required if `ThroughputMode` is set to `provisioned` . Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact AWS Support . For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide* .
	ProvisionedThroughputInMibps *float64 `field:"optional" json:"provisionedThroughputInMibps" yaml:"provisionedThroughputInMibps"`
	// Specifies the throughput mode for the file system.
	//
	// The mode can be `bursting` , `provisioned` , or `elastic` . If you set `ThroughputMode` to `provisioned` , you must also set a value for `ProvisionedThroughputInMibps` . After you create the file system, you can decrease your file system's throughput in Provisioned Throughput mode or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide* .
	//
	// Default is `elastic` .
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
}

