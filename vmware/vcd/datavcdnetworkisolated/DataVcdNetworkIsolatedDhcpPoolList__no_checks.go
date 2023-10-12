//go:build no_runtime_type_checking

package datavcdnetworkisolated

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdNetworkIsolatedDhcpPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdNetworkIsolatedDhcpPoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdNetworkIsolatedDhcpPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdNetworkIsolatedDhcpPoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdNetworkIsolatedDhcpPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdNetworkIsolatedDhcpPoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

