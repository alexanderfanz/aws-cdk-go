package awsdatabrew

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DataBrew::Dataset`.
//
// TODO: EXAMPLE
//
type CfnDataset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Format() *string
	SetFormat(val *string)
	FormatOptions() interface{}
	SetFormatOptions(val interface{})
	Input() interface{}
	SetInput(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	PathOptions() interface{}
	SetPathOptions(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDataset
type jsiiProxy_CfnDataset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) FormatOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) PathOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Dataset`.
func NewCfnDataset(scope constructs.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope constructs.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataset) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetFormatOptions(val interface{}) {
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetInput(val interface{}) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetPathOptions(val interface{}) {
	_jsii_.Set(
		j,
		"pathOptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDataset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDataset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDataset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnDataset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDataset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDataset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDataset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDataset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDataset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDataset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDataset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDataset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnDataset_CsvOptionsProperty struct {
	// `CfnDataset.CsvOptionsProperty.Delimiter`.
	Delimiter *string `json:"delimiter"`
	// `CfnDataset.CsvOptionsProperty.HeaderRow`.
	HeaderRow interface{} `json:"headerRow"`
}

// TODO: EXAMPLE
//
type CfnDataset_DataCatalogInputDefinitionProperty struct {
	// `CfnDataset.DataCatalogInputDefinitionProperty.CatalogId`.
	CatalogId *string `json:"catalogId"`
	// `CfnDataset.DataCatalogInputDefinitionProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnDataset.DataCatalogInputDefinitionProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnDataset.DataCatalogInputDefinitionProperty.TempDirectory`.
	TempDirectory interface{} `json:"tempDirectory"`
}

// TODO: EXAMPLE
//
type CfnDataset_DatabaseInputDefinitionProperty struct {
	// `CfnDataset.DatabaseInputDefinitionProperty.DatabaseTableName`.
	DatabaseTableName *string `json:"databaseTableName"`
	// `CfnDataset.DatabaseInputDefinitionProperty.GlueConnectionName`.
	GlueConnectionName *string `json:"glueConnectionName"`
	// `CfnDataset.DatabaseInputDefinitionProperty.QueryString`.
	QueryString *string `json:"queryString"`
	// `CfnDataset.DatabaseInputDefinitionProperty.TempDirectory`.
	TempDirectory interface{} `json:"tempDirectory"`
}

// TODO: EXAMPLE
//
type CfnDataset_DatasetParameterProperty struct {
	// `CfnDataset.DatasetParameterProperty.CreateColumn`.
	CreateColumn interface{} `json:"createColumn"`
	// `CfnDataset.DatasetParameterProperty.DatetimeOptions`.
	DatetimeOptions interface{} `json:"datetimeOptions"`
	// `CfnDataset.DatasetParameterProperty.Filter`.
	Filter interface{} `json:"filter"`
	// `CfnDataset.DatasetParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnDataset.DatasetParameterProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnDataset_DatetimeOptionsProperty struct {
	// `CfnDataset.DatetimeOptionsProperty.Format`.
	Format *string `json:"format"`
	// `CfnDataset.DatetimeOptionsProperty.LocaleCode`.
	LocaleCode *string `json:"localeCode"`
	// `CfnDataset.DatetimeOptionsProperty.TimezoneOffset`.
	TimezoneOffset *string `json:"timezoneOffset"`
}

// TODO: EXAMPLE
//
type CfnDataset_ExcelOptionsProperty struct {
	// `CfnDataset.ExcelOptionsProperty.HeaderRow`.
	HeaderRow interface{} `json:"headerRow"`
	// `CfnDataset.ExcelOptionsProperty.SheetIndexes`.
	SheetIndexes interface{} `json:"sheetIndexes"`
	// `CfnDataset.ExcelOptionsProperty.SheetNames`.
	SheetNames *[]*string `json:"sheetNames"`
}

// TODO: EXAMPLE
//
type CfnDataset_FilesLimitProperty struct {
	// `CfnDataset.FilesLimitProperty.MaxFiles`.
	MaxFiles *float64 `json:"maxFiles"`
	// `CfnDataset.FilesLimitProperty.Order`.
	Order *string `json:"order"`
	// `CfnDataset.FilesLimitProperty.OrderedBy`.
	OrderedBy *string `json:"orderedBy"`
}

// TODO: EXAMPLE
//
type CfnDataset_FilterExpressionProperty struct {
	// `CfnDataset.FilterExpressionProperty.Expression`.
	Expression *string `json:"expression"`
	// `CfnDataset.FilterExpressionProperty.ValuesMap`.
	ValuesMap interface{} `json:"valuesMap"`
}

// TODO: EXAMPLE
//
type CfnDataset_FilterValueProperty struct {
	// `CfnDataset.FilterValueProperty.Value`.
	Value *string `json:"value"`
	// `CfnDataset.FilterValueProperty.ValueReference`.
	ValueReference *string `json:"valueReference"`
}

// TODO: EXAMPLE
//
type CfnDataset_FormatOptionsProperty struct {
	// `CfnDataset.FormatOptionsProperty.Csv`.
	Csv interface{} `json:"csv"`
	// `CfnDataset.FormatOptionsProperty.Excel`.
	Excel interface{} `json:"excel"`
	// `CfnDataset.FormatOptionsProperty.Json`.
	Json interface{} `json:"json"`
}

// TODO: EXAMPLE
//
type CfnDataset_InputProperty struct {
	// `CfnDataset.InputProperty.DatabaseInputDefinition`.
	DatabaseInputDefinition interface{} `json:"databaseInputDefinition"`
	// `CfnDataset.InputProperty.DataCatalogInputDefinition`.
	DataCatalogInputDefinition interface{} `json:"dataCatalogInputDefinition"`
	// `CfnDataset.InputProperty.Metadata`.
	Metadata interface{} `json:"metadata"`
	// `CfnDataset.InputProperty.S3InputDefinition`.
	S3InputDefinition interface{} `json:"s3InputDefinition"`
}

// TODO: EXAMPLE
//
type CfnDataset_JsonOptionsProperty struct {
	// `CfnDataset.JsonOptionsProperty.MultiLine`.
	MultiLine interface{} `json:"multiLine"`
}

// TODO: EXAMPLE
//
type CfnDataset_MetadataProperty struct {
	// `CfnDataset.MetadataProperty.SourceArn`.
	SourceArn *string `json:"sourceArn"`
}

// TODO: EXAMPLE
//
type CfnDataset_PathOptionsProperty struct {
	// `CfnDataset.PathOptionsProperty.FilesLimit`.
	FilesLimit interface{} `json:"filesLimit"`
	// `CfnDataset.PathOptionsProperty.LastModifiedDateCondition`.
	LastModifiedDateCondition interface{} `json:"lastModifiedDateCondition"`
	// `CfnDataset.PathOptionsProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
}

// TODO: EXAMPLE
//
type CfnDataset_PathParameterProperty struct {
	// `CfnDataset.PathParameterProperty.DatasetParameter`.
	DatasetParameter interface{} `json:"datasetParameter"`
	// `CfnDataset.PathParameterProperty.PathParameterName`.
	PathParameterName *string `json:"pathParameterName"`
}

// TODO: EXAMPLE
//
type CfnDataset_S3LocationProperty struct {
	// `CfnDataset.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDataset.S3LocationProperty.Key`.
	Key *string `json:"key"`
}

// Properties for defining a `AWS::DataBrew::Dataset`.
//
// TODO: EXAMPLE
//
type CfnDatasetProps struct {
	// `AWS::DataBrew::Dataset.Format`.
	Format *string `json:"format"`
	// `AWS::DataBrew::Dataset.FormatOptions`.
	FormatOptions interface{} `json:"formatOptions"`
	// `AWS::DataBrew::Dataset.Input`.
	Input interface{} `json:"input"`
	// `AWS::DataBrew::Dataset.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Dataset.PathOptions`.
	PathOptions interface{} `json:"pathOptions"`
	// `AWS::DataBrew::Dataset.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::DataBrew::Job`.
//
// TODO: EXAMPLE
//
type CfnJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseOutputs() interface{}
	SetDatabaseOutputs(val interface{})
	DataCatalogOutputs() interface{}
	SetDataCatalogOutputs(val interface{})
	DatasetName() *string
	SetDatasetName(val *string)
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	JobSample() interface{}
	SetJobSample(val interface{})
	LogicalId() *string
	LogSubscription() *string
	SetLogSubscription(val *string)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	OutputLocation() interface{}
	SetOutputLocation(val interface{})
	Outputs() interface{}
	SetOutputs(val interface{})
	ProfileConfiguration() interface{}
	SetProfileConfiguration(val interface{})
	ProjectName() *string
	SetProjectName(val *string)
	Recipe() interface{}
	SetRecipe(val interface{})
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Timeout() *float64
	SetTimeout(val *float64)
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	ValidationConfigurations() interface{}
	SetValidationConfigurations(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnJob
type jsiiProxy_CfnJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatabaseOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DataCatalogOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCatalogOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) JobSample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogSubscription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) OutputLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Outputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProfileConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Recipe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recipe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ValidationConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Job`.
func NewCfnJob(scope constructs.Construct, id *string, props *CfnJobProps) CfnJob {
	_init_.Initialize()

	j := jsiiProxy_CfnJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Job`.
func NewCfnJob_Override(c CfnJob, scope constructs.Construct, id *string, props *CfnJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJob) SetDatabaseOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"databaseOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDataCatalogOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"dataCatalogOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetEncryptionKeyArn(val *string) {
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetJobSample(val interface{}) {
	_jsii_.Set(
		j,
		"jobSample",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetLogSubscription(val *string) {
	_jsii_.Set(
		j,
		"logSubscription",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetOutputLocation(val interface{}) {
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"outputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetProfileConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"profileConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetRecipe(val interface{}) {
	_jsii_.Set(
		j,
		"recipe",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetValidationConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"validationConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJob_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnJob) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnJob) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnJob) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnJob) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnJob) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnJob) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnJob) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnJob_AllowedStatisticsProperty struct {
	// `CfnJob.AllowedStatisticsProperty.Statistics`.
	Statistics *[]*string `json:"statistics"`
}

// TODO: EXAMPLE
//
type CfnJob_ColumnSelectorProperty struct {
	// `CfnJob.ColumnSelectorProperty.Name`.
	Name *string `json:"name"`
	// `CfnJob.ColumnSelectorProperty.Regex`.
	Regex *string `json:"regex"`
}

// TODO: EXAMPLE
//
type CfnJob_ColumnStatisticsConfigurationProperty struct {
	// `CfnJob.ColumnStatisticsConfigurationProperty.Selectors`.
	Selectors interface{} `json:"selectors"`
	// `CfnJob.ColumnStatisticsConfigurationProperty.Statistics`.
	Statistics interface{} `json:"statistics"`
}

// TODO: EXAMPLE
//
type CfnJob_CsvOutputOptionsProperty struct {
	// `CfnJob.CsvOutputOptionsProperty.Delimiter`.
	Delimiter *string `json:"delimiter"`
}

// TODO: EXAMPLE
//
type CfnJob_DataCatalogOutputProperty struct {
	// `CfnJob.DataCatalogOutputProperty.CatalogId`.
	CatalogId *string `json:"catalogId"`
	// `CfnJob.DataCatalogOutputProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnJob.DataCatalogOutputProperty.DatabaseOptions`.
	DatabaseOptions interface{} `json:"databaseOptions"`
	// `CfnJob.DataCatalogOutputProperty.Overwrite`.
	Overwrite interface{} `json:"overwrite"`
	// `CfnJob.DataCatalogOutputProperty.S3Options`.
	S3Options interface{} `json:"s3Options"`
	// `CfnJob.DataCatalogOutputProperty.TableName`.
	TableName *string `json:"tableName"`
}

// TODO: EXAMPLE
//
type CfnJob_DatabaseOutputProperty struct {
	// `CfnJob.DatabaseOutputProperty.DatabaseOptions`.
	DatabaseOptions interface{} `json:"databaseOptions"`
	// `CfnJob.DatabaseOutputProperty.DatabaseOutputMode`.
	DatabaseOutputMode *string `json:"databaseOutputMode"`
	// `CfnJob.DatabaseOutputProperty.GlueConnectionName`.
	GlueConnectionName *string `json:"glueConnectionName"`
}

// TODO: EXAMPLE
//
type CfnJob_DatabaseTableOutputOptionsProperty struct {
	// `CfnJob.DatabaseTableOutputOptionsProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnJob.DatabaseTableOutputOptionsProperty.TempDirectory`.
	TempDirectory interface{} `json:"tempDirectory"`
}

// TODO: EXAMPLE
//
type CfnJob_EntityDetectorConfigurationProperty struct {
	// `CfnJob.EntityDetectorConfigurationProperty.AllowedStatistics`.
	AllowedStatistics interface{} `json:"allowedStatistics"`
	// `CfnJob.EntityDetectorConfigurationProperty.EntityTypes`.
	EntityTypes *[]*string `json:"entityTypes"`
}

// TODO: EXAMPLE
//
type CfnJob_JobSampleProperty struct {
	// `CfnJob.JobSampleProperty.Mode`.
	Mode *string `json:"mode"`
	// `CfnJob.JobSampleProperty.Size`.
	Size *float64 `json:"size"`
}

// TODO: EXAMPLE
//
type CfnJob_OutputFormatOptionsProperty struct {
	// `CfnJob.OutputFormatOptionsProperty.Csv`.
	Csv interface{} `json:"csv"`
}

// TODO: EXAMPLE
//
type CfnJob_OutputLocationProperty struct {
	// `CfnJob.OutputLocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnJob.OutputLocationProperty.Key`.
	Key *string `json:"key"`
}

// TODO: EXAMPLE
//
type CfnJob_OutputProperty struct {
	// `CfnJob.OutputProperty.CompressionFormat`.
	CompressionFormat *string `json:"compressionFormat"`
	// `CfnJob.OutputProperty.Format`.
	Format *string `json:"format"`
	// `CfnJob.OutputProperty.FormatOptions`.
	FormatOptions interface{} `json:"formatOptions"`
	// `CfnJob.OutputProperty.Location`.
	Location interface{} `json:"location"`
	// `CfnJob.OutputProperty.Overwrite`.
	Overwrite interface{} `json:"overwrite"`
	// `CfnJob.OutputProperty.PartitionColumns`.
	PartitionColumns *[]*string `json:"partitionColumns"`
}

// TODO: EXAMPLE
//
type CfnJob_ParameterMapProperty struct {
}

// TODO: EXAMPLE
//
type CfnJob_ProfileConfigurationProperty struct {
	// `CfnJob.ProfileConfigurationProperty.ColumnStatisticsConfigurations`.
	ColumnStatisticsConfigurations interface{} `json:"columnStatisticsConfigurations"`
	// `CfnJob.ProfileConfigurationProperty.DatasetStatisticsConfiguration`.
	DatasetStatisticsConfiguration interface{} `json:"datasetStatisticsConfiguration"`
	// `CfnJob.ProfileConfigurationProperty.EntityDetectorConfiguration`.
	EntityDetectorConfiguration interface{} `json:"entityDetectorConfiguration"`
	// `CfnJob.ProfileConfigurationProperty.ProfileColumns`.
	ProfileColumns interface{} `json:"profileColumns"`
}

// TODO: EXAMPLE
//
type CfnJob_RecipeProperty struct {
	// `CfnJob.RecipeProperty.Name`.
	Name *string `json:"name"`
	// `CfnJob.RecipeProperty.Version`.
	Version *string `json:"version"`
}

// TODO: EXAMPLE
//
type CfnJob_S3LocationProperty struct {
	// `CfnJob.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnJob.S3LocationProperty.Key`.
	Key *string `json:"key"`
}

// TODO: EXAMPLE
//
type CfnJob_S3TableOutputOptionsProperty struct {
	// `CfnJob.S3TableOutputOptionsProperty.Location`.
	Location interface{} `json:"location"`
}

// TODO: EXAMPLE
//
type CfnJob_StatisticOverrideProperty struct {
	// `CfnJob.StatisticOverrideProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `CfnJob.StatisticOverrideProperty.Statistic`.
	Statistic *string `json:"statistic"`
}

// TODO: EXAMPLE
//
type CfnJob_StatisticsConfigurationProperty struct {
	// `CfnJob.StatisticsConfigurationProperty.IncludedStatistics`.
	IncludedStatistics *[]*string `json:"includedStatistics"`
	// `CfnJob.StatisticsConfigurationProperty.Overrides`.
	Overrides interface{} `json:"overrides"`
}

// TODO: EXAMPLE
//
type CfnJob_ValidationConfigurationProperty struct {
	// `CfnJob.ValidationConfigurationProperty.RulesetArn`.
	RulesetArn *string `json:"rulesetArn"`
	// `CfnJob.ValidationConfigurationProperty.ValidationMode`.
	ValidationMode *string `json:"validationMode"`
}

// Properties for defining a `AWS::DataBrew::Job`.
//
// TODO: EXAMPLE
//
type CfnJobProps struct {
	// `AWS::DataBrew::Job.DatabaseOutputs`.
	DatabaseOutputs interface{} `json:"databaseOutputs"`
	// `AWS::DataBrew::Job.DataCatalogOutputs`.
	DataCatalogOutputs interface{} `json:"dataCatalogOutputs"`
	// `AWS::DataBrew::Job.DatasetName`.
	DatasetName *string `json:"datasetName"`
	// `AWS::DataBrew::Job.EncryptionKeyArn`.
	EncryptionKeyArn *string `json:"encryptionKeyArn"`
	// `AWS::DataBrew::Job.EncryptionMode`.
	EncryptionMode *string `json:"encryptionMode"`
	// `AWS::DataBrew::Job.JobSample`.
	JobSample interface{} `json:"jobSample"`
	// `AWS::DataBrew::Job.LogSubscription`.
	LogSubscription *string `json:"logSubscription"`
	// `AWS::DataBrew::Job.MaxCapacity`.
	MaxCapacity *float64 `json:"maxCapacity"`
	// `AWS::DataBrew::Job.MaxRetries`.
	MaxRetries *float64 `json:"maxRetries"`
	// `AWS::DataBrew::Job.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Job.OutputLocation`.
	OutputLocation interface{} `json:"outputLocation"`
	// `AWS::DataBrew::Job.Outputs`.
	Outputs interface{} `json:"outputs"`
	// `AWS::DataBrew::Job.ProfileConfiguration`.
	ProfileConfiguration interface{} `json:"profileConfiguration"`
	// `AWS::DataBrew::Job.ProjectName`.
	ProjectName *string `json:"projectName"`
	// `AWS::DataBrew::Job.Recipe`.
	Recipe interface{} `json:"recipe"`
	// `AWS::DataBrew::Job.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::DataBrew::Job.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DataBrew::Job.Timeout`.
	Timeout *float64 `json:"timeout"`
	// `AWS::DataBrew::Job.Type`.
	Type *string `json:"type"`
	// `AWS::DataBrew::Job.ValidationConfigurations`.
	ValidationConfigurations interface{} `json:"validationConfigurations"`
}

// A CloudFormation `AWS::DataBrew::Project`.
//
// TODO: EXAMPLE
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatasetName() *string
	SetDatasetName(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	RecipeName() *string
	SetRecipeName(val *string)
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Sample() interface{}
	SetSample(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) RecipeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Sample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Project`.
func NewCfnProject(scope constructs.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Project`.
func NewCfnProject_Override(c CfnProject, scope constructs.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetRecipeName(val *string) {
	_jsii_.Set(
		j,
		"recipeName",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSample(val interface{}) {
	_jsii_.Set(
		j,
		"sample",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnProject_SampleProperty struct {
	// `CfnProject.SampleProperty.Size`.
	Size *float64 `json:"size"`
	// `CfnProject.SampleProperty.Type`.
	Type *string `json:"type"`
}

// Properties for defining a `AWS::DataBrew::Project`.
//
// TODO: EXAMPLE
//
type CfnProjectProps struct {
	// `AWS::DataBrew::Project.DatasetName`.
	DatasetName *string `json:"datasetName"`
	// `AWS::DataBrew::Project.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Project.RecipeName`.
	RecipeName *string `json:"recipeName"`
	// `AWS::DataBrew::Project.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::DataBrew::Project.Sample`.
	Sample interface{} `json:"sample"`
	// `AWS::DataBrew::Project.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::DataBrew::Recipe`.
//
// TODO: EXAMPLE
//
type CfnRecipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Steps() interface{}
	SetSteps(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRecipe
type jsiiProxy_CfnRecipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRecipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Steps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Recipe`.
func NewCfnRecipe(scope constructs.Construct, id *string, props *CfnRecipeProps) CfnRecipe {
	_init_.Initialize()

	j := jsiiProxy_CfnRecipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Recipe`.
func NewCfnRecipe_Override(c CfnRecipe, scope constructs.Construct, id *string, props *CfnRecipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecipe) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRecipe) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRecipe) SetSteps(val interface{}) {
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRecipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRecipe_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRecipe) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnRecipe) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecipe) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnRecipe) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRecipe) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRecipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnRecipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnRecipe) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRecipe) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRecipe) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRecipe) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnRecipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnRecipe_ActionProperty struct {
	// `CfnRecipe.ActionProperty.Operation`.
	Operation *string `json:"operation"`
	// `CfnRecipe.ActionProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
}

// TODO: EXAMPLE
//
type CfnRecipe_ConditionExpressionProperty struct {
	// `CfnRecipe.ConditionExpressionProperty.Condition`.
	Condition *string `json:"condition"`
	// `CfnRecipe.ConditionExpressionProperty.TargetColumn`.
	TargetColumn *string `json:"targetColumn"`
	// `CfnRecipe.ConditionExpressionProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnRecipe_DataCatalogInputDefinitionProperty struct {
	// `CfnRecipe.DataCatalogInputDefinitionProperty.CatalogId`.
	CatalogId *string `json:"catalogId"`
	// `CfnRecipe.DataCatalogInputDefinitionProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnRecipe.DataCatalogInputDefinitionProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnRecipe.DataCatalogInputDefinitionProperty.TempDirectory`.
	TempDirectory interface{} `json:"tempDirectory"`
}

// TODO: EXAMPLE
//
type CfnRecipe_ParameterMapProperty struct {
}

// TODO: EXAMPLE
//
type CfnRecipe_RecipeParametersProperty struct {
	// `CfnRecipe.RecipeParametersProperty.AggregateFunction`.
	AggregateFunction *string `json:"aggregateFunction"`
	// `CfnRecipe.RecipeParametersProperty.Base`.
	Base *string `json:"base"`
	// `CfnRecipe.RecipeParametersProperty.CaseStatement`.
	CaseStatement *string `json:"caseStatement"`
	// `CfnRecipe.RecipeParametersProperty.CategoryMap`.
	CategoryMap *string `json:"categoryMap"`
	// `CfnRecipe.RecipeParametersProperty.CharsToRemove`.
	CharsToRemove *string `json:"charsToRemove"`
	// `CfnRecipe.RecipeParametersProperty.CollapseConsecutiveWhitespace`.
	CollapseConsecutiveWhitespace *string `json:"collapseConsecutiveWhitespace"`
	// `CfnRecipe.RecipeParametersProperty.ColumnDataType`.
	ColumnDataType *string `json:"columnDataType"`
	// `CfnRecipe.RecipeParametersProperty.ColumnRange`.
	ColumnRange *string `json:"columnRange"`
	// `CfnRecipe.RecipeParametersProperty.Count`.
	Count *string `json:"count"`
	// `CfnRecipe.RecipeParametersProperty.CustomCharacters`.
	CustomCharacters *string `json:"customCharacters"`
	// `CfnRecipe.RecipeParametersProperty.CustomStopWords`.
	CustomStopWords *string `json:"customStopWords"`
	// `CfnRecipe.RecipeParametersProperty.CustomValue`.
	CustomValue *string `json:"customValue"`
	// `CfnRecipe.RecipeParametersProperty.DatasetsColumns`.
	DatasetsColumns *string `json:"datasetsColumns"`
	// `CfnRecipe.RecipeParametersProperty.DateAddValue`.
	DateAddValue *string `json:"dateAddValue"`
	// `CfnRecipe.RecipeParametersProperty.DateTimeFormat`.
	DateTimeFormat *string `json:"dateTimeFormat"`
	// `CfnRecipe.RecipeParametersProperty.DateTimeParameters`.
	DateTimeParameters *string `json:"dateTimeParameters"`
	// `CfnRecipe.RecipeParametersProperty.DeleteOtherRows`.
	DeleteOtherRows *string `json:"deleteOtherRows"`
	// `CfnRecipe.RecipeParametersProperty.Delimiter`.
	Delimiter *string `json:"delimiter"`
	// `CfnRecipe.RecipeParametersProperty.EndPattern`.
	EndPattern *string `json:"endPattern"`
	// `CfnRecipe.RecipeParametersProperty.EndPosition`.
	EndPosition *string `json:"endPosition"`
	// `CfnRecipe.RecipeParametersProperty.EndValue`.
	EndValue *string `json:"endValue"`
	// `CfnRecipe.RecipeParametersProperty.ExpandContractions`.
	ExpandContractions *string `json:"expandContractions"`
	// `CfnRecipe.RecipeParametersProperty.Exponent`.
	Exponent *string `json:"exponent"`
	// `CfnRecipe.RecipeParametersProperty.FalseString`.
	FalseString *string `json:"falseString"`
	// `CfnRecipe.RecipeParametersProperty.GroupByAggFunctionOptions`.
	GroupByAggFunctionOptions *string `json:"groupByAggFunctionOptions"`
	// `CfnRecipe.RecipeParametersProperty.GroupByColumns`.
	GroupByColumns *string `json:"groupByColumns"`
	// `CfnRecipe.RecipeParametersProperty.HiddenColumns`.
	HiddenColumns *string `json:"hiddenColumns"`
	// `CfnRecipe.RecipeParametersProperty.IgnoreCase`.
	IgnoreCase *string `json:"ignoreCase"`
	// `CfnRecipe.RecipeParametersProperty.IncludeInSplit`.
	IncludeInSplit *string `json:"includeInSplit"`
	// `CfnRecipe.RecipeParametersProperty.Input`.
	Input interface{} `json:"input"`
	// `CfnRecipe.RecipeParametersProperty.Interval`.
	Interval *string `json:"interval"`
	// `CfnRecipe.RecipeParametersProperty.IsText`.
	IsText *string `json:"isText"`
	// `CfnRecipe.RecipeParametersProperty.JoinKeys`.
	JoinKeys *string `json:"joinKeys"`
	// `CfnRecipe.RecipeParametersProperty.JoinType`.
	JoinType *string `json:"joinType"`
	// `CfnRecipe.RecipeParametersProperty.LeftColumns`.
	LeftColumns *string `json:"leftColumns"`
	// `CfnRecipe.RecipeParametersProperty.Limit`.
	Limit *string `json:"limit"`
	// `CfnRecipe.RecipeParametersProperty.LowerBound`.
	LowerBound *string `json:"lowerBound"`
	// `CfnRecipe.RecipeParametersProperty.MapType`.
	MapType *string `json:"mapType"`
	// `CfnRecipe.RecipeParametersProperty.ModeType`.
	ModeType *string `json:"modeType"`
	// `CfnRecipe.RecipeParametersProperty.MultiLine`.
	MultiLine interface{} `json:"multiLine"`
	// `CfnRecipe.RecipeParametersProperty.NumRows`.
	NumRows *string `json:"numRows"`
	// `CfnRecipe.RecipeParametersProperty.NumRowsAfter`.
	NumRowsAfter *string `json:"numRowsAfter"`
	// `CfnRecipe.RecipeParametersProperty.NumRowsBefore`.
	NumRowsBefore *string `json:"numRowsBefore"`
	// `CfnRecipe.RecipeParametersProperty.OrderByColumn`.
	OrderByColumn *string `json:"orderByColumn"`
	// `CfnRecipe.RecipeParametersProperty.OrderByColumns`.
	OrderByColumns *string `json:"orderByColumns"`
	// `CfnRecipe.RecipeParametersProperty.Other`.
	Other *string `json:"other"`
	// `CfnRecipe.RecipeParametersProperty.Pattern`.
	Pattern *string `json:"pattern"`
	// `CfnRecipe.RecipeParametersProperty.PatternOption1`.
	PatternOption1 *string `json:"patternOption1"`
	// `CfnRecipe.RecipeParametersProperty.PatternOption2`.
	PatternOption2 *string `json:"patternOption2"`
	// `CfnRecipe.RecipeParametersProperty.PatternOptions`.
	PatternOptions *string `json:"patternOptions"`
	// `CfnRecipe.RecipeParametersProperty.Period`.
	Period *string `json:"period"`
	// `CfnRecipe.RecipeParametersProperty.Position`.
	Position *string `json:"position"`
	// `CfnRecipe.RecipeParametersProperty.RemoveAllPunctuation`.
	RemoveAllPunctuation *string `json:"removeAllPunctuation"`
	// `CfnRecipe.RecipeParametersProperty.RemoveAllQuotes`.
	RemoveAllQuotes *string `json:"removeAllQuotes"`
	// `CfnRecipe.RecipeParametersProperty.RemoveAllWhitespace`.
	RemoveAllWhitespace *string `json:"removeAllWhitespace"`
	// `CfnRecipe.RecipeParametersProperty.RemoveCustomCharacters`.
	RemoveCustomCharacters *string `json:"removeCustomCharacters"`
	// `CfnRecipe.RecipeParametersProperty.RemoveCustomValue`.
	RemoveCustomValue *string `json:"removeCustomValue"`
	// `CfnRecipe.RecipeParametersProperty.RemoveLeadingAndTrailingPunctuation`.
	RemoveLeadingAndTrailingPunctuation *string `json:"removeLeadingAndTrailingPunctuation"`
	// `CfnRecipe.RecipeParametersProperty.RemoveLeadingAndTrailingQuotes`.
	RemoveLeadingAndTrailingQuotes *string `json:"removeLeadingAndTrailingQuotes"`
	// `CfnRecipe.RecipeParametersProperty.RemoveLeadingAndTrailingWhitespace`.
	RemoveLeadingAndTrailingWhitespace *string `json:"removeLeadingAndTrailingWhitespace"`
	// `CfnRecipe.RecipeParametersProperty.RemoveLetters`.
	RemoveLetters *string `json:"removeLetters"`
	// `CfnRecipe.RecipeParametersProperty.RemoveNumbers`.
	RemoveNumbers *string `json:"removeNumbers"`
	// `CfnRecipe.RecipeParametersProperty.RemoveSourceColumn`.
	RemoveSourceColumn *string `json:"removeSourceColumn"`
	// `CfnRecipe.RecipeParametersProperty.RemoveSpecialCharacters`.
	RemoveSpecialCharacters *string `json:"removeSpecialCharacters"`
	// `CfnRecipe.RecipeParametersProperty.RightColumns`.
	RightColumns *string `json:"rightColumns"`
	// `CfnRecipe.RecipeParametersProperty.SampleSize`.
	SampleSize *string `json:"sampleSize"`
	// `CfnRecipe.RecipeParametersProperty.SampleType`.
	SampleType *string `json:"sampleType"`
	// `CfnRecipe.RecipeParametersProperty.SecondaryInputs`.
	SecondaryInputs interface{} `json:"secondaryInputs"`
	// `CfnRecipe.RecipeParametersProperty.SecondInput`.
	SecondInput *string `json:"secondInput"`
	// `CfnRecipe.RecipeParametersProperty.SheetIndexes`.
	SheetIndexes interface{} `json:"sheetIndexes"`
	// `CfnRecipe.RecipeParametersProperty.SheetNames`.
	SheetNames *[]*string `json:"sheetNames"`
	// `CfnRecipe.RecipeParametersProperty.SourceColumn`.
	SourceColumn *string `json:"sourceColumn"`
	// `CfnRecipe.RecipeParametersProperty.SourceColumn1`.
	SourceColumn1 *string `json:"sourceColumn1"`
	// `CfnRecipe.RecipeParametersProperty.SourceColumn2`.
	SourceColumn2 *string `json:"sourceColumn2"`
	// `CfnRecipe.RecipeParametersProperty.SourceColumns`.
	SourceColumns *string `json:"sourceColumns"`
	// `CfnRecipe.RecipeParametersProperty.StartColumnIndex`.
	StartColumnIndex *string `json:"startColumnIndex"`
	// `CfnRecipe.RecipeParametersProperty.StartPattern`.
	StartPattern *string `json:"startPattern"`
	// `CfnRecipe.RecipeParametersProperty.StartPosition`.
	StartPosition *string `json:"startPosition"`
	// `CfnRecipe.RecipeParametersProperty.StartValue`.
	StartValue *string `json:"startValue"`
	// `CfnRecipe.RecipeParametersProperty.StemmingMode`.
	StemmingMode *string `json:"stemmingMode"`
	// `CfnRecipe.RecipeParametersProperty.StepCount`.
	StepCount *string `json:"stepCount"`
	// `CfnRecipe.RecipeParametersProperty.StepIndex`.
	StepIndex *string `json:"stepIndex"`
	// `CfnRecipe.RecipeParametersProperty.StopWordsMode`.
	StopWordsMode *string `json:"stopWordsMode"`
	// `CfnRecipe.RecipeParametersProperty.Strategy`.
	Strategy *string `json:"strategy"`
	// `CfnRecipe.RecipeParametersProperty.TargetColumn`.
	TargetColumn *string `json:"targetColumn"`
	// `CfnRecipe.RecipeParametersProperty.TargetColumnNames`.
	TargetColumnNames *string `json:"targetColumnNames"`
	// `CfnRecipe.RecipeParametersProperty.TargetDateFormat`.
	TargetDateFormat *string `json:"targetDateFormat"`
	// `CfnRecipe.RecipeParametersProperty.TargetIndex`.
	TargetIndex *string `json:"targetIndex"`
	// `CfnRecipe.RecipeParametersProperty.TimeZone`.
	TimeZone *string `json:"timeZone"`
	// `CfnRecipe.RecipeParametersProperty.TokenizerPattern`.
	TokenizerPattern *string `json:"tokenizerPattern"`
	// `CfnRecipe.RecipeParametersProperty.TrueString`.
	TrueString *string `json:"trueString"`
	// `CfnRecipe.RecipeParametersProperty.UdfLang`.
	UdfLang *string `json:"udfLang"`
	// `CfnRecipe.RecipeParametersProperty.Units`.
	Units *string `json:"units"`
	// `CfnRecipe.RecipeParametersProperty.UnpivotColumn`.
	UnpivotColumn *string `json:"unpivotColumn"`
	// `CfnRecipe.RecipeParametersProperty.UpperBound`.
	UpperBound *string `json:"upperBound"`
	// `CfnRecipe.RecipeParametersProperty.UseNewDataFrame`.
	UseNewDataFrame *string `json:"useNewDataFrame"`
	// `CfnRecipe.RecipeParametersProperty.Value`.
	Value *string `json:"value"`
	// `CfnRecipe.RecipeParametersProperty.Value1`.
	Value1 *string `json:"value1"`
	// `CfnRecipe.RecipeParametersProperty.Value2`.
	Value2 *string `json:"value2"`
	// `CfnRecipe.RecipeParametersProperty.ValueColumn`.
	ValueColumn *string `json:"valueColumn"`
	// `CfnRecipe.RecipeParametersProperty.ViewFrame`.
	ViewFrame *string `json:"viewFrame"`
}

// TODO: EXAMPLE
//
type CfnRecipe_RecipeStepProperty struct {
	// `CfnRecipe.RecipeStepProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnRecipe.RecipeStepProperty.ConditionExpressions`.
	ConditionExpressions interface{} `json:"conditionExpressions"`
}

// TODO: EXAMPLE
//
type CfnRecipe_S3LocationProperty struct {
	// `CfnRecipe.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnRecipe.S3LocationProperty.Key`.
	Key *string `json:"key"`
}

// TODO: EXAMPLE
//
type CfnRecipe_SecondaryInputProperty struct {
	// `CfnRecipe.SecondaryInputProperty.DataCatalogInputDefinition`.
	DataCatalogInputDefinition interface{} `json:"dataCatalogInputDefinition"`
	// `CfnRecipe.SecondaryInputProperty.S3InputDefinition`.
	S3InputDefinition interface{} `json:"s3InputDefinition"`
}

// Properties for defining a `AWS::DataBrew::Recipe`.
//
// TODO: EXAMPLE
//
type CfnRecipeProps struct {
	// `AWS::DataBrew::Recipe.Description`.
	Description *string `json:"description"`
	// `AWS::DataBrew::Recipe.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Recipe.Steps`.
	Steps interface{} `json:"steps"`
	// `AWS::DataBrew::Recipe.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::DataBrew::Ruleset`.
//
// TODO: EXAMPLE
//
type CfnRuleset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Rules() interface{}
	SetRules(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TargetArn() *string
	SetTargetArn(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRuleset
type jsiiProxy_CfnRuleset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRuleset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Ruleset`.
func NewCfnRuleset(scope constructs.Construct, id *string, props *CfnRulesetProps) CfnRuleset {
	_init_.Initialize()

	j := jsiiProxy_CfnRuleset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Ruleset`.
func NewCfnRuleset_Override(c CfnRuleset, scope constructs.Construct, id *string, props *CfnRulesetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRuleset) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetTargetArn(val *string) {
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRuleset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRuleset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRuleset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRuleset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnRuleset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRuleset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnRuleset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRuleset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRuleset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnRuleset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnRuleset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRuleset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRuleset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRuleset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRuleset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnRuleset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRuleset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnRuleset_ColumnSelectorProperty struct {
	// `CfnRuleset.ColumnSelectorProperty.Name`.
	Name *string `json:"name"`
	// `CfnRuleset.ColumnSelectorProperty.Regex`.
	Regex *string `json:"regex"`
}

// TODO: EXAMPLE
//
type CfnRuleset_RuleProperty struct {
	// `CfnRuleset.RuleProperty.CheckExpression`.
	CheckExpression *string `json:"checkExpression"`
	// `CfnRuleset.RuleProperty.ColumnSelectors`.
	ColumnSelectors interface{} `json:"columnSelectors"`
	// `CfnRuleset.RuleProperty.Disabled`.
	Disabled interface{} `json:"disabled"`
	// `CfnRuleset.RuleProperty.Name`.
	Name *string `json:"name"`
	// `CfnRuleset.RuleProperty.SubstitutionMap`.
	SubstitutionMap interface{} `json:"substitutionMap"`
	// `CfnRuleset.RuleProperty.Threshold`.
	Threshold interface{} `json:"threshold"`
}

// TODO: EXAMPLE
//
type CfnRuleset_SubstitutionValueProperty struct {
	// `CfnRuleset.SubstitutionValueProperty.Value`.
	Value *string `json:"value"`
	// `CfnRuleset.SubstitutionValueProperty.ValueReference`.
	ValueReference *string `json:"valueReference"`
}

// TODO: EXAMPLE
//
type CfnRuleset_ThresholdProperty struct {
	// `CfnRuleset.ThresholdProperty.Type`.
	Type *string `json:"type"`
	// `CfnRuleset.ThresholdProperty.Unit`.
	Unit *string `json:"unit"`
	// `CfnRuleset.ThresholdProperty.Value`.
	Value *float64 `json:"value"`
}

// Properties for defining a `AWS::DataBrew::Ruleset`.
//
// TODO: EXAMPLE
//
type CfnRulesetProps struct {
	// `AWS::DataBrew::Ruleset.Description`.
	Description *string `json:"description"`
	// `AWS::DataBrew::Ruleset.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Ruleset.Rules`.
	Rules interface{} `json:"rules"`
	// `AWS::DataBrew::Ruleset.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DataBrew::Ruleset.TargetArn`.
	TargetArn *string `json:"targetArn"`
}

// A CloudFormation `AWS::DataBrew::Schedule`.
//
// TODO: EXAMPLE
//
type CfnSchedule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CronExpression() *string
	SetCronExpression(val *string)
	JobNames() *[]*string
	SetJobNames(val *[]*string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSchedule
type jsiiProxy_CfnSchedule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchedule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) JobNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jobNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Schedule`.
func NewCfnSchedule(scope constructs.Construct, id *string, props *CfnScheduleProps) CfnSchedule {
	_init_.Initialize()

	j := jsiiProxy_CfnSchedule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Schedule`.
func NewCfnSchedule_Override(c CfnSchedule, scope constructs.Construct, id *string, props *CfnScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchedule) SetCronExpression(val *string) {
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_CfnSchedule) SetJobNames(val *[]*string) {
	_jsii_.Set(
		j,
		"jobNames",
		val,
	)
}

func (j *jsiiProxy_CfnSchedule) SetName(val *string) {
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
func CfnSchedule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchedule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchedule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSchedule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnSchedule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchedule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnSchedule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSchedule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSchedule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnSchedule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchedule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSchedule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSchedule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchedule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnSchedule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::DataBrew::Schedule`.
//
// TODO: EXAMPLE
//
type CfnScheduleProps struct {
	// `AWS::DataBrew::Schedule.CronExpression`.
	CronExpression *string `json:"cronExpression"`
	// `AWS::DataBrew::Schedule.JobNames`.
	JobNames *[]*string `json:"jobNames"`
	// `AWS::DataBrew::Schedule.Name`.
	Name *string `json:"name"`
	// `AWS::DataBrew::Schedule.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

