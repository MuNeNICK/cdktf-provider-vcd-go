//go:build no_runtime_type_checking

package datavcdvappvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVappVmInternalDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVappVmInternalDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmInternalDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmInternalDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmInternalDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVappVmInternalDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

