package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
)

// A regional grouping of one or more container instances on which you can run tasks and services.
// Experimental.
type ICluster interface {
	awscdk.IResource
	// The autoscaling group added to the cluster if capacity is associated to the cluster.
	// Experimental.
	AutoscalingGroup() awsautoscaling.IAutoScalingGroup
	// The Amazon Resource Name (ARN) that identifies the cluster.
	// Experimental.
	ClusterArn() *string
	// The name of the cluster.
	// Experimental.
	ClusterName() *string
	// Manage the allowed network connections for the cluster with Security Groups.
	// Experimental.
	Connections() awsec2.Connections
	// The AWS Cloud Map namespace to associate with the cluster.
	// Experimental.
	DefaultCloudMapNamespace() awsservicediscovery.INamespace
	// The execute command configuration for the cluster.
	// Experimental.
	ExecuteCommandConfiguration() *ExecuteCommandConfiguration
	// Specifies whether the cluster has EC2 instance capacity.
	// Experimental.
	HasEc2Capacity() *bool
	// The VPC associated with the cluster.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICluster) AutoscalingGroup() awsautoscaling.IAutoScalingGroup {
	var returns awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoscalingGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) DefaultCloudMapNamespace() awsservicediscovery.INamespace {
	var returns awsservicediscovery.INamespace
	_jsii_.Get(
		j,
		"defaultCloudMapNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ExecuteCommandConfiguration() *ExecuteCommandConfiguration {
	var returns *ExecuteCommandConfiguration
	_jsii_.Get(
		j,
		"executeCommandConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) HasEc2Capacity() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasEc2Capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

