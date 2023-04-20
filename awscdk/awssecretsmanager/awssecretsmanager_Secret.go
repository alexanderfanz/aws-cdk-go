package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Creates a new secret in AWS SecretsManager.
//
// Example:
//   // Creates a new IAM user, access and secret keys, and stores the secret access key in a Secret.
//   user := iam.NewUser(this, jsii.String("User"))
//   accessKey := iam.NewAccessKey(this, jsii.String("AccessKey"), &AccessKeyProps{
//   	User: User,
//   })
//   secretValue := secretsmanager.SecretStringValueBeta1_FromToken(accessKey.SecretAccessKey.ToString())
//   secretsmanager.NewSecret(this, jsii.String("Secret"), &SecretProps{
//   	SecretStringBeta1: secretValue,
//   })
//
// Experimental.
type Secret interface {
	awscdk.Resource
	ISecret
	// Provides an identifier for this secret for use in IAM policies.
	//
	// If there is a full ARN, this is just the ARN;
	// if we have a partial ARN -- due to either importing by secret name or partial ARN --
	// then we need to add a suffix to capture the full ARN's format.
	// Experimental.
	ArnForPolicies() *string
	// Experimental.
	AutoCreatePolicy() *bool
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	// Experimental.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	// Experimental.
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	// Experimental.
	SecretName() *string
	// Retrieve the value of the stored secret as a `SecretValue`.
	// Experimental.
	SecretValue() awscdk.SecretValue
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a replica region for the secret.
	// Experimental.
	AddReplicaRegion(region *string, encryptionKey awskms.IKey)
	// Adds a rotation schedule to the secret.
	// Experimental.
	AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule
	// Adds a target attachment to the secret.
	//
	// Returns: an AttachedSecret.
	// Deprecated: use `attach()` instead.
	AddTargetAttachment(id *string, options *AttachedSecretOptions) SecretTargetAttachment
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Attach a target to this secret.
	//
	// Returns: An attached secret.
	// Experimental.
	Attach(target ISecretAttachmentTarget) ISecret
	// Denies the `DeleteSecret` action to all principals within the current account.
	// Experimental.
	DenyAccountRootDelete()
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants reading the secret value to some role.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	// Experimental.
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecret
}

func (j *jsiiProxy_Secret) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecret(scope constructs.Construct, id *string, props *SecretProps) Secret {
	_init_.Initialize()

	if err := validateNewSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Secret{}

	_jsii_.Create(
		"monocdk.aws_secretsmanager.Secret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecret_Override(s Secret, scope constructs.Construct, id *string, props *SecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_secretsmanager.Secret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Deprecated: use `fromSecretCompleteArn` or `fromSecretPartialArn`.
func Secret_FromSecretArn(scope constructs.Construct, id *string, secretArn *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretArnParameters(scope, id, secretArn); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretArn",
		[]interface{}{scope, id, secretArn},
		&returns,
	)

	return returns
}

// Import an existing secret into the Stack.
// Experimental.
func Secret_FromSecretAttributes(scope constructs.Construct, id *string, attrs *SecretAttributes) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a secret by complete ARN.
//
// The complete ARN is the ARN with the Secrets Manager-supplied suffix.
// Experimental.
func Secret_FromSecretCompleteArn(scope constructs.Construct, id *string, secretCompleteArn *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretCompleteArnParameters(scope, id, secretCompleteArn); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretCompleteArn",
		[]interface{}{scope, id, secretCompleteArn},
		&returns,
	)

	return returns
}

// Imports a secret by secret name;
//
// the ARN of the Secret will be set to the secret name.
// A secret with this name must exist in the same account & region.
// Deprecated: use `fromSecretNameV2`.
func Secret_FromSecretName(scope constructs.Construct, id *string, secretName *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretNameParameters(scope, id, secretName); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretName",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by secret name.
//
// A secret with this name must exist in the same account & region.
// Replaces the deprecated `fromSecretName`.
// Experimental.
func Secret_FromSecretNameV2(scope constructs.Construct, id *string, secretName *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretNameV2Parameters(scope, id, secretName); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretNameV2",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by partial ARN.
//
// The partial ARN is the ARN without the Secrets Manager-supplied suffix.
// Experimental.
func Secret_FromSecretPartialArn(scope constructs.Construct, id *string, secretPartialArn *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromSecretPartialArnParameters(scope, id, secretPartialArn); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"fromSecretPartialArn",
		[]interface{}{scope, id, secretPartialArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Secret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Secret_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecret_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_secretsmanager.Secret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	if err := s.validateAddReplicaRegionParameters(region); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

func (s *jsiiProxy_Secret) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
	if err := s.validateAddRotationScheduleParameters(id, options); err != nil {
		panic(err)
	}
	var returns RotationSchedule

	_jsii_.Invoke(
		s,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) AddTargetAttachment(id *string, options *AttachedSecretOptions) SecretTargetAttachment {
	if err := s.validateAddTargetAttachmentParameters(id, options); err != nil {
		panic(err)
	}
	var returns SecretTargetAttachment

	_jsii_.Invoke(
		s,
		"addTargetAttachment",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := s.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Secret) Attach(target ISecretAttachmentTarget) ISecret {
	if err := s.validateAttachParameters(target); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.Invoke(
		s,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		s,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Secret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (s *jsiiProxy_Secret) GetResourceNameAttribute(nameAttr *string) *string {
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

func (s *jsiiProxy_Secret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
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

func (s *jsiiProxy_Secret) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Secret) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Secret) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Secret) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	if err := s.validateSecretValueFromJsonParameters(jsonField); err != nil {
		panic(err)
	}
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		s,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Secret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

