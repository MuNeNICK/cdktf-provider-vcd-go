//go:build no_runtime_type_checking

package edgegateway

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EdgegatewayExternalNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EdgegatewayExternalNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEdgegatewayExternalNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

