package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::QuickSight::DataSource`.
//
// Creates a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_quicksight.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &cfnDataSourceProps{
//   	alternateDataSourceParameters: []interface{}{
//   		&dataSourceParametersProperty{
//   			amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			athenaParameters: &athenaParametersProperty{
//   				workGroup: jsii.String("workGroup"),
//   			},
//   			auroraParameters: &auroraParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mariaDbParameters: &mariaDbParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mySqlParameters: &mySqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			oracleParameters: &oracleParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			postgreSqlParameters: &postgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			prestoParameters: &prestoParametersProperty{
//   				catalog: jsii.String("catalog"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			rdsParameters: &rdsParametersProperty{
//   				database: jsii.String("database"),
//   				instanceId: jsii.String("instanceId"),
//   			},
//   			redshiftParameters: &redshiftParametersProperty{
//   				database: jsii.String("database"),
//
//   				// the properties below are optional
//   				clusterId: jsii.String("clusterId"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			s3Parameters: &s3ParametersProperty{
//   				manifestFileLocation: &manifestFileLocationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			snowflakeParameters: &snowflakeParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				warehouse: jsii.String("warehouse"),
//   			},
//   			sparkParameters: &sparkParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			sqlServerParameters: &sqlServerParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			teradataParameters: &teradataParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	awsAccountId: jsii.String("awsAccountId"),
//   	credentials: &dataSourceCredentialsProperty{
//   		copySourceArn: jsii.String("copySourceArn"),
//   		credentialPair: &credentialPairProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			alternateDataSourceParameters: []interface{}{
//   				&dataSourceParametersProperty{
//   					amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					athenaParameters: &athenaParametersProperty{
//   						workGroup: jsii.String("workGroup"),
//   					},
//   					auroraParameters: &auroraParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mariaDbParameters: &mariaDbParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mySqlParameters: &mySqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					oracleParameters: &oracleParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					postgreSqlParameters: &postgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					prestoParameters: &prestoParametersProperty{
//   						catalog: jsii.String("catalog"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					rdsParameters: &rdsParametersProperty{
//   						database: jsii.String("database"),
//   						instanceId: jsii.String("instanceId"),
//   					},
//   					redshiftParameters: &redshiftParametersProperty{
//   						database: jsii.String("database"),
//
//   						// the properties below are optional
//   						clusterId: jsii.String("clusterId"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					s3Parameters: &s3ParametersProperty{
//   						manifestFileLocation: &manifestFileLocationProperty{
//   							bucket: jsii.String("bucket"),
//   							key: jsii.String("key"),
//   						},
//   					},
//   					snowflakeParameters: &snowflakeParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						warehouse: jsii.String("warehouse"),
//   					},
//   					sparkParameters: &sparkParametersProperty{
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					sqlServerParameters: &sqlServerParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					teradataParameters: &teradataParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	dataSourceId: jsii.String("dataSourceId"),
//   	dataSourceParameters: &dataSourceParametersProperty{
//   		amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		athenaParameters: &athenaParametersProperty{
//   			workGroup: jsii.String("workGroup"),
//   		},
//   		auroraParameters: &auroraParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mariaDbParameters: &mariaDbParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mySqlParameters: &mySqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		oracleParameters: &oracleParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		postgreSqlParameters: &postgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		prestoParameters: &prestoParametersProperty{
//   			catalog: jsii.String("catalog"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		rdsParameters: &rdsParametersProperty{
//   			database: jsii.String("database"),
//   			instanceId: jsii.String("instanceId"),
//   		},
//   		redshiftParameters: &redshiftParametersProperty{
//   			database: jsii.String("database"),
//
//   			// the properties below are optional
//   			clusterId: jsii.String("clusterId"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		s3Parameters: &s3ParametersProperty{
//   			manifestFileLocation: &manifestFileLocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		snowflakeParameters: &snowflakeParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			warehouse: jsii.String("warehouse"),
//   		},
//   		sparkParameters: &sparkParametersProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		sqlServerParameters: &sqlServerParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		teradataParameters: &teradataParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	errorInfo: &dataSourceErrorInfoProperty{
//   		message: jsii.String("message"),
//   		type: jsii.String("type"),
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	sslProperties: &sslPropertiesProperty{
//   		disableSsl: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	vpcConnectionProperties: &vpcConnectionPropertiesProperty{
//   		vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   })
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	AlternateDataSourceParameters() interface{}
	SetAlternateDataSourceParameters(val interface{})
	// The Amazon Resource Name (ARN) of the dataset.
	AttrArn() *string
	// The time that this data source was created.
	AttrCreatedTime() *string
	// The last time that this data source was updated.
	AttrLastUpdatedTime() *string
	// The HTTP status of the request.
	AttrStatus() *string
	// The AWS account ID.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	Credentials() interface{}
	SetCredentials(val interface{})
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSourceId() *string
	SetDataSourceId(val *string)
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters() interface{}
	SetDataSourceParameters(val interface{})
	// Error information from the last update or the creation of the data source.
	ErrorInfo() interface{}
	SetErrorInfo(val interface{})
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
	// A display name for the data source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions on the data source.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties() interface{}
	SetSslProperties(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	Tags() awscdk.TagManager
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
	Type() *string
	SetType(val *string)
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
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties() interface{}
	SetVpcConnectionProperties(val interface{})
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) AlternateDataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternateDataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Credentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ErrorInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) SslProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) VpcConnectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConnectionProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetAlternateDataSourceParameters(val interface{}) {
	_jsii_.Set(
		j,
		"alternateDataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceId(val *string) {
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceParameters(val interface{}) {
	_jsii_.Set(
		j,
		"dataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetErrorInfo(val interface{}) {
	_jsii_.Set(
		j,
		"errorInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetSslProperties(val interface{}) {
	_jsii_.Set(
		j,
		"sslProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetVpcConnectionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConnectionProperties",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
