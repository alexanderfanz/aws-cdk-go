package awsquicksight


// The set parameter operation that sets parameters in custom action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionSetParametersOperationProperty := &CustomActionSetParametersOperationProperty{
//   	ParameterValueConfigurations: []interface{}{
//   		&SetParameterValueConfigurationProperty{
//   			DestinationParameterName: jsii.String("destinationParameterName"),
//   			Value: &DestinationParameterValueConfigurationProperty{
//   				CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   					CustomValues: &CustomParameterValuesProperty{
//   						DateTimeValues: []*string{
//   							jsii.String("dateTimeValues"),
//   						},
//   						DecimalValues: []interface{}{
//   							jsii.Number(123),
//   						},
//   						IntegerValues: []interface{}{
//   							jsii.Number(123),
//   						},
//   						StringValues: []*string{
//   							jsii.String("stringValues"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					IncludeNullValue: jsii.Boolean(false),
//   				},
//   				SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   				SourceField: jsii.String("sourceField"),
//   				SourceParameterName: jsii.String("sourceParameterName"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_CustomActionSetParametersOperationProperty struct {
	// The parameter that determines the value configuration.
	ParameterValueConfigurations interface{} `field:"required" json:"parameterValueConfigurations" yaml:"parameterValueConfigurations"`
}
