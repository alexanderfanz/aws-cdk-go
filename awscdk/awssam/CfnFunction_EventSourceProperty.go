package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &EventSourceProperty{
//   	Properties: &S3EventProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnFunction_EventSourceProperty struct {
	// `CfnFunction.EventSourceProperty.Properties`.
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// `CfnFunction.EventSourceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

