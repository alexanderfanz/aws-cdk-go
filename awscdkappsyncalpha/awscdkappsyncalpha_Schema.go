// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
)

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.addHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &httpDataSourceOptions{
//   	name: jsii.String("httpDsWithStepF"),
//   	description: jsii.String("from appsync to StepFunctions Workflow"),
//   	authorizationConfig: &awsIamConfig{
//   		signingRegion: jsii.String("us-east-1"),
//   		signingServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("callStepFunction"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type Schema interface {
	// The definition for this schema.
	// Experimental.
	Definition() *string
	// Experimental.
	SetDefinition(val *string)
	// Add a mutation field to the schema's Mutation. CDK will create an Object Type called 'Mutation'. For example,.
	//
	// type Mutation {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddMutation(fieldName *string, field ResolvableField) ObjectType
	// Add a query field to the schema's Query. CDK will create an Object Type called 'Query'. For example,.
	//
	// type Query {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddQuery(fieldName *string, field ResolvableField) ObjectType
	// Add a subscription field to the schema's Subscription. CDK will create an Object Type called 'Subscription'. For example,.
	//
	// type Subscription {
	//    fieldName: Field.returnType
	// }.
	// Experimental.
	AddSubscription(fieldName *string, field Field) ObjectType
	// Escape hatch to add to Schema as desired.
	//
	// Will always result
	// in a newline.
	// Experimental.
	AddToSchema(addition *string, delimiter *string)
	// Add type to the schema.
	// Experimental.
	AddType(type_ IIntermediateType) IIntermediateType
	// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
	// Experimental.
	Bind(api GraphqlApi) awsappsync.CfnGraphQLSchema
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

func (j *jsiiProxy_Schema) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


// Experimental.
func NewSchema(options *SchemaOptions) Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Schema",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewSchema_Override(s Schema, options *SchemaOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Schema",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_Schema) SetDefinition(val *string) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

// Generate a Schema from file.
//
// Returns: `SchemaAsset` with immutable schema defintion.
// Experimental.
func Schema_FromAsset(filePath *string) Schema {
	_init_.Initialize()

	var returns Schema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Schema",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddMutation(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addMutation",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddQuery(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addQuery",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddSubscription(fieldName *string, field Field) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addSubscription",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) AddToSchema(addition *string, delimiter *string) {
	_jsii_.InvokeVoid(
		s,
		"addToSchema",
		[]interface{}{addition, delimiter},
	)
}

func (s *jsiiProxy_Schema) AddType(type_ IIntermediateType) IIntermediateType {
	var returns IIntermediateType

	_jsii_.Invoke(
		s,
		"addType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) Bind(api GraphqlApi) awsappsync.CfnGraphQLSchema {
	var returns awsappsync.CfnGraphQLSchema

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api},
		&returns,
	)

	return returns
}
