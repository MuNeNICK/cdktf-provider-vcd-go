//go:build no_runtime_type_checking

package vappnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappNetworkStaticIpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappNetworkStaticIpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappNetworkStaticIpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappNetworkStaticIpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappNetworkStaticIpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappNetworkStaticIpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappNetworkStaticIpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

