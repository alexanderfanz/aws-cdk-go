package awsservicecatalog


// Properties for defining a `CfnLaunchTemplateConstraint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchTemplateConstraintProps := &cfnLaunchTemplateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	rules: jsii.String("rules"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnLaunchTemplateConstraintProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The constraint rules.
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
