package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An AWS service for an interface VPC endpoint.
//
// Example:
//   // Add gateway endpoints when creating the VPC
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
//   	gatewayEndpoints: map[string]gatewayVpcEndpointOptions{
//   		"S3": &gatewayVpcEndpointOptions{
//   			"service": ec2.GatewayVpcEndpointAwsService_S3(),
//   		},
//   	},
//   })
//
//   // Alternatively gateway endpoints can be added on the VPC
//   dynamoDbEndpoint := vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &gatewayVpcEndpointOptions{
//   	service: ec2.gatewayVpcEndpointAwsService_DYNAMODB(),
//   })
//
//   // This allows to customize the endpoint policy
//   dynamoDbEndpoint.addToPolicy(
//   iam.NewPolicyStatement(&policyStatementProps{
//   	 // Restrict to listing and describing tables
//   	principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	actions: []*string{
//   		jsii.String("dynamodb:DescribeTable"),
//   		jsii.String("dynamodb:ListTables"),
//   	},
//   	resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
//   // Add an interface endpoint
//   vpc.addInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &interfaceVpcEndpointOptions{
//   	service: ec2.interfaceVpcEndpointAwsService_ECR_DOCKER(),
//   })
//
type InterfaceVpcEndpointAwsService interface {
	IInterfaceVpcEndpointService
	// The name of the service.
	//
	// e.g. com.amazonaws.us-east-1.ecs
	Name() *string
	// The port of the service.
	Port() *float64
	// Whether Private DNS is supported by default.
	PrivateDnsDefault() *bool
	// The short name of the service.
	//
	// e.g. ecs
	ShortName() *string
}

// The jsii proxy struct for InterfaceVpcEndpointAwsService
type jsiiProxy_InterfaceVpcEndpointAwsService struct {
	jsiiProxy_IInterfaceVpcEndpointService
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) PrivateDnsDefault() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateDnsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}


func NewInterfaceVpcEndpointAwsService(name *string, prefix *string, port *float64) InterfaceVpcEndpointAwsService {
	_init_.Initialize()

	j := jsiiProxy_InterfaceVpcEndpointAwsService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		[]interface{}{name, prefix, port},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpointAwsService_Override(i InterfaceVpcEndpointAwsService, name *string, prefix *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		[]interface{}{name, prefix, port},
		i,
	)
}

func InterfaceVpcEndpointAwsService_APIGATEWAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APIGATEWAY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPLICATION_AUTOSCALING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPLICATION_AUTOSCALING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ATHENA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ATHENA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AUTOSCALING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AUTOSCALING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AUTOSCALING_PLANS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AUTOSCALING_PLANS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BATCH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BATCH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDFORMATION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDFORMATION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDTRAIL() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_EVENTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_EVENTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_LOGS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_LOGS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEARTIFACT_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEARTIFACT_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEARTIFACT_REPOSITORIES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEARTIFACT_REPOSITORIES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEBUILD() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEBUILD",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEBUILD_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEBUILD_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_GIT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_GIT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_GIT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_GIT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEGURU_PROFILER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEGURU_PROFILER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEGURU_REVIEWER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEGURU_REVIEWER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEPIPELINE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEPIPELINE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONFIG() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONFIG",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EC2() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EC2",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EC2_MESSAGES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EC2_MESSAGES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECR_DOCKER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECR_DOCKER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS_AGENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS_AGENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS_TELEMETRY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS_TELEMETRY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_FILESYSTEM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_FILESYSTEM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_FILESYSTEM_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_FILESYSTEM_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_INFERENCE_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_INFERENCE_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_LOAD_BALANCING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_LOAD_BALANCING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GLUE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GLUE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KEYSPACES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KEYSPACES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KINESIS_FIREHOSE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KINESIS_FIREHOSE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KINESIS_STREAMS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KINESIS_STREAMS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KMS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KMS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LAMBDA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LAMBDA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_RDS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"RDS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_RDS_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"RDS_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_S3() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"S3",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_NOTEBOOK() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_NOTEBOOK",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_RUNTIME_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_RUNTIME_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SECRETS_MANAGER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SECRETS_MANAGER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVICE_CATALOG() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVICE_CATALOG",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SNS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SNS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SQS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SQS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM_MESSAGES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM_MESSAGES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STEP_FUNCTIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STEP_FUNCTIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STORAGE_GATEWAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STORAGE_GATEWAY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TEXTRACT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TEXTRACT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TEXTRACT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TEXTRACT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSCRIBE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSCRIBE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSFER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSFER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_XRAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"XRAY",
		&returns,
	)
	return returns
}
