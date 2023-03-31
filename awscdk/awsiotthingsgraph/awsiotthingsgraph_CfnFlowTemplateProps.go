package awsiotthingsgraph


// Properties for defining a `CfnFlowTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowTemplateProps := &cfnFlowTemplateProps{
//   	definition: &definitionDocumentProperty{
//   		language: jsii.String("language"),
//   		text: jsii.String("text"),
//   	},
//
//   	// the properties below are optional
//   	compatibleNamespaceVersion: jsii.Number(123),
//   }
//
type CfnFlowTemplateProps struct {
	// `AWS::IoTThingsGraph::FlowTemplate.Definition`.
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// `AWS::IoTThingsGraph::FlowTemplate.CompatibleNamespaceVersion`.
	CompatibleNamespaceVersion *float64 `field:"optional" json:"compatibleNamespaceVersion" yaml:"compatibleNamespaceVersion"`
}

