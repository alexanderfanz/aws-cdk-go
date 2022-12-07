package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPricingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPricingRuleProps := &cfnPricingRuleProps{
//   	modifierPercentage: jsii.Number(123),
//   	name: jsii.String("name"),
//   	scope: jsii.String("scope"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	billingEntity: jsii.String("billingEntity"),
//   	description: jsii.String("description"),
//   	service: jsii.String("service"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPricingRuleProps struct {
	// A percentage modifier applied on the public pricing rates.
	ModifierPercentage *float64 `field:"required" json:"modifierPercentage" yaml:"modifierPercentage"`
	// The name of a pricing rule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scope of pricing rule that indicates if it is globally applicable, or if it is service-specific.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The type of pricing rule.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The seller of services provided by AWS , their affiliates, or third-party providers selling services via AWS Marketplace .
	BillingEntity *string `field:"optional" json:"billingEntity" yaml:"billingEntity"`
	// The pricing rule description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If the `Scope` attribute is `SERVICE` , this attribute indicates which service the `PricingRule` is applicable for.
	Service *string `field:"optional" json:"service" yaml:"service"`
	// A map that contains tag keys and tag values that are attached to a pricing rule.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

