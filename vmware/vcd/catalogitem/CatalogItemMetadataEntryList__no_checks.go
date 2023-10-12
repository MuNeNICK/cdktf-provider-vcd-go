//go:build no_runtime_type_checking

package catalogitem

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogItemMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogItemMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogItemMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogItemMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogItemMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogItemMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogItemMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

