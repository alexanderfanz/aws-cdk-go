package awsiot


// Properties for defining an AWS IoT Rule.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &S3PutObjectActionProps{
//   			AccessControl: s3.BucketAccessControl_PUBLIC_READ,
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

