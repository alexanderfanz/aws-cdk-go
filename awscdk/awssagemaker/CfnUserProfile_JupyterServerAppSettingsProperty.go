package awssagemaker


// The JupyterServer app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jupyterServerAppSettingsProperty := &JupyterServerAppSettingsProperty{
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterserverappsettings.html
//
type CfnUserProfile_JupyterServerAppSettingsProperty struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterserverappsettings.html#cfn-sagemaker-userprofile-jupyterserverappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

