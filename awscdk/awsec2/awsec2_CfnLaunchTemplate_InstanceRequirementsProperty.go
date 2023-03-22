package awsec2


// The attributes for the instance types.
//
// When you specify instance attributes, Amazon EC2 will identify instance types with these attributes.
//
// When you specify multiple attributes, you get instance types that satisfy all of the specified attributes. If you specify multiple values for an attribute, you get instance types that satisfy any of the specified values.
//
// To limit the list of instance types from which Amazon EC2 can identify matching instance types, you can use one of the following parameters, but not both in the same request:
//
// - `AllowedInstanceTypes` - The instance types to include in the list. All other instance types are ignored, even if they match your specified attributes.
// - `ExcludedInstanceTypes` - The instance types to exclude from the list, even if they match your specified attributes.
//
// > You must specify `VCpuCount` and `MemoryMiB` . All other attributes are optional. Any unspecified optional attribute is set to its default.
//
// For more information, see [Attribute-based instance type selection for EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html) , [Attribute-based instance type selection for Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html) , and [Spot placement score](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceRequirementsProperty := &instanceRequirementsProperty{
//   	acceleratorCount: &acceleratorCountProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	acceleratorManufacturers: []*string{
//   		jsii.String("acceleratorManufacturers"),
//   	},
//   	acceleratorNames: []*string{
//   		jsii.String("acceleratorNames"),
//   	},
//   	acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	acceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	allowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
//   	},
//   	bareMetal: jsii.String("bareMetal"),
//   	baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	burstablePerformance: jsii.String("burstablePerformance"),
//   	cpuManufacturers: []*string{
//   		jsii.String("cpuManufacturers"),
//   	},
//   	excludedInstanceTypes: []*string{
//   		jsii.String("excludedInstanceTypes"),
//   	},
//   	instanceGenerations: []*string{
//   		jsii.String("instanceGenerations"),
//   	},
//   	localStorage: jsii.String("localStorage"),
//   	localStorageTypes: []*string{
//   		jsii.String("localStorageTypes"),
//   	},
//   	memoryGiBPerVCpu: &memoryGiBPerVCpuProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	memoryMiB: &memoryMiBProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	networkBandwidthGbps: &networkBandwidthGbpsProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	networkInterfaceCount: &networkInterfaceCountProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	requireHibernateSupport: jsii.Boolean(false),
//   	spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	totalLocalStorageGb: &totalLocalStorageGBProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	vCpuCount: &vCpuCountProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   }
//
type CfnLaunchTemplate_InstanceRequirementsProperty struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on an instance.
	//
	// To exclude accelerator-enabled instance types, set `Max` to `0` .
	//
	// Default: No minimum or maximum limits.
	AcceleratorCount interface{} `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// - For instance types with NVIDIA devices, specify `nvidia` .
	// - For instance types with AMD devices, specify `amd` .
	// - For instance types with AWS devices, specify `amazon-web-services` .
	// - For instance types with Xilinx devices, specify `xilinx` .
	//
	// Default: Any manufacturer.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// The accelerators that must be on the instance type.
	//
	// - For instance types with NVIDIA A100 GPUs, specify `a100` .
	// - For instance types with NVIDIA V100 GPUs, specify `v100` .
	// - For instance types with NVIDIA K80 GPUs, specify `k80` .
	// - For instance types with NVIDIA T4 GPUs, specify `t4` .
	// - For instance types with NVIDIA M60 GPUs, specify `m60` .
	// - For instance types with AMD Radeon Pro V520 GPUs, specify `radeon-pro-v520` .
	// - For instance types with Xilinx VU9P FPGAs, specify `vu9p` .
	// - For instance types with AWS Inferentia chips, specify `inferentia` .
	// - For instance types with NVIDIA GRID K520 GPUs, specify `k520` .
	//
	// Default: Any accelerator.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The minimum and maximum amount of total accelerator memory, in MiB.
	//
	// Default: No minimum or maximum limits.
	AcceleratorTotalMemoryMiB interface{} `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// The accelerator types that must be on the instance type.
	//
	// - For instance types with GPU accelerators, specify `gpu` .
	// - For instance types with FPGA accelerators, specify `fpga` .
	// - For instance types with inference accelerators, specify `inference` .
	//
	// Default: Any accelerator type.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The instance types to apply your specified attributes against.
	//
	// All other instance types are ignored, even if they match your specified attributes.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ), to allow an instance type, size, or generation. The following are examples: `m5.8xlarge` , `c5*.*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` ,Amazon EC2 will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , Amazon EC2 will allow all the M5a instance types, but not the M5n instance types.
	//
	// > If you specify `AllowedInstanceTypes` , you can't specify `ExcludedInstanceTypes` .
	//
	// Default: All instance types.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types must be included, excluded, or required.
	//
	// - To include bare metal instance types, specify `included` .
	// - To require only bare metal instance types, specify `required` .
	// - To exclude bare metal instance types, specify `excluded` .
	//
	// Default: `excluded`.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
	//
	// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide* .
	//
	// Default: No minimum or maximum limits.
	BaselineEbsBandwidthMbps interface{} `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Indicates whether burstable performance T instance types are included, excluded, or required.
	//
	// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) .
	//
	// - To include burstable performance instance types, specify `included` .
	// - To require only burstable performance instance types, specify `required` .
	// - To exclude burstable performance instance types, specify `excluded` .
	//
	// Default: `excluded`.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// The CPU manufacturers to include.
	//
	// - For instance types with Intel CPUs, specify `intel` .
	// - For instance types with AMD CPUs, specify `amd` .
	// - For instance types with AWS CPUs, specify `amazon-web-services` .
	//
	// > Don't confuse the CPU manufacturer with the CPU architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	//
	// Default: Any manufacturer.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// The instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ), to exclude an instance type, size, or generation. The following are examples: `m5.8xlarge` , `c5*.*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` ,Amazon EC2 will exclude the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , Amazon EC2 will exclude all the M5a instance types, but not the M5n instance types.
	//
	// > If you specify `ExcludedInstanceTypes` , you can't specify `AllowedInstanceTypes` .
	//
	// Default: No excluded instance types.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// The current generation instance types are recommended for use. Current generation instance types are typically the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	//
	// For current generation instance types, specify `current` .
	//
	// For previous generation instance types, specify `previous` .
	//
	// Default: Current and previous generation instance types.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide* .
	//
	// - To include instance types with instance store volumes, specify `included` .
	// - To require only instance types with instance store volumes, specify `required` .
	// - To exclude instance types with instance store volumes, specify `excluded` .
	//
	// Default: `included`.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// The type of local storage that is required.
	//
	// - For instance types with hard disk drive (HDD) storage, specify `hdd` .
	// - For instance types with solid state drive (SSD) storage, specify `ssd` .
	//
	// Default: `hdd` and `ssd`.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// The minimum and maximum amount of memory per vCPU, in GiB.
	//
	// Default: No minimum or maximum limits.
	MemoryGiBPerVCpu interface{} `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// The minimum and maximum amount of memory, in MiB.
	MemoryMiB interface{} `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
	//
	// Default: No minimum or maximum limits.
	NetworkBandwidthGbps interface{} `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// The minimum and maximum number of network interfaces.
	//
	// Default: No minimum or maximum limits.
	NetworkInterfaceCount interface{} `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// The price protection threshold for On-Demand Instances.
	//
	// This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage above the least expensive current generation M, C, or R instance type with your specified attributes. When Amazon EC2 selects instance types with your attributes, it excludes instance types priced above your threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//
	// To turn off price protection, specify a high value, such as `999999` .
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html) .
	//
	// > If you set `TargetCapacityUnitType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//
	// Default: `20`.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must support hibernation for On-Demand Instances.
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) .
	//
	// Default: `false`.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// The price protection threshold for Spot Instances.
	//
	// This is the maximum you’ll pay for a Spot Instance, expressed as a percentage above the least expensive current generation M, C, or R instance type with your specified attributes. When Amazon EC2 selects instance types with your attributes, it excludes instance types priced above your threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//
	// To turn off price protection, specify a high value, such as `999999` .
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html) .
	//
	// > If you set `TargetCapacityUnitType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//
	// Default: `100`.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum amount of total local storage, in GB.
	//
	// Default: No minimum or maximum limits.
	TotalLocalStorageGb interface{} `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs.
	VCpuCount interface{} `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

