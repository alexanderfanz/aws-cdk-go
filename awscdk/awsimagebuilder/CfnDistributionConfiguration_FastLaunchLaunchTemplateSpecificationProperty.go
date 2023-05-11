package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchLaunchTemplateSpecificationProperty := &FastLaunchLaunchTemplateSpecificationProperty{
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   }
//
type CfnDistributionConfiguration_FastLaunchLaunchTemplateSpecificationProperty struct {
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateId`.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateName`.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateVersion`.
	LaunchTemplateVersion *string `field:"optional" json:"launchTemplateVersion" yaml:"launchTemplateVersion"`
}
