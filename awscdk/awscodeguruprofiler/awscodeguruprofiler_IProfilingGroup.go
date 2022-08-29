package awscodeguruprofiler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// IResource represents a Profiling Group.
type IProfilingGroup interface {
	awscdk.IResource
	// Grant access to publish profiling information to the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//   - codeguru-profiler:ConfigureAgent
	// - codeguru-profiler:PostAgentProfile.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to read profiling information from the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//   - codeguru-profiler:GetProfile
	// - codeguru-profiler:DescribeProfilingGroup.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// A name for the profiling group.
	ProfilingGroupName() *string
}

// The jsii proxy for IProfilingGroup
type jsiiProxy_IProfilingGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProfilingGroup) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProfilingGroup) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IProfilingGroup) ProfilingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profilingGroupName",
		&returns,
	)
	return returns
}
