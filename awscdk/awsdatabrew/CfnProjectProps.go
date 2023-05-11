package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &CfnProjectProps{
//   	DatasetName: jsii.String("datasetName"),
//   	Name: jsii.String("name"),
//   	RecipeName: jsii.String("recipeName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Sample: &SampleProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Size: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The dataset that the project is to act upon.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The unique name of a project.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of a recipe that will be developed during a project session.
	RecipeName *string `field:"required" json:"recipeName" yaml:"recipeName"`
	// The Amazon Resource Name (ARN) of the role that will be assumed for this project.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The sample size and sampling type to apply to the data.
	//
	// If this parameter isn't specified, then the sample consists of the first 500 rows from the dataset.
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// Metadata tags that have been applied to the project.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
