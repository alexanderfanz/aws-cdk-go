package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties to define an application listener.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationListenerProps struct {
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// Name of the listener.
	// Experimental.
	Name *string `json:"name"`
	// The port on which the listener listens for requests.
	// Experimental.
	Port *float64 `json:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `json:"protocol"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `json:"sslPolicy"`
}

// An EC2 service running on an ECS cluster fronted by an application load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedEc2Service interface {
	ApplicationLoadBalancedServiceBase
	Certificate() awscertificatemanager.ICertificate
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	Service() awsecs.Ec2Service
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	TaskDefinition() awsecs.Ec2TaskDefinition
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancedEc2Service
type jsiiProxy_ApplicationLoadBalancedEc2Service struct {
	jsiiProxy_ApplicationLoadBalancedServiceBase
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) RedirectListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"redirectListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationLoadBalancedEc2Service class.
// Experimental.
func NewApplicationLoadBalancedEc2Service(scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) ApplicationLoadBalancedEc2Service {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancedEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationLoadBalancedEc2Service class.
// Experimental.
func NewApplicationLoadBalancedEc2Service_Override(a ApplicationLoadBalancedEc2Service, scope constructs.Construct, id *string, props *ApplicationLoadBalancedEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationLoadBalancedEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		a,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ApplicationLoadBalancedEc2Service service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedEc2ServiceProps struct {
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the application load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The application load balancer that will serve traffic to the service.
	//
	// The VPC attribute of a load balancer must be specified for it to be used
	// to create a new service with this pattern.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `json:"loadBalancer"`
	// Name of the load balancer.
	// Experimental.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Determines whether or not the Security Group for the Load Balancer's Listener will be open to all traffic by default.
	// Experimental.
	OpenListener *bool `json:"openListener"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `json:"protocol"`
	// The protocol version to use.
	// Experimental.
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `json:"protocolVersion"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType ApplicationLoadBalancedServiceRecordType `json:"recordType"`
	// Specifies whether the load balancer should redirect traffic on port 80 to port 443 to support HTTP->HTTPS redirects This is only valid if the protocol of the ALB is HTTPS.
	// Experimental.
	RedirectHTTP *bool `json:"redirectHTTP"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `json:"sslPolicy"`
	// The protocol for connections from the load balancer to the ECS tasks.
	//
	// The default target port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).
	// Experimental.
	TargetProtocol awselasticloadbalancingv2.ApplicationProtocol `json:"targetProtocol"`
	// The properties required to create a new task definition.
	//
	// TaskDefinition or TaskImageOptions must be specified, but not both.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under contention, Docker attempts to keep the
	// container memory within the limit. If the container requires more memory,
	// it can consume up to the value specified by the Memory property or all of
	// the available memory on the container instance—whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both..
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `json:"taskDefinition"`
}

// A Fargate service running on an ECS cluster fronted by an application load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedFargateService interface {
	ApplicationLoadBalancedServiceBase
	AssignPublicIp() *bool
	Certificate() awscertificatemanager.ICertificate
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	Service() awsecs.FargateService
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	TaskDefinition() awsecs.FargateTaskDefinition
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancedFargateService
type jsiiProxy_ApplicationLoadBalancedFargateService struct {
	jsiiProxy_ApplicationLoadBalancedServiceBase
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) RedirectListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"redirectListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationLoadBalancedFargateService class.
// Experimental.
func NewApplicationLoadBalancedFargateService(scope constructs.Construct, id *string, props *ApplicationLoadBalancedFargateServiceProps) ApplicationLoadBalancedFargateService {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancedFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationLoadBalancedFargateService class.
// Experimental.
func NewApplicationLoadBalancedFargateService_Override(a ApplicationLoadBalancedFargateService, scope constructs.Construct, id *string, props *ApplicationLoadBalancedFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationLoadBalancedFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedFargateService) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		a,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ApplicationLoadBalancedFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedFargateServiceProps struct {
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the application load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The application load balancer that will serve traffic to the service.
	//
	// The VPC attribute of a load balancer must be specified for it to be used
	// to create a new service with this pattern.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `json:"loadBalancer"`
	// Name of the load balancer.
	// Experimental.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Determines whether or not the Security Group for the Load Balancer's Listener will be open to all traffic by default.
	// Experimental.
	OpenListener *bool `json:"openListener"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `json:"protocol"`
	// The protocol version to use.
	// Experimental.
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `json:"protocolVersion"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType ApplicationLoadBalancedServiceRecordType `json:"recordType"`
	// Specifies whether the load balancer should redirect traffic on port 80 to port 443 to support HTTP->HTTPS redirects This is only valid if the protocol of the ALB is HTTPS.
	// Experimental.
	RedirectHTTP *bool `json:"redirectHTTP"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `json:"sslPolicy"`
	// The protocol for connections from the load balancer to the ECS tasks.
	//
	// The default target port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).
	// Experimental.
	TargetProtocol awselasticloadbalancingv2.ApplicationProtocol `json:"targetProtocol"`
	// The properties required to create a new task definition.
	//
	// TaskDefinition or TaskImageOptions must be specified, but not both.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `json:"taskDefinition"`
	// The subnets to associate with the service.
	// Experimental.
	TaskSubnets *awsec2.SubnetSelection `json:"taskSubnets"`
}

// The base class for ApplicationLoadBalancedEc2Service and ApplicationLoadBalancedFargateService services.
// Experimental.
type ApplicationLoadBalancedServiceBase interface {
	constructs.Construct
	Certificate() awscertificatemanager.ICertificate
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	Node() constructs.Node
	RedirectListener() awselasticloadbalancingv2.ApplicationListener
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancedServiceBase
type jsiiProxy_ApplicationLoadBalancedServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) RedirectListener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"redirectListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancedServiceBase) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationLoadBalancedServiceBase class.
// Experimental.
func NewApplicationLoadBalancedServiceBase_Override(a ApplicationLoadBalancedServiceBase, scope constructs.Construct, id *string, props *ApplicationLoadBalancedServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedServiceBase",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationLoadBalancedServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationLoadBalancedServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		a,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base ApplicationLoadBalancedEc2Service or ApplicationLoadBalancedFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedServiceBaseProps struct {
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the application load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The application load balancer that will serve traffic to the service.
	//
	// The VPC attribute of a load balancer must be specified for it to be used
	// to create a new service with this pattern.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `json:"loadBalancer"`
	// Name of the load balancer.
	// Experimental.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Determines whether or not the Security Group for the Load Balancer's Listener will be open to all traffic by default.
	// Experimental.
	OpenListener *bool `json:"openListener"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `json:"protocol"`
	// The protocol version to use.
	// Experimental.
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `json:"protocolVersion"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType ApplicationLoadBalancedServiceRecordType `json:"recordType"`
	// Specifies whether the load balancer should redirect traffic on port 80 to port 443 to support HTTP->HTTPS redirects This is only valid if the protocol of the ALB is HTTPS.
	// Experimental.
	RedirectHTTP *bool `json:"redirectHTTP"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `json:"sslPolicy"`
	// The protocol for connections from the load balancer to the ECS tasks.
	//
	// The default target port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).
	// Experimental.
	TargetProtocol awselasticloadbalancingv2.ApplicationProtocol `json:"targetProtocol"`
	// The properties required to create a new task definition.
	//
	// TaskDefinition or TaskImageOptions must be specified, but not both.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Describes the type of DNS record the service should create.
// Experimental.
type ApplicationLoadBalancedServiceRecordType string

const (
	ApplicationLoadBalancedServiceRecordType_ALIAS ApplicationLoadBalancedServiceRecordType = "ALIAS"
	ApplicationLoadBalancedServiceRecordType_CNAME ApplicationLoadBalancedServiceRecordType = "CNAME"
	ApplicationLoadBalancedServiceRecordType_NONE ApplicationLoadBalancedServiceRecordType = "NONE"
)

// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedTaskImageOptions struct {
	// The container name value to be specified in the task definition.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	// Experimental.
	ContainerPort *float64 `json:"containerPort"`
	// A key/value map of labels to add to the container.
	// Experimental.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// Options for configuring a new container.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancedTaskImageProps struct {
	// The container name value to be specified in the task definition.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// A list of port numbers on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	// Experimental.
	ContainerPorts *[]*float64 `json:"containerPorts"`
	// A key/value map of labels to add to the container.
	// Experimental.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secrets to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// Properties to define an application load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationLoadBalancerProps struct {
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Listeners (at least one listener) attached to this load balancer.
	// Experimental.
	Listeners *[]*ApplicationListenerProps `json:"listeners"`
	// Name of the load balancer.
	// Experimental.
	Name *string `json:"name"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
}

// An EC2 service running on an ECS cluster fronted by an application load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationMultipleTargetGroupsEc2Service interface {
	ApplicationMultipleTargetGroupsServiceBase
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	Listeners() *[]awselasticloadbalancingv2.ApplicationListener
	SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener)
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	Service() awsecs.Ec2Service
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup)
	TaskDefinition() awsecs.Ec2TaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
	ToString() *string
}

// The jsii proxy struct for ApplicationMultipleTargetGroupsEc2Service
type jsiiProxy_ApplicationMultipleTargetGroupsEc2Service struct {
	jsiiProxy_ApplicationMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) Listeners() *[]awselasticloadbalancingv2.ApplicationListener {
	var returns *[]awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns *[]awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationMultipleTargetGroupsEc2Service class.
// Experimental.
func NewApplicationMultipleTargetGroupsEc2Service(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsEc2ServiceProps) ApplicationMultipleTargetGroupsEc2Service {
	_init_.Initialize()

	j := jsiiProxy_ApplicationMultipleTargetGroupsEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationMultipleTargetGroupsEc2Service class.
// Experimental.
func NewApplicationMultipleTargetGroupsEc2Service_Override(a ApplicationMultipleTargetGroupsEc2Service, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsEc2Service",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationMultipleTargetGroupsEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) {
	_jsii_.InvokeVoid(
		a,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) FindListener(name *string) awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener

	_jsii_.Invoke(
		a,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ApplicationMultipleTargetGroupsEc2Service service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationMultipleTargetGroupsEc2ServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The application load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*ApplicationLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify ALB target groups.
	// Experimental.
	TargetGroups *[]*ApplicationTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The minimum number of CPU units to reserve for the container.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	//
	// Note that this setting will be ignored if TaskImagesOptions is specified
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The task definition to use for tasks in the service. Only one of TaskDefinition or TaskImageOptions must be specified.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `json:"taskDefinition"`
}

// A Fargate service running on an ECS cluster fronted by an application load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationMultipleTargetGroupsFargateService interface {
	ApplicationMultipleTargetGroupsServiceBase
	AssignPublicIp() *bool
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	Listeners() *[]awselasticloadbalancingv2.ApplicationListener
	SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener)
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	Service() awsecs.FargateService
	TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup)
	TaskDefinition() awsecs.FargateTaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
	ToString() *string
}

// The jsii proxy struct for ApplicationMultipleTargetGroupsFargateService
type jsiiProxy_ApplicationMultipleTargetGroupsFargateService struct {
	jsiiProxy_ApplicationMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Listeners() *[]awselasticloadbalancingv2.ApplicationListener {
	var returns *[]awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns *[]awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationMultipleTargetGroupsFargateService class.
// Experimental.
func NewApplicationMultipleTargetGroupsFargateService(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) ApplicationMultipleTargetGroupsFargateService {
	_init_.Initialize()

	j := jsiiProxy_ApplicationMultipleTargetGroupsFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ApplicationMultipleTargetGroupsFargateService class.
// Experimental.
func NewApplicationMultipleTargetGroupsFargateService_Override(a ApplicationMultipleTargetGroupsFargateService, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationMultipleTargetGroupsFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) {
	_jsii_.InvokeVoid(
		a,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) FindListener(name *string) awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener

	_jsii_.Invoke(
		a,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ApplicationMultipleTargetGroupsFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationMultipleTargetGroupsFargateServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The application load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*ApplicationLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify ALB target groups.
	// Experimental.
	TargetGroups *[]*ApplicationTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The task definition to use for tasks in the service. Only one of TaskDefinition or TaskImageOptions must be specified.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `json:"taskDefinition"`
}

// The base class for ApplicationMultipleTargetGroupsEc2Service and ApplicationMultipleTargetGroupsFargateService classes.
// Experimental.
type ApplicationMultipleTargetGroupsServiceBase interface {
	constructs.Construct
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.ApplicationListener
	Listeners() *[]awselasticloadbalancingv2.ApplicationListener
	SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener)
	LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup)
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.ApplicationListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup
	ToString() *string
}

// The jsii proxy struct for ApplicationMultipleTargetGroupsServiceBase
type jsiiProxy_ApplicationMultipleTargetGroupsServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Listener() awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Listeners() *[]awselasticloadbalancingv2.ApplicationListener {
	var returns *[]awselasticloadbalancingv2.ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) LoadBalancer() awselasticloadbalancingv2.ApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.ApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) TargetGroups() *[]awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns *[]awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ApplicationMultipleTargetGroupsServiceBase class.
// Experimental.
func NewApplicationMultipleTargetGroupsServiceBase_Override(a ApplicationMultipleTargetGroupsServiceBase, scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsServiceBase",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) SetListeners(val *[]awselasticloadbalancingv2.ApplicationListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) SetTargetGroups(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationMultipleTargetGroupsServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ApplicationMultipleTargetGroupsServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) {
	_jsii_.InvokeVoid(
		a,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		a,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) FindListener(name *string) awselasticloadbalancingv2.ApplicationListener {
	var returns awselasticloadbalancingv2.ApplicationListener

	_jsii_.Invoke(
		a,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		a,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base ApplicationMultipleTargetGroupsEc2Service or ApplicationMultipleTargetGroupsFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationMultipleTargetGroupsServiceBaseProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The application load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*ApplicationLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify ALB target groups.
	// Experimental.
	TargetGroups *[]*ApplicationTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Properties to define an application target group.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationTargetProps struct {
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	// Experimental.
	ContainerPort *float64 `json:"containerPort"`
	// Rule applies if the requested host matches the indicated host.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#host-conditions
	//
	// Experimental.
	HostHeader *string `json:"hostHeader"`
	// Name of the listener the target group attached to.
	// Experimental.
	Listener *string `json:"listener"`
	// Rule applies if the requested path matches the given path pattern.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Experimental.
	PathPattern *string `json:"pathPattern"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	// Experimental.
	Priority *float64 `json:"priority"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	// Experimental.
	Protocol awsecs.Protocol `json:"protocol"`
}

// Properties to define an network listener.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkListenerProps struct {
	// Name of the listener.
	// Experimental.
	Name *string `json:"name"`
	// The port on which the listener listens for requests.
	// Experimental.
	Port *float64 `json:"port"`
}

// An EC2 service running on an ECS cluster fronted by a network load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedEc2Service interface {
	NetworkLoadBalancedServiceBase
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	Node() constructs.Node
	Service() awsecs.Ec2Service
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	TaskDefinition() awsecs.Ec2TaskDefinition
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancedEc2Service
type jsiiProxy_NetworkLoadBalancedEc2Service struct {
	jsiiProxy_NetworkLoadBalancedServiceBase
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkLoadBalancedEc2Service class.
// Experimental.
func NewNetworkLoadBalancedEc2Service(scope constructs.Construct, id *string, props *NetworkLoadBalancedEc2ServiceProps) NetworkLoadBalancedEc2Service {
	_init_.Initialize()

	j := jsiiProxy_NetworkLoadBalancedEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkLoadBalancedEc2Service class.
// Experimental.
func NewNetworkLoadBalancedEc2Service_Override(n NetworkLoadBalancedEc2Service, scope constructs.Construct, id *string, props *NetworkLoadBalancedEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedEc2Service",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkLoadBalancedEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedEc2Service) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		n,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedEc2Service) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the NetworkLoadBalancedEc2Service service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedEc2ServiceProps struct {
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the network load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The network load balancer that will serve traffic to the service.
	//
	// If the load balancer has been imported, the vpc attribute must be specified
	// in the call to fromNetworkLoadBalancerAttributes().
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.INetworkLoadBalancer `json:"loadBalancer"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType NetworkLoadBalancedServiceRecordType `json:"recordType"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The properties required to create a new task definition.
	//
	// One of taskImageOptions or taskDefinition must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under contention, Docker attempts to keep the
	// container memory within the limit. If the container requires more memory,
	// it can consume up to the value specified by the Memory property or all of
	// the available memory on the container instance—whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both..
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `json:"taskDefinition"`
}

// A Fargate service running on an ECS cluster fronted by a network load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedFargateService interface {
	NetworkLoadBalancedServiceBase
	AssignPublicIp() *bool
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	Node() constructs.Node
	Service() awsecs.FargateService
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	TaskDefinition() awsecs.FargateTaskDefinition
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancedFargateService
type jsiiProxy_NetworkLoadBalancedFargateService struct {
	jsiiProxy_NetworkLoadBalancedServiceBase
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkLoadBalancedFargateService class.
// Experimental.
func NewNetworkLoadBalancedFargateService(scope constructs.Construct, id *string, props *NetworkLoadBalancedFargateServiceProps) NetworkLoadBalancedFargateService {
	_init_.Initialize()

	j := jsiiProxy_NetworkLoadBalancedFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkLoadBalancedFargateService class.
// Experimental.
func NewNetworkLoadBalancedFargateService_Override(n NetworkLoadBalancedFargateService, scope constructs.Construct, id *string, props *NetworkLoadBalancedFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkLoadBalancedFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedFargateService) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		n,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the NetworkLoadBalancedFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedFargateServiceProps struct {
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the network load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The network load balancer that will serve traffic to the service.
	//
	// If the load balancer has been imported, the vpc attribute must be specified
	// in the call to fromNetworkLoadBalancerAttributes().
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.INetworkLoadBalancer `json:"loadBalancer"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType NetworkLoadBalancedServiceRecordType `json:"recordType"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The properties required to create a new task definition.
	//
	// One of taskImageOptions or taskDefinition must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `json:"taskDefinition"`
	// The subnets to associate with the service.
	// Experimental.
	TaskSubnets *awsec2.SubnetSelection `json:"taskSubnets"`
}

// The base class for NetworkLoadBalancedEc2Service and NetworkLoadBalancedFargateService services.
// Experimental.
type NetworkLoadBalancedServiceBase interface {
	constructs.Construct
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	Node() constructs.Node
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancedServiceBase
type jsiiProxy_NetworkLoadBalancedServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkLoadBalancedServiceBase class.
// Experimental.
func NewNetworkLoadBalancedServiceBase_Override(n NetworkLoadBalancedServiceBase, scope constructs.Construct, id *string, props *NetworkLoadBalancedServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedServiceBase",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkLoadBalancedServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds service as a target of the target group.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedServiceBase) AddServiceAsTarget(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		n,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkLoadBalancedServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base NetworkLoadBalancedEc2Service or NetworkLoadBalancedFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedServiceBaseProps struct {
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// Listener port of the network load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// The network load balancer that will serve traffic to the service.
	//
	// If the load balancer has been imported, the vpc attribute must be specified
	// in the call to fromNetworkLoadBalancerAttributes().
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.INetworkLoadBalancer `json:"loadBalancer"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType NetworkLoadBalancedServiceRecordType `json:"recordType"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// The properties required to create a new task definition.
	//
	// One of taskImageOptions or taskDefinition must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageOptions `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Describes the type of DNS record the service should create.
// Experimental.
type NetworkLoadBalancedServiceRecordType string

const (
	NetworkLoadBalancedServiceRecordType_ALIAS NetworkLoadBalancedServiceRecordType = "ALIAS"
	NetworkLoadBalancedServiceRecordType_CNAME NetworkLoadBalancedServiceRecordType = "CNAME"
	NetworkLoadBalancedServiceRecordType_NONE NetworkLoadBalancedServiceRecordType = "NONE"
)

// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedTaskImageOptions struct {
	// The container name value to be specified in the task definition.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	// Experimental.
	ContainerPort *float64 `json:"containerPort"`
	// A key/value map of labels to add to the container.
	// Experimental.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// Options for configuring a new container.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancedTaskImageProps struct {
	// The container name value to be specified in the task definition.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// A list of port numbers on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	// Experimental.
	ContainerPorts *[]*float64 `json:"containerPorts"`
	// A key/value map of labels to add to the container.
	// Experimental.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Experimental.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secrets to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// Properties to define an network load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkLoadBalancerProps struct {
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `json:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `json:"domainZone"`
	// Listeners (at least one listener) attached to this load balancer.
	// Experimental.
	Listeners *[]*NetworkListenerProps `json:"listeners"`
	// Name of the load balancer.
	// Experimental.
	Name *string `json:"name"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `json:"publicLoadBalancer"`
}

// An EC2 service running on an ECS cluster fronted by a network load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMultipleTargetGroupsEc2Service interface {
	NetworkMultipleTargetGroupsServiceBase
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	Listeners() *[]awselasticloadbalancingv2.NetworkListener
	SetListeners(val *[]awselasticloadbalancingv2.NetworkListener)
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	Service() awsecs.Ec2Service
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup)
	TaskDefinition() awsecs.Ec2TaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.NetworkListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup
	ToString() *string
}

// The jsii proxy struct for NetworkMultipleTargetGroupsEc2Service
type jsiiProxy_NetworkMultipleTargetGroupsEc2Service struct {
	jsiiProxy_NetworkMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) Listeners() *[]awselasticloadbalancingv2.NetworkListener {
	var returns *[]awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup {
	var returns *[]awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkMultipleTargetGroupsEc2Service class.
// Experimental.
func NewNetworkMultipleTargetGroupsEc2Service(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsEc2ServiceProps) NetworkMultipleTargetGroupsEc2Service {
	_init_.Initialize()

	j := jsiiProxy_NetworkMultipleTargetGroupsEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkMultipleTargetGroupsEc2Service class.
// Experimental.
func NewNetworkMultipleTargetGroupsEc2Service_Override(n NetworkMultipleTargetGroupsEc2Service, scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsEc2Service",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) SetListeners(val *[]awselasticloadbalancingv2.NetworkListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkMultipleTargetGroupsEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) {
	_jsii_.InvokeVoid(
		n,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) FindListener(name *string) awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener

	_jsii_.Invoke(
		n,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup

	_jsii_.Invoke(
		n,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the NetworkMultipleTargetGroupsEc2Service service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMultipleTargetGroupsEc2ServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The network load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*NetworkLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify NLB target groups.
	// Experimental.
	TargetGroups *[]*NetworkTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The minimum number of CPU units to reserve for the container.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	//
	// Note that this setting will be ignored if TaskImagesOptions is specified.
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The task definition to use for tasks in the service. Only one of TaskDefinition or TaskImageOptions must be specified.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `json:"taskDefinition"`
}

// A Fargate service running on an ECS cluster fronted by a network load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMultipleTargetGroupsFargateService interface {
	NetworkMultipleTargetGroupsServiceBase
	AssignPublicIp() *bool
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	Listeners() *[]awselasticloadbalancingv2.NetworkListener
	SetListeners(val *[]awselasticloadbalancingv2.NetworkListener)
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	Service() awsecs.FargateService
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup)
	TaskDefinition() awsecs.FargateTaskDefinition
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.NetworkListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup
	ToString() *string
}

// The jsii proxy struct for NetworkMultipleTargetGroupsFargateService
type jsiiProxy_NetworkMultipleTargetGroupsFargateService struct {
	jsiiProxy_NetworkMultipleTargetGroupsServiceBase
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Listeners() *[]awselasticloadbalancingv2.NetworkListener {
	var returns *[]awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup {
	var returns *[]awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkMultipleTargetGroupsFargateService class.
// Experimental.
func NewNetworkMultipleTargetGroupsFargateService(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) NetworkMultipleTargetGroupsFargateService {
	_init_.Initialize()

	j := jsiiProxy_NetworkMultipleTargetGroupsFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the NetworkMultipleTargetGroupsFargateService class.
// Experimental.
func NewNetworkMultipleTargetGroupsFargateService_Override(n NetworkMultipleTargetGroupsFargateService, scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) SetListeners(val *[]awselasticloadbalancingv2.NetworkListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkMultipleTargetGroupsFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) {
	_jsii_.InvokeVoid(
		n,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) FindListener(name *string) awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener

	_jsii_.Invoke(
		n,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup

	_jsii_.Invoke(
		n,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the NetworkMultipleTargetGroupsFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMultipleTargetGroupsFargateServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The network load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*NetworkLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify NLB target groups.
	// Experimental.
	TargetGroups *[]*NetworkTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Determines whether the service will be assigned a public IP address.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The task definition to use for tasks in the service. Only one of TaskDefinition or TaskImageOptions must be specified.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `json:"taskDefinition"`
}

// The base class for NetworkMultipleTargetGroupsEc2Service and NetworkMultipleTargetGroupsFargateService classes.
// Experimental.
type NetworkMultipleTargetGroupsServiceBase interface {
	constructs.Construct
	Cluster() awsecs.ICluster
	DesiredCount() *float64
	InternalDesiredCount() *float64
	Listener() awselasticloadbalancingv2.NetworkListener
	Listeners() *[]awselasticloadbalancingv2.NetworkListener
	SetListeners(val *[]awselasticloadbalancingv2.NetworkListener)
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	LogDriver() awsecs.LogDriver
	SetLogDriver(val awsecs.LogDriver)
	Node() constructs.Node
	TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup
	SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup)
	AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	FindListener(name *string) awselasticloadbalancingv2.NetworkListener
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup
	ToString() *string
}

// The jsii proxy struct for NetworkMultipleTargetGroupsServiceBase
type jsiiProxy_NetworkMultipleTargetGroupsServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) Listeners() *[]awselasticloadbalancingv2.NetworkListener {
	var returns *[]awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) TargetGroups() *[]awselasticloadbalancingv2.NetworkTargetGroup {
	var returns *[]awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkMultipleTargetGroupsServiceBase class.
// Experimental.
func NewNetworkMultipleTargetGroupsServiceBase_Override(n NetworkMultipleTargetGroupsServiceBase, scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsServiceBase",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) SetListeners(val *[]awselasticloadbalancingv2.NetworkListener) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) SetLogDriver(val awsecs.LogDriver) {
	_jsii_.Set(
		j,
		"logDriver",
		val,
	)
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) SetTargetGroups(val *[]awselasticloadbalancingv2.NetworkTargetGroup) {
	_jsii_.Set(
		j,
		"targetGroups",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkMultipleTargetGroupsServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkMultipleTargetGroupsServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) AddPortMappingForTargets(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) {
	_jsii_.InvokeVoid(
		n,
		"addPortMappingForTargets",
		[]interface{}{container, targets},
	)
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) FindListener(name *string) awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener

	_jsii_.Invoke(
		n,
		"findListener",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) RegisterECSTargets(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup

	_jsii_.Invoke(
		n,
		"registerECSTargets",
		[]interface{}{service, container, targets},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base NetworkMultipleTargetGroupsEc2Service or NetworkMultipleTargetGroupsFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMultipleTargetGroupsServiceBaseProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `json:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1
	// Experimental.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The network load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*NetworkLoadBalancerProps `json:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// Name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Properties to specify NLB target groups.
	// Experimental.
	TargetGroups *[]*NetworkTargetProps `json:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `json:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Properties to define a network load balancer target group.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkTargetProps struct {
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	// Experimental.
	ContainerPort *float64 `json:"containerPort"`
	// Name of the listener the target group attached to.
	// Experimental.
	Listener *string `json:"listener"`
}

// Class to create a queue processing EC2 service.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueProcessingEc2Service interface {
	QueueProcessingServiceBase
	Cluster() awsecs.ICluster
	DeadLetterQueue() awssqs.IQueue
	DesiredCount() *float64
	Environment() *map[string]*string
	LogDriver() awsecs.LogDriver
	MaxCapacity() *float64
	MinCapacity() *float64
	Node() constructs.Node
	ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval
	Secrets() *map[string]awsecs.Secret
	Service() awsecs.Ec2Service
	SqsQueue() awssqs.IQueue
	TaskDefinition() awsecs.Ec2TaskDefinition
	ConfigureAutoscalingForService(service awsecs.BaseService)
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	GrantPermissionsToService(service awsecs.BaseService)
	ToString() *string
}

// The jsii proxy struct for QueueProcessingEc2Service
type jsiiProxy_QueueProcessingEc2Service struct {
	jsiiProxy_QueueProcessingServiceBase
}

func (j *jsiiProxy_QueueProcessingEc2Service) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval {
	var returns *[]*awsapplicationautoscaling.ScalingInterval
	_jsii_.Get(
		j,
		"scalingSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Secrets() *map[string]awsecs.Secret {
	var returns *map[string]awsecs.Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) Service() awsecs.Ec2Service {
	var returns awsecs.Ec2Service
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) SqsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingEc2Service) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the QueueProcessingEc2Service class.
// Experimental.
func NewQueueProcessingEc2Service(scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) QueueProcessingEc2Service {
	_init_.Initialize()

	j := jsiiProxy_QueueProcessingEc2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the QueueProcessingEc2Service class.
// Experimental.
func NewQueueProcessingEc2Service_Override(q QueueProcessingEc2Service, scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		[]interface{}{scope, id, props},
		q,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func QueueProcessingEc2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingEc2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
// Experimental.
func (q *jsiiProxy_QueueProcessingEc2Service) ConfigureAutoscalingForService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"configureAutoscalingForService",
		[]interface{}{service},
	)
}

// Returns the default cluster.
// Experimental.
func (q *jsiiProxy_QueueProcessingEc2Service) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		q,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Grant SQS permissions to an ECS service.
// Experimental.
func (q *jsiiProxy_QueueProcessingEc2Service) GrantPermissionsToService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"grantPermissionsToService",
		[]interface{}{service},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (q *jsiiProxy_QueueProcessingEc2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the QueueProcessingEc2Service service.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueProcessingEc2ServiceProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Experimental.
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `minScalingCapacity` or a literal object instead.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	//
	// The variable `QUEUE_NAME` with value `queue.queueName` will
	// always be passed.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of a family that the task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The maximum number of times that a message can be received by consumers.
	//
	// When this value is exceeded for a message the message will be automatically sent to the Dead Letter Queue.
	// Experimental.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Maximum capacity to scale to.
	// Experimental.
	MaxScalingCapacity *float64 `json:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Experimental.
	MinScalingCapacity *float64 `json:"minScalingCapacity"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// A queue for which to process items from.
	//
	// If specified and this is a FIFO queue, the queue name must end in the string '.fifo'. See
	// [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)
	// Experimental.
	Queue awssqs.IQueue `json:"queue"`
	// The number of seconds that Dead Letter Queue retains a message.
	// Experimental.
	RetentionPeriod awscdk.Duration `json:"retentionPeriod"`
	// The intervals for scaling based on the SQS queue's ApproximateNumberOfMessagesVisible metric.
	//
	// Maps a range of metric values to a particular scaling behavior. See
	// [Simple and Step Scaling Policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
	// Experimental.
	ScalingSteps *[]*awsapplicationautoscaling.ScalingInterval `json:"scalingSteps"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message and delete it from the queue
	// before it becomes visible again for dequeueing by another processor. Values must be between 0 and (12 hours).
	// Experimental.
	VisibilityTimeout awscdk.Duration `json:"visibilityTimeout"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Optional name for the container added.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// Gpu count for container in task definition.
	//
	// Set this if you want to use gpu based instances.
	// Experimental.
	GpuCount *float64 `json:"gpuCount"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under contention, Docker attempts to keep the
	// container memory within the limit. If the container requires more memory,
	// it can consume up to the value specified by the Memory property or all of
	// the available memory on the container instance—whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
}

// Class to create a queue processing Fargate service.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueProcessingFargateService interface {
	QueueProcessingServiceBase
	Cluster() awsecs.ICluster
	DeadLetterQueue() awssqs.IQueue
	DesiredCount() *float64
	Environment() *map[string]*string
	LogDriver() awsecs.LogDriver
	MaxCapacity() *float64
	MinCapacity() *float64
	Node() constructs.Node
	ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval
	Secrets() *map[string]awsecs.Secret
	Service() awsecs.FargateService
	SqsQueue() awssqs.IQueue
	TaskDefinition() awsecs.FargateTaskDefinition
	ConfigureAutoscalingForService(service awsecs.BaseService)
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	GrantPermissionsToService(service awsecs.BaseService)
	ToString() *string
}

// The jsii proxy struct for QueueProcessingFargateService
type jsiiProxy_QueueProcessingFargateService struct {
	jsiiProxy_QueueProcessingServiceBase
}

func (j *jsiiProxy_QueueProcessingFargateService) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval {
	var returns *[]*awsapplicationautoscaling.ScalingInterval
	_jsii_.Get(
		j,
		"scalingSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) Secrets() *map[string]awsecs.Secret {
	var returns *map[string]awsecs.Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) Service() awsecs.FargateService {
	var returns awsecs.FargateService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) SqsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingFargateService) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the QueueProcessingFargateService class.
// Experimental.
func NewQueueProcessingFargateService(scope constructs.Construct, id *string, props *QueueProcessingFargateServiceProps) QueueProcessingFargateService {
	_init_.Initialize()

	j := jsiiProxy_QueueProcessingFargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingFargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the QueueProcessingFargateService class.
// Experimental.
func NewQueueProcessingFargateService_Override(q QueueProcessingFargateService, scope constructs.Construct, id *string, props *QueueProcessingFargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingFargateService",
		[]interface{}{scope, id, props},
		q,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func QueueProcessingFargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingFargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
// Experimental.
func (q *jsiiProxy_QueueProcessingFargateService) ConfigureAutoscalingForService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"configureAutoscalingForService",
		[]interface{}{service},
	)
}

// Returns the default cluster.
// Experimental.
func (q *jsiiProxy_QueueProcessingFargateService) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		q,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Grant SQS permissions to an ECS service.
// Experimental.
func (q *jsiiProxy_QueueProcessingFargateService) GrantPermissionsToService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"grantPermissionsToService",
		[]interface{}{service},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (q *jsiiProxy_QueueProcessingFargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the QueueProcessingFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueProcessingFargateServiceProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Experimental.
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `minScalingCapacity` or a literal object instead.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	//
	// The variable `QUEUE_NAME` with value `queue.queueName` will
	// always be passed.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of a family that the task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The maximum number of times that a message can be received by consumers.
	//
	// When this value is exceeded for a message the message will be automatically sent to the Dead Letter Queue.
	// Experimental.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Maximum capacity to scale to.
	// Experimental.
	MaxScalingCapacity *float64 `json:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Experimental.
	MinScalingCapacity *float64 `json:"minScalingCapacity"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// A queue for which to process items from.
	//
	// If specified and this is a FIFO queue, the queue name must end in the string '.fifo'. See
	// [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)
	// Experimental.
	Queue awssqs.IQueue `json:"queue"`
	// The number of seconds that Dead Letter Queue retains a message.
	// Experimental.
	RetentionPeriod awscdk.Duration `json:"retentionPeriod"`
	// The intervals for scaling based on the SQS queue's ApproximateNumberOfMessagesVisible metric.
	//
	// Maps a range of metric values to a particular scaling behavior. See
	// [Simple and Step Scaling Policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
	// Experimental.
	ScalingSteps *[]*awsapplicationautoscaling.ScalingInterval `json:"scalingSteps"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message and delete it from the queue
	// before it becomes visible again for dequeueing by another processor. Values must be between 0 and (12 hours).
	// Experimental.
	VisibilityTimeout awscdk.Duration `json:"visibilityTimeout"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// Optional name for the container added.
	// Experimental.
	ContainerName *string `json:"containerName"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 0.5GB, 1GB, 2GB - Available cpu values: 256 (.25 vCPU)
	//
	// 1GB, 2GB, 3GB, 4GB - Available cpu values: 512 (.5 vCPU)
	//
	// 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4GB and 16GB in 1GB increments - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8GB and 30GB in 1GB increments - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The subnets to associate with the service.
	// Experimental.
	TaskSubnets *awsec2.SubnetSelection `json:"taskSubnets"`
}

// The base class for QueueProcessingEc2Service and QueueProcessingFargateService services.
// Experimental.
type QueueProcessingServiceBase interface {
	constructs.Construct
	Cluster() awsecs.ICluster
	DeadLetterQueue() awssqs.IQueue
	DesiredCount() *float64
	Environment() *map[string]*string
	LogDriver() awsecs.LogDriver
	MaxCapacity() *float64
	MinCapacity() *float64
	Node() constructs.Node
	ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval
	Secrets() *map[string]awsecs.Secret
	SqsQueue() awssqs.IQueue
	ConfigureAutoscalingForService(service awsecs.BaseService)
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	GrantPermissionsToService(service awsecs.BaseService)
	ToString() *string
}

// The jsii proxy struct for QueueProcessingServiceBase
type jsiiProxy_QueueProcessingServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_QueueProcessingServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) LogDriver() awsecs.LogDriver {
	var returns awsecs.LogDriver
	_jsii_.Get(
		j,
		"logDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) ScalingSteps() *[]*awsapplicationautoscaling.ScalingInterval {
	var returns *[]*awsapplicationautoscaling.ScalingInterval
	_jsii_.Get(
		j,
		"scalingSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) Secrets() *map[string]awsecs.Secret {
	var returns *map[string]awsecs.Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueProcessingServiceBase) SqsQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


// Constructs a new instance of the QueueProcessingServiceBase class.
// Experimental.
func NewQueueProcessingServiceBase_Override(q QueueProcessingServiceBase, scope constructs.Construct, id *string, props *QueueProcessingServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingServiceBase",
		[]interface{}{scope, id, props},
		q,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func QueueProcessingServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.QueueProcessingServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Configure autoscaling based off of CPU utilization as well as the number of messages visible in the SQS queue.
// Experimental.
func (q *jsiiProxy_QueueProcessingServiceBase) ConfigureAutoscalingForService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"configureAutoscalingForService",
		[]interface{}{service},
	)
}

// Returns the default cluster.
// Experimental.
func (q *jsiiProxy_QueueProcessingServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		q,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Grant SQS permissions to an ECS service.
// Experimental.
func (q *jsiiProxy_QueueProcessingServiceBase) GrantPermissionsToService(service awsecs.BaseService) {
	_jsii_.InvokeVoid(
		q,
		"grantPermissionsToService",
		[]interface{}{service},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (q *jsiiProxy_QueueProcessingServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base QueueProcessingEc2Service or QueueProcessingFargateService service.
//
// TODO: EXAMPLE
//
// Experimental.
type QueueProcessingServiceBaseProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Experimental.
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `minScalingCapacity` or a literal object instead.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// The environment variables to pass to the container.
	//
	// The variable `QUEUE_NAME` with value `queue.queueName` will
	// always be passed.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The name of a family that the task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `json:"family"`
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The maximum number of times that a message can be received by consumers.
	//
	// When this value is exceeded for a message the message will be automatically sent to the Dead Letter Queue.
	// Experimental.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Maximum capacity to scale to.
	// Experimental.
	MaxScalingCapacity *float64 `json:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Experimental.
	MinScalingCapacity *float64 `json:"minScalingCapacity"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `json:"propagateTags"`
	// A queue for which to process items from.
	//
	// If specified and this is a FIFO queue, the queue name must end in the string '.fifo'. See
	// [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)
	// Experimental.
	Queue awssqs.IQueue `json:"queue"`
	// The number of seconds that Dead Letter Queue retains a message.
	// Experimental.
	RetentionPeriod awscdk.Duration `json:"retentionPeriod"`
	// The intervals for scaling based on the SQS queue's ApproximateNumberOfMessagesVisible metric.
	//
	// Maps a range of metric values to a particular scaling behavior. See
	// [Simple and Step Scaling Policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
	// Experimental.
	ScalingSteps *[]*awsapplicationautoscaling.ScalingInterval `json:"scalingSteps"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The name of the service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message and delete it from the queue
	// before it becomes visible again for dequeueing by another processor. Values must be between 0 and (12 hours).
	// Experimental.
	VisibilityTimeout awscdk.Duration `json:"visibilityTimeout"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// A scheduled EC2 task that will be initiated off of CloudWatch Events.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledEc2Task interface {
	ScheduledTaskBase
	Cluster() awsecs.ICluster
	DesiredTaskCount() *float64
	EventRule() awsevents.Rule
	Node() constructs.Node
	SubnetSelection() *awsec2.SubnetSelection
	Task() awseventstargets.EcsTask
	TaskDefinition() awsecs.Ec2TaskDefinition
	AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask)
	AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ScheduledEc2Task
type jsiiProxy_ScheduledEc2Task struct {
	jsiiProxy_ScheduledTaskBase
}

func (j *jsiiProxy_ScheduledEc2Task) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) DesiredTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) EventRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) Task() awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledEc2Task) TaskDefinition() awsecs.Ec2TaskDefinition {
	var returns awsecs.Ec2TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScheduledEc2Task class.
// Experimental.
func NewScheduledEc2Task(scope constructs.Construct, id *string, props *ScheduledEc2TaskProps) ScheduledEc2Task {
	_init_.Initialize()

	j := jsiiProxy_ScheduledEc2Task{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledEc2Task",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScheduledEc2Task class.
// Experimental.
func NewScheduledEc2Task_Override(s ScheduledEc2Task, scope constructs.Construct, id *string, props *ScheduledEc2TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledEc2Task",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ScheduledEc2Task_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledEc2Task",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds task as a target of the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledEc2Task) AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask) {
	_jsii_.InvokeVoid(
		s,
		"addTaskAsTarget",
		[]interface{}{ecsTaskTarget},
	)
}

// Create an ECS task using the task definition provided and add it to the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledEc2Task) AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask

	_jsii_.Invoke(
		s,
		"addTaskDefinitionToEventTarget",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

// Create an AWS Log Driver with the provided streamPrefix.
// Experimental.
func (s *jsiiProxy_ScheduledEc2Task) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		s,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (s *jsiiProxy_ScheduledEc2Task) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		s,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ScheduledEc2Task) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ScheduledEc2Task using a task definition.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledEc2TaskDefinitionOptions struct {
	// The task definition to use for tasks in the service. One of image or taskDefinition must be specified.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `json:"taskDefinition"`
}

// The properties for the ScheduledEc2Task using an image.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledEc2TaskImageOptions struct {
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The minimum number of CPU units to reserve for the container.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under contention, Docker attempts to keep the
	// container memory within the limit. If the container requires more memory,
	// it can consume up to the value specified by the Memory property or all of
	// the available memory on the container instance—whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Experimental.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
}

// The properties for the ScheduledEc2Task task.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledEc2TaskProps struct {
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Indicates whether the rule is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// A name for the rule.
	// Experimental.
	RuleName *string `json:"ruleName"`
	// The schedule or rate (frequency) that determines when CloudWatch Events runs the rule.
	//
	// For more information, see
	// [Schedule Expression Syntax for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch User Guide.
	// Experimental.
	Schedule awsapplicationautoscaling.Schedule `json:"schedule"`
	// Existing security groups to use for your service.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The properties to define if using an existing TaskDefinition in this construct.
	//
	// ScheduledEc2TaskDefinitionOptions or ScheduledEc2TaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledEc2TaskDefinitionOptions *ScheduledEc2TaskDefinitionOptions `json:"scheduledEc2TaskDefinitionOptions"`
	// The properties to define if the construct is to create a TaskDefinition.
	//
	// ScheduledEc2TaskDefinitionOptions or ScheduledEc2TaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledEc2TaskImageOptions *ScheduledEc2TaskImageOptions `json:"scheduledEc2TaskImageOptions"`
}

// A scheduled Fargate task that will be initiated off of CloudWatch Events.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledFargateTask interface {
	ScheduledTaskBase
	Cluster() awsecs.ICluster
	DesiredTaskCount() *float64
	EventRule() awsevents.Rule
	Node() constructs.Node
	SubnetSelection() *awsec2.SubnetSelection
	Task() awseventstargets.EcsTask
	TaskDefinition() awsecs.FargateTaskDefinition
	AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask)
	AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ScheduledFargateTask
type jsiiProxy_ScheduledFargateTask struct {
	jsiiProxy_ScheduledTaskBase
}

func (j *jsiiProxy_ScheduledFargateTask) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) DesiredTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) EventRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) Task() awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledFargateTask) TaskDefinition() awsecs.FargateTaskDefinition {
	var returns awsecs.FargateTaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScheduledFargateTask class.
// Experimental.
func NewScheduledFargateTask(scope constructs.Construct, id *string, props *ScheduledFargateTaskProps) ScheduledFargateTask {
	_init_.Initialize()

	j := jsiiProxy_ScheduledFargateTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledFargateTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScheduledFargateTask class.
// Experimental.
func NewScheduledFargateTask_Override(s ScheduledFargateTask, scope constructs.Construct, id *string, props *ScheduledFargateTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledFargateTask",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ScheduledFargateTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledFargateTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds task as a target of the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledFargateTask) AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask) {
	_jsii_.InvokeVoid(
		s,
		"addTaskAsTarget",
		[]interface{}{ecsTaskTarget},
	)
}

// Create an ECS task using the task definition provided and add it to the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledFargateTask) AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask

	_jsii_.Invoke(
		s,
		"addTaskDefinitionToEventTarget",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

// Create an AWS Log Driver with the provided streamPrefix.
// Experimental.
func (s *jsiiProxy_ScheduledFargateTask) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		s,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (s *jsiiProxy_ScheduledFargateTask) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		s,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ScheduledFargateTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ScheduledFargateTask using a task definition.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledFargateTaskDefinitionOptions struct {
	// The task definition to use for tasks in the service. Image or taskDefinition must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	TaskDefinition awsecs.FargateTaskDefinition `json:"taskDefinition"`
}

// The properties for the ScheduledFargateTask using an image.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledFargateTaskImageOptions struct {
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	// Experimental.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
}

// The properties for the ScheduledFargateTask task.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledFargateTaskProps struct {
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Indicates whether the rule is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// A name for the rule.
	// Experimental.
	RuleName *string `json:"ruleName"`
	// The schedule or rate (frequency) that determines when CloudWatch Events runs the rule.
	//
	// For more information, see
	// [Schedule Expression Syntax for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch User Guide.
	// Experimental.
	Schedule awsapplicationautoscaling.Schedule `json:"schedule"`
	// Existing security groups to use for your service.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// The properties to define if using an existing TaskDefinition in this construct.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledFargateTaskDefinitionOptions *ScheduledFargateTaskDefinitionOptions `json:"scheduledFargateTaskDefinitionOptions"`
	// The properties to define if the construct is to create a TaskDefinition.
	//
	// ScheduledFargateTaskDefinitionOptions or ScheduledFargateTaskImageOptions must be defined, but not both.
	// Experimental.
	ScheduledFargateTaskImageOptions *ScheduledFargateTaskImageOptions `json:"scheduledFargateTaskImageOptions"`
}

// The base class for ScheduledEc2Task and ScheduledFargateTask tasks.
// Experimental.
type ScheduledTaskBase interface {
	constructs.Construct
	Cluster() awsecs.ICluster
	DesiredTaskCount() *float64
	EventRule() awsevents.Rule
	Node() constructs.Node
	SubnetSelection() *awsec2.SubnetSelection
	AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask)
	AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	ToString() *string
}

// The jsii proxy struct for ScheduledTaskBase
type jsiiProxy_ScheduledTaskBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ScheduledTaskBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) DesiredTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) EventRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledTaskBase) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScheduledTaskBase class.
// Experimental.
func NewScheduledTaskBase_Override(s ScheduledTaskBase, scope constructs.Construct, id *string, props *ScheduledTaskBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledTaskBase",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ScheduledTaskBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.ScheduledTaskBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds task as a target of the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledTaskBase) AddTaskAsTarget(ecsTaskTarget awseventstargets.EcsTask) {
	_jsii_.InvokeVoid(
		s,
		"addTaskAsTarget",
		[]interface{}{ecsTaskTarget},
	)
}

// Create an ECS task using the task definition provided and add it to the scheduled event rule.
// Experimental.
func (s *jsiiProxy_ScheduledTaskBase) AddTaskDefinitionToEventTarget(taskDefinition awsecs.TaskDefinition) awseventstargets.EcsTask {
	var returns awseventstargets.EcsTask

	_jsii_.Invoke(
		s,
		"addTaskDefinitionToEventTarget",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

// Create an AWS Log Driver with the provided streamPrefix.
// Experimental.
func (s *jsiiProxy_ScheduledTaskBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		s,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Returns the default cluster.
// Experimental.
func (s *jsiiProxy_ScheduledTaskBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	var returns awsecs.Cluster

	_jsii_.Invoke(
		s,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ScheduledTaskBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base ScheduledEc2Task or ScheduledFargateTask task.
//
// TODO: EXAMPLE
//
// Experimental.
type ScheduledTaskBaseProps struct {
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredTaskCount *float64 `json:"desiredTaskCount"`
	// Indicates whether the rule is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// A name for the rule.
	// Experimental.
	RuleName *string `json:"ruleName"`
	// The schedule or rate (frequency) that determines when CloudWatch Events runs the rule.
	//
	// For more information, see
	// [Schedule Expression Syntax for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch User Guide.
	// Experimental.
	Schedule awsapplicationautoscaling.Schedule `json:"schedule"`
	// Existing security groups to use for your service.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// TODO: EXAMPLE
//
// Experimental.
type ScheduledTaskImageProps struct {
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `json:"image"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `json:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `json:"secrets"`
}

