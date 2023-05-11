package awssagemaker


// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckBiasProperty := &DriftCheckBiasProperty{
//   	ConfigFile: &FileSourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   	},
//   	PostTrainingConstraints: &MetricsSourceProperty{
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   	},
//   	PreTrainingConstraints: &MetricsSourceProperty{
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_DriftCheckBiasProperty struct {
	// The bias config file for a model.
	ConfigFile interface{} `field:"optional" json:"configFile" yaml:"configFile"`
	// The post-training constraints.
	PostTrainingConstraints interface{} `field:"optional" json:"postTrainingConstraints" yaml:"postTrainingConstraints"`
	// The pre-training constraints.
	PreTrainingConstraints interface{} `field:"optional" json:"preTrainingConstraints" yaml:"preTrainingConstraints"`
}
