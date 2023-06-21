package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for a FargateComputeEnvironment.
//
// Example:
//   var vpc iVpc
//
//   sharedComputeEnv := batch.NewFargateComputeEnvironment(this, jsii.String("spotEnv"), &FargateComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	Spot: jsii.Boolean(true),
//   })
//   lowPriorityQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	Priority: jsii.Number(1),
//   })
//   highPriorityQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	Priority: jsii.Number(10),
//   })
//   lowPriorityQueue.AddComputeEnvironment(sharedComputeEnv, jsii.Number(1))
//   highPriorityQueue.AddComputeEnvironment(sharedComputeEnv, jsii.Number(1))
//
// Experimental.
type FargateComputeEnvironmentProps struct {
	// The name of the ComputeEnvironment.
	// Experimental.
	ComputeEnvironmentName *string `field:"optional" json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// Whether or not this ComputeEnvironment can accept jobs from a Queue.
	//
	// Enabled ComputeEnvironments can accept jobs from a Queue and
	// can scale instances up or down.
	// Disabled ComputeEnvironments cannot accept jobs from a Queue or
	// scale instances up or down.
	//
	// If you change a ComputeEnvironment from enabled to disabled while it is executing jobs,
	// Jobs in the `STARTED` or `RUNNING` states will not
	// be interrupted. As jobs complete, the ComputeEnvironment will scale instances down to `minvCpus`.
	//
	// To ensure you aren't billed for unused capacity, set `minvCpus` to `0`.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The role Batch uses to perform actions on your behalf in your account, such as provision instances to run your jobs.
	// Experimental.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// VPC in which this Compute Environment will launch Instances.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The maximum vCpus this `ManagedComputeEnvironment` can scale up to. Each vCPU is equivalent to 1024 CPU shares.
	//
	// *Note*: if this Compute Environment uses EC2 resources (not Fargate) with either `AllocationStrategy.BEST_FIT_PROGRESSIVE` or
	// `AllocationStrategy.SPOT_CAPACITY_OPTIMIZED`, or `AllocationStrategy.BEST_FIT` with Spot instances,
	// The scheduler may exceed this number by at most one of the instances specified in `instanceTypes`
	// or `instanceClasses`.
	// Experimental.
	MaxvCpus *float64 `field:"optional" json:"maxvCpus" yaml:"maxvCpus"`
	// Specifies whether this Compute Environment is replaced if an update is made that requires replacing its instances.
	//
	// To enable more properties to be updated,
	// set this property to `false`. When changing the value of this property to false,
	// do not change any other properties at the same time.
	// If other properties are changed at the same time,
	// and the change needs to be rolled back but it can't,
	// it's possible for the stack to go into the UPDATE_ROLLBACK_FAILED state.
	// You can't update a stack that is in the UPDATE_ROLLBACK_FAILED state.
	// However, if you can continue to roll it back,
	// you can return the stack to its original settings and then try to update it again.
	//
	// The properties which require a replacement of the Compute Environment are:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html
	//
	// Experimental.
	ReplaceComputeEnvironment *bool `field:"optional" json:"replaceComputeEnvironment" yaml:"replaceComputeEnvironment"`
	// The security groups this Compute Environment will launch instances in.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Whether or not to use spot instances.
	//
	// Spot instances are less expensive EC2 instances that can be
	// reclaimed by EC2 at any time; your job will be given two minutes
	// of notice before reclamation.
	// Experimental.
	Spot *bool `field:"optional" json:"spot" yaml:"spot"`
	// Whether or not any running jobs will be immediately terminated when an infrastructure update occurs.
	//
	// If this is enabled, any terminated jobs may be retried, depending on the job's
	// retry policy.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	//
	// Experimental.
	TerminateOnUpdate *bool `field:"optional" json:"terminateOnUpdate" yaml:"terminateOnUpdate"`
	// Only meaningful if `terminateOnUpdate` is `false`.
	//
	// If so,
	// when an infrastructure update is triggered, any running jobs
	// will be allowed to run until `updateTimeout` has expired.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	//
	// Experimental.
	UpdateTimeout awscdk.Duration `field:"optional" json:"updateTimeout" yaml:"updateTimeout"`
	// Whether or not the AMI is updated to the latest one supported by Batch when an infrastructure update occurs.
	//
	// If you specify a specific AMI, this property will be ignored.
	// Experimental.
	UpdateToLatestImageVersion *bool `field:"optional" json:"updateToLatestImageVersion" yaml:"updateToLatestImageVersion"`
	// The VPC Subnets this Compute Environment will launch instances in.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

