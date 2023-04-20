package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3notifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Use an SNS topic as a bucket notification destination.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
//
// Experimental.
type SnsDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	// Experimental.
	Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

// Experimental.
func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	if err := validateNewSnsDestinationParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsDestination) Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	if err := s.validateBindParameters(_scope, bucket); err != nil {
		panic(err)
	}
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

