package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoT::TopicRule`.
//
// Use the `AWS::IoT::TopicRule` resource to declare an AWS IoT rule. For information about working with AWS IoT rules, see [Rules for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicRule := awscdk.Aws_iot.NewCfnTopicRule(this, jsii.String("MyCfnTopicRule"), &CfnTopicRuleProps{
//   	TopicRulePayload: &TopicRulePayloadProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   					AlarmName: jsii.String("alarmName"),
//   					RoleArn: jsii.String("roleArn"),
//   					StateReason: jsii.String("stateReason"),
//   					StateValue: jsii.String("stateValue"),
//   				},
//   				CloudwatchLogs: &CloudwatchLogsActionProperty{
//   					LogGroupName: jsii.String("logGroupName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   				},
//   				CloudwatchMetric: &CloudwatchMetricActionProperty{
//   					MetricName: jsii.String("metricName"),
//   					MetricNamespace: jsii.String("metricNamespace"),
//   					MetricUnit: jsii.String("metricUnit"),
//   					MetricValue: jsii.String("metricValue"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					MetricTimestamp: jsii.String("metricTimestamp"),
//   				},
//   				DynamoDb: &DynamoDBActionProperty{
//   					HashKeyField: jsii.String("hashKeyField"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					RoleArn: jsii.String("roleArn"),
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					HashKeyType: jsii.String("hashKeyType"),
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				DynamoDBv2: &DynamoDBv2ActionProperty{
//   					PutItem: &PutItemInputProperty{
//   						TableName: jsii.String("tableName"),
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Elasticsearch: &ElasticsearchActionProperty{
//   					Endpoint: jsii.String("endpoint"),
//   					Id: jsii.String("id"),
//   					Index: jsii.String("index"),
//   					RoleArn: jsii.String("roleArn"),
//   					Type: jsii.String("type"),
//   				},
//   				Firehose: &FirehoseActionProperty{
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   					Separator: jsii.String("separator"),
//   				},
//   				Http: &HttpActionProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					Auth: &HttpAuthorizationProperty{
//   						Sigv4: &SigV4AuthorizationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							ServiceName: jsii.String("serviceName"),
//   							SigningRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					ConfirmationUrl: jsii.String("confirmationUrl"),
//   					Headers: []interface{}{
//   						&HttpActionHeaderProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				IotAnalytics: &IotAnalyticsActionProperty{
//   					ChannelName: jsii.String("channelName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   				},
//   				IotEvents: &IotEventsActionProperty{
//   					InputName: jsii.String("inputName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   					MessageId: jsii.String("messageId"),
//   				},
//   				IotSiteWise: &IotSiteWiseActionProperty{
//   					PutAssetPropertyValueEntries: []interface{}{
//   						&PutAssetPropertyValueEntryProperty{
//   							PropertyValues: []interface{}{
//   								&AssetPropertyValueProperty{
//   									Timestamp: &AssetPropertyTimestampProperty{
//   										TimeInSeconds: jsii.String("timeInSeconds"),
//
//   										// the properties below are optional
//   										OffsetInNanos: jsii.String("offsetInNanos"),
//   									},
//   									Value: &AssetPropertyVariantProperty{
//   										BooleanValue: jsii.String("booleanValue"),
//   										DoubleValue: jsii.String("doubleValue"),
//   										IntegerValue: jsii.String("integerValue"),
//   										StringValue: jsii.String("stringValue"),
//   									},
//
//   									// the properties below are optional
//   									Quality: jsii.String("quality"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   						},
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Kafka: &KafkaActionProperty{
//   					ClientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					DestinationArn: jsii.String("destinationArn"),
//   					Topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					Key: jsii.String("key"),
//   					Partition: jsii.String("partition"),
//   				},
//   				Kinesis: &KinesisActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					StreamName: jsii.String("streamName"),
//
//   					// the properties below are optional
//   					PartitionKey: jsii.String("partitionKey"),
//   				},
//   				Lambda: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   				Location: &LocationActionProperty{
//   					DeviceId: jsii.String("deviceId"),
//   					Latitude: jsii.String("latitude"),
//   					Longitude: jsii.String("longitude"),
//   					RoleArn: jsii.String("roleArn"),
//   					TrackerName: jsii.String("trackerName"),
//
//   					// the properties below are optional
//   					Timestamp: NewDate(),
//   				},
//   				OpenSearch: &OpenSearchActionProperty{
//   					Endpoint: jsii.String("endpoint"),
//   					Id: jsii.String("id"),
//   					Index: jsii.String("index"),
//   					RoleArn: jsii.String("roleArn"),
//   					Type: jsii.String("type"),
//   				},
//   				Republish: &RepublishActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					Topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					Headers: &RepublishActionHeadersProperty{
//   						ContentType: jsii.String("contentType"),
//   						CorrelationData: jsii.String("correlationData"),
//   						MessageExpiry: jsii.String("messageExpiry"),
//   						PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   						ResponseTopic: jsii.String("responseTopic"),
//   						UserProperties: []interface{}{
//   							&UserPropertyProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Qos: jsii.Number(123),
//   				},
//   				S3: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//   					Key: jsii.String("key"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					CannedAcl: jsii.String("cannedAcl"),
//   				},
//   				Sns: &SnsActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					TargetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					MessageFormat: jsii.String("messageFormat"),
//   				},
//   				Sqs: &SqsActionProperty{
//   					QueueUrl: jsii.String("queueUrl"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					UseBase64: jsii.Boolean(false),
//   				},
//   				StepFunctions: &StepFunctionsActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					StateMachineName: jsii.String("stateMachineName"),
//
//   					// the properties below are optional
//   					ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   				},
//   				Timestream: &TimestreamActionProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					Dimensions: []interface{}{
//   						&TimestreamDimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					Timestamp: &TimestreamTimestampProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Sql: jsii.String("sql"),
//
//   		// the properties below are optional
//   		AwsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   		Description: jsii.String("description"),
//   		ErrorAction: &ActionProperty{
//   			CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   				AlarmName: jsii.String("alarmName"),
//   				RoleArn: jsii.String("roleArn"),
//   				StateReason: jsii.String("stateReason"),
//   				StateValue: jsii.String("stateValue"),
//   			},
//   			CloudwatchLogs: &CloudwatchLogsActionProperty{
//   				LogGroupName: jsii.String("logGroupName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   			},
//   			CloudwatchMetric: &CloudwatchMetricActionProperty{
//   				MetricName: jsii.String("metricName"),
//   				MetricNamespace: jsii.String("metricNamespace"),
//   				MetricUnit: jsii.String("metricUnit"),
//   				MetricValue: jsii.String("metricValue"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				MetricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			DynamoDb: &DynamoDBActionProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				HashKeyType: jsii.String("hashKeyType"),
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			DynamoDBv2: &DynamoDBv2ActionProperty{
//   				PutItem: &PutItemInputProperty{
//   					TableName: jsii.String("tableName"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Elasticsearch: &ElasticsearchActionProperty{
//   				Endpoint: jsii.String("endpoint"),
//   				Id: jsii.String("id"),
//   				Index: jsii.String("index"),
//   				RoleArn: jsii.String("roleArn"),
//   				Type: jsii.String("type"),
//   			},
//   			Firehose: &FirehoseActionProperty{
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   				Separator: jsii.String("separator"),
//   			},
//   			Http: &HttpActionProperty{
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				Auth: &HttpAuthorizationProperty{
//   					Sigv4: &SigV4AuthorizationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						ServiceName: jsii.String("serviceName"),
//   						SigningRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				ConfirmationUrl: jsii.String("confirmationUrl"),
//   				Headers: []interface{}{
//   					&HttpActionHeaderProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			IotAnalytics: &IotAnalyticsActionProperty{
//   				ChannelName: jsii.String("channelName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   			},
//   			IotEvents: &IotEventsActionProperty{
//   				InputName: jsii.String("inputName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   				MessageId: jsii.String("messageId"),
//   			},
//   			IotSiteWise: &IotSiteWiseActionProperty{
//   				PutAssetPropertyValueEntries: []interface{}{
//   					&PutAssetPropertyValueEntryProperty{
//   						PropertyValues: []interface{}{
//   							&AssetPropertyValueProperty{
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								Quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						AssetId: jsii.String("assetId"),
//   						EntryId: jsii.String("entryId"),
//   						PropertyAlias: jsii.String("propertyAlias"),
//   						PropertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Kafka: &KafkaActionProperty{
//   				ClientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				DestinationArn: jsii.String("destinationArn"),
//   				Topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				Key: jsii.String("key"),
//   				Partition: jsii.String("partition"),
//   			},
//   			Kinesis: &KinesisActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				StreamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				PartitionKey: jsii.String("partitionKey"),
//   			},
//   			Lambda: &LambdaActionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   			},
//   			Location: &LocationActionProperty{
//   				DeviceId: jsii.String("deviceId"),
//   				Latitude: jsii.String("latitude"),
//   				Longitude: jsii.String("longitude"),
//   				RoleArn: jsii.String("roleArn"),
//   				TrackerName: jsii.String("trackerName"),
//
//   				// the properties below are optional
//   				Timestamp: NewDate(),
//   			},
//   			OpenSearch: &OpenSearchActionProperty{
//   				Endpoint: jsii.String("endpoint"),
//   				Id: jsii.String("id"),
//   				Index: jsii.String("index"),
//   				RoleArn: jsii.String("roleArn"),
//   				Type: jsii.String("type"),
//   			},
//   			Republish: &RepublishActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				Topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				Headers: &RepublishActionHeadersProperty{
//   					ContentType: jsii.String("contentType"),
//   					CorrelationData: jsii.String("correlationData"),
//   					MessageExpiry: jsii.String("messageExpiry"),
//   					PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   					ResponseTopic: jsii.String("responseTopic"),
//   					UserProperties: []interface{}{
//   						&UserPropertyProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Qos: jsii.Number(123),
//   			},
//   			S3: &S3ActionProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				CannedAcl: jsii.String("cannedAcl"),
//   			},
//   			Sns: &SnsActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				TargetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				MessageFormat: jsii.String("messageFormat"),
//   			},
//   			Sqs: &SqsActionProperty{
//   				QueueUrl: jsii.String("queueUrl"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				UseBase64: jsii.Boolean(false),
//   			},
//   			StepFunctions: &StepFunctionsActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				StateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			Timestream: &TimestreamActionProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   				Dimensions: []interface{}{
//   					&TimestreamDimensionProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				Timestamp: &TimestreamTimestampProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		RuleDisabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	RuleName: jsii.String("ruleName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnTopicRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
	AttrArn() *string
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
	// The name of the rule.
	RuleName() *string
	SetRuleName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags() awscdk.TagManager
	// The rule payload.
	TopicRulePayload() interface{}
	SetTopicRulePayload(val interface{})
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

// The jsii proxy struct for CfnTopicRule
type jsiiProxy_CfnTopicRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopicRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) TopicRulePayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicRulePayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule(scope awscdk.Construct, id *string, props *CfnTopicRuleProps) CfnTopicRule {
	_init_.Initialize()

	if err := validateNewCfnTopicRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicRule{}

	_jsii_.Create(
		"monocdk.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule_Override(c CfnTopicRule, scope awscdk.Construct, id *string, props *CfnTopicRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopicRule)SetRuleName(val *string) {
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_CfnTopicRule)SetTopicRulePayload(val interface{}) {
	if err := j.validateSetTopicRulePayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicRulePayload",
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
func CfnTopicRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iot.CfnTopicRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTopicRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iot.CfnTopicRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iot.CfnTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iot.CfnTopicRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicRule) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopicRule) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnTopicRule) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTopicRule) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopicRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopicRule) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopicRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopicRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTopicRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTopicRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

