//go:build no_runtime_type_checking

package datavcdvappvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVappVmDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVappVmDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVappVmDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

