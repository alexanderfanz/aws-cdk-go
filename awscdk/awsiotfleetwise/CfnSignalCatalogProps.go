package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSignalCatalog`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSignalCatalogProps := &CfnSignalCatalogProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NodeCounts: &NodeCountsProperty{
//   		TotalActuators: jsii.Number(123),
//   		TotalAttributes: jsii.Number(123),
//   		TotalBranches: jsii.Number(123),
//   		TotalNodes: jsii.Number(123),
//   		TotalSensors: jsii.Number(123),
//   	},
//   	Nodes: []interface{}{
//   		&NodeProperty{
//   			Actuator: &ActuatorProperty{
//   				DataType: jsii.String("dataType"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				Description: jsii.String("description"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Attribute: &AttributeProperty{
//   				DataType: jsii.String("dataType"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Branch: &BranchProperty{
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   			Sensor: &SensorProperty{
//   				DataType: jsii.String("dataType"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				Description: jsii.String("description"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSignalCatalogProps struct {
	// (Optional) A brief description of the signal catalog.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// (Optional) The name of the signal catalog.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// (Optional) Information about the number of nodes and node types in a vehicle network.
	NodeCounts interface{} `field:"optional" json:"nodeCounts" yaml:"nodeCounts"`
	// (Optional) A list of information about nodes, which are a general abstraction of signals.
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
	// (Optional) Metadata that can be used to manage the signal catalog.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

