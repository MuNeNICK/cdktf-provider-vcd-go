//go:build no_runtime_type_checking

package vappnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappNetworkDhcpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappNetworkDhcpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappNetworkDhcpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappNetworkDhcpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappNetworkDhcpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappNetworkDhcpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappNetworkDhcpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

