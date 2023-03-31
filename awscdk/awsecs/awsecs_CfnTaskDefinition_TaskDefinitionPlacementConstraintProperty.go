package awsecs


// An object representing a constraint on task placement in the task definition.
//
// For more information, see [Task placement constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > Task placement constraints aren't supported for tasks run on AWS Fargate .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskDefinitionPlacementConstraintProperty := &taskDefinitionPlacementConstraintProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	expression: jsii.String("expression"),
//   }
//
type CfnTaskDefinition_TaskDefinitionPlacementConstraintProperty struct {
	// The type of constraint.
	//
	// The `MemberOf` constraint restricts selection to be from a group of valid candidates.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A cluster query language expression to apply to the constraint.
	//
	// For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

