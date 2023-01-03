package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   challengeActionProperty := &challengeActionProperty{
//   	customRequestHandling: &customRequestHandlingProperty{
//   		insertHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_ChallengeActionProperty struct {
	// `CfnWebACL.ChallengeActionProperty.CustomRequestHandling`.
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

