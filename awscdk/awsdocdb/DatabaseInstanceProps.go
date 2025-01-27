package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Construction properties for a DatabaseInstanceNew.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var databaseCluster databaseCluster
//   var instanceType instanceType
//
//   databaseInstanceProps := &DatabaseInstanceProps{
//   	Cluster: databaseCluster,
//   	InstanceType: instanceType,
//
//   	// the properties below are optional
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	DbInstanceName: jsii.String("dbInstanceName"),
//   	EnablePerformanceInsights: jsii.Boolean(false),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   }
//
type DatabaseInstanceProps struct {
	// The DocumentDB database cluster the instance should launch into.
	Cluster IDatabaseCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the compute and memory capacity classes.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Default: true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Default: - no preference.
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Default: - a CloudFormation generated name.
	//
	DbInstanceName *string `field:"optional" json:"dbInstanceName" yaml:"dbInstanceName"`
	// A value that indicates whether to enable Performance Insights for the DB Instance.
	// Default: - false.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window.
	// Default: - a 30-minute window selected at random from an 8-hour block of
	// time for each AWS Region, occurring on a random day of the week. To see
	// the time blocks available, see https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-maintain.html#maintenance-window
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Default: RemovalPolicy.Retain
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

