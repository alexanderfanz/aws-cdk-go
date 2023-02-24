package awsssm


// Attributes for parameters of string list type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listParameterAttributes := &ListParameterAttributes{
//   	ParameterName: jsii.String("parameterName"),
//
//   	// the properties below are optional
//   	ElementType: awscdk.Aws_ssm.ParameterValueType_STRING,
//   	SimpleName: jsii.Boolean(false),
//   	Version: jsii.Number(123),
//   }
//
// See: ParameterType.
//
type ListParameterAttributes struct {
	// The name of the parameter store value.
	//
	// This value can be a token or a concrete string. If it is a concrete string
	// and includes "/" it must also be prefixed with a "/" (fully-qualified).
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The type of the string list parameter value.
	//
	// Using specific types can be helpful in catching invalid values
	// at the start of creating or updating a stack. CloudFormation validates
	// the values against existing values in the account.
	//
	// Note - if you want to allow values from different AWS accounts, use
	// ParameterValueType.STRING
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-ssm-parameter-types
	//
	ElementType ParameterValueType `field:"optional" json:"elementType" yaml:"elementType"`
	// The version number of the value you wish to retrieve.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

