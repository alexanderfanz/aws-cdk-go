package awssupportapp


// Properties for defining a `CfnSlackWorkspaceConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackWorkspaceConfigurationProps := &cfnSlackWorkspaceConfigurationProps{
//   	teamId: jsii.String("teamId"),
//
//   	// the properties below are optional
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnSlackWorkspaceConfigurationProps struct {
	// The team ID in Slack.
	//
	// This ID uniquely identifies a Slack workspace, such as `T012ABCDEFG` .
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// An identifier used to update an existing Slack workspace configuration in AWS CloudFormation , such as `100` .
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

