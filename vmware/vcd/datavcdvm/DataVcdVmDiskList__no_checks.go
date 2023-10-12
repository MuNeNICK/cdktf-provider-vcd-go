//go:build no_runtime_type_checking

package datavcdvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVmDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVmDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVmDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

