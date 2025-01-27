package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A sourced configuration represents configuration stored in an Amazon S3 bucket, AWS Secrets Manager secret, Systems Manager (SSM) Parameter Store parameter, SSM document, or AWS CodePipeline.
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
type SourcedConfiguration interface {
	constructs.Construct
	IConfiguration
	IExtensible
	// The application associated with the configuration.
	// Experimental.
	Application() IApplication
	// Experimental.
	ApplicationId() *string
	// Experimental.
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the configuration profile.
	// Experimental.
	ConfigurationProfileArn() *string
	// The ID of the configuration profile.
	// Experimental.
	ConfigurationProfileId() *string
	// The deployment key for the configuration.
	// Experimental.
	DeploymentKey() awskms.IKey
	// The deployment strategy for the configuration.
	// Experimental.
	DeploymentStrategy() IDeploymentStrategy
	// The environments to deploy to.
	// Experimental.
	DeployTo() *[]IEnvironment
	// The description of the configuration.
	// Experimental.
	Description() *string
	// Experimental.
	Extensible() ExtensibleBase
	// Experimental.
	SetExtensible(val ExtensibleBase)
	// The location where the configuration is stored.
	// Experimental.
	Location() ConfigurationSource
	// The name of the configuration.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The IAM role to retrieve the configuration.
	// Experimental.
	RetrievalRole() awsiam.IRole
	// The key to decrypt the configuration if applicable.
	//
	// This key
	// can be used when storing configuration in AWS Secrets Manager, Systems Manager Parameter Store,
	// or Amazon S3.
	// Experimental.
	SourceKey() awskms.IKey
	// The configuration type.
	// Experimental.
	Type() ConfigurationType
	// The validators for the configuration.
	// Experimental.
	Validators() *[]IValidator
	// The version number of the configuration to deploy.
	// Experimental.
	VersionNumber() *string
	// Experimental.
	AddExistingEnvironmentsToApplication()
	// Adds an extension association to the configuration profile.
	// Experimental.
	AddExtension(extension IExtension)
	// Deploys the configuration to the specified environment.
	// Experimental.
	Deploy(environment IEnvironment)
	// Experimental.
	DeployConfigToEnvironments()
	// Experimental.
	GetDeploymentHash(environment IEnvironment) *string
	// Adds an extension defined by the action point and event destination and also creates an extension association to the configuration profile.
	// Experimental.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to the configuration profile.
	// Experimental.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SourcedConfiguration
type jsiiProxy_SourcedConfiguration struct {
	internal.Type__constructsConstruct
	jsiiProxy_IConfiguration
	jsiiProxy_IExtensible
}

func (j *jsiiProxy_SourcedConfiguration) Application() IApplication {
	var returns IApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) ConfigurationProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) DeploymentKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"deploymentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) DeploymentStrategy() IDeploymentStrategy {
	var returns IDeploymentStrategy
	_jsii_.Get(
		j,
		"deploymentStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) DeployTo() *[]IEnvironment {
	var returns *[]IEnvironment
	_jsii_.Get(
		j,
		"deployTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Extensible() ExtensibleBase {
	var returns ExtensibleBase
	_jsii_.Get(
		j,
		"extensible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Location() ConfigurationSource {
	var returns ConfigurationSource
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) RetrievalRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"retrievalRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) SourceKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"sourceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Type() ConfigurationType {
	var returns ConfigurationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) Validators() *[]IValidator {
	var returns *[]IValidator
	_jsii_.Get(
		j,
		"validators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourcedConfiguration) VersionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}


// Experimental.
func NewSourcedConfiguration(scope constructs.Construct, id *string, props *SourcedConfigurationProps) SourcedConfiguration {
	_init_.Initialize()

	if err := validateNewSourcedConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SourcedConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SourcedConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSourcedConfiguration_Override(s SourcedConfiguration, scope constructs.Construct, id *string, props *SourcedConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SourcedConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SourcedConfiguration)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_SourcedConfiguration)SetExtensible(val ExtensibleBase) {
	if err := j.validateSetExtensibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensible",
		val,
	)
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
// Experimental.
func SourcedConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSourcedConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.SourcedConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourcedConfiguration) AddExistingEnvironmentsToApplication() {
	_jsii_.InvokeVoid(
		s,
		"addExistingEnvironmentsToApplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourcedConfiguration) AddExtension(extension IExtension) {
	if err := s.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addExtension",
		[]interface{}{extension},
	)
}

func (s *jsiiProxy_SourcedConfiguration) Deploy(environment IEnvironment) {
	if err := s.validateDeployParameters(environment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"deploy",
		[]interface{}{environment},
	)
}

func (s *jsiiProxy_SourcedConfiguration) DeployConfigToEnvironments() {
	_jsii_.InvokeVoid(
		s,
		"deployConfigToEnvironments",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourcedConfiguration) GetDeploymentHash(environment IEnvironment) *string {
	if err := s.validateGetDeploymentHashParameters(environment); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getDeploymentHash",
		[]interface{}{environment},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourcedConfiguration) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := s.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

func (s *jsiiProxy_SourcedConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

