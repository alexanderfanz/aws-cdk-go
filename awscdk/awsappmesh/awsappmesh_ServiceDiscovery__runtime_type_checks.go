//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
)

func (s *jsiiProxy_ServiceDiscovery) validateBindParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateServiceDiscovery_CloudMapParameters(service awsservicediscovery.IService) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateServiceDiscovery_DnsParameters(hostname *string) error {
	if hostname == nil {
		return fmt.Errorf("parameter hostname is required, but nil was provided")
	}

	return nil
}

