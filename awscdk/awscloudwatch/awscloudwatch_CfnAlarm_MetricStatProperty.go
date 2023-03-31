package awscloudwatch


// This structure defines the metric to be returned, along with the statistics, period, and units.
//
// `MetricStat` is a property of the [MetricDataQuery](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStatProperty := &metricStatProperty{
//   	metric: &metricProperty{
//   		dimensions: []interface{}{
//   			&dimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   	},
//   	period: jsii.Number(123),
//   	stat: jsii.String("stat"),
//
//   	// the properties below are optional
//   	unit: jsii.String("unit"),
//   }
//
type CfnAlarm_MetricStatProperty struct {
	// The metric to return, including the metric name, namespace, and dimensions.
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a `PutMetricData` call that includes a `StorageResolution` of 1 second.
	//
	// If the `StartTime` parameter specifies a time stamp that is greater than 3 hours ago, you must specify the period as follows or no data points in that time range is returned:
	//
	// - Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).
	// - Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).
	// - Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide* .
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// Valid values are: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

