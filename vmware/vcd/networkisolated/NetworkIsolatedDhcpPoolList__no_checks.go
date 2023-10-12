//go:build no_runtime_type_checking

package networkisolated

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIsolatedDhcpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedDhcpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedDhcpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedDhcpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedDhcpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedDhcpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkIsolatedDhcpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

