//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEcsEc2ContainerDefinition) validateAddUlimitParameters(ulimit *Ulimit) error {
	return nil
}
