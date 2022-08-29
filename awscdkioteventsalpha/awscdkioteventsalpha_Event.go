// The CDK Construct Library for AWS::IoTEvents
package awscdkioteventsalpha


// Specifies the actions to be performed when the condition evaluates to `true`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_alpha "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//
//   var action iAction
//   var expression expression
//
//   event := &event{
//   	eventName: jsii.String("eventName"),
//
//   	// the properties below are optional
//   	actions: []*iAction{
//   		action,
//   	},
//   	condition: expression,
//   }
//
// Experimental.
type Event struct {
	// The name of the event.
	// Experimental.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The actions to be performed.
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The Boolean expression that, when `true`, causes the actions to be performed.
	// Experimental.
	Condition Expression `field:"optional" json:"condition" yaml:"condition"`
}
