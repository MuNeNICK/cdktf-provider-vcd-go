//go:build no_runtime_type_checking

package vappvm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappVmNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappVmNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappVmNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappVmNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappVmNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappVmNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappVmNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

