package awsimagebuilder


// The distribution configuration distribution defines the settings for a specific Region in the Distribution Configuration.
//
// You must specify whether the distribution is for an AMI or a container image. To do so, include exactly one of the following data types for your distribution:
//
// - amiDistributionConfiguration
// - containerDistributionConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   distributionProperty := &DistributionProperty{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	AmiDistributionConfiguration: amiDistributionConfiguration,
//   	ContainerDistributionConfiguration: containerDistributionConfiguration,
//   	FastLaunchConfigurations: []interface{}{
//   		&FastLaunchConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Enabled: jsii.Boolean(false),
//   			LaunchTemplate: &FastLaunchLaunchTemplateSpecificationProperty{
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   				LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   			},
//   			MaxParallelLaunches: jsii.Number(123),
//   			SnapshotConfiguration: &FastLaunchSnapshotConfigurationProperty{
//   				TargetResourceCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LaunchTemplateConfigurations: []interface{}{
//   		&LaunchTemplateConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			SetDefaultVersion: jsii.Boolean(false),
//   		},
//   	},
//   	LicenseConfigurationArns: []*string{
//   		jsii.String("licenseConfigurationArns"),
//   	},
//   }
//
type CfnDistributionConfiguration_DistributionProperty struct {
	// The target Region for the Distribution Configuration.
	//
	// For example, `eu-west-1` .
	Region *string `field:"required" json:"region" yaml:"region"`
	// The specific AMI settings, such as launch permissions and AMI tags.
	//
	// For details, see example schema below.
	AmiDistributionConfiguration interface{} `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// Container distribution settings for encryption, licensing, and sharing in a specific Region.
	//
	// For details, see example schema below.
	ContainerDistributionConfiguration interface{} `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// `CfnDistributionConfiguration.DistributionProperty.FastLaunchConfigurations`.
	FastLaunchConfigurations interface{} `field:"optional" json:"fastLaunchConfigurations" yaml:"fastLaunchConfigurations"`
	// A group of launchTemplateConfiguration settings that apply to image distribution for specified accounts.
	LaunchTemplateConfigurations interface{} `field:"optional" json:"launchTemplateConfigurations" yaml:"launchTemplateConfigurations"`
	// The License Manager Configuration to associate with the AMI in the specified Region.
	//
	// For more information, see the [LicenseConfiguration API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_LicenseConfiguration.html) .
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
}
