package awssagemaker


// The metadata of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointMetadataProperty := &endpointMetadataProperty{
//   	endpointName: jsii.String("endpointName"),
//
//   	// the properties below are optional
//   	endpointConfigName: jsii.String("endpointConfigName"),
//   	endpointStatus: jsii.String("endpointStatus"),
//   }
//
type CfnInferenceExperiment_EndpointMetadataProperty struct {
	// The name of the endpoint.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The name of the endpoint configuration.
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The status of the endpoint.
	//
	// For possible values of the status of an endpoint, see `EndpointSummary$EndpointStatus` .
	EndpointStatus *string `field:"optional" json:"endpointStatus" yaml:"endpointStatus"`
}

