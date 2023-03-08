package awscloudwatch


// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetectorProps := &CfnAnomalyDetectorProps{
//   	Configuration: &ConfigurationProperty{
//   		ExcludedTimeRanges: []interface{}{
//   			&RangeProperty{
//   				EndTime: jsii.String("endTime"),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		MetricTimeZone: jsii.String("metricTimeZone"),
//   	},
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricMathAnomalyDetector: &MetricMathAnomalyDetectorProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Period: jsii.Number(123),
//   					Stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				Period: jsii.Number(123),
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	SingleMetricAnomalyDetector: &SingleMetricAnomalyDetectorProperty{
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Stat: jsii.String("stat"),
//   	},
//   	Stat: jsii.String("stat"),
//   }
//
type CfnAnomalyDetectorProps struct {
	// Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.
	//
	// The configuration can also include the time zone to use for the metric.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The dimensions of the metric associated with the anomaly detection band.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The CloudWatch metric math expression for this anomaly detector.
	MetricMathAnomalyDetector interface{} `field:"optional" json:"metricMathAnomalyDetector" yaml:"metricMathAnomalyDetector"`
	// The name of the metric associated with the anomaly detection band.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric associated with the anomaly detection band.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The CloudWatch metric and statistic for this anomaly detector.
	SingleMetricAnomalyDetector interface{} `field:"optional" json:"singleMetricAnomalyDetector" yaml:"singleMetricAnomalyDetector"`
	// The statistic of the metric associated with the anomaly detection band.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

