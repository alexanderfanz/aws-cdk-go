package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a SageMaker Model.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//   import "github.com/aws-samples/dummy/path"
//
//
//   image := sagemaker.ContainerImage_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("Dockerfile"), jsii.String("directory")))
//   modelData := sagemaker.ModelData_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("artifact"), jsii.String("file.tar.gz")))
//
//   model := sagemaker.NewModel(this, jsii.String("PrimaryContainerModel"), &ModelProps{
//   	Containers: []containerDefinition{
//   		&containerDefinition{
//   			Image: image,
//   			ModelData: modelData,
//   		},
//   	},
//   })
//
// Experimental.
type ModelProps struct {
	// Whether to allow the SageMaker Model to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// SageMaker Model to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	// Default: true.
	//
	// Experimental.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Specifies the container definitions for this model, consisting of either a single primary container or an inference pipeline of multiple containers.
	// Default: - none.
	//
	// Experimental.
	Containers *[]*ContainerDefinition `field:"optional" json:"containers" yaml:"containers"`
	// Name of the SageMaker Model.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that ID for the model's
	// name.
	//
	// Experimental.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The IAM role that the Amazon SageMaker service assumes.
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html#sagemaker-roles-createmodel-perms
	//
	// Default: - a new IAM role will be created with the `AmazonSageMakerFullAccess` policy attached.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The security groups to associate to the Model.
	//
	// If no security groups are provided and 'vpc' is
	// configured, one security group will be created automatically.
	// Default: - A security group will be automatically created if 'vpc' is supplied.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC to deploy model containers to.
	// Default: - none.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The VPC subnets to use when deploying model containers.
	// Default: - none.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

