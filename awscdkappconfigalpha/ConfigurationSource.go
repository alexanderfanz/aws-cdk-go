package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Defines the integrated configuration sources.
//
// Example:
//   var application application
//   var bucket bucket
//
//
//   appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
//   	Application: Application,
//   	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   	Name: jsii.String("MyConfig"),
//   	Description: jsii.String("This is my sourced configuration from CDK."),
//   })
//
// Experimental.
type ConfigurationSource interface {
	// The KMS Key that encrypts the configuration.
	// Experimental.
	Key() awskms.IKey
	// The URI of the configuration source.
	// Experimental.
	LocationUri() *string
	// The type of the configuration source.
	// Experimental.
	Type() ConfigurationSourceType
}

// The jsii proxy struct for ConfigurationSource
type jsiiProxy_ConfigurationSource struct {
	_ byte // padding
}

func (j *jsiiProxy_ConfigurationSource) Key() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSource) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSource) Type() ConfigurationSourceType {
	var returns ConfigurationSourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewConfigurationSource_Override(c ConfigurationSource) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		nil, // no parameters
		c,
	)
}

// Defines configuration content from an Amazon S3 bucket.
// Experimental.
func ConfigurationSource_FromBucket(bucket awss3.IBucket, objectKey *string, key awskms.IKey) ConfigurationSource {
	_init_.Initialize()

	if err := validateConfigurationSource_FromBucketParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns ConfigurationSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		"fromBucket",
		[]interface{}{bucket, objectKey, key},
		&returns,
	)

	return returns
}

// Defines configuration content from a Systems Manager (SSM) document.
// Experimental.
func ConfigurationSource_FromCfnDocument(document awsssm.CfnDocument) ConfigurationSource {
	_init_.Initialize()

	if err := validateConfigurationSource_FromCfnDocumentParameters(document); err != nil {
		panic(err)
	}
	var returns ConfigurationSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		"fromCfnDocument",
		[]interface{}{document},
		&returns,
	)

	return returns
}

// Defines configuration content from a Systems Manager (SSM) Parameter Store parameter.
// Experimental.
func ConfigurationSource_FromParameter(parameter awsssm.IParameter, key awskms.IKey) ConfigurationSource {
	_init_.Initialize()

	if err := validateConfigurationSource_FromParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns ConfigurationSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		"fromParameter",
		[]interface{}{parameter, key},
		&returns,
	)

	return returns
}

// Defines configuration content from AWS CodePipeline.
// Experimental.
func ConfigurationSource_FromPipeline(pipeline awscodepipeline.IPipeline) ConfigurationSource {
	_init_.Initialize()

	if err := validateConfigurationSource_FromPipelineParameters(pipeline); err != nil {
		panic(err)
	}
	var returns ConfigurationSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		"fromPipeline",
		[]interface{}{pipeline},
		&returns,
	)

	return returns
}

// Defines configuration content from an AWS Secrets Manager secret.
// Experimental.
func ConfigurationSource_FromSecret(secret awssecretsmanager.ISecret) ConfigurationSource {
	_init_.Initialize()

	if err := validateConfigurationSource_FromSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns ConfigurationSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
		"fromSecret",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

