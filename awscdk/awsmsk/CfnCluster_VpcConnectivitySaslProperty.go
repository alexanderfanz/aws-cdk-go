package awsmsk


// Not currently supported by AWS CloudFormation .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivitySaslProperty := &VpcConnectivitySaslProperty{
//   	Iam: &VpcConnectivityIamProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Scram: &VpcConnectivityScramProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_VpcConnectivitySaslProperty struct {
	// `CfnCluster.VpcConnectivitySaslProperty.Iam`.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// `CfnCluster.VpcConnectivitySaslProperty.Scram`.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

