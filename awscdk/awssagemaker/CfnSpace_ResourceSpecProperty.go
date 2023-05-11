package awssagemaker


// Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecProperty := &ResourceSpecProperty{
//   	InstanceType: jsii.String("instanceType"),
//   	SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   	SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   }
//
type CfnSpace_ResourceSpecProperty struct {
	// The instance type that the image version runs on.
	//
	// > *JupyterServer apps* only support the `system` value.
	// >
	// > For *KernelGateway apps* , the `system` value is translated to `ml.t3.medium` . KernelGateway apps also support all other values for available instance types.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the SageMaker image that the image version belongs to.
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}
