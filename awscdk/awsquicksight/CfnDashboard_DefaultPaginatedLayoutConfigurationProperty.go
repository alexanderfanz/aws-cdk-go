package awsquicksight


// The options that determine the default settings for a paginated layout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultPaginatedLayoutConfigurationProperty := &DefaultPaginatedLayoutConfigurationProperty{
//   	SectionBased: &DefaultSectionBasedLayoutConfigurationProperty{
//   		CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   			PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   				PaperMargin: &SpacingProperty{
//   					Bottom: jsii.String("bottom"),
//   					Left: jsii.String("left"),
//   					Right: jsii.String("right"),
//   					Top: jsii.String("top"),
//   				},
//   				PaperOrientation: jsii.String("paperOrientation"),
//   				PaperSize: jsii.String("paperSize"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_DefaultPaginatedLayoutConfigurationProperty struct {
	// The options that determine the default settings for a section-based layout configuration.
	SectionBased interface{} `field:"optional" json:"sectionBased" yaml:"sectionBased"`
}
