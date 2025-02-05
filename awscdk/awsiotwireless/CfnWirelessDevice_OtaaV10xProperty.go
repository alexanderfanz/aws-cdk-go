package awsiotwireless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otaaV10xProperty := &OtaaV10xProperty{
//   	AppEui: jsii.String("appEui"),
//   	AppKey: jsii.String("appKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html
//
type CfnWirelessDevice_OtaaV10xProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html#cfn-iotwireless-wirelessdevice-otaav10x-appeui
	//
	AppEui *string `field:"required" json:"appEui" yaml:"appEui"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html#cfn-iotwireless-wirelessdevice-otaav10x-appkey
	//
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
}

