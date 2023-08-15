package awswafv2


// Not currently supported by AWS CloudFormation .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestInspectionACFPProperty := &RequestInspectionACFPProperty{
//   	PayloadType: jsii.String("payloadType"),
//
//   	// the properties below are optional
//   	AddressFields: []interface{}{
//   		&FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
//   	EmailField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   	PasswordField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   	PhoneNumberFields: []interface{}{
//   		&FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
//   	UsernameField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html
//
type CfnWebACL_RequestInspectionACFPProperty struct {
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-payloadtype
	//
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-addressfields
	//
	AddressFields interface{} `field:"optional" json:"addressFields" yaml:"addressFields"`
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-emailfield
	//
	EmailField interface{} `field:"optional" json:"emailField" yaml:"emailField"`
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-passwordfield
	//
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-phonenumberfields
	//
	PhoneNumberFields interface{} `field:"optional" json:"phoneNumberFields" yaml:"phoneNumberFields"`
	// Not currently supported by AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-usernamefield
	//
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}
