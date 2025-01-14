package awsemr


// Properties for defining a `CfnInstanceFleetConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceFleetConfigProps := &CfnInstanceFleetConfigProps{
//   	ClusterId: jsii.String("clusterId"),
//   	InstanceFleetType: jsii.String("instanceFleetType"),
//
//   	// the properties below are optional
//   	InstanceTypeConfigs: []interface{}{
//   		&InstanceTypeConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			BidPrice: jsii.String("bidPrice"),
//   			BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			Configurations: []interface{}{
//   				&configurationProperty{
//   					Classification: jsii.String("classification"),
//   					ConfigurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					Configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			CustomAmiId: jsii.String("customAmiId"),
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []interface{}{
//   					&EbsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							SizeInGb: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//
//   							// the properties below are optional
//   							Iops: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   		OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   		},
//   		SpotSpecification: &SpotProvisioningSpecificationProperty{
//   			TimeoutAction: jsii.String("timeoutAction"),
//   			TimeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			BlockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	TargetOnDemandCapacity: jsii.Number(123),
//   	TargetSpotCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html
//
type CfnInstanceFleetConfigProps struct {
	// The unique identifier of the EMR cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-clusterid
	//
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The node type that the instance fleet hosts.
	//
	// *Allowed Values* : TASK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-instancefleettype
	//
	InstanceFleetType *string `field:"required" json:"instanceFleetType" yaml:"instanceFleetType"`
	// `InstanceTypeConfigs` determine the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-instancetypeconfigs
	//
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-launchspecifications
	//
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-targetondemandcapacity
	//
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancefleetconfig.html#cfn-emr-instancefleetconfig-targetspotcapacity
	//
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

