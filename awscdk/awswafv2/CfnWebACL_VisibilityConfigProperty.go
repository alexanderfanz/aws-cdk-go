package awswafv2


// Defines and enables Amazon CloudWatch metrics and web request sample collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visibilityConfigProperty := &VisibilityConfigProperty{
//   	CloudWatchMetricsEnabled: jsii.Boolean(false),
//   	MetricName: jsii.String("metricName"),
//   	SampledRequestsEnabled: jsii.Boolean(false),
//   }
//
type CfnWebACL_VisibilityConfigProperty struct {
	// A boolean indicating whether the associated resource sends metrics to Amazon CloudWatch.
	//
	// For the list of available metrics, see [AWS WAF Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics) in the *AWS WAF Developer Guide* .
	CloudWatchMetricsEnabled interface{} `field:"required" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// A name of the Amazon CloudWatch metric dimension.
	//
	// The name can contain only the characters: A-Z, a-z, 0-9, - (hyphen), and _ (underscore). The name can be from one to 128 characters long. It can't contain whitespace or metric names that are reserved for AWS WAF , for example `All` and `Default_Action` .
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules.
	//
	// You can view the sampled requests through the AWS WAF console.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}
