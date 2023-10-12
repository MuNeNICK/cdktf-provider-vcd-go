//go:build no_runtime_type_checking

package datavcdvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVmCustomizationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVmCustomizationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmCustomizationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmCustomizationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmCustomizationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVmCustomizationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

