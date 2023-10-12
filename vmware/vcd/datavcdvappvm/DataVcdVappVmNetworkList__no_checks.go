//go:build no_runtime_type_checking

package datavcdvappvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVappVmNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVappVmNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVappVmNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

