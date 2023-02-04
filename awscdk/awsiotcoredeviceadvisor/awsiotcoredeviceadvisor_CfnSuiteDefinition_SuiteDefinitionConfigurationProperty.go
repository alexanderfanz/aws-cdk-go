package awsiotcoredeviceadvisor


// Gets the suite definition configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suiteDefinitionConfigurationProperty := &suiteDefinitionConfigurationProperty{
//   	devicePermissionRoleArn: jsii.String("devicePermissionRoleArn"),
//   	rootGroup: jsii.String("rootGroup"),
//
//   	// the properties below are optional
//   	devices: []interface{}{
//   		&deviceUnderTestProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			thingArn: jsii.String("thingArn"),
//   		},
//   	},
//   	intendedForQualification: jsii.Boolean(false),
//   	suiteDefinitionName: jsii.String("suiteDefinitionName"),
//   }
//
type CfnSuiteDefinition_SuiteDefinitionConfigurationProperty struct {
	// Gets the device permission ARN.
	//
	// This is a required parameter.
	DevicePermissionRoleArn *string `field:"required" json:"devicePermissionRoleArn" yaml:"devicePermissionRoleArn"`
	// Gets the test suite root group.
	//
	// This is a required parameter.
	RootGroup *string `field:"required" json:"rootGroup" yaml:"rootGroup"`
	// Gets the devices configured.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Gets the tests intended for qualification in a suite.
	IntendedForQualification interface{} `field:"optional" json:"intendedForQualification" yaml:"intendedForQualification"`
	// Gets the suite definition name.
	//
	// This is a required parameter.
	SuiteDefinitionName *string `field:"optional" json:"suiteDefinitionName" yaml:"suiteDefinitionName"`
}
