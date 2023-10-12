//go:build no_runtime_type_checking

package datavcdvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVmNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVmNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVmNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

