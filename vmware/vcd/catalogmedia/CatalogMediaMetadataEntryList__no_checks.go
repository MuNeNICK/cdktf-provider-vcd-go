//go:build no_runtime_type_checking

package catalogmedia

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogMediaMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogMediaMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogMediaMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogMediaMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogMediaMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogMediaMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogMediaMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

