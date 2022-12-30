package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v3"
)

// This creates a service using the Fargate launch type on an ECS cluster.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
// Experimental.
type FargateService interface {
	BaseService
	IFargateService
	// The details of the AWS Cloud Map service.
	// Experimental.
	CloudmapService() awsservicediscovery.Service
	// Experimental.
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	// Experimental.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	// Experimental.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
	// Experimental.
	Connections() awsec2.Connections
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
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	// Experimental.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	// Experimental.
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	// Experimental.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	// Experimental.
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
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
	// The Amazon Resource Name (ARN) of the service.
	// Experimental.
	ServiceArn() *string
	// The name of the service.
	// Experimental.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	// Experimental.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	// Experimental.
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	// Experimental.
	TaskDefinition() TaskDefinition
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
	// Associates this service with a CloudMap service.
	// Experimental.
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	// This method is called to attach this service to an Application Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	// Experimental.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	// Experimental.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	// Experimental.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
	// Experimental.
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// This method is called to create a networkConfiguration.
	// Deprecated: use configureAwsVpcNetworkingWithSecurityGroups instead.
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
	// This method is called to create a networkConfiguration.
	// Experimental.
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	// Enable CloudMap service discovery for the service.
	//
	// Returns: The created CloudMap service.
	// Experimental.
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
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
	// Return a load balancing target for a specific container and port.
	//
	// Use this function to create a load balancer target if you want to load balance to
	// another container than the first essential container or the first mapped port on
	// the container.
	//
	// Use the return value of this function where you would normally use a load balancer
	// target, instead of the `Service` object itself.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   listener.addTargets('ECS', {
	//     port: 80,
	//     targets: [service.loadBalancerTarget({
	//       containerName: 'MyContainer',
	//       containerPort: 1234,
	//     })],
	//   });
	//
	// Experimental.
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	// Experimental.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	// Experimental.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
	//
	// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   service.registerLoadBalancerTargets(
	//     {
	//       containerName: 'web',
	//       containerPort: 80,
	//       newTargetGroupId: 'ECS',
	//       listener: ecs.ListenerConfig.applicationListener(listener, {
	//         protocol: elbv2.ApplicationProtocol.HTTPS
	//       }),
	//     },
	//   )
	//
	// Experimental.
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
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
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for FargateService
type jsiiProxy_FargateService struct {
	jsiiProxy_BaseService
	jsiiProxy_IFargateService
}

func (j *jsiiProxy_FargateService) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the FargateService class.
// Experimental.
func NewFargateService(scope constructs.Construct, id *string, props *FargateServiceProps) FargateService {
	_init_.Initialize()

	if err := validateNewFargateServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FargateService{}

	_jsii_.Create(
		"monocdk.aws_ecs.FargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FargateService class.
// Experimental.
func NewFargateService_Override(f FargateService, scope constructs.Construct, id *string, props *FargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.FargateService",
		[]interface{}{scope, id, props},
		f,
	)
}

func (j *jsiiProxy_FargateService)SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_FargateService)SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_FargateService)SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_FargateService)SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	if err := j.validateSetServiceRegistriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Imports from the specified service ARN.
// Experimental.
func FargateService_FromFargateServiceArn(scope constructs.Construct, id *string, fargateServiceArn *string) IFargateService {
	_init_.Initialize()

	if err := validateFargateService_FromFargateServiceArnParameters(scope, id, fargateServiceArn); err != nil {
		panic(err)
	}
	var returns IFargateService

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FargateService",
		"fromFargateServiceArn",
		[]interface{}{scope, id, fargateServiceArn},
		&returns,
	)

	return returns
}

// Imports from the specified service attributes.
// Experimental.
func FargateService_FromFargateServiceAttributes(scope constructs.Construct, id *string, attrs *FargateServiceAttributes) IBaseService {
	_init_.Initialize()

	if err := validateFargateService_FromFargateServiceAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IBaseService

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FargateService",
		"fromFargateServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
// Experimental.
func FargateService_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	if err := validateFargateService_FromServiceArnWithClusterParameters(scope, id, serviceArn); err != nil {
		panic(err)
	}
	var returns IBaseService

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FargateService",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func FargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFargateService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FargateService_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFargateService_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FargateService",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FargateService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	if err := f.validateAssociateCloudMapServiceParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

func (f *jsiiProxy_FargateService) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := f.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		f,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	if err := f.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

func (f *jsiiProxy_FargateService) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := f.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		f,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	if err := f.validateAutoScaleTaskCountParameters(props); err != nil {
		panic(err)
	}
	var returns ScalableTaskCount

	_jsii_.Invoke(
		f,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	if err := f.validateConfigureAwsVpcNetworkingParameters(vpc, vpcSubnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroup},
	)
}

func (f *jsiiProxy_FargateService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	if err := f.validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc, vpcSubnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

func (f *jsiiProxy_FargateService) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	if err := f.validateEnableCloudMapParameters(options); err != nil {
		panic(err)
	}
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		f,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	if err := f.validateLoadBalancerTargetParameters(options); err != nil {
		panic(err)
	}
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		f,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateService) OnSynthesize(session constructs.ISynthesisSession) {
	if err := f.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FargateService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateService) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	if err := f.validateRegisterLoadBalancerTargetsParameters(&targets); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"registerLoadBalancerTargets",
		args,
	)
}

func (f *jsiiProxy_FargateService) Synthesize(session awscdk.ISynthesisSession) {
	if err := f.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

