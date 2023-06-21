package awscdk


// Properties for defining a `CfnTypeActivation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTypeActivationProps := &CfnTypeActivationProps{
//   	AutoUpdate: jsii.Boolean(false),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	LoggingConfig: &LoggingConfigProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogRoleArn: jsii.String("logRoleArn"),
//   	},
//   	MajorVersion: jsii.String("majorVersion"),
//   	PublicTypeArn: jsii.String("publicTypeArn"),
//   	PublisherId: jsii.String("publisherId"),
//   	Type: jsii.String("type"),
//   	TypeName: jsii.String("typeName"),
//   	TypeNameAlias: jsii.String("typeNameAlias"),
//   	VersionBump: jsii.String("versionBump"),
//   }
//
type CfnTypeActivationProps struct {
	// Whether to automatically update the extension in this account and Region when a new *minor* version is published by the extension publisher.
	//
	// Major versions released by the publisher must be manually updated.
	//
	// The default is `true` .
	AutoUpdate interface{} `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// The name of the IAM execution role to use to activate the extension.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies logging configuration information for an extension.
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The major version of this extension you want to activate, if multiple major versions are available.
	//
	// The default is the latest major version. CloudFormation uses the latest available *minor* version of the major version selected.
	//
	// You can specify `MajorVersion` or `VersionBump` , but not both.
	MajorVersion *string `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The Amazon Resource Number (ARN) of the public extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublicTypeArn *string `field:"optional" json:"publicTypeArn" yaml:"publicTypeArn"`
	// The ID of the extension publisher.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	PublisherId *string `field:"optional" json:"publisherId" yaml:"publisherId"`
	// The extension type.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The name of the extension.
	//
	// Conditional: You must specify `PublicTypeArn` , or `TypeName` , `Type` , and `PublisherId` .
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// An alias to assign to the public extension, in this account and Region.
	//
	// If you specify an alias for the extension, CloudFormation treats the alias as the extension type name within this account and Region. You must use the alias to refer to the extension in your templates, API calls, and CloudFormation console.
	//
	// An extension alias must be unique within a given account and Region. You can activate the same public resource multiple times in the same account and Region, using different type name aliases.
	TypeNameAlias *string `field:"optional" json:"typeNameAlias" yaml:"typeNameAlias"`
	// Manually updates a previously-activated type to a new major or minor version, if available.
	//
	// You can also use this parameter to update the value of `AutoUpdate` .
	//
	// - `MAJOR` : CloudFormation updates the extension to the newest major version, if one is available.
	// - `MINOR` : CloudFormation updates the extension to the newest minor version, if one is available.
	VersionBump *string `field:"optional" json:"versionBump" yaml:"versionBump"`
}

