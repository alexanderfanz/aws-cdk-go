# Amazon EventBridge Scheduler Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[Amazon EventBridge Scheduler](https://aws.amazon.com/blogs/compute/introducing-amazon-eventbridge-scheduler/) is a feature from Amazon EventBridge
that allows you to create, run, and manage scheduled tasks at scale. With EventBridge Scheduler, you can schedule one-time or recurrently tens
of millions of tasks across many AWS services without provisioning or managing underlying infrastructure.

This library contains integration classes for Amazon EventBridge Scheduler to call any
number of supported AWS Services.

The following targets are supported:

1. `targets.LambdaInvoke`: [Invoke an AWS Lambda function](#invoke-a-lambda-function))
2. `targets.StepFunctionsStartExecution`: [Start an AWS Step Function](#start-an-aws-step-function)
3. `targets.CodeBuildStartBuild`: [Start a CodeBuild job](#start-a-codebuild-job)
4. `targets.SqsSendMessage`: [Send a Message to an Amazon SQS Queue](#send-a-message-to-sqs-queue)
5. `targets.SnsPublish`: [Publish messages to an Amazon SNS topic](#publish-messages-to-an-amazon-sns-topic)
6. `targets.EventBridgePutEvents`: [Put Events on EventBridge](#send-events-to-an-eventbridge-event-bus)
7. `targets.InspectorStartAssessmentRun`: [Start an Amazon Inspector assessment run](#start-an-amazon-inspector-assessment-run)
8. `targets.KinesisStreamPutRecord`: [Put a record to an Amazon Kinesis Data Streams](#put-a-record-to-an-amazon-kinesis-data-streams)
9. `targets.KinesisDataFirehosePutRecord`: [Put a record to a Kinesis Data Firehose](#put-a-record-to-a-kinesis-data-firehose)
10. `targets.CodePipelineStartPipelineExecution`: [Start a CodePipeline execution](#start-a-codepipeline-execution)

## Invoke a Lambda function

Use the `LambdaInvoke` target to invoke a lambda function.

The code snippet below creates an event rule with a Lambda function as a target
called every hour by Event Bridge Scheduler with custom payload. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
})

dlq := sqs.NewQueue(this, jsii.String("DLQ"), &QueueProps{
	QueueName: jsii.String("MyDLQ"),
})

target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	DeadLetterQueue: dlq,
	MaxEventAge: awscdk.Duration_Minutes(jsii.Number(1)),
	RetryAttempts: jsii.Number(3),
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"payload": jsii.String("useful"),
	}),
})

schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: Target,
})
```

## Start an AWS Step Function

Use the `StepFunctionsStartExecution` target to start a new execution on a StepFunction.

The code snippet below creates an event rule with a Step Function as a target
called every hour by Event Bridge Scheduler with a custom payload.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import tasks "github.com/aws/aws-cdk-go/awscdk"


payload := map[string]*string{
	"Name": jsii.String("MyParameter"),
	"Value": jsii.String("🌥️"),
}

putParameterStep := tasks.NewCallAwsService(this, jsii.String("PutParameter"), &CallAwsServiceProps{
	Service: jsii.String("ssm"),
	Action: jsii.String("putParameter"),
	IamResources: []*string{
		jsii.String("*"),
	},
	Parameters: map[string]interface{}{
		"Name.$": jsii.String("$.Name"),
		"Value.$": jsii.String("$.Value"),
		"Type": jsii.String("String"),
		"Overwrite": jsii.Boolean(true),
	},
})

stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(putParameterStep),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: targets.NewStepFunctionsStartExecution(stateMachine, &ScheduleTargetBaseProps{
		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
	}),
})
```

## Start a CodeBuild job

Use the `CodeBuildStartBuild` target to start a new build run on a CodeBuild project.

The code snippet below creates an event rule with a CodeBuild project as target which is
called every hour by Event Bridge Scheduler.

```go
import codebuild "github.com/aws/aws-cdk-go/awscdk"

var project project


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewCodeBuildStartBuild(project),
})
```

## Send A Message To SQS Queue

Use the `SqsSendMessage` target to send a message to SQS Queue.

The code snippet below creates an event rule with a SQS Queue as a target
called every hour by Event Bridge Scheduler with a custom payload.

Contains the `messageGroupId` to use when the target is a FIFO queue. If you specify
a FIFO queue as a target, the queue must have content-based deduplication enabled.

```go
payload := "test"
messageGroupId := "id"
queue := sqs.NewQueue(this, jsii.String("MyQueue"), &QueueProps{
	Fifo: jsii.Boolean(true),
	ContentBasedDeduplication: jsii.Boolean(true),
})

target := targets.NewSqsSendMessage(queue, &SqsSendMessageProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromText(payload),
	MessageGroupId: jsii.String(MessageGroupId),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
	Target: Target,
})
```

## Publish messages to an Amazon SNS topic

Use the `SnsPublish` target to publish messages to an Amazon SNS topic.

The code snippets below create an event rule with a Amazon SNS topic as a target.
It's called every hour by Amazon Event Bridge Scheduler with custom payload.

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


topic := sns.NewTopic(this, jsii.String("Topic"))

payload := map[string]*string{
	"message": jsii.String("Hello scheduler!"),
}

target := targets.NewSnsPublish(topic, &ScheduleTargetBaseProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: Target,
})
```

## Send events to an EventBridge event bus

Use the `EventBridgePutEvents` target to send events to an EventBridge event bus.

The code snippet below creates an event rule with an EventBridge event bus as a target
called every hour by Event Bridge Scheduler with a custom event payload.

```go
import events "github.com/aws/aws-cdk-go/awscdk"


eventBus := events.NewEventBus(this, jsii.String("EventBus"), &EventBusProps{
	EventBusName: jsii.String("DomainEvents"),
})

eventEntry := &EventBridgePutEventsEntry{
	EventBus: EventBus,
	Source: jsii.String("PetService"),
	Detail: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"Name": jsii.String("Fluffy"),
	}),
	DetailType: jsii.String("🐶"),
}

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: targets.NewEventBridgePutEvents(eventEntry, &ScheduleTargetBaseProps{
	}),
})
```

## Start an Amazon Inspector assessment run

Use the `InspectorStartAssessmentRun` target to start an Inspector assessment run.

The code snippet below creates an event rule with an assessment template as target which is
called every hour by Event Bridge Scheduler.

```go
import inspector "github.com/aws/aws-cdk-go/awscdk"

var assessmentTemplate cfnAssessmentTemplate


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewInspectorStartAssessmentRun(assessmentTemplate),
})
```

## Put a record to an Amazon Kinesis Data Streams

Use the `KinesisStreamPutRecord` target to put a record to an Amazon Kinesis Data Streams.

The code snippet below creates an event rule with a stream as target which is
called every hour by Event Bridge Scheduler.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"


stream := kinesis.NewStream(this, jsii.String("MyStream"))

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewKinesisStreamPutRecord(stream, &KinesisStreamPutRecordProps{
		PartitionKey: jsii.String("key"),
	}),
})
```

## Put a record to a Kinesis Data Firehose

Use the `KinesisDataFirehosePutRecord` target to put a record to a Kinesis Data Firehose delivery stream.

The code snippet below creates an event rule with a delivery stream as a target
called every hour by Event Bridge Scheduler with a custom payload.

```go
import firehose "github.com/aws/aws-cdk-go/awscdk"
var deliveryStream cfnDeliveryStream


payload := map[string]*string{
	"Data": jsii.String("record"),
}

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewKinesisDataFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
	}),
})
```

## Start a CodePipeline execution

Use the `CodePipelineStartPipelineExecution` target to start a new execution for a CodePipeline pipeline.

The code snippet below creates an event rule with a CodePipeline pipeline as target which is
called every hour by Event Bridge Scheduler.

```go
import codepipeline "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewCodePipelineStartPipelineExecution(pipeline),
})
```
