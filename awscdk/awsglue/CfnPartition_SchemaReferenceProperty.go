package awsglue


// An object that references a schema stored in the AWS Glue Schema Registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaReferenceProperty := &SchemaReferenceProperty{
//   	SchemaId: &SchemaIdProperty{
//   		RegistryName: jsii.String("registryName"),
//   		SchemaArn: jsii.String("schemaArn"),
//   		SchemaName: jsii.String("schemaName"),
//   	},
//   	SchemaVersionId: jsii.String("schemaVersionId"),
//   	SchemaVersionNumber: jsii.Number(123),
//   }
//
type CfnPartition_SchemaReferenceProperty struct {
	// A structure that contains schema identity fields.
	//
	// Either this or the `SchemaVersionId` has to be
	// provided.
	SchemaId interface{} `field:"optional" json:"schemaId" yaml:"schemaId"`
	// The unique ID assigned to a version of the schema.
	//
	// Either this or the `SchemaId` has to be provided.
	SchemaVersionId *string `field:"optional" json:"schemaVersionId" yaml:"schemaVersionId"`
	// The version number of the schema.
	SchemaVersionNumber *float64 `field:"optional" json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
}

