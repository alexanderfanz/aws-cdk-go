package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A Kubernetes EmptyDir volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var size size
//
//   emptyDirVolume := batch_alpha.NewEmptyDirVolume(&EmptyDirVolumeOptions{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Medium: batch_alpha.EmptyDirMediumType_DISK,
//   	MountPath: jsii.String("mountPath"),
//   	Readonly: jsii.Boolean(false),
//   	SizeLimit: size,
//   })
//
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
// Experimental.
type EmptyDirVolume interface {
	EksVolume
	// The path on the container where the container is mounted.
	// Experimental.
	ContainerPath() *string
	// The storage type to use for this Volume.
	// Experimental.
	Medium() EmptyDirMediumType
	// The name of this volume.
	//
	// The name must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	// Experimental.
	Name() *string
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Experimental.
	Readonly() *bool
	// The maximum size for this Volume.
	// Experimental.
	SizeLimit() awscdk.Size
}

// The jsii proxy struct for EmptyDirVolume
type jsiiProxy_EmptyDirVolume struct {
	jsiiProxy_EksVolume
}

func (j *jsiiProxy_EmptyDirVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmptyDirVolume) Medium() EmptyDirMediumType {
	var returns EmptyDirMediumType
	_jsii_.Get(
		j,
		"medium",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmptyDirVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmptyDirVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmptyDirVolume) SizeLimit() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"sizeLimit",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmptyDirVolume(options *EmptyDirVolumeOptions) EmptyDirVolume {
	_init_.Initialize()

	if err := validateNewEmptyDirVolumeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmptyDirVolume{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewEmptyDirVolume_Override(e EmptyDirVolume, options *EmptyDirVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		[]interface{}{options},
		e,
	)
}

// Creates a Kubernetes EmptyDir volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
// Experimental.
func EmptyDirVolume_EmptyDir(options *EmptyDirVolumeOptions) EmptyDirVolume {
	_init_.Initialize()

	if err := validateEmptyDirVolume_EmptyDirParameters(options); err != nil {
		panic(err)
	}
	var returns EmptyDirVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		"emptyDir",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Kubernetes HostPath volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
// Experimental.
func EmptyDirVolume_HostPath(options *HostPathVolumeOptions) HostPathVolume {
	_init_.Initialize()

	if err := validateEmptyDirVolume_HostPathParameters(options); err != nil {
		panic(err)
	}
	var returns HostPathVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		"hostPath",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns `true` if `x` is an EmptyDirVolume, `false` otherwise.
// Experimental.
func EmptyDirVolume_IsEmptyDirVolume(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmptyDirVolume_IsEmptyDirVolumeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		"isEmptyDirVolume",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Creates a Kubernetes Secret volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
// Experimental.
func EmptyDirVolume_Secret(options *SecretPathVolumeOptions) SecretPathVolume {
	_init_.Initialize()

	if err := validateEmptyDirVolume_SecretParameters(options); err != nil {
		panic(err)
	}
	var returns SecretPathVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		"secret",
		[]interface{}{options},
		&returns,
	)

	return returns
}

