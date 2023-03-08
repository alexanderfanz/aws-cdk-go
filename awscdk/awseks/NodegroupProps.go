package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// NodeGroup properties interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   nodegroupProps := &NodegroupProps{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	AmiType: awscdk.Aws_eks.NodegroupAmiType_AL2_X86_64,
//   	CapacityType: awscdk.*Aws_eks.CapacityType_SPOT,
//   	DesiredSize: jsii.Number(123),
//   	DiskSize: jsii.Number(123),
//   	ForceUpdate: jsii.Boolean(false),
//   	InstanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	LaunchTemplateSpec: &LaunchTemplateSpec{
//   		Id: jsii.String("id"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	NodegroupName: jsii.String("nodegroupName"),
//   	NodeRole: role,
//   	ReleaseVersion: jsii.String("releaseVersion"),
//   	RemoteAccess: &NodegroupRemoteAccess{
//   		SshKeyName: jsii.String("sshKeyName"),
//
//   		// the properties below are optional
//   		SourceSecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   	},
//   	Subnets: &SubnetSelection{
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
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Taints: []taintSpec{
//   		&taintSpec{
//   			Effect: awscdk.*Aws_eks.TaintEffect_NO_SCHEDULE,
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type NodegroupProps struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	AmiType NodegroupAmiType `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	CapacityType CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	LaunchTemplateSpec *LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	NodeRole awsiam.IRole `field:"optional" json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	RemoteAccess *NodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints *[]*TaintSpec `field:"optional" json:"taints" yaml:"taints"`
	// Cluster resource.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

