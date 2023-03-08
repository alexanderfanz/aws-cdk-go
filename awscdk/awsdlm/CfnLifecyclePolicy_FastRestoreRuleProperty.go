package awsdlm


// *[Snapshot policies only]* Specifies a rule for enabling fast snapshot restore for snapshots created by snapshot policies.
//
// You can enable fast snapshot restore based on either a count or a time interval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastRestoreRuleProperty := &FastRestoreRuleProperty{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_FastRestoreRuleProperty struct {
	// The Availability Zones in which to enable fast snapshot restore.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The number of snapshots to be enabled with fast snapshot restore.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The amount of time to enable fast snapshot restore.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for enabling fast snapshot restore.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

