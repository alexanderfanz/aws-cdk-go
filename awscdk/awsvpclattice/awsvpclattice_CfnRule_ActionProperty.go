package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	forward: &forwardProperty{
//   		targetGroups: []interface{}{
//   			&weightedTargetGroupProperty{
//   				targetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   				// the properties below are optional
//   				weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnRule_ActionProperty struct {
	// `CfnRule.ActionProperty.Forward`.
	Forward interface{} `field:"required" json:"forward" yaml:"forward"`
}

