package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkteam`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkteamProps := &CfnWorkteamProps{
//   	Description: jsii.String("description"),
//   	MemberDefinitions: []interface{}{
//   		&MemberDefinitionProperty{
//   			CognitoMemberDefinition: &CognitoMemberDefinitionProperty{
//   				CognitoClientId: jsii.String("cognitoClientId"),
//   				CognitoUserGroup: jsii.String("cognitoUserGroup"),
//   				CognitoUserPool: jsii.String("cognitoUserPool"),
//   			},
//   			OidcMemberDefinition: &OidcMemberDefinitionProperty{
//   				OidcGroups: []*string{
//   					jsii.String("oidcGroups"),
//   				},
//   			},
//   		},
//   	},
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkforceName: jsii.String("workforceName"),
//   	WorkteamName: jsii.String("workteamName"),
//   }
//
type CfnWorkteamProps struct {
	// A description of the work team.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.
	//
	// Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `CognitoMemberDefinition` . For workforces created using your own OIDC identity provider (IdP) use `OidcMemberDefinition` .
	MemberDefinitions interface{} `field:"optional" json:"memberDefinitions" yaml:"memberDefinitions"`
	// Configures SNS notifications of available or expiring work items for work teams.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// An array of key-value pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::SageMaker::Workteam.WorkforceName`.
	WorkforceName *string `field:"optional" json:"workforceName" yaml:"workforceName"`
	// The name of the work team.
	WorkteamName *string `field:"optional" json:"workteamName" yaml:"workteamName"`
}

