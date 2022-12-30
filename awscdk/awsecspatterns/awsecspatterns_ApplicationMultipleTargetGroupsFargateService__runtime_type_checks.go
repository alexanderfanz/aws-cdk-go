//go:build !no_runtime_type_checking

package awsecspatterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v3"
)

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if targets == nil {
		return fmt.Errorf("parameter targets is required, but nil was provided")
	}
	for idx_26cafb, v := range *targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targets[%#v]", idx_26cafb) }); err != nil {
			return err
		}
	}

	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if targets == nil {
		return fmt.Errorf("parameter targets is required, but nil was provided")
	}
	for idx_26cafb, v := range *targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targets[%#v]", idx_26cafb) }); err != nil {
			return err
		}
	}

	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateApplicationMultipleTargetGroupsFargateService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateSetListenersParameters(val *[]awselasticloadbalancingv2.ApplicationListener) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateSetTargetGroupsParameters(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationMultipleTargetGroupsFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

