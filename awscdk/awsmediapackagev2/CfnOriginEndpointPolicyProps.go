package awsmediapackagev2


// Properties for defining a `CfnOriginEndpointPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnOriginEndpointPolicyProps := &CfnOriginEndpointPolicyProps{
//   	Policy: policy,
//
//   	// the properties below are optional
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html
//
type CfnOriginEndpointPolicyProps struct {
	// The policy associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The name of the channel group associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-channelgroupname
	//
	ChannelGroupName *string `field:"optional" json:"channelGroupName" yaml:"channelGroupName"`
	// The channel name associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The name of the origin endpoint associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-originendpointname
	//
	OriginEndpointName *string `field:"optional" json:"originEndpointName" yaml:"originEndpointName"`
}

