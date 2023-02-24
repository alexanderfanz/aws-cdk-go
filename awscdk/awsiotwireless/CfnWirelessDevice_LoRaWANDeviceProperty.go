package awsiotwireless


// LoRaWAN object for create functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANDeviceProperty := &LoRaWANDeviceProperty{
//   	AbpV10X: &AbpV10xProperty{
//   		DevAddr: jsii.String("devAddr"),
//   		SessionKeys: &SessionKeysAbpV10xProperty{
//   			AppSKey: jsii.String("appSKey"),
//   			NwkSKey: jsii.String("nwkSKey"),
//   		},
//   	},
//   	AbpV11: &AbpV11Property{
//   		DevAddr: jsii.String("devAddr"),
//   		SessionKeys: &SessionKeysAbpV11Property{
//   			AppSKey: jsii.String("appSKey"),
//   			FNwkSIntKey: jsii.String("fNwkSIntKey"),
//   			NwkSEncKey: jsii.String("nwkSEncKey"),
//   			SNwkSIntKey: jsii.String("sNwkSIntKey"),
//   		},
//   	},
//   	DevEui: jsii.String("devEui"),
//   	DeviceProfileId: jsii.String("deviceProfileId"),
//   	OtaaV10X: &OtaaV10xProperty{
//   		AppEui: jsii.String("appEui"),
//   		AppKey: jsii.String("appKey"),
//   	},
//   	OtaaV11: &OtaaV11Property{
//   		AppKey: jsii.String("appKey"),
//   		JoinEui: jsii.String("joinEui"),
//   		NwkKey: jsii.String("nwkKey"),
//   	},
//   	ServiceProfileId: jsii.String("serviceProfileId"),
//   }
//
type CfnWirelessDevice_LoRaWANDeviceProperty struct {
	// LoRaWAN object for create APIs.
	AbpV10X interface{} `field:"optional" json:"abpV10X" yaml:"abpV10X"`
	// ABP device object for create APIs for v1.1.
	AbpV11 interface{} `field:"optional" json:"abpV11" yaml:"abpV11"`
	// The DevEUI value.
	DevEui *string `field:"optional" json:"devEui" yaml:"devEui"`
	// The ID of the device profile for the new wireless device.
	DeviceProfileId *string `field:"optional" json:"deviceProfileId" yaml:"deviceProfileId"`
	// OTAA device object for create APIs for v1.0.x.
	OtaaV10X interface{} `field:"optional" json:"otaaV10X" yaml:"otaaV10X"`
	// OTAA device object for v1.1 for create APIs.
	OtaaV11 interface{} `field:"optional" json:"otaaV11" yaml:"otaaV11"`
	// The ID of the service profile.
	ServiceProfileId *string `field:"optional" json:"serviceProfileId" yaml:"serviceProfileId"`
}

