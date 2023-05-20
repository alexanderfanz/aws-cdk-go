package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removalPolicyOptions := &RemovalPolicyOptions{
//   	ApplyToUpdateReplacePolicy: jsii.Boolean(false),
//   	Default: cdk.RemovalPolicy_DESTROY,
//   }
//
type RemovalPolicyOptions struct {
	// Apply the same deletion policy to the resource's "UpdateReplacePolicy".
	ApplyToUpdateReplacePolicy *bool `field:"optional" json:"applyToUpdateReplacePolicy" yaml:"applyToUpdateReplacePolicy"`
	// The default policy to apply in case the removal policy is not defined.
	Default RemovalPolicy `field:"optional" json:"default" yaml:"default"`
}

