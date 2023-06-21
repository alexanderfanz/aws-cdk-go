package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Creates a Host volume.
//
// This volume will persist on the host at the specified `hostPath`.
// If the `hostPath` is not specified, Docker will choose the host path. In this case,
// the data may not persist after the containers that use it stop running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   hostVolume := batch_alpha.NewHostVolume(&HostVolumeOptions{
//   	ContainerPath: jsii.String("containerPath"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	HostPath: jsii.String("hostPath"),
//   	Readonly: jsii.Boolean(false),
//   })
//
// Experimental.
type HostVolume interface {
	EcsVolume
	// The path on the container that this volume will be mounted to.
	// Experimental.
	ContainerPath() *string
	// The path on the host machine this container will have access to.
	// Experimental.
	HostPath() *string
	// The name of this volume.
	// Experimental.
	Name() *string
	// Whether or not the container has readonly access to this volume.
	// Experimental.
	Readonly() *bool
}

// The jsii proxy struct for HostVolume
type jsiiProxy_HostVolume struct {
	jsiiProxy_EcsVolume
}

func (j *jsiiProxy_HostVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVolume) HostPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewHostVolume(options *HostVolumeOptions) HostVolume {
	_init_.Initialize()

	if err := validateNewHostVolumeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostVolume{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewHostVolume_Override(h HostVolume, options *HostVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		[]interface{}{options},
		h,
	)
}

// Creates a Volume that uses an AWS Elastic File System (EFS);
//
// this volume can grow and shrink as needed.
// See: https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html
//
// Experimental.
func HostVolume_Efs(options *EfsVolumeOptions) EfsVolume {
	_init_.Initialize()

	if err := validateHostVolume_EfsParameters(options); err != nil {
		panic(err)
	}
	var returns EfsVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		"efs",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Host volume.
//
// This volume will persist on the host at the specified `hostPath`.
// If the `hostPath` is not specified, Docker will choose the host path. In this case,
// the data may not persist after the containers that use it stop running.
// Experimental.
func HostVolume_Host(options *HostVolumeOptions) HostVolume {
	_init_.Initialize()

	if err := validateHostVolume_HostParameters(options); err != nil {
		panic(err)
	}
	var returns HostVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		"host",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// returns `true` if `x` is a `HostVolume`, `false` otherwise.
// Experimental.
func HostVolume_IsHostVolume(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostVolume_IsHostVolumeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		"isHostVolume",
		[]interface{}{x},
		&returns,
	)

	return returns
}

