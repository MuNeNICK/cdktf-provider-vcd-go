//go:build no_runtime_type_checking

package datavcdindependentdisk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdIndependentDiskMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdIndependentDiskMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdIndependentDiskMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdIndependentDiskMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdIndependentDiskMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdIndependentDiskMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

