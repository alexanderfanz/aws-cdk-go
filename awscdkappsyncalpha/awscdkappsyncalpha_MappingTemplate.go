// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MappingTemplates for AppSync resolvers.
//
// Example:
//   var dummyRequest mappingTemplate
//   var dummyResponse mappingTemplate
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &graphqlApiProps{
//   	name: jsii.String("demo"),
//   })
//
//   objectTypes := []interfaceType{
//   	node,
//   	filmNode,
//   }
//
//   filmConnections := generateEdgeAndConnection(filmNode)
//
//   api.addQuery(jsii.String("allFilms"), appsync.NewResolvableField(&resolvableFieldOptions{
//   	returnType: filmConnections.connection.attribute(),
//   	args: args,
//   	dataSource: api.addNoneDataSource(jsii.String("none")),
//   	requestMappingTemplate: dummyRequest,
//   	responseMappingTemplate: dummyResponse,
//   }))
//
//   api.addType(node)
//   api.addType(filmNode)
//   api.addType(filmConnections.edge)
//   api.addType(filmConnections.connection)
//
// Experimental.
type MappingTemplate interface {
	// this is called to render the mapping template to a VTL string.
	// Experimental.
	RenderTemplate() *string
}

// The jsii proxy struct for MappingTemplate
type jsiiProxy_MappingTemplate struct {
	_ byte // padding
}

// Experimental.
func NewMappingTemplate_Override(m MappingTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		nil, // no parameters
		m,
	)
}

// Mapping template to delete a single item from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbDeleteItem(keyName *string, idArg *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbDeleteItem",
		[]interface{}{keyName, idArg},
		&returns,
	)

	return returns
}

// Mapping template to get a single item from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbGetItem(keyName *string, idArg *string, consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbGetItem",
		[]interface{}{keyName, idArg, consistentRead},
		&returns,
	)

	return returns
}

// Mapping template to save a single item to a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbPutItem(key PrimaryKey, values AttributeValues) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbPutItem",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Mapping template to query a set of items from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbQuery(cond KeyCondition, indexName *string, consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbQuery",
		[]interface{}{cond, indexName, consistentRead},
		&returns,
	)

	return returns
}

// Mapping template for a single result item from DynamoDB.
// Experimental.
func MappingTemplate_DynamoDbResultItem() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbResultItem",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template for a result list from DynamoDB.
// Experimental.
func MappingTemplate_DynamoDbResultList() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbResultList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template to scan a DynamoDB table to fetch all entries.
// Experimental.
func MappingTemplate_DynamoDbScanTable(consistentRead *bool) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"dynamoDbScanTable",
		[]interface{}{consistentRead},
		&returns,
	)

	return returns
}

// Create a mapping template from the given file.
// Experimental.
func MappingTemplate_FromFile(fileName *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"fromFile",
		[]interface{}{fileName},
		&returns,
	)

	return returns
}

// Create a mapping template from the given string.
// Experimental.
func MappingTemplate_FromString(template *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Mapping template to invoke a Lambda function.
// Experimental.
func MappingTemplate_LambdaRequest(payload *string, operation *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"lambdaRequest",
		[]interface{}{payload, operation},
		&returns,
	)

	return returns
}

// Mapping template to return the Lambda result to the caller.
// Experimental.
func MappingTemplate_LambdaResult() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.MappingTemplate",
		"lambdaResult",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MappingTemplate) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}
