package awsopensearchservice


// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
//
// For more information, see [Launching your Amazon OpenSearch Service domains using a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCOptionsProperty := &vPCOptionsProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDomain_VPCOptionsProperty struct {
	// The list of security group IDs that are associated with the VPC endpoints for the domain.
	//
	// If you don't provide a security group ID, OpenSearch Service uses the default security group for the VPC. To learn more, see [Security groups for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon VPC User Guide* .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Provide one subnet ID for each Availability Zone that your domain uses.
	//
	// For example, you must specify three subnet IDs for a three Availability Zone domain. To learn more, see [VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the *Amazon VPC User Guide* .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

