package awssso


// Properties for defining a `CfnInstanceAccessControlAttributeConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceAccessControlAttributeConfigurationProps := &CfnInstanceAccessControlAttributeConfigurationProps{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	AccessControlAttributes: []interface{}{
//   		&AccessControlAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: &AccessControlAttributeValueProperty{
//   				Source: []*string{
//   					jsii.String("source"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnInstanceAccessControlAttributeConfigurationProps struct {
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Lists the attributes that are configured for ABAC in the specified IAM Identity Center instance.
	AccessControlAttributes interface{} `field:"optional" json:"accessControlAttributes" yaml:"accessControlAttributes"`
}
