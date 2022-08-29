package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Pinpoint::Segment`.
//
// Updates the configuration, dimension, and other settings for an existing segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//   var tags interface{}
//   var userAttributes interface{}
//
//   cfnSegment := awscdk.Aws_pinpoint.NewCfnSegment(this, jsii.String("MyCfnSegment"), &cfnSegmentProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensions: &segmentDimensionsProperty{
//   		attributes: attributes,
//   		behavior: &behaviorProperty{
//   			recency: &recencyProperty{
//   				duration: jsii.String("duration"),
//   				recencyType: jsii.String("recencyType"),
//   			},
//   		},
//   		demographic: &demographicProperty{
//   			appVersion: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			channel: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			deviceType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			make: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			model: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			platform: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		location: &locationProperty{
//   			country: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			gpsPoint: &gPSPointProperty{
//   				coordinates: &coordinatesProperty{
//   					latitude: jsii.Number(123),
//   					longitude: jsii.Number(123),
//   				},
//   				rangeInKilometers: jsii.Number(123),
//   			},
//   		},
//   		metrics: metrics,
//   		userAttributes: userAttributes,
//   	},
//   	segmentGroups: &segmentGroupsProperty{
//   		groups: []interface{}{
//   			&groupsProperty{
//   				dimensions: []interface{}{
//   					&segmentDimensionsProperty{
//   						attributes: attributes,
//   						behavior: &behaviorProperty{
//   							recency: &recencyProperty{
//   								duration: jsii.String("duration"),
//   								recencyType: jsii.String("recencyType"),
//   							},
//   						},
//   						demographic: &demographicProperty{
//   							appVersion: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							channel: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							deviceType: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							make: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							model: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							platform: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   						location: &locationProperty{
//   							country: &setDimensionProperty{
//   								dimensionType: jsii.String("dimensionType"),
//   								values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   							gpsPoint: &gPSPointProperty{
//   								coordinates: &coordinatesProperty{
//   									latitude: jsii.Number(123),
//   									longitude: jsii.Number(123),
//   								},
//   								rangeInKilometers: jsii.Number(123),
//   							},
//   						},
//   						metrics: metrics,
//   						userAttributes: userAttributes,
//   					},
//   				},
//   				sourceSegments: []interface{}{
//   					&sourceSegmentsProperty{
//   						id: jsii.String("id"),
//
//   						// the properties below are optional
//   						version: jsii.Number(123),
//   					},
//   				},
//   				sourceType: jsii.String("sourceType"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		include: jsii.String("include"),
//   	},
//   	tags: tags,
//   })
//
type CfnSegment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the Amazon Pinpoint application that the segment is associated with.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the segment.
	AttrArn() *string
	// The unique identifier for the segment.
	AttrSegmentId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The criteria that define the dimensions for the segment.
	Dimensions() interface{}
	SetDimensions(val interface{})
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
	// The name of the segment.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The segment group to use and the dimensions to apply to the group's base segments in order to build the segment.
	//
	// A segment group can consist of zero or more base segments. Your request can include only one segment group.
	SegmentGroups() interface{}
	SetSegmentGroups(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnSegment
type jsiiProxy_CfnSegment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSegment) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) AttrSegmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSegmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Dimensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) SegmentGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSegment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::Segment`.
func NewCfnSegment(scope constructs.Construct, id *string, props *CfnSegmentProps) CfnSegment {
	_init_.Initialize()

	j := jsiiProxy_CfnSegment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::Segment`.
func NewCfnSegment_Override(c CfnSegment, scope constructs.Construct, id *string, props *CfnSegmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSegment) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetDimensions(val interface{}) {
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSegment) SetSegmentGroups(val interface{}) {
	_jsii_.Set(
		j,
		"segmentGroups",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSegment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSegment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
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
func CfnSegment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSegment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnSegment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSegment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSegment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSegment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSegment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSegment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSegment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSegment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSegment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSegment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSegment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSegment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
