package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMulticastGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMulticastGroupProps := &CfnMulticastGroupProps{
//   	LoRaWan: &LoRaWANProperty{
//   		DlClass: jsii.String("dlClass"),
//   		RfRegion: jsii.String("rfRegion"),
//
//   		// the properties below are optional
//   		NumberOfDevicesInGroup: jsii.Number(123),
//   		NumberOfDevicesRequested: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	AssociateWirelessDevice: jsii.String("associateWirelessDevice"),
//   	Description: jsii.String("description"),
//   	DisassociateWirelessDevice: jsii.String("disassociateWirelessDevice"),
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMulticastGroupProps struct {
	// The LoRaWAN information that is to be used with the multicast group.
	LoRaWan interface{} `field:"required" json:"loRaWan" yaml:"loRaWan"`
	// The ID of the wireless device to associate with a multicast group.
	AssociateWirelessDevice *string `field:"optional" json:"associateWirelessDevice" yaml:"associateWirelessDevice"`
	// The description of the multicast group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the wireless device to disassociate from a multicast group.
	DisassociateWirelessDevice *string `field:"optional" json:"disassociateWirelessDevice" yaml:"disassociateWirelessDevice"`
	// The name of the multicast group.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

