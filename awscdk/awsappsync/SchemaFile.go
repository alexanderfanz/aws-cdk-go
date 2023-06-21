package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   api := appsync.NewGraphqlApi(stack, jsii.String("EventBridgeApi"), &GraphqlApiProps{
//   	Name: jsii.String("EventBridgeApi"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
//   })
//
//   bus := events.NewEventBus(stack, jsii.String("DestinationEventBus"), map[string]interface{}{
//   })
//
//   dataSource := api.AddEventBridgeDataSource(jsii.String("NoneDS"), bus)
//
//   dataSource.CreateResolver(jsii.String("EventResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("emitEvent"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type SchemaFile interface {
	ISchema
	// The definition for this schema.
	Definition() *string
	SetDefinition(val *string)
	// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
	Bind(api IGraphqlApi, _options *SchemaBindOptions) ISchemaConfig
}

// The jsii proxy struct for SchemaFile
type jsiiProxy_SchemaFile struct {
	jsiiProxy_ISchema
}

func (j *jsiiProxy_SchemaFile) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


func NewSchemaFile(options *SchemaProps) SchemaFile {
	_init_.Initialize()

	if err := validateNewSchemaFileParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchemaFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewSchemaFile_Override(s SchemaFile, options *SchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SchemaFile)SetDefinition(val *string) {
	if err := j.validateSetDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

// Generate a Schema from file.
//
// Returns: `SchemaAsset` with immutable schema defintion.
func SchemaFile_FromAsset(filePath *string) SchemaFile {
	_init_.Initialize()

	if err := validateSchemaFile_FromAssetParameters(filePath); err != nil {
		panic(err)
	}
	var returns SchemaFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFile) Bind(api IGraphqlApi, _options *SchemaBindOptions) ISchemaConfig {
	if err := s.validateBindParameters(api, _options); err != nil {
		panic(err)
	}
	var returns ISchemaConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api, _options},
		&returns,
	)

	return returns
}

