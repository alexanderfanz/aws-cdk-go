package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gelfLogDriver := awscdk.Aws_ecs.NewGelfLogDriver(&gelfLogDriverProps{
//   	address: jsii.String("address"),
//
//   	// the properties below are optional
//   	compressionLevel: jsii.Number(123),
//   	compressionType: awscdk.*Aws_ecs.gelfCompressionType_GZIP,
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tcpMaxReconnect: jsii.Number(123),
//   	tcpReconnectDelay: cdk.duration.minutes(jsii.Number(30)),
//   })
//
type GelfLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for GelfLogDriver
type jsiiProxy_GelfLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the GelfLogDriver class.
func NewGelfLogDriver(props *GelfLogDriverProps) GelfLogDriver {
	_init_.Initialize()

	j := jsiiProxy_GelfLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the GelfLogDriver class.
func NewGelfLogDriver_Override(g GelfLogDriver, props *GelfLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		[]interface{}{props},
		g,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func GelfLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GelfLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}
