package awsmedialive


// Configuration information for an output.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputDestinationProperty := &OutputDestinationProperty{
//   	Id: jsii.String("id"),
//   	MediaPackageSettings: []interface{}{
//   		&MediaPackageOutputDestinationSettingsProperty{
//   			ChannelId: jsii.String("channelId"),
//   		},
//   	},
//   	MultiplexSettings: &MultiplexProgramChannelDestinationSettingsProperty{
//   		MultiplexId: jsii.String("multiplexId"),
//   		ProgramName: jsii.String("programName"),
//   	},
//   	Settings: []interface{}{
//   		&OutputDestinationSettingsProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			StreamName: jsii.String("streamName"),
//   			Url: jsii.String("url"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   }
//
type CfnChannel_OutputDestinationProperty struct {
	// The ID for this destination.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The destination settings for a MediaPackage output.
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// Destination settings for a Multiplex output;
	//
	// one destination for both encoders.
	MultiplexSettings interface{} `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The destination settings for an output.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

