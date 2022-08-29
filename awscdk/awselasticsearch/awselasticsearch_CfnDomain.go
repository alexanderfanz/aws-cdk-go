package awselasticsearch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Elasticsearch::Domain`.
//
// The AWS::Elasticsearch::Domain resource creates an Amazon OpenSearch Service domain.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and legacy Elasticsearch. For instructions to upgrade domains defined within CloudFormation from Elasticsearch to OpenSearch, see [Remarks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--remarks) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicies interface{}
//
//   cfnDomain := awscdk.Aws_elasticsearch.NewCfnDomain(this, jsii.String("MyCfnDomain"), &cfnDomainProps{
//   	accessPolicies: accessPolicies,
//   	advancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	advancedSecurityOptions: &advancedSecurityOptionsInputProperty{
//   		anonymousAuthEnabled: jsii.Boolean(false),
//   		enabled: jsii.Boolean(false),
//   		internalUserDatabaseEnabled: jsii.Boolean(false),
//   		masterUserOptions: &masterUserOptionsProperty{
//   			masterUserArn: jsii.String("masterUserArn"),
//   			masterUserName: jsii.String("masterUserName"),
//   			masterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   	},
//   	cognitoOptions: &cognitoOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		identityPoolId: jsii.String("identityPoolId"),
//   		roleArn: jsii.String("roleArn"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	domainEndpointOptions: &domainEndpointOptionsProperty{
//   		customEndpoint: jsii.String("customEndpoint"),
//   		customEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   		customEndpointEnabled: jsii.Boolean(false),
//   		enforceHttps: jsii.Boolean(false),
//   		tlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	ebsOptions: &eBSOptionsProperty{
//   		ebsEnabled: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	elasticsearchClusterConfig: &elasticsearchClusterConfigProperty{
//   		coldStorageOptions: &coldStorageOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		dedicatedMasterCount: jsii.Number(123),
//   		dedicatedMasterEnabled: jsii.Boolean(false),
//   		dedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		warmCount: jsii.Number(123),
//   		warmEnabled: jsii.Boolean(false),
//   		warmType: jsii.String("warmType"),
//   		zoneAwarenessConfig: &zoneAwarenessConfigProperty{
//   			availabilityZoneCount: jsii.Number(123),
//   		},
//   		zoneAwarenessEnabled: jsii.Boolean(false),
//   	},
//   	elasticsearchVersion: jsii.String("elasticsearchVersion"),
//   	encryptionAtRestOptions: &encryptionAtRestOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	logPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	nodeToNodeEncryptionOptions: &nodeToNodeEncryptionOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	snapshotOptions: &snapshotOptionsProperty{
//   		automatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcOptions: &vPCOptionsProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guid* e.
	AccessPolicies() interface{}
	SetAccessPolicies(val interface{})
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [Advanced cluster parameters](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) in the *Amazon OpenSearch Service Developer Guide* .
	AdvancedOptions() interface{}
	SetAdvancedOptions(val interface{})
	// Specifies options for fine-grained access control.
	AdvancedSecurityOptions() interface{}
	SetAdvancedSecurityOptions(val interface{})
	// The Amazon Resource Name (ARN) of the domain, such as `arn:aws:es:us-west-2:123456789012:domain/mystack-elasti-1ab2cdefghij` .
	//
	// This returned value is the same as the one returned by `AWS::Elasticsearch::Domain.DomainArn` .
	AttrArn() *string
	// The domain-specific endpoint that's used for requests to the OpenSearch APIs, such as `search-mystack-elasti-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com` .
	AttrDomainEndpoint() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoOptions() interface{}
	SetCognitoOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	DomainEndpointOptions() interface{}
	SetDomainEndpointOptions(val interface{})
	// A name for the OpenSearch Service domain.
	//
	// For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the *Amazon OpenSearch Service Developer Guide* . If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DomainName() *string
	SetDomainName(val *string)
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	EbsOptions() interface{}
	SetEbsOptions(val interface{})
	// ElasticsearchClusterConfig is a property of the AWS::Elasticsearch::Domain resource that configures the cluster of an Amazon OpenSearch Service domain.
	ElasticsearchClusterConfig() interface{}
	SetElasticsearchClusterConfig(val interface{})
	// The version of Elasticsearch to use, such as 2.3. If not specified, 1.5 is used as the default. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `ElasticsearchVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `ElasticsearchVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	ElasticsearchVersion() *string
	SetElasticsearchVersion(val *string)
	// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	EncryptionAtRestOptions() interface{}
	SetEncryptionAtRestOptions(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	//
	// Each key needs a valid `LogPublishingOption` value.
	LogPublishingOptions() interface{}
	SetLogPublishingOptions(val interface{})
	// The tree node.
	Node() constructs.Node
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	NodeToNodeEncryptionOptions() interface{}
	SetNodeToNodeEncryptionOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indices.
	SnapshotOptions() interface{}
	SetSnapshotOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	VpcOptions() interface{}
	SetVpcOptions(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomain) AccessPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AdvancedOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AdvancedSecurityOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CognitoOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainEndpointOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EbsOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) ElasticsearchClusterConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EncryptionAtRestOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogPublishingOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) NodeToNodeEncryptionOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SnapshotOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) VpcOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}


// Create a new `AWS::Elasticsearch::Domain`.
func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Elasticsearch::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetAccessPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAdvancedOptions(val interface{}) {
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAdvancedSecurityOptions(val interface{}) {
	_jsii_.Set(
		j,
		"advancedSecurityOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetCognitoOptions(val interface{}) {
	_jsii_.Set(
		j,
		"cognitoOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainEndpointOptions(val interface{}) {
	_jsii_.Set(
		j,
		"domainEndpointOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetEbsOptions(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetElasticsearchClusterConfig(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchClusterConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetElasticsearchVersion(val *string) {
	_jsii_.Set(
		j,
		"elasticsearchVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetEncryptionAtRestOptions(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionAtRestOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetLogPublishingOptions(val interface{}) {
	_jsii_.Set(
		j,
		"logPublishingOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetNodeToNodeEncryptionOptions(val interface{}) {
	_jsii_.Set(
		j,
		"nodeToNodeEncryptionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetSnapshotOptions(val interface{}) {
	_jsii_.Set(
		j,
		"snapshotOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetVpcOptions(val interface{}) {
	_jsii_.Set(
		j,
		"vpcOptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
