package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new String SSM Parameter.
//
// Example:
//   ssmParameter := ssm.NewStringParameter(this, jsii.String("mySsmParameter"), &StringParameterProps{
//   	ParameterName: jsii.String("mySsmParameter"),
//   	StringValue: jsii.String("mySsmParameterValue"),
//   })
//
type StringParameter interface {
	awscdk.Resource
	IParameter
	IStringParameter
	// The encryption key that is used to encrypt this parameter.
	//
	// * @default - default master key.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The ARN of the SSM Parameter resource.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	ParameterName() *string
	// The type of the SSM Parameter resource.
	ParameterType() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value.
	StringValue() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants read (DescribeParameter, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StringParameter
type jsiiProxy_StringParameter struct {
	internal.Type__awscdkResource
	jsiiProxy_IParameter
	jsiiProxy_IStringParameter
}

func (j *jsiiProxy_StringParameter) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringParameter) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}


func NewStringParameter(scope constructs.Construct, id *string, props *StringParameterProps) StringParameter {
	_init_.Initialize()

	if err := validateNewStringParameterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringParameter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ssm.StringParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStringParameter_Override(s StringParameter, scope constructs.Construct, id *string, props *StringParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ssm.StringParameter",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports a secure string parameter from the SSM parameter store.
func StringParameter_FromSecureStringParameterAttributes(scope constructs.Construct, id *string, attrs *SecureStringParameterAttributes) IStringParameter {
	_init_.Initialize()

	if err := validateStringParameter_FromSecureStringParameterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IStringParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"fromSecureStringParameterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an external string parameter with name and optional version.
func StringParameter_FromStringParameterAttributes(scope constructs.Construct, id *string, attrs *StringParameterAttributes) IStringParameter {
	_init_.Initialize()

	if err := validateStringParameter_FromStringParameterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IStringParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"fromStringParameterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an external string parameter by name.
func StringParameter_FromStringParameterName(scope constructs.Construct, id *string, stringParameterName *string) IStringParameter {
	_init_.Initialize()

	if err := validateStringParameter_FromStringParameterNameParameters(scope, id, stringParameterName); err != nil {
		panic(err)
	}
	var returns IStringParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"fromStringParameterName",
		[]interface{}{scope, id, stringParameterName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StringParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStringParameter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func StringParameter_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStringParameter_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func StringParameter_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateStringParameter_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment).
// Deprecated: Use `SecretValue.ssmSecure()` instead, it will correctly type the imported value as a `SecretValue` and allow importing without version.
func StringParameter_ValueForSecureStringParameter(scope constructs.Construct, parameterName *string, version *float64) *string {
	_init_.Initialize()

	if err := validateStringParameter_ValueForSecureStringParameterParameters(scope, parameterName, version); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"valueForSecureStringParameter",
		[]interface{}{scope, parameterName, version},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment) to the string value of an SSM string parameter.
func StringParameter_ValueForStringParameter(scope constructs.Construct, parameterName *string, version *float64) *string {
	_init_.Initialize()

	if err := validateStringParameter_ValueForStringParameterParameters(scope, parameterName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"valueForStringParameter",
		[]interface{}{scope, parameterName, version},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment) to the string value of an SSM string parameter.
// Deprecated: - use valueForTypedStringParameterV2 instead.
func StringParameter_ValueForTypedStringParameter(scope constructs.Construct, parameterName *string, type_ ParameterType, version *float64) *string {
	_init_.Initialize()

	if err := validateStringParameter_ValueForTypedStringParameterParameters(scope, parameterName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"valueForTypedStringParameter",
		[]interface{}{scope, parameterName, type_, version},
		&returns,
	)

	return returns
}

// Returns a token that will resolve (during deployment) to the string value of an SSM string parameter.
func StringParameter_ValueForTypedStringParameterV2(scope constructs.Construct, parameterName *string, type_ ParameterValueType, version *float64) *string {
	_init_.Initialize()

	if err := validateStringParameter_ValueForTypedStringParameterV2Parameters(scope, parameterName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"valueForTypedStringParameterV2",
		[]interface{}{scope, parameterName, type_, version},
		&returns,
	)

	return returns
}

// Reads the value of an SSM parameter during synthesis through an environmental context provider.
//
// Requires that the stack this scope is defined in will have explicit
// account/region information. Otherwise, it will fail during synthesis.
func StringParameter_ValueFromLookup(scope constructs.Construct, parameterName *string) *string {
	_init_.Initialize()

	if err := validateStringParameter_ValueFromLookupParameters(scope, parameterName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.StringParameter",
		"valueFromLookup",
		[]interface{}{scope, parameterName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_StringParameter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
