package awsservicecatalog


// Properties for provisoning rule constraint.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   portfolio.constrainCloudFormationParameters(product, &CloudFormationRuleConstraintOptions{
//   	Rule: &TemplateRule{
//   		RuleName: jsii.String("testInstanceType"),
//   		Condition: cdk.Fn_ConditionEquals(cdk.Fn_Ref(jsii.String("Environment")), jsii.String("test")),
//   		Assertions: []templateRuleAssertion{
//   			&templateRuleAssertion{
//   				Assert: cdk.Fn_ConditionContains([]*string{
//   					jsii.String("t2.micro"),
//   					jsii.String("t2.small"),
//   				}, cdk.Fn_*Ref(jsii.String("InstanceType"))),
//   				Description: jsii.String("For test environment, the instance type should be small"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type CloudFormationRuleConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// The rule with condition and assertions to apply to template.
	// Experimental.
	Rule *TemplateRule `field:"required" json:"rule" yaml:"rule"`
}

