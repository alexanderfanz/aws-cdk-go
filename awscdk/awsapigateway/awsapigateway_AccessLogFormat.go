package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// factory methods for access log format.
//
// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		accessLogFormat: apigateway.accessLogFormat.custom(
//   		fmt.Sprintf("%v %v %v", apigateway.accessLogField.contextRequestId(), apigateway.*accessLogField.contextErrorMessage(), apigateway.*accessLogField.contextErrorMessageString())),
//   	},
//   })
//
type AccessLogFormat interface {
	// Output a format string to be used with CloudFormation.
	ToString() *string
}

// The jsii proxy struct for AccessLogFormat
type jsiiProxy_AccessLogFormat struct {
	_ byte // padding
}

// Generate Common Log Format.
func AccessLogFormat_Clf() AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"clf",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Custom log format.
//
// You can create any log format string. You can easily get the $ context variable by using the methods of AccessLogField.
//
// Example:
//   apigateway.accessLogFormat.custom(jSON.stringify(map[string]interface{}{
//   	"requestId": apigateway.AccessLogField.contextRequestId(),
//   	"sourceIp": apigateway.AccessLogField.contextIdentitySourceIp(),
//   	"method": apigateway.AccessLogField.contextHttpMethod(),
//   	"userContext": map[string]*string{
//   		"sub": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("sub")),
//   		"email": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("email")),
//   	},
//   }))
//
func AccessLogFormat_Custom(format *string) AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"custom",
		[]interface{}{format},
		&returns,
	)

	return returns
}

// Access log will be produced in the JSON format with a set of fields most useful in the access log.
//
// All fields are turned on by default with the
// option to turn off specific fields.
func AccessLogFormat_JsonWithStandardFields(fields *JsonWithStandardFieldProps) AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"jsonWithStandardFields",
		[]interface{}{fields},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessLogFormat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
