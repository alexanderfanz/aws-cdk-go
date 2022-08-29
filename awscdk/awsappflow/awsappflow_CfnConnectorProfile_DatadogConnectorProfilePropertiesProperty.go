package awsappflow


// The connector-specific profile properties required by Datadog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datadogConnectorProfilePropertiesProperty := &datadogConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_DatadogConnectorProfilePropertiesProperty struct {
	// The location of the Datadog resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}
