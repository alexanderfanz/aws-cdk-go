package awspipes


// The parameters required to set up a target for your pipe.
//
// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetParametersProperty := &PipeTargetParametersProperty{
//   	BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   		JobDefinition: jsii.String("jobDefinition"),
//   		JobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		ArrayProperties: &BatchArrayPropertiesProperty{
//   			Size: jsii.Number(123),
//   		},
//   		ContainerOverrides: &BatchContainerOverridesProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Environment: []interface{}{
//   				&BatchEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			ResourceRequirements: []interface{}{
//   				&BatchResourceRequirementProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		DependsOn: []interface{}{
//   			&BatchJobDependencyProperty{
//   				JobId: jsii.String("jobId"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		RetryStrategy: &BatchRetryStrategyProperty{
//   			Attempts: jsii.Number(123),
//   		},
//   	},
//   	CloudWatchLogsParameters: &PipeTargetCloudWatchLogsParametersProperty{
//   		LogStreamName: jsii.String("logStreamName"),
//   		Timestamp: jsii.String("timestamp"),
//   	},
//   	EcsTaskParameters: &PipeTargetEcsTaskParametersProperty{
//   		TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		CapacityProviderStrategy: []interface{}{
//   			&CapacityProviderStrategyItemProperty{
//   				CapacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				Base: jsii.Number(123),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		EnableEcsManagedTags: jsii.Boolean(false),
//   		EnableExecuteCommand: jsii.Boolean(false),
//   		Group: jsii.String("group"),
//   		LaunchType: jsii.String("launchType"),
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   			},
//   		},
//   		Overrides: &EcsTaskOverrideProperty{
//   			ContainerOverrides: []interface{}{
//   				&EcsContainerOverrideProperty{
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Cpu: jsii.Number(123),
//   					Environment: []interface{}{
//   						&EcsEnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					EnvironmentFiles: []interface{}{
//   						&EcsEnvironmentFileProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Memory: jsii.Number(123),
//   					MemoryReservation: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					ResourceRequirements: []interface{}{
//   						&EcsResourceRequirementProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			Cpu: jsii.String("cpu"),
//   			EphemeralStorage: &EcsEphemeralStorageProperty{
//   				SizeInGiB: jsii.Number(123),
//   			},
//   			ExecutionRoleArn: jsii.String("executionRoleArn"),
//   			InferenceAcceleratorOverrides: []interface{}{
//   				&EcsInferenceAcceleratorOverrideProperty{
//   					DeviceName: jsii.String("deviceName"),
//   					DeviceType: jsii.String("deviceType"),
//   				},
//   			},
//   			Memory: jsii.String("memory"),
//   			TaskRoleArn: jsii.String("taskRoleArn"),
//   		},
//   		PlacementConstraints: []interface{}{
//   			&PlacementConstraintProperty{
//   				Expression: jsii.String("expression"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlacementStrategy: []interface{}{
//   			&PlacementStrategyProperty{
//   				Field: jsii.String("field"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlatformVersion: jsii.String("platformVersion"),
//   		PropagateTags: jsii.String("propagateTags"),
//   		ReferenceId: jsii.String("referenceId"),
//   		Tags: []cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TaskCount: jsii.Number(123),
//   	},
//   	EventBridgeEventBusParameters: &PipeTargetEventBridgeEventBusParametersProperty{
//   		DetailType: jsii.String("detailType"),
//   		EndpointId: jsii.String("endpointId"),
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Source: jsii.String("source"),
//   		Time: jsii.String("time"),
//   	},
//   	HttpParameters: &PipeTargetHttpParametersProperty{
//   		HeaderParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		PathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		QueryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	InputTemplate: jsii.String("inputTemplate"),
//   	KinesisStreamParameters: &PipeTargetKinesisStreamParametersProperty{
//   		PartitionKey: jsii.String("partitionKey"),
//   	},
//   	LambdaFunctionParameters: &PipeTargetLambdaFunctionParametersProperty{
//   		InvocationType: jsii.String("invocationType"),
//   	},
//   	RedshiftDataParameters: &PipeTargetRedshiftDataParametersProperty{
//   		Database: jsii.String("database"),
//   		Sqls: []*string{
//   			jsii.String("sqls"),
//   		},
//
//   		// the properties below are optional
//   		DbUser: jsii.String("dbUser"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		StatementName: jsii.String("statementName"),
//   		WithEvent: jsii.Boolean(false),
//   	},
//   	SageMakerPipelineParameters: &PipeTargetSageMakerPipelineParametersProperty{
//   		PipelineParameterList: []interface{}{
//   			&SageMakerPipelineParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SqsQueueParameters: &PipeTargetSqsQueueParametersProperty{
//   		MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   		MessageGroupId: jsii.String("messageGroupId"),
//   	},
//   	StepFunctionStateMachineParameters: &PipeTargetStateMachineParametersProperty{
//   		InvocationType: jsii.String("invocationType"),
//   	},
//   }
//
type CfnPipe_PipeTargetParametersProperty struct {
	// The parameters for using an AWS Batch job as a target.
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// The parameters for using an CloudWatch Logs log stream as a target.
	CloudWatchLogsParameters interface{} `field:"optional" json:"cloudWatchLogsParameters" yaml:"cloudWatchLogsParameters"`
	// The parameters for using an Amazon ECS task as a target.
	EcsTaskParameters interface{} `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// The parameters for using an EventBridge event bus as a target.
	EventBridgeEventBusParameters interface{} `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
	// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// The parameters for using a Kinesis stream as a source.
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// The parameters for using a Lambda function as a target.
	LambdaFunctionParameters interface{} `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement.
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The parameters for using a SageMaker pipeline as a target.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// The parameters for using a Amazon SQS stream as a source.
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// The parameters for using a Step Functions state machine as a target.
	StepFunctionStateMachineParameters interface{} `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

