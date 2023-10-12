//go:build no_runtime_type_checking

package networkrouted

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkRoutedDhcpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkRoutedDhcpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedDhcpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedDhcpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedDhcpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedDhcpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkRoutedDhcpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

