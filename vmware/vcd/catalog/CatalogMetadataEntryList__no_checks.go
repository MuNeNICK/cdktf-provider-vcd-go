//go:build no_runtime_type_checking

package catalog

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

