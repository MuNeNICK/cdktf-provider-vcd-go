//go:build no_runtime_type_checking

package datavcdcatalog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdCatalogFilterMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdCatalogFilterMetadataList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogFilterMetadataList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogFilterMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogFilterMetadataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdCatalogFilterMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdCatalogFilterMetadataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

