package awsiotevents


// When entering this state, perform these `actions` if the `condition` is TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onEnterProperty := &onEnterProperty{
//   	events: []interface{}{
//   		&eventProperty{
//   			eventName: jsii.String("eventName"),
//
//   			// the properties below are optional
//   			actions: []interface{}{
//   				&actionProperty{
//   					clearTimer: &clearTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					dynamoDb: &dynamoDBProperty{
//   						hashKeyField: jsii.String("hashKeyField"),
//   						hashKeyValue: jsii.String("hashKeyValue"),
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						hashKeyType: jsii.String("hashKeyType"),
//   						operation: jsii.String("operation"),
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						payloadField: jsii.String("payloadField"),
//   						rangeKeyField: jsii.String("rangeKeyField"),
//   						rangeKeyType: jsii.String("rangeKeyType"),
//   						rangeKeyValue: jsii.String("rangeKeyValue"),
//   					},
//   					dynamoDBv2: &dynamoDBv2Property{
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					firehose: &firehoseProperty{
//   						deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						separator: jsii.String("separator"),
//   					},
//   					iotEvents: &iotEventsProperty{
//   						inputName: jsii.String("inputName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					iotSiteWise: &iotSiteWiseProperty{
//   						propertyValue: &assetPropertyValueProperty{
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   					iotTopicPublish: &iotTopicPublishProperty{
//   						mqttTopic: jsii.String("mqttTopic"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					lambda: &lambdaProperty{
//   						functionArn: jsii.String("functionArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					resetTimer: &resetTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					setTimer: &setTimerProperty{
//   						timerName: jsii.String("timerName"),
//
//   						// the properties below are optional
//   						durationExpression: jsii.String("durationExpression"),
//   						seconds: jsii.Number(123),
//   					},
//   					setVariable: &setVariableProperty{
//   						value: jsii.String("value"),
//   						variableName: jsii.String("variableName"),
//   					},
//   					sns: &snsProperty{
//   						targetArn: jsii.String("targetArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					sqs: &sqsProperty{
//   						queueUrl: jsii.String("queueUrl"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						useBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			condition: jsii.String("condition"),
//   		},
//   	},
//   }
//
type CfnDetectorModel_OnEnterProperty struct {
	// Specifies the actions that are performed when the state is entered and the `condition` is `TRUE` .
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}
