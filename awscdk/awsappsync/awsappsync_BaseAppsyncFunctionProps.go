package awsappsync


// the base properties for AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mappingTemplate mappingTemplate
//
//   baseAppsyncFunctionProps := &BaseAppsyncFunctionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RequestMappingTemplate: mappingTemplate,
//   	ResponseMappingTemplate: mappingTemplate,
//   }
//
// Experimental.
type BaseAppsyncFunctionProps struct {
	// the name of the AppSync Function.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// the description for this AppSync Function.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// the request mapping template for the AppSync Function.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
}

