//go:build no_runtime_type_checking

package datavcdvappvm

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdVappVmMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdVappVmMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdVappVmMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdVappVmMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

