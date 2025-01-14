package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
)

// Experimental.
type IExtension interface {
	awscdk.IResource
	// The actions for the extension.
	// Experimental.
	Actions() *[]Action
	// The description of the extension.
	// Experimental.
	Description() *string
	// The Amazon Resource Name (ARN) of the extension.
	// Experimental.
	ExtensionArn() *string
	// The ID of the extension.
	// Experimental.
	ExtensionId() *string
	// The version number of the extension.
	// Experimental.
	ExtensionVersionNumber() *float64
	// The latest version number of the extension.
	// Experimental.
	LatestVersionNumber() *float64
	// The name of the extension.
	// Experimental.
	Name() *string
	// The parameters of the extension.
	// Experimental.
	Parameters() *[]Parameter
}

// The jsii proxy for IExtension
type jsiiProxy_IExtension struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IExtension) Actions() *[]Action {
	var returns *[]Action
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extensionVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) LatestVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Parameters() *[]Parameter {
	var returns *[]Parameter
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

