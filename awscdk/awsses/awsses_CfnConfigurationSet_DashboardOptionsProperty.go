package awsses


// Settings for your VDM configuration as applicable to the Dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardOptionsProperty := &dashboardOptionsProperty{
//   	engagementMetrics: jsii.String("engagementMetrics"),
//   }
//
type CfnConfigurationSet_DashboardOptionsProperty struct {
	// Specifies the status of your VDM engagement metrics collection. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables engagement metrics for the configuration set.
	// - `DISABLED` – Amazon SES disables engagement metrics for the configuration set.
	EngagementMetrics *string `field:"required" json:"engagementMetrics" yaml:"engagementMetrics"`
}
