//go:build no_runtime_type_checking

package networkrouted

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkRoutedStaticIpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkRoutedStaticIpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedStaticIpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedStaticIpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedStaticIpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedStaticIpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkRoutedStaticIpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

