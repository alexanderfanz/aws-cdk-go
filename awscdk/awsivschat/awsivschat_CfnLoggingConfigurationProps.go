package awsivschat

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingConfigurationProps := &CfnLoggingConfigurationProps{
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		CloudWatchLogs: &CloudWatchLogsDestinationConfigurationProperty{
//   			LogGroupName: jsii.String("logGroupName"),
//   		},
//   		Firehose: &FirehoseDestinationConfigurationProperty{
//   			DeliveryStreamName: jsii.String("deliveryStreamName"),
//   		},
//   		S3: &S3DestinationConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLoggingConfigurationProps struct {
	// The DestinationConfiguration is a complex type that contains information about where chat content will be logged.
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Logging-configuration name.
	//
	// The value does not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
