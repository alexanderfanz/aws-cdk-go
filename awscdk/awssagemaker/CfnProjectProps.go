package awssagemaker

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
//   var serviceCatalogProvisioningDetails interface{}
//
//   cfnProjectProps := &CfnProjectProps{
//   	ProjectName: jsii.String("projectName"),
//   	ServiceCatalogProvisioningDetails: serviceCatalogProvisioningDetails,
//
//   	// the properties below are optional
//   	ProjectDescription: jsii.String("projectDescription"),
//   	ServiceCatalogProvisionedProductDetails: &ServiceCatalogProvisionedProductDetailsProperty{
//   		ProvisionedProductId: jsii.String("provisionedProductId"),
//   		ProvisionedProductStatusMessage: jsii.String("provisionedProductStatusMessage"),
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
	// The name of the project.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The product ID and provisioning artifact ID to provision a service catalog.
	//
	// For information, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
	ServiceCatalogProvisioningDetails interface{} `field:"required" json:"serviceCatalogProvisioningDetails" yaml:"serviceCatalogProvisioningDetails"`
	// The description of the project.
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// `AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails`.
	ServiceCatalogProvisionedProductDetails interface{} `field:"optional" json:"serviceCatalogProvisionedProductDetails" yaml:"serviceCatalogProvisionedProductDetails"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

