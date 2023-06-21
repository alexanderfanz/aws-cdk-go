package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for a ManagedEc2EksComputeEnvironment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var launchTemplate launchTemplate
//   var machineImage iMachineImage
//   var placementGroup placementGroup
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   managedEc2EksComputeEnvironmentProps := &ManagedEc2EksComputeEnvironmentProps{
//   	EksCluster: cluster,
//   	KubernetesNamespace: jsii.String("kubernetesNamespace"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	AllocationStrategy: batch_alpha.AllocationStrategy_BEST_FIT,
//   	ComputeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	Enabled: jsii.Boolean(false),
//   	Images: []eksMachineImage{
//   		&eksMachineImage{
//   			Image: machineImage,
//   			ImageType: batch_alpha.EksMachineImageType_EKS_AL2,
//   		},
//   	},
//   	InstanceClasses: []instanceClass{
//   		awscdk.Aws_ec2.*instanceClass_STANDARD3,
//   	},
//   	InstanceRole: role,
//   	InstanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	LaunchTemplate: launchTemplate,
//   	MaxvCpus: jsii.Number(123),
//   	MinvCpus: jsii.Number(123),
//   	PlacementGroup: placementGroup,
//   	ReplaceComputeEnvironment: jsii.Boolean(false),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	ServiceRole: role,
//   	Spot: jsii.Boolean(false),
//   	SpotBidPercentage: jsii.Number(123),
//   	TerminateOnUpdate: jsii.Boolean(false),
//   	UpdateTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	UpdateToLatestImageVersion: jsii.Boolean(false),
//   	UseOptimalInstanceClasses: jsii.Boolean(false),
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.*Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
// Experimental.
type ManagedEc2EksComputeEnvironmentProps struct {
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
	// The cluster that backs this Compute Environment. Required for Compute Environments running Kubernetes jobs.
	//
	// Please ensure that you have followed the steps at
	//
	// https://docs.aws.amazon.com/batch/latest/userguide/getting-started-eks.html
	//
	// before attempting to deploy a `ManagedEc2EksComputeEnvironment` that uses this cluster.
	// If you do not follow the steps in the link, the deployment fail with a message that the
	// compute environment did not stabilize.
	// Experimental.
	EksCluster awseks.ICluster `field:"required" json:"eksCluster" yaml:"eksCluster"`
	// The namespace of the Cluster.
	// Experimental.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
	// The allocation strategy to use if not enough instances of the best fitting instance type can be allocated.
	// Experimental.
	AllocationStrategy AllocationStrategy `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Configure which AMIs this Compute Environment can launch.
	// Experimental.
	Images *[]*EksMachineImage `field:"optional" json:"images" yaml:"images"`
	// The instance types that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	// Batch will automatically choose the instance size.
	// Experimental.
	InstanceClasses *[]awsec2.InstanceClass `field:"optional" json:"instanceClasses" yaml:"instanceClasses"`
	// The execution Role that instances launched by this Compute Environment will use.
	// Experimental.
	InstanceRole awsiam.IRole `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The instance types that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Launch Template that this Compute Environment will use to provision EC2 Instances.
	//
	// *Note*: if `securityGroups` is specified on both your
	// launch template and this Compute Environment, **the
	// `securityGroup`s on the Compute Environment override the
	// ones on the launch template.**
	// Experimental.
	LaunchTemplate awsec2.ILaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The minimum vCPUs that an environment should maintain, even if the compute environment is DISABLED.
	// Experimental.
	MinvCpus *float64 `field:"optional" json:"minvCpus" yaml:"minvCpus"`
	// The EC2 placement group to associate with your compute resources.
	//
	// If you intend to submit multi-node parallel jobs to this Compute Environment,
	// you should consider creating a cluster placement group and associate it with your compute resources.
	// This keeps your multi-node parallel job on a logical grouping of instances
	// within a single Availability Zone with high network flow potential.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
	//
	// Experimental.
	PlacementGroup awsec2.IPlacementGroup `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for that instance type before instances are launched.
	//
	// For example, if your maximum percentage is 20%, the Spot price must be
	// less than 20% of the current On-Demand price for that Instance.
	// You always pay the lowest market price and never more than your maximum percentage.
	// For most use cases, Batch recommends leaving this field empty.
	//
	// Implies `spot == true` if set.
	// Experimental.
	SpotBidPercentage *float64 `field:"optional" json:"spotBidPercentage" yaml:"spotBidPercentage"`
	// Whether or not to use batch's optimal instance type.
	//
	// The optimal instance type is equivalent to adding the
	// C4, M4, and R4 instance classes. You can specify other instance classes
	// (of the same architecture) in addition to the optimal instance classes.
	// Experimental.
	UseOptimalInstanceClasses *bool `field:"optional" json:"useOptimalInstanceClasses" yaml:"useOptimalInstanceClasses"`
}

