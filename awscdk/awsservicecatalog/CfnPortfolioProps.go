package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPortfolio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioProps := &CfnPortfolioProps{
//   	DisplayName: jsii.String("displayName"),
//   	ProviderName: jsii.String("providerName"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPortfolioProps struct {
	// The name to use for display purposes.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the portfolio provider.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the portfolio.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
