package awsec2


// Options to add a flow log to a VPC.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   vpc.addFlowLog(jsii.String("FlowLogS3"), &FlowLogOptions{
//   	Destination: ec2.FlowLogDestination_ToS3(),
//   })
//
//   vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &FlowLogOptions{
//   	TrafficType: ec2.FlowLogTrafficType_REJECT,
//   })
//
// Experimental.
type FlowLogOptions struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	// Experimental.
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	// Experimental.
	TrafficType FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
}

