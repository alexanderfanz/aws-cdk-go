package awsemr


// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
//
// > The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR releases 5.12.1 and later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onDemandProvisioningSpecificationProperty := &OnDemandProvisioningSpecificationProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandprovisioningspecification.html
//
type CfnCluster_OnDemandProvisioningSpecificationProperty struct {
	// Specifies the strategy to use in launching On-Demand instance fleets.
	//
	// Currently, the only option is `lowest-price` (the default), which launches the lowest price first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandprovisioningspecification.html#cfn-emr-cluster-ondemandprovisioningspecification-allocationstrategy
	//
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
}

