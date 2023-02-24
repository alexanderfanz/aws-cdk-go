// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Exposed secret for log configuration.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &JobDefinitionProps{
//   	Container: &JobDefinitionContainer{
//   		Image: ecs.EcrImage_FromRegistry(jsii.String("docker/whalesay")),
//   		LogConfiguration: &LogConfiguration{
//   			LogDriver: batch.LogDriver_AWSLOGS,
//   			Options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			SecretOptions: []exposedSecret{
//   				batch.*exposedSecret_FromParametersStore(jsii.String("xyz"), ssm.StringParameter_FromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ExposedSecret interface {
	// Name of the option.
	// Experimental.
	OptionName() *string
	// Experimental.
	SetOptionName(val *string)
	// ARN of the secret option.
	// Experimental.
	SecretArn() *string
	// Experimental.
	SetSecretArn(val *string)
}

// The jsii proxy struct for ExposedSecret
type jsiiProxy_ExposedSecret struct {
	_ byte // padding
}

func (j *jsiiProxy_ExposedSecret) OptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExposedSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewExposedSecret(optionName *string, secretArn *string) ExposedSecret {
	_init_.Initialize()

	if err := validateNewExposedSecretParameters(optionName, secretArn); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExposedSecret{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.ExposedSecret",
		[]interface{}{optionName, secretArn},
		&j,
	)

	return &j
}

// Experimental.
func NewExposedSecret_Override(e ExposedSecret, optionName *string, secretArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.ExposedSecret",
		[]interface{}{optionName, secretArn},
		e,
	)
}

func (j *jsiiProxy_ExposedSecret)SetOptionName(val *string) {
	if err := j.validateSetOptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionName",
		val,
	)
}

func (j *jsiiProxy_ExposedSecret)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

// User Parameters Store Parameter.
// Experimental.
func ExposedSecret_FromParametersStore(optionName *string, parameter awsssm.IParameter) ExposedSecret {
	_init_.Initialize()

	if err := validateExposedSecret_FromParametersStoreParameters(optionName, parameter); err != nil {
		panic(err)
	}
	var returns ExposedSecret

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.ExposedSecret",
		"fromParametersStore",
		[]interface{}{optionName, parameter},
		&returns,
	)

	return returns
}

// Use Secrets Manager Secret.
// Experimental.
func ExposedSecret_FromSecretsManager(optionName *string, secret awssecretsmanager.ISecret) ExposedSecret {
	_init_.Initialize()

	if err := validateExposedSecret_FromSecretsManagerParameters(optionName, secret); err != nil {
		panic(err)
	}
	var returns ExposedSecret

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.ExposedSecret",
		"fromSecretsManager",
		[]interface{}{optionName, secret},
		&returns,
	)

	return returns
}

