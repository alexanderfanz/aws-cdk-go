// The CDK Construct Library for AWS::IoT
package awscdkiotalpha


// Properties for defining an AWS IoT Rule.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type TopicRuleProps struct {
	// A simplified SQL syntax to filter messages received on an MQTT topic and push the data elsewhere.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html
	//
	// Experimental.
	Sql IotSql `field:"required" json:"sql" yaml:"sql"`
	// The actions associated with the topic rule.
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// A textual description of the topic rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the rule is enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The action AWS IoT performs when it is unable to perform a rule's action.
	// Experimental.
	ErrorAction IAction `field:"optional" json:"errorAction" yaml:"errorAction"`
	// The name of the topic rule.
	// Experimental.
	TopicRuleName *string `field:"optional" json:"topicRuleName" yaml:"topicRuleName"`
}
