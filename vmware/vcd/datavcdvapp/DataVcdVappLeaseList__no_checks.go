//go:build no_runtime_type_checking

package datavcdvapp

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVappLeaseList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVappLeaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappLeaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappLeaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappLeaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVappLeaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

