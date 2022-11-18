package cxapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Return the appropriate values for the environment placeholders.
// Experimental.
type IEnvironmentPlaceholderProvider interface {
	// Return the account.
	// Experimental.
	AccountId() *string
	// Return the partition.
	// Experimental.
	Partition() *string
	// Return the region.
	// Experimental.
	Region() *string
}

// The jsii proxy for IEnvironmentPlaceholderProvider
type jsiiProxy_IEnvironmentPlaceholderProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) AccountId() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"accountId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) Partition() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"partition",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) Region() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"region",
		nil, // no parameters
		&returns,
	)

	return returns
}

