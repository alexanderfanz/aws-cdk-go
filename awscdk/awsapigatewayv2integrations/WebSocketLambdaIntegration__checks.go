//go:build !no_runtime_type_checking

package awsapigatewayv2integrations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (w *jsiiProxy_WebSocketLambdaIntegration) validateBindParameters(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewWebSocketLambdaIntegrationParameters(id *string, handler awslambda.IFunction) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	return nil
}

