package awsappconfig


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Monitors: []interface{}{
//   		&MonitorsProperty{
//   			AlarmArn: jsii.String("alarmArn"),
//   			AlarmRoleArn: jsii.String("alarmRoleArn"),
//   		},
//   	},
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// The application ID.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A name for the environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Amazon CloudWatch alarms to monitor during the deployment process.
	Monitors interface{} `field:"optional" json:"monitors" yaml:"monitors"`
	// Metadata to assign to the environment.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*CfnEnvironment_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

