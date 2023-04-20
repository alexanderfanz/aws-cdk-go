package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTEvents::DetectorModel`.
//
// The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states* . For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide* .
//
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or AWS CloudFormation ) all detector instances created by the model are reset to their initial states. (The detector's `state` , and the values of any variables and timers are reset.)
// >
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or AWS CloudFormation ) the version number of the detector model is incremented. (A detector model with version number 1 before the update has version number 2 after the update succeeds.)
// >
// > If you attempt to update a detector model using AWS CloudFormation and the update does not succeed, the system may, in some cases, restore the original detector model. When this occurs, the detector model's version is incremented twice (for example, from version 1 to version 3) and the detector instances are reset.
// >
// > Also, be aware that if you attempt to update several detector models at once using AWS CloudFormation , some updates may succeed and others fail. In this case, the effects on each detector model's detector instances and version number depend on whether the update succeeded or failed, with the results as stated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorModel := awscdk.Aws_iotevents.NewCfnDetectorModel(this, jsii.String("MyCfnDetectorModel"), &CfnDetectorModelProps{
//   	DetectorModelDefinition: &DetectorModelDefinitionProperty{
//   		InitialStateName: jsii.String("initialStateName"),
//   		States: []interface{}{
//   			&StateProperty{
//   				StateName: jsii.String("stateName"),
//
//   				// the properties below are optional
//   				OnEnter: &OnEnterProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				OnExit: &OnExitProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				OnInput: &OnInputProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   					TransitionEvents: []interface{}{
//   						&TransitionEventProperty{
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   							NextState: jsii.String("nextState"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	DetectorModelDescription: jsii.String("detectorModelDescription"),
//   	DetectorModelName: jsii.String("detectorModelName"),
//   	EvaluationMethod: jsii.String("evaluationMethod"),
//   	Key: jsii.String("key"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDetectorModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Information that defines how a detector operates.
	DetectorModelDefinition() interface{}
	SetDetectorModelDefinition(val interface{})
	// A brief description of the detector model.
	DetectorModelDescription() *string
	SetDetectorModelDescription(val *string)
	// The name of the detector model.
	DetectorModelName() *string
	SetDetectorModelName(val *string)
	// Information about the order in which events are evaluated and how actions are executed.
	EvaluationMethod() *string
	SetEvaluationMethod(val *string)
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	Key() *string
	SetKey(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDetectorModel
type jsiiProxy_CfnDetectorModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDetectorModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectorModelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) EvaluationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel(scope awscdk.Construct, id *string, props *CfnDetectorModelProps) CfnDetectorModel {
	_init_.Initialize()

	if err := validateNewCfnDetectorModelParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDetectorModel{}

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel_Override(c CfnDetectorModel, scope awscdk.Construct, id *string, props *CfnDetectorModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetDetectorModelDefinition(val interface{}) {
	if err := j.validateSetDetectorModelDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectorModelDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetDetectorModelDescription(val *string) {
	_jsii_.Set(
		j,
		"detectorModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetDetectorModelName(val *string) {
	_jsii_.Set(
		j,
		"detectorModelName",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetEvaluationMethod(val *string) {
	_jsii_.Set(
		j,
		"evaluationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDetectorModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDetectorModel_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDetectorModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDetectorModel_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDetectorModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDetectorModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDetectorModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDetectorModel) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDetectorModel) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDetectorModel) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDetectorModel) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDetectorModel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDetectorModel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDetectorModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDetectorModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

