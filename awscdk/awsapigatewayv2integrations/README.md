# AWS APIGatewayv2 Integrations

## Table of Contents

* [HTTP APIs](#http-apis)

  * [Lambda Integration](#lambda)
  * [HTTP Proxy Integration](#http-proxy)
  * [Private Integration](#private-integration)
  * [Request Parameters](#request-parameters)
* [WebSocket APIs](#websocket-apis)

  * [Lambda WebSocket Integration](#lambda-websocket-integration)

## HTTP APIs

Integrations connect a route to backend resources. HTTP APIs support Lambda proxy, AWS service, and HTTP proxy integrations. HTTP proxy integrations are also known as private integrations.

### Lambda

Lambda integrations enable integrating an HTTP API route with a Lambda function. When a client invokes the route, the
API Gateway service forwards the request to the Lambda function and returns the function's response to the client.

The API Gateway service will invoke the Lambda function with an event payload of a specific format. The service expects
the function to respond in a specific format. The details on this format are available at [Working with AWS Lambda
proxy integrations](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html).

The following code configures a route `GET /books` with a Lambda proxy integration.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var booksDefaultFn function

booksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)

httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

httpApi.AddRoutes(&AddRoutesOptions{
	Path: jsii.String("/books"),
	Methods: []httpMethod{
		apigwv2.*httpMethod_GET,
	},
	Integration: booksIntegration,
})
```

### HTTP Proxy

HTTP Proxy integrations enables connecting an HTTP API route to a publicly routable HTTP endpoint. When a client
invokes the route, the API Gateway service forwards the entire request and response between the API Gateway endpoint
and the integrating HTTP endpoint. More information can be found at [Working with HTTP proxy
integrations](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-http.html).

The following code configures a route `GET /books` with an HTTP proxy integration to an HTTP endpoint
`get-books-proxy.example.com`.

```go
import "github.com/aws/aws-cdk-go/awscdk"


booksIntegration := awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com"))

httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

httpApi.AddRoutes(&AddRoutesOptions{
	Path: jsii.String("/books"),
	Methods: []httpMethod{
		apigwv2.*httpMethod_GET,
	},
	Integration: booksIntegration,
})
```

### Private Integration

Private integrations enable integrating an HTTP API route with private resources in a VPC, such as Application Load Balancers or
Amazon ECS container-based applications.  Using private integrations, resources in a VPC can be exposed for access by
clients outside of the VPC.

The following integrations are supported for private resources in a VPC.

#### Application Load Balancer

The following code is a basic application load balancer private integration of HTTP API:

```go
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("lb"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
})
listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
})

httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener),
})
```

When an imported load balancer is used, the `vpc` option must be specified for `HttpAlbIntegration`.

#### Network Load Balancer

The following code is a basic network load balancer private integration of HTTP API:

```go
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
})
listener := lb.AddListener(jsii.String("listener"), &BaseNetworkListenerProps{
	Port: jsii.Number(80),
})
listener.AddTargets(jsii.String("target"), &AddNetworkTargetsProps{
	Port: jsii.Number(80),
})

httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
})
```

When an imported load balancer is used, the `vpc` option must be specified for `HttpNlbIntegration`.

#### Cloud Map Service Discovery

The following code is a basic discovery service private integration of HTTP API:

```go
import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
	Vpc: Vpc,
})
namespace := servicediscovery.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
	Name: jsii.String("boobar.com"),
	Vpc: Vpc,
})
service := namespace.CreateService(jsii.String("Service"))

httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &HttpServiceDiscoveryIntegrationProps{
		VpcLink: *VpcLink,
	}),
})
```

### Request Parameters

Request parameter mapping allows API requests from clients to be modified before they reach backend integrations.
Parameter mapping can be used to specify modifications to request parameters. See [Transforming API requests and
responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).

The following example creates a new header - `header2` - as a copy of `header1` and removes `header1`.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var lb applicationLoadBalancer

listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
})

httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &HttpAlbIntegrationProps{
		ParameterMapping: apigwv2.NewParameterMapping().AppendHeader(jsii.String("header2"), apigwv2.MappingValue_RequestHeader(jsii.String("header1"))).RemoveHeader(jsii.String("header1")),
	}),
})
```

To add mapping keys and values not yet supported by the CDK, use the `custom()` method:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var lb applicationLoadBalancer

listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
})

httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &HttpAlbIntegrationProps{
		ParameterMapping: apigwv2.NewParameterMapping().Custom(jsii.String("myKey"), jsii.String("myValue")),
	}),
})
```

## WebSocket APIs

WebSocket integrations connect a route to backend resources. The following integrations are supported in the CDK.

### Lambda WebSocket Integration

Lambda integrations enable integrating a WebSocket API route with a Lambda function. When a client connects/disconnects
or sends a message specific to a route, the API Gateway service forwards the request to the Lambda function

The API Gateway service will invoke the Lambda function with an event payload of a specific format.

The following code configures a `sendMessage` route with a Lambda integration

```go
import "github.com/aws/aws-cdk-go/awscdk"

var messageHandler function


webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
	WebSocketApi: WebSocketApi,
	StageName: jsii.String("dev"),
	AutoDeploy: jsii.Boolean(true),
})
webSocketApi.AddRoute(jsii.String("sendMessage"), &WebSocketRouteOptions{
	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
})
```
