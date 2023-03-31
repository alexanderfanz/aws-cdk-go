package awsivschat

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRoom`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoomProps := &cfnRoomProps{
//   	loggingConfigurationIdentifiers: []*string{
//   		jsii.String("loggingConfigurationIdentifiers"),
//   	},
//   	maximumMessageLength: jsii.Number(123),
//   	maximumMessageRatePerSecond: jsii.Number(123),
//   	messageReviewHandler: &messageReviewHandlerProperty{
//   		fallbackResult: jsii.String("fallbackResult"),
//   		uri: jsii.String("uri"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRoomProps struct {
	// List of logging-configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers *[]*string `field:"optional" json:"loggingConfigurationIdentifiers" yaml:"loggingConfigurationIdentifiers"`
	// Maximum number of characters in a single message.
	//
	// Messages are expected to be UTF-8 encoded and this limit applies specifically to rune/code-point count, not number of bytes.
	MaximumMessageLength *float64 `field:"optional" json:"maximumMessageLength" yaml:"maximumMessageLength"`
	// Maximum number of messages per second that can be sent to the room (by all clients).
	MaximumMessageRatePerSecond *float64 `field:"optional" json:"maximumMessageRatePerSecond" yaml:"maximumMessageRatePerSecond"`
	// Configuration information for optional review of messages.
	MessageReviewHandler interface{} `field:"optional" json:"messageReviewHandler" yaml:"messageReviewHandler"`
	// Room name.
	//
	// The value does not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

