package awsimagebuilder


// Properties for defining a `CfnInfrastructureConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfrastructureConfigurationProps := &CfnInfrastructureConfigurationProps{
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InstanceMetadataOptions: &InstanceMetadataOptionsProperty{
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   	},
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	KeyPair: jsii.String("keyPair"),
//   	Logging: &LoggingProperty{
//   		S3Logs: &S3LogsProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	ResourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TerminateInstanceOnFailure: jsii.Boolean(false),
//   }
//
type CfnInfrastructureConfigurationProps struct {
	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string `field:"required" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The name of the infrastructure configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the infrastructure configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The instance metadata option settings for the infrastructure configuration.
	InstanceMetadataOptions interface{} `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// The instance types of the infrastructure configuration.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Amazon EC2 key pair of the infrastructure configuration.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The logging configuration defines where Image Builder uploads your logs.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The tags attached to the resource created by Image Builder.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) of the SNS topic for the infrastructure configuration.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The subnet ID of the infrastructure configuration.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags of the infrastructure configuration.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The terminate instance on failure configuration of the infrastructure configuration.
	TerminateInstanceOnFailure interface{} `field:"optional" json:"terminateInstanceOnFailure" yaml:"terminateInstanceOnFailure"`
}

