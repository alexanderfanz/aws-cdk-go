package awsapplicationautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a scalable target.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_7(),
//   	Handler: jsii.String("index.handler"),
//   	Code: Code,
//
//   	ReservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.currentVersion
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &ScalableTargetProps{
//   	ServiceNamespace: appscaling.ServiceNamespace_LAMBDA,
//   	MaxCapacity: jsii.Number(100),
//   	MinCapacity: jsii.Number(10),
//   	ResourceId: fmt.Sprintf("function:%v:%v", handler.FunctionName, fnVer.Version),
//   	ScalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.ScaleToTrackMetric(jsii.String("PceTracking"), &BasicTargetTrackingScalingPolicyProps{
//   	TargetValue: jsii.Number(0.9),
//   	PredefinedMetric: appscaling.PredefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
type ScalableTargetProps struct {
	// The maximum value that Application Auto Scaling can use to scale a target during a scaling activity.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum value that Application Auto Scaling can use to scale a target during a scaling activity.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// The resource identifier to associate with this scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// Example value: `service/ecsStack-MyECSCluster-AB12CDE3F4GH/ecsStack-MyECSService-AB12CDE3F4GH`.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension that's associated with the scalable target.
	//
	// Specify the service namespace, resource type, and scaling property.
	//
	// Example value: `ecs:service:DesiredCount`.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html
	//
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service.
	//
	// For valid AWS service namespace values, see the RegisterScalableTarget
	// action in the Application Auto Scaling API Reference.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html
	//
	ServiceNamespace ServiceNamespace `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// Role that allows Application Auto Scaling to modify your scalable target.
	// Default: A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

