package awsquicksight


// Determines the border options for a table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalTableBorderOptionsProperty := &GlobalTableBorderOptionsProperty{
//   	SideSpecificBorder: &TableSideBorderOptionsProperty{
//   		Bottom: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		InnerHorizontal: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		InnerVertical: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Left: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Right: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Top: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   	},
//   	UniformBorder: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_GlobalTableBorderOptionsProperty struct {
	// Determines the options for side specific border.
	SideSpecificBorder interface{} `field:"optional" json:"sideSpecificBorder" yaml:"sideSpecificBorder"`
	// Determines the options for uniform border.
	UniformBorder interface{} `field:"optional" json:"uniformBorder" yaml:"uniformBorder"`
}
