package awsguardduty


// Properties for defining a `CfnMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberProps := &CfnMemberProps{
//   	DetectorId: jsii.String("detectorId"),
//   	Email: jsii.String("email"),
//   	MemberId: jsii.String("memberId"),
//
//   	// the properties below are optional
//   	DisableEmailNotification: jsii.Boolean(false),
//   	Message: jsii.String("message"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html
//
type CfnMemberProps struct {
	// The ID of the detector associated with the GuardDuty service to add the member to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-detectorid
	//
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The email address associated with the member account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-email
	//
	Email *string `field:"required" json:"email" yaml:"email"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-memberid
	//
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
	// Specifies whether or not to disable email notification for the member account that you invite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-disableemailnotification
	//
	DisableEmailNotification interface{} `field:"optional" json:"disableEmailNotification" yaml:"disableEmailNotification"`
	// The invitation message that you want to send to the accounts that you're inviting to GuardDuty as members.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// You can use the `Status` property to update the status of the relationship between the member account and its administrator account.
	//
	// Valid values are `Created` and `Invited` when using an `AWS::GuardDuty::Member` resource. If the value for this property is not provided or set to `Created` , a member account is created but not invited. If the value of this property is set to `Invited` , a member account is created and invited.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

