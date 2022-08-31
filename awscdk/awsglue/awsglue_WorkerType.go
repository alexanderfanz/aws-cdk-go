package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerType := awscdk.Aws_glue.workerType_G_1X()
//
// Experimental.
type WorkerType interface {
	// The name of this WorkerType, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for WorkerType
type jsiiProxy_WorkerType struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkerType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom worker type.
// Experimental.
func WorkerType_Of(workerType *string) WorkerType {
	_init_.Initialize()

	if err := validateWorkerType_OfParameters(workerType); err != nil {
		panic(err)
	}
	var returns WorkerType

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.WorkerType",
		"of",
		[]interface{}{workerType},
		&returns,
	)

	return returns
}

func WorkerType_G_1X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"monocdk.aws_glue.WorkerType",
		"G_1X",
		&returns,
	)
	return returns
}

func WorkerType_G_2X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"monocdk.aws_glue.WorkerType",
		"G_2X",
		&returns,
	)
	return returns
}

func WorkerType_STANDARD() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"monocdk.aws_glue.WorkerType",
		"STANDARD",
		&returns,
	)
	return returns
}
