//go:build no_runtime_type_checking

package externalnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalNetworkVsphereNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalNetworkVsphereNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkVsphereNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkVsphereNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkVsphereNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalNetworkVsphereNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalNetworkVsphereNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

