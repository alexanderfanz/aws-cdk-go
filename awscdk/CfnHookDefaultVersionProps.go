package awscdk


// Properties for defining a `CfnHookDefaultVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHookDefaultVersionProps := &CfnHookDefaultVersionProps{
//   	TypeName: jsii.String("typeName"),
//   	TypeVersionArn: jsii.String("typeVersionArn"),
//   	VersionId: jsii.String("versionId"),
//   }
//
type CfnHookDefaultVersionProps struct {
	// The name of the hook.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The version ID of the type configuration.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn *string `field:"optional" json:"typeVersionArn" yaml:"typeVersionArn"`
	// The version ID of the type specified.
	//
	// You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

