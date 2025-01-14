package awsconnect


// Contains information about the phone configuration settings for a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPhoneConfigProperty := &UserPhoneConfigProperty{
//   	PhoneType: jsii.String("phoneType"),
//
//   	// the properties below are optional
//   	AfterContactWorkTimeLimit: jsii.Number(123),
//   	AutoAccept: jsii.Boolean(false),
//   	DeskPhoneNumber: jsii.String("deskPhoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userphoneconfig.html
//
type CfnUser_UserPhoneConfigProperty struct {
	// The phone type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userphoneconfig.html#cfn-connect-user-userphoneconfig-phonetype
	//
	PhoneType *string `field:"required" json:"phoneType" yaml:"phoneType"`
	// The After Call Work (ACW) timeout setting, in seconds.
	//
	// > When returned by a `SearchUsers` call, `AfterContactWorkTimeLimit` is returned in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userphoneconfig.html#cfn-connect-user-userphoneconfig-aftercontactworktimelimit
	//
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
	// The Auto accept setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userphoneconfig.html#cfn-connect-user-userphoneconfig-autoaccept
	//
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// The phone number for the user's desk phone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userphoneconfig.html#cfn-connect-user-userphoneconfig-deskphonenumber
	//
	DeskPhoneNumber *string `field:"optional" json:"deskPhoneNumber" yaml:"deskPhoneNumber"`
}

