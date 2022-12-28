package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SessionPinningFilter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionPinningFilter := awscdk.Aws_rds.sessionPinningFilter.of(jsii.String("filterName"))
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy-pinning
//
// Experimental.
type SessionPinningFilter interface {
	// Filter name.
	// Experimental.
	FilterName() *string
}

// The jsii proxy struct for SessionPinningFilter
type jsiiProxy_SessionPinningFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_SessionPinningFilter) FilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterName",
		&returns,
	)
	return returns
}


// custom filter.
// Experimental.
func SessionPinningFilter_Of(filterName *string) SessionPinningFilter {
	_init_.Initialize()

	if err := validateSessionPinningFilter_OfParameters(filterName); err != nil {
		panic(err)
	}
	var returns SessionPinningFilter

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.SessionPinningFilter",
		"of",
		[]interface{}{filterName},
		&returns,
	)

	return returns
}

func SessionPinningFilter_EXCLUDE_VARIABLE_SETS() SessionPinningFilter {
	_init_.Initialize()
	var returns SessionPinningFilter
	_jsii_.StaticGet(
		"monocdk.aws_rds.SessionPinningFilter",
		"EXCLUDE_VARIABLE_SETS",
		&returns,
	)
	return returns
}

