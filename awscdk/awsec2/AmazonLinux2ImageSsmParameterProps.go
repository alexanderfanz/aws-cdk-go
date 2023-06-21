package awsec2


// Properties specific to amzn2 images.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   // Amazon Linux 2
//   // Amazon Linux 2
//   ec2.NewInstance(this, jsii.String("Instance2"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   })
//
//   // Amazon Linux 2 with kernel 5.x
//   // Amazon Linux 2 with kernel 5.x
//   ec2.NewInstance(this, jsii.String("Instance3"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux2(&AmazonLinux2ImageSsmParameterProps{
//   		Kernel: ec2.AmazonLinux2Kernel_KERNEL_5_10(),
//   	}),
//   })
//
//   // AWS Linux 2022
//   // AWS Linux 2022
//   ec2.NewInstance(this, jsii.String("Instance4"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2022(),
//   })
//
//   // Graviton 3 Processor
//   // Graviton 3 Processor
//   ec2.NewInstance(this, jsii.String("Instance5"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.*instanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux2022(&AmazonLinux2022ImageSsmParameterProps{
//   		CpuType: ec2.AmazonLinuxCpuType_ARM_64,
//   	}),
//   })
//
type AmazonLinux2ImageSsmParameterProps struct {
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// Initial user data.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// CPU Type.
	CpuType AmazonLinuxCpuType `field:"optional" json:"cpuType" yaml:"cpuType"`
	// What edition of Amazon Linux to use.
	Edition AmazonLinuxEdition `field:"optional" json:"edition" yaml:"edition"`
	// What kernel version of Amazon Linux to use.
	Kernel AmazonLinux2Kernel `field:"optional" json:"kernel" yaml:"kernel"`
	// What storage backed image to use.
	Storage AmazonLinuxStorage `field:"optional" json:"storage" yaml:"storage"`
	// Virtualization type.
	Virtualization AmazonLinuxVirt `field:"optional" json:"virtualization" yaml:"virtualization"`
}

