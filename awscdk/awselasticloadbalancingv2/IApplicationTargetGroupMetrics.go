package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Contains all metrics for a Target Group of a Application Load Balancer.
type IApplicationTargetGroupMetrics interface {
	// Return the given named metric for this Network Target Group.
	Custom(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of healthy hosts in the target group.
	HealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in this target group.
	//
	// This does not include any response codes generated by the load balancer.
	HttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of IPv6 requests received by the target group.
	Ipv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of requests processed over IPv4 and IPv6.
	//
	// This count includes only the requests with a response generated by a target of the load balancer.
	RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of requests received by each target in a target group.
	//
	// The only valid statistic is Sum. Note that this represents the average not the sum.
	RequestCountPerTarget(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were not successfully established between the load balancer and target.
	TargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
	TargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
	//
	// Possible causes include a mismatch of ciphers or protocols.
	TargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of unhealthy hosts in the target group.
	UnhealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy for IApplicationTargetGroupMetrics
type jsiiProxy_IApplicationTargetGroupMetrics struct {
	_ byte // padding
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) Custom(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateCustomParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"custom",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) HealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateHealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"healthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) HttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateHttpCodeTargetParameters(code, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"httpCodeTarget",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) Ipv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateIpv6RequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"ipv6RequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"requestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) RequestCountPerTarget(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateRequestCountPerTargetParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"requestCountPerTarget",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) TargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateTargetConnectionErrorCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"targetConnectionErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) TargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateTargetResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"targetResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) TargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateTargetTLSNegotiationErrorCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"targetTLSNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationTargetGroupMetrics) UnhealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateUnhealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"unhealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

