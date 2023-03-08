package awsekslegacy


// Options for adding an AutoScalingGroup as capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupOptions := &autoScalingGroupOptions{
//   	bootstrapEnabled: jsii.Boolean(false),
//   	bootstrapOptions: &bootstrapOptions{
//   		additionalArgs: jsii.String("additionalArgs"),
//   		awsApiRetryAttempts: jsii.Number(123),
//   		dockerConfigJson: jsii.String("dockerConfigJson"),
//   		enableDockerBridge: jsii.Boolean(false),
//   		kubeletExtraArgs: jsii.String("kubeletExtraArgs"),
//   		useMaxPods: jsii.Boolean(false),
//   	},
//   	mapRole: jsii.Boolean(false),
//   }
//
// Experimental.
type AutoScalingGroupOptions struct {
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Experimental.
	BootstrapEnabled *bool `field:"optional" json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// Allows options for node bootstrapping through EC2 user data.
	// Experimental.
	BootstrapOptions *BootstrapOptions `field:"optional" json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Experimental.
	MapRole *bool `field:"optional" json:"mapRole" yaml:"mapRole"`
}
