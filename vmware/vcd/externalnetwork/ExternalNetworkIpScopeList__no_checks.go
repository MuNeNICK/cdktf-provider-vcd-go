//go:build no_runtime_type_checking

package externalnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalNetworkIpScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalNetworkIpScopeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkIpScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalNetworkIpScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

