package awscdkscheduleralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The text, or well-formed JSON, passed to the target of the schedule.
//
// Example:
//   input := awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   	"QueueName": jsii.String("MyQueue"),
//   })
//
// Experimental.
type ScheduleTargetInput interface {
	// Return the input properties for this input object.
	// Experimental.
	Bind(schedule ISchedule) *string
}

// The jsii proxy struct for ScheduleTargetInput
type jsiiProxy_ScheduleTargetInput struct {
	_ byte // padding
}

// Experimental.
func NewScheduleTargetInput_Override(s ScheduleTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetInput",
		nil, // no parameters
		s,
	)
}

// Pass a JSON object to the target, it is possible to embed `ContextAttributes` and other cdk references.
// Experimental.
func ScheduleTargetInput_FromObject(obj interface{}) ScheduleTargetInput {
	_init_.Initialize()

	if err := validateScheduleTargetInput_FromObjectParameters(obj); err != nil {
		panic(err)
	}
	var returns ScheduleTargetInput

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetInput",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Pass text to the target, it is possible to embed `ContextAttributes` that will be resolved to actual values while the CloudFormation is deployed or cdk Tokens that will be resolved when the CloudFormation templates are generated by CDK.
//
// The target input value will be a single string that you pass.
// For passing complex values like JSON object to a target use method
// `ScheduleTargetInput.fromObject()` instead.
// Experimental.
func ScheduleTargetInput_FromText(text *string) ScheduleTargetInput {
	_init_.Initialize()

	if err := validateScheduleTargetInput_FromTextParameters(text); err != nil {
		panic(err)
	}
	var returns ScheduleTargetInput

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetInput",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleTargetInput) Bind(schedule ISchedule) *string {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}
