package awsevents


// Contains additional parameters for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionHttpParametersProperty := &ConnectionHttpParametersProperty{
//   	BodyParameters: []interface{}{
//   		&ParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			IsValueSecret: jsii.Boolean(false),
//   		},
//   	},
//   	HeaderParameters: []interface{}{
//   		&ParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			IsValueSecret: jsii.Boolean(false),
//   		},
//   	},
//   	QueryStringParameters: []interface{}{
//   		&ParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			IsValueSecret: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnConnection_ConnectionHttpParametersProperty struct {
	// Contains additional body string parameters for the connection.
	BodyParameters interface{} `field:"optional" json:"bodyParameters" yaml:"bodyParameters"`
	// Contains additional header parameters for the connection.
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Contains additional query string parameters for the connection.
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

