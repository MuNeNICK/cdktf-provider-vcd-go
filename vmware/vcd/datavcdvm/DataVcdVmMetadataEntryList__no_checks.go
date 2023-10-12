//go:build no_runtime_type_checking

package datavcdvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVmMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVmMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVmMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVmMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

