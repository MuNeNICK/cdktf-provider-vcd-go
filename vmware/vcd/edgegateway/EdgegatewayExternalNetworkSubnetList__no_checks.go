//go:build no_runtime_type_checking

package edgegateway

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEdgegatewayExternalNetworkSubnetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

