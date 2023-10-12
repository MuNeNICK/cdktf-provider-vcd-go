//go:build no_runtime_type_checking

package datavcdcatalog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdCatalogMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdCatalogMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdCatalogMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

