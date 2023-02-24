package awsappsync


// Properties for defining a `CfnResolver`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverProps := &CfnResolverProps{
//   	ApiId: jsii.String("apiId"),
//   	FieldName: jsii.String("fieldName"),
//   	TypeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	CachingConfig: &CachingConfigProperty{
//   		Ttl: jsii.Number(123),
//
//   		// the properties below are optional
//   		CachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	Code: jsii.String("code"),
//   	CodeS3Location: jsii.String("codeS3Location"),
//   	DataSourceName: jsii.String("dataSourceName"),
//   	Kind: jsii.String("kind"),
//   	MaxBatchSize: jsii.Number(123),
//   	PipelineConfig: &PipelineConfigProperty{
//   		Functions: []*string{
//   			jsii.String("functions"),
//   		},
//   	},
//   	RequestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	RequestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	ResponseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	ResponseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	Runtime: &AppSyncRuntimeProperty{
//   		Name: jsii.String("name"),
//   		RuntimeVersion: jsii.String("runtimeVersion"),
//   	},
//   	SyncConfig: &SyncConfigProperty{
//   		ConflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		ConflictHandler: jsii.String("conflictHandler"),
//   		LambdaConflictHandlerConfig: &LambdaConflictHandlerConfigProperty{
//   			LambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   }
//
type CfnResolverProps struct {
	// The AWS AppSync GraphQL API to which you want to attach this resolver.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The GraphQL field on a type that invokes the resolver.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The GraphQL type that invokes this resolver.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for the resolver.
	CachingConfig interface{} `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The `resolver` code that contains the request and response functions.
	//
	// When code is used, the `runtime` is required. The runtime value must be `APPSYNC_JS` .
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Amazon S3 endpoint.
	CodeS3Location *string `field:"optional" json:"codeS3Location" yaml:"codeS3Location"`
	// The resolver data source name.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// The resolver type.
	//
	// - *UNIT* : A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
	// - *PIPELINE* : A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of `Function` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// Functions linked with the pipeline resolver.
	PipelineConfig interface{} `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template.
	//
	// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The location of a request mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The response mapping template.
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function.
	//
	// Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
	Runtime interface{} `field:"optional" json:"runtime" yaml:"runtime"`
	// The `SyncConfig` for a resolver attached to a versioned data source.
	SyncConfig interface{} `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

