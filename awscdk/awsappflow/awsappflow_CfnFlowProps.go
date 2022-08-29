package awsappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customProperties interface{}
//
//   cfnFlowProps := &cfnFlowProps{
//   	destinationFlowConfigList: []interface{}{
//   		&destinationFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   				customConnector: &customConnectorDestinationPropertiesProperty{
//   					entityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					customProperties: customProperties,
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				eventBridge: &eventBridgeDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				marketo: &marketoDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				redshift: &redshiftDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				s3: &s3DestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//   						preserveSourceDataTyping: jsii.Boolean(false),
//   					},
//   				},
//   				salesforce: &salesforceDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				sapoData: &sAPODataDestinationPropertiesProperty{
//   					objectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				snowflake: &snowflakeDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				upsolver: &upsolverDestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//   					s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   						prefixConfig: &prefixConfigProperty{
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   						},
//   						fileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				zendesk: &zendeskDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			apiVersion: jsii.String("apiVersion"),
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	flowName: jsii.String("flowName"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			amplitude: &amplitudeSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			customConnector: &customConnectorSourcePropertiesProperty{
//   				entityName: jsii.String("entityName"),
//
//   				// the properties below are optional
//   				customProperties: customProperties,
//   			},
//   			datadog: &datadogSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			dynatrace: &dynatraceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			inforNexus: &inforNexusSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				s3InputFormatConfig: &s3InputFormatConfigProperty{
//   					s3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataSourcePropertiesProperty{
//   				objectPath: jsii.String("objectPath"),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			singular: &singularSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			slack: &slackSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			trendmicro: &trendmicroSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			veeva: &veevaSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				documentType: jsii.String("documentType"),
//   				includeAllVersions: jsii.Boolean(false),
//   				includeRenditions: jsii.Boolean(false),
//   				includeSourceFiles: jsii.Boolean(false),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		apiVersion: jsii.String("apiVersion"),
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				amplitude: jsii.String("amplitude"),
//   				customConnector: jsii.String("customConnector"),
//   				datadog: jsii.String("datadog"),
//   				dynatrace: jsii.String("dynatrace"),
//   				googleAnalytics: jsii.String("googleAnalytics"),
//   				inforNexus: jsii.String("inforNexus"),
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				sapoData: jsii.String("sapoData"),
//   				serviceNow: jsii.String("serviceNow"),
//   				singular: jsii.String("singular"),
//   				slack: jsii.String("slack"),
//   				trendmicro: jsii.String("trendmicro"),
//   				veeva: jsii.String("veeva"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesObjectProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			firstExecutionFrom: jsii.Number(123),
//   			flowErrorDeactivationThreshold: jsii.Number(123),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsArn: jsii.String("kmsArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFlowProps struct {
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList interface{} `field:"required" json:"destinationFlowConfigList" yaml:"destinationFlowConfigList"`
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig interface{} `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks interface{} `field:"required" json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig interface{} `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// A user-entered description of the flow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// The tags used to organize, track, or control access for your flow.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
