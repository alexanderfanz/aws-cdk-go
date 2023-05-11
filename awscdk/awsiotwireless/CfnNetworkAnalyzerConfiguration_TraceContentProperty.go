package awsiotwireless


// Trace content for your wireless gateway and wireless device resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   traceContentProperty := &TraceContentProperty{
//   	LogLevel: jsii.String("logLevel"),
//   	WirelessDeviceFrameInfo: jsii.String("wirelessDeviceFrameInfo"),
//   }
//
type CfnNetworkAnalyzerConfiguration_TraceContentProperty struct {
	// The log level for a log message.
	//
	// The log levels can be disabled, or set to `ERROR` to display less verbose logs containing only error information, or to `INFO` for more detailed logs.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// `FrameInfo` of your wireless device resources for the trace content.
	//
	// Use FrameInfo to debug the communication between your LoRaWAN end devices and the network server.
	WirelessDeviceFrameInfo *string `field:"optional" json:"wirelessDeviceFrameInfo" yaml:"wirelessDeviceFrameInfo"`
}
