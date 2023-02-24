package awsamplifyuibuilder


// The `Predicate` property specifies information for generating Amplify DataStore queries.
//
// Use `Predicate` to retrieve a subset of the data in a collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   predicateProperty := &predicateProperty{
//   	And: []interface{}{
//   		&predicateProperty{
//   			And: []interface{}{
//   				predicateProperty_,
//   			},
//   			Field: jsii.String("field"),
//   			Operand: jsii.String("operand"),
//   			Operator: jsii.String("operator"),
//   			Or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   	Field: jsii.String("field"),
//   	Operand: jsii.String("operand"),
//   	Operator: jsii.String("operator"),
//   	Or: []interface{}{
//   		&predicateProperty{
//   			And: []interface{}{
//   				predicateProperty_,
//   			},
//   			Field: jsii.String("field"),
//   			Operand: jsii.String("operand"),
//   			Operator: jsii.String("operator"),
//   			Or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   }
//
type CfnComponent_PredicateProperty struct {
	// A list of predicates to combine logically.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// The field to query.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value to use when performing the evaluation.
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The operator to use to perform the evaluation.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// A list of predicates to combine logically.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
}

