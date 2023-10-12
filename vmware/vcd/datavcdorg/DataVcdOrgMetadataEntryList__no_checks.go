//go:build no_runtime_type_checking

package datavcdorg

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdOrgMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdOrgMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdOrgMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdOrgMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdOrgMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdOrgMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

