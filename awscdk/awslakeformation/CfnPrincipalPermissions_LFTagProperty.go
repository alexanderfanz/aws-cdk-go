package awslakeformation


// The LF-tag key and values attached to a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagProperty := &LFTagProperty{
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagProperty struct {
	// The key-name for the LF-tag.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}
