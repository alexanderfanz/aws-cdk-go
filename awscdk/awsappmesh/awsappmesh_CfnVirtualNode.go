package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppMesh::VirtualNode`.
//
// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment. When you create a virtual node, you can specify the service discovery information for your task group, and whether the proxy running in a task group will communicate with other proxies using Transport Layer Security (TLS).
//
// You define a `listener` for any inbound traffic that your virtual node expects. Any virtual service that your virtual node expects to communicate to is specified as a `backend` .
//
// The response metadata for your new virtual node contains the `arn` that is associated with the virtual node. Set this value to the full ARN; for example, `arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp` ) as the `APPMESH_RESOURCE_ARN` environment variable for your task group's Envoy proxy container in your task definition or pod spec. This is then mapped to the `node.id` and `node.cluster` Envoy parameters.
//
// > By default, App Mesh uses the name of the resource you specified in `APPMESH_RESOURCE_ARN` when Envoy is referring to itself in metrics and traces. You can override this behavior by setting the `APPMESH_RESOURCE_CLUSTER` environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html) . You must be using `1.15.0` or later of the Envoy image when setting these variables. For more information about App Mesh Envoy variables, see [Envoy image](https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html) in the AWS App Mesh User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualNode := awscdk.Aws_appmesh.NewCfnVirtualNode(this, jsii.String("MyCfnVirtualNode"), &cfnVirtualNodeProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &virtualNodeSpecProperty{
//   		backendDefaults: &backendDefaultsProperty{
//   			clientPolicy: &clientPolicyProperty{
//   				tls: &clientPolicyTlsProperty{
//   					validation: &tlsValidationContextProperty{
//   						trust: &tlsValidationContextTrustProperty{
//   							acm: &tlsValidationContextAcmTrustProperty{
//   								certificateAuthorityArns: []*string{
//   									jsii.String("certificateAuthorityArns"),
//   								},
//   							},
//   							file: &tlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &tlsValidationContextSdsTrustProperty{
//   								secretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   							match: &subjectAlternativeNameMatchersProperty{
//   								exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					certificate: &clientTlsCertificateProperty{
//   						file: &listenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &listenerTlsSdsCertificateProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//   					enforce: jsii.Boolean(false),
//   					ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		backends: []interface{}{
//   			&backendProperty{
//   				virtualService: &virtualServiceBackendProperty{
//   					virtualServiceName: jsii.String("virtualServiceName"),
//
//   					// the properties below are optional
//   					clientPolicy: &clientPolicyProperty{
//   						tls: &clientPolicyTlsProperty{
//   							validation: &tlsValidationContextProperty{
//   								trust: &tlsValidationContextTrustProperty{
//   									acm: &tlsValidationContextAcmTrustProperty{
//   										certificateAuthorityArns: []*string{
//   											jsii.String("certificateAuthorityArns"),
//   										},
//   									},
//   									file: &tlsValidationContextFileTrustProperty{
//   										certificateChain: jsii.String("certificateChain"),
//   									},
//   									sds: &tlsValidationContextSdsTrustProperty{
//   										secretName: jsii.String("secretName"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   									match: &subjectAlternativeNameMatchersProperty{
//   										exact: []*string{
//   											jsii.String("exact"),
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							certificate: &clientTlsCertificateProperty{
//   								file: &listenerTlsFileCertificateProperty{
//   									certificateChain: jsii.String("certificateChain"),
//   									privateKey: jsii.String("privateKey"),
//   								},
//   								sds: &listenerTlsSdsCertificateProperty{
//   									secretName: jsii.String("secretName"),
//   								},
//   							},
//   							enforce: jsii.Boolean(false),
//   							ports: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		listeners: []interface{}{
//   			&listenerProperty{
//   				portMapping: &portMappingProperty{
//   					port: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   				},
//
//   				// the properties below are optional
//   				connectionPool: &virtualNodeConnectionPoolProperty{
//   					grpc: &virtualNodeGrpcConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   					http: &virtualNodeHttpConnectionPoolProperty{
//   						maxConnections: jsii.Number(123),
//
//   						// the properties below are optional
//   						maxPendingRequests: jsii.Number(123),
//   					},
//   					http2: &virtualNodeHttp2ConnectionPoolProperty{
//   						maxRequests: jsii.Number(123),
//   					},
//   					tcp: &virtualNodeTcpConnectionPoolProperty{
//   						maxConnections: jsii.Number(123),
//   					},
//   				},
//   				healthCheck: &healthCheckProperty{
//   					healthyThreshold: jsii.Number(123),
//   					intervalMillis: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   					timeoutMillis: jsii.Number(123),
//   					unhealthyThreshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					path: jsii.String("path"),
//   					port: jsii.Number(123),
//   				},
//   				outlierDetection: &outlierDetectionProperty{
//   					baseEjectionDuration: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					interval: &durationProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.Number(123),
//   					},
//   					maxEjectionPercent: jsii.Number(123),
//   					maxServerErrors: jsii.Number(123),
//   				},
//   				timeout: &listenerTimeoutProperty{
//   					grpc: &grpcTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					http: &httpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					http2: &httpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   						perRequest: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   					tcp: &tcpTimeoutProperty{
//   						idle: &durationProperty{
//   							unit: jsii.String("unit"),
//   							value: jsii.Number(123),
//   						},
//   					},
//   				},
//   				tls: &listenerTlsProperty{
//   					certificate: &listenerTlsCertificateProperty{
//   						acm: &listenerTlsAcmCertificateProperty{
//   							certificateArn: jsii.String("certificateArn"),
//   						},
//   						file: &listenerTlsFileCertificateProperty{
//   							certificateChain: jsii.String("certificateChain"),
//   							privateKey: jsii.String("privateKey"),
//   						},
//   						sds: &listenerTlsSdsCertificateProperty{
//   							secretName: jsii.String("secretName"),
//   						},
//   					},
//   					mode: jsii.String("mode"),
//
//   					// the properties below are optional
//   					validation: &listenerTlsValidationContextProperty{
//   						trust: &listenerTlsValidationContextTrustProperty{
//   							file: &tlsValidationContextFileTrustProperty{
//   								certificateChain: jsii.String("certificateChain"),
//   							},
//   							sds: &tlsValidationContextSdsTrustProperty{
//   								secretName: jsii.String("secretName"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						subjectAlternativeNames: &subjectAlternativeNamesProperty{
//   							match: &subjectAlternativeNameMatchersProperty{
//   								exact: []*string{
//   									jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		logging: &loggingProperty{
//   			accessLog: &accessLogProperty{
//   				file: &fileAccessLogProperty{
//   					path: jsii.String("path"),
//   				},
//   			},
//   		},
//   		serviceDiscovery: &serviceDiscoveryProperty{
//   			awsCloudMap: &awsCloudMapServiceDiscoveryProperty{
//   				namespaceName: jsii.String("namespaceName"),
//   				serviceName: jsii.String("serviceName"),
//
//   				// the properties below are optional
//   				attributes: []interface{}{
//   					&awsCloudMapInstanceAttributeProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				ipPreference: jsii.String("ipPreference"),
//   			},
//   			dns: &dnsServiceDiscoveryProperty{
//   				hostname: jsii.String("hostname"),
//
//   				// the properties below are optional
//   				ipPreference: jsii.String("ipPreference"),
//   				responseType: jsii.String("responseType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	meshOwner: jsii.String("meshOwner"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	virtualNodeName: jsii.String("virtualNodeName"),
//   })
//
type CfnVirtualNode interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The full Amazon Resource Name (ARN) for the virtual node.
	AttrArn() *string
	// The name of the service mesh that the virtual node resides in.
	AttrMeshName() *string
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrMeshOwner() *string
	// The AWS IAM account ID of the resource owner.
	//
	// If the account ID is not your own, then it's the ID of the mesh owner or of another account that the mesh is shared with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	AttrResourceOwner() *string
	// The unique identifier for the virtual node.
	AttrUid() *string
	// The name of the virtual node.
	AttrVirtualNodeName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The name of the service mesh to create the virtual node in.
	MeshName() *string
	SetMeshName(val *string)
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner() *string
	SetMeshOwner(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The virtual node specification to apply.
	Spec() interface{}
	SetSpec(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Optional metadata that you can apply to the virtual node to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
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
	// The name to use for the virtual node.
	VirtualNodeName() *string
	SetVirtualNodeName(val *string)
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

// The jsii proxy struct for CfnVirtualNode
type jsiiProxy_CfnVirtualNode struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVirtualNode) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrMeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrMeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMeshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) AttrVirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVirtualNodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) MeshOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Spec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualNode) VirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNode(scope constructs.Construct, id *string, props *CfnVirtualNodeProps) CfnVirtualNode {
	_init_.Initialize()

	j := jsiiProxy_CfnVirtualNode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppMesh::VirtualNode`.
func NewCfnVirtualNode_Override(c CfnVirtualNode, scope constructs.Construct, id *string, props *CfnVirtualNodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetMeshName(val *string) {
	_jsii_.Set(
		j,
		"meshName",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetMeshOwner(val *string) {
	_jsii_.Set(
		j,
		"meshOwner",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetSpec(val interface{}) {
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_CfnVirtualNode) SetVirtualNodeName(val *string) {
	_jsii_.Set(
		j,
		"virtualNodeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVirtualNode_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnVirtualNode_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
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
func CfnVirtualNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualNode_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appmesh.CfnVirtualNode",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualNode) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVirtualNode) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVirtualNode) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVirtualNode) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVirtualNode) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualNode) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
