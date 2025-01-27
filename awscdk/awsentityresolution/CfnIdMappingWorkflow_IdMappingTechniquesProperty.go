package awsentityresolution


// An object which defines the ID mapping techniques and provider configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingTechniquesProperty := &IdMappingTechniquesProperty{
//   	IdMappingType: jsii.String("idMappingType"),
//   	ProviderProperties: &ProviderPropertiesProperty{
//   		ProviderServiceArn: jsii.String("providerServiceArn"),
//
//   		// the properties below are optional
//   		IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   			IntermediateS3Path: jsii.String("intermediateS3Path"),
//   		},
//   		ProviderConfiguration: map[string]*string{
//   			"providerConfigurationKey": jsii.String("providerConfiguration"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html
//
type CfnIdMappingWorkflow_IdMappingTechniquesProperty struct {
	// The type of ID mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-idmappingtype
	//
	IdMappingType *string `field:"optional" json:"idMappingType" yaml:"idMappingType"`
	// An object which defines any additional configurations required by the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingtechniques.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques-providerproperties
	//
	ProviderProperties interface{} `field:"optional" json:"providerProperties" yaml:"providerProperties"`
}

