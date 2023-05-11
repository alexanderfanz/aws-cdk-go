package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPermissionSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inlinePolicy interface{}
//
//   cfnPermissionSetProps := &CfnPermissionSetProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CustomerManagedPolicyReferences: []interface{}{
//   		&CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	InlinePolicy: inlinePolicy,
//   	ManagedPolicies: []*string{
//   		jsii.String("managedPolicies"),
//   	},
//   	PermissionsBoundary: &PermissionsBoundaryProperty{
//   		CustomerManagedPolicyReference: &CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Path: jsii.String("path"),
//   		},
//   		ManagedPolicyArn: jsii.String("managedPolicyArn"),
//   	},
//   	RelayStateType: jsii.String("relayStateType"),
//   	SessionDuration: jsii.String("sessionDuration"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPermissionSetProps struct {
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the permission set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
	CustomerManagedPolicyReferences interface{} `field:"optional" json:"customerManagedPolicyReferences" yaml:"customerManagedPolicyReferences"`
	// The description of the `PermissionSet` .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The inline policy that is attached to the permission set.
	//
	// > For `Length Constraints` , if a valid ARN is provided for a permission set, it is possible for an empty inline policy to be returned.
	InlinePolicy interface{} `field:"optional" json:"inlinePolicy" yaml:"inlinePolicy"`
	// A structure that stores the details of the AWS managed policy.
	ManagedPolicies *[]*string `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary.
	//
	// Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
	//
	// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
	PermissionsBoundary interface{} `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Used to redirect users within the application during the federation authentication process.
	RelayStateType *string `field:"optional" json:"relayStateType" yaml:"relayStateType"`
	// The length of time that the application user sessions are valid for in the ISO-8601 standard.
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// The tags to attach to the new `PermissionSet` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
