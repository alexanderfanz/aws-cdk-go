package awselasticloadbalancingv2


// Options for ApplicationListener lookup.
//
// Example:
//   listener := elbv2.applicationListener.fromLookup(this, jsii.String("ALBListener"), &applicationListenerLookupOptions{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   	listenerProtocol: elbv2.applicationProtocol_HTTPS,
//   	listenerPort: jsii.Number(443),
//   })
//
// Experimental.
type ApplicationListenerLookupOptions struct {
	// Filter listeners by listener port.
	// Experimental.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	// Experimental.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// ARN of the listener to look up.
	// Experimental.
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener protocol.
	// Experimental.
	ListenerProtocol ApplicationProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

