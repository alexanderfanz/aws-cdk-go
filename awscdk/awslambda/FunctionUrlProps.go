package awslambda


// Properties for a FunctionUrl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   functionUrlProps := &FunctionUrlProps{
//   	Function: function_,
//
//   	// the properties below are optional
//   	AuthType: awscdk.Aws_lambda.FunctionUrlAuthType_AWS_IAM,
//   	Cors: &FunctionUrlCorsOptions{
//   		AllowCredentials: jsii.Boolean(false),
//   		AllowedHeaders: []*string{
//   			jsii.String("allowedHeaders"),
//   		},
//   		AllowedMethods: []httpMethod{
//   			awscdk.*Aws_lambda.*httpMethod_GET,
//   		},
//   		AllowedOrigins: []*string{
//   			jsii.String("allowedOrigins"),
//   		},
//   		ExposedHeaders: []*string{
//   			jsii.String("exposedHeaders"),
//   		},
//   		MaxAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	},
//   }
//
type FunctionUrlProps struct {
	// The type of authentication that your function URL uses.
	AuthType FunctionUrlAuthType `field:"optional" json:"authType" yaml:"authType"`
	// The cross-origin resource sharing (CORS) settings for your function URL.
	Cors *FunctionUrlCorsOptions `field:"optional" json:"cors" yaml:"cors"`
	// The function to which this url refers.
	//
	// It can also be an `Alias` but not a `Version`.
	Function IFunction `field:"required" json:"function" yaml:"function"`
}

