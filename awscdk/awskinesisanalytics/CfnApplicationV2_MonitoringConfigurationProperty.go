package awskinesisanalytics


// Describes configuration parameters for Amazon CloudWatch logging for a Java-based Kinesis Data Analytics application.
//
// For more information about CloudWatch logging, see [Monitoring](https://docs.aws.amazon.com/kinesisanalytics/latest/java/monitoring-overview) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	ConfigurationType: jsii.String("configurationType"),
//
//   	// the properties below are optional
//   	LogLevel: jsii.String("logLevel"),
//   	MetricsLevel: jsii.String("metricsLevel"),
//   }
//
type CfnApplicationV2_MonitoringConfigurationProperty struct {
	// Describes whether to use the default CloudWatch logging configuration for an application.
	//
	// You must set this property to `CUSTOM` in order to set the `LogLevel` or `MetricsLevel` parameters.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes the verbosity of the CloudWatch Logs for an application.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Describes the granularity of the CloudWatch Logs for an application.
	//
	// The `Parallelism` level is not recommended for applications with a Parallelism over 64 due to excessive costs.
	MetricsLevel *string `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
}
