//go:build no_runtime_type_checking

package externalnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeStaticIpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalNetworkIpScopeStaticIpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

