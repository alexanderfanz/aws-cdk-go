package awsautoscalinghooktargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscalinghooktargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Use a Lambda Function as a hook target.
//
// Internally creates a Topic to make the connection.
//
// TODO: EXAMPLE
//
// Experimental.
type FunctionHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for FunctionHook
type jsiiProxy_FunctionHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewFunctionHook(fn awslambda.IFunction, encryptionKey awskms.IKey) FunctionHook {
	_init_.Initialize()

	j := jsiiProxy_FunctionHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		&j,
	)

	return &j
}

// Experimental.
func NewFunctionHook_Override(f FunctionHook, fn awslambda.IFunction, encryptionKey awskms.IKey) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		f,
	)
}

// Called when this object is used as the target of a lifecycle hook.
// Experimental.
func (f *jsiiProxy_FunctionHook) Bind(scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, lifecycleHook},
		&returns,
	)

	return returns
}

// Use an SQS queue as a hook target.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(_scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for QueueHook
type jsiiProxy_QueueHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewQueueHook(queue awssqs.IQueue) QueueHook {
	_init_.Initialize()

	j := jsiiProxy_QueueHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewQueueHook_Override(q QueueHook, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		q,
	)
}

// Called when this object is used as the target of a lifecycle hook.
// Experimental.
func (q *jsiiProxy_QueueHook) Bind(_scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		q,
		"bind",
		[]interface{}{_scope, lifecycleHook},
		&returns,
	)

	return returns
}

// Use an SNS topic as a hook target.
//
// TODO: EXAMPLE
//
// Experimental.
type TopicHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(_scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for TopicHook
type jsiiProxy_TopicHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewTopicHook(topic awssns.ITopic) TopicHook {
	_init_.Initialize()

	j := jsiiProxy_TopicHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicHook_Override(t TopicHook, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		t,
	)
}

// Called when this object is used as the target of a lifecycle hook.
// Experimental.
func (t *jsiiProxy_TopicHook) Bind(_scope awscdk.Construct, lifecycleHook awsautoscaling.ILifecycleHook) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope, lifecycleHook},
		&returns,
	)

	return returns
}

