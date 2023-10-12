//go:build no_runtime_type_checking

package networkisolated

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedStaticIpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkIsolatedStaticIpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

