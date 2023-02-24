package awsiot


// Defines the type of action and the parameters for that action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionParamsProperty := &ActionParamsProperty{
//   	AddThingsToThingGroupParams: &AddThingsToThingGroupParamsProperty{
//   		ThingGroupNames: []*string{
//   			jsii.String("thingGroupNames"),
//   		},
//
//   		// the properties below are optional
//   		OverrideDynamicGroups: jsii.Boolean(false),
//   	},
//   	EnableIoTLoggingParams: &EnableIoTLoggingParamsProperty{
//   		LogLevel: jsii.String("logLevel"),
//   		RoleArnForLogging: jsii.String("roleArnForLogging"),
//   	},
//   	PublishFindingToSnsParams: &PublishFindingToSnsParamsProperty{
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	ReplaceDefaultPolicyVersionParams: &ReplaceDefaultPolicyVersionParamsProperty{
//   		TemplateName: jsii.String("templateName"),
//   	},
//   	UpdateCaCertificateParams: &UpdateCACertificateParamsProperty{
//   		Action: jsii.String("action"),
//   	},
//   	UpdateDeviceCertificateParams: &UpdateDeviceCertificateParamsProperty{
//   		Action: jsii.String("action"),
//   	},
//   }
//
type CfnMitigationAction_ActionParamsProperty struct {
	// Specifies the group to which you want to add the devices.
	AddThingsToThingGroupParams interface{} `field:"optional" json:"addThingsToThingGroupParams" yaml:"addThingsToThingGroupParams"`
	// Specifies the logging level and the role with permissions for logging.
	//
	// You cannot specify a logging level of `DISABLED` .
	EnableIoTLoggingParams interface{} `field:"optional" json:"enableIoTLoggingParams" yaml:"enableIoTLoggingParams"`
	// Specifies the topic to which the finding should be published.
	PublishFindingToSnsParams interface{} `field:"optional" json:"publishFindingToSnsParams" yaml:"publishFindingToSnsParams"`
	// Replaces the policy version with a default or blank policy.
	//
	// You specify the template name. Only a value of `BLANK_POLICY` is currently supported.
	ReplaceDefaultPolicyVersionParams interface{} `field:"optional" json:"replaceDefaultPolicyVersionParams" yaml:"replaceDefaultPolicyVersionParams"`
	// Specifies the new state for the CA certificate.
	//
	// Only a value of `DEACTIVATE` is currently supported.
	UpdateCaCertificateParams interface{} `field:"optional" json:"updateCaCertificateParams" yaml:"updateCaCertificateParams"`
	// Specifies the new state for a device certificate.
	//
	// Only a value of `DEACTIVATE` is currently supported.
	UpdateDeviceCertificateParams interface{} `field:"optional" json:"updateDeviceCertificateParams" yaml:"updateDeviceCertificateParams"`
}

