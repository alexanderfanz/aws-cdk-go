package awsinspectorv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::InspectorV2::Filter`.
//
// Details about a filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilter := awscdk.Aws_inspectorv2.NewCfnFilter(this, jsii.String("MyCfnFilter"), &cfnFilterProps{
//   	filterAction: jsii.String("filterAction"),
//   	filterCriteria: &filterCriteriaProperty{
//   		awsAccountId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceImageId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceSubnetId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceVpcId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageArchitecture: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageHash: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImagePushedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		ecrImageRegistry: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageRepositoryName: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageTags: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingArn: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingStatus: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		firstObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		inspectorScore: []interface{}{
//   			&numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   		},
//   		lastObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		networkProtocol: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		portRange: []interface{}{
//   			&portRangeFilterProperty{
//   				beginInclusive: jsii.Number(123),
//   				endInclusive: jsii.Number(123),
//   			},
//   		},
//   		relatedVulnerabilities: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceTags: []interface{}{
//   			&mapFilterProperty{
//   				comparison: jsii.String("comparison"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		severity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		title: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		updatedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		vendorSeverity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilityId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilitySource: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerablePackages: []interface{}{
//   			&packageFilterProperty{
//   				architecture: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				epoch: &numberFilterProperty{
//   					lowerInclusive: jsii.Number(123),
//   					upperInclusive: jsii.Number(123),
//   				},
//   				name: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				release: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				sourceLayerHash: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				version: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   })
//
type CfnFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Number (ARN) associated with this filter.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the filter.
	Description() *string
	SetDescription(val *string)
	// The action that is to be applied to the findings that match the filter.
	FilterAction() *string
	SetFilterAction(val *string)
	// Details on the filter criteria associated with this filter.
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
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
	// The name of the filter.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnFilter
type jsiiProxy_CfnFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFilter) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterCriteria() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter(scope constructs.Construct, id *string, props *CfnFilterProps) CfnFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter_Override(c CfnFilter, scope constructs.Construct, id *string, props *CfnFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFilter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterAction(val *string) {
	_jsii_.Set(
		j,
		"filterAction",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterCriteria(val interface{}) {
	_jsii_.Set(
		j,
		"filterCriteria",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
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
func CfnFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFilter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFilter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFilter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFilter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
