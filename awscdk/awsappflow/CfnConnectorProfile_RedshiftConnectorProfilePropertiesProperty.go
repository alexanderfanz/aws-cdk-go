package awsappflow


// The connector-specific profile properties when using Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftConnectorProfilePropertiesProperty := &RedshiftConnectorProfilePropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	DataApiRoleArn: jsii.String("dataApiRoleArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DatabaseUrl: jsii.String("databaseUrl"),
//   	IsRedshiftServerless: jsii.Boolean(false),
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty struct {
	// A name for the associated Amazon S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The Amazon Resource Name (ARN) of IAM role that grants Amazon Redshift read-only access to Amazon S3.
	//
	// For more information, and for the polices that you attach to this role, see [Allow Amazon Redshift to access your Amazon AppFlow data in Amazon S3](https://docs.aws.amazon.com/appflow/latest/userguide/security_iam_service-role-policies.html#redshift-access-s3) .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.ClusterIdentifier`.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.DataApiRoleArn`.
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.DatabaseName`.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.IsRedshiftServerless`.
	IsRedshiftServerless interface{} `field:"optional" json:"isRedshiftServerless" yaml:"isRedshiftServerless"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.WorkgroupName`.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

