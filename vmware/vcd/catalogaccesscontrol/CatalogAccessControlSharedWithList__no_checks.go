//go:build no_runtime_type_checking

package catalogaccesscontrol

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogAccessControlSharedWithList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogAccessControlSharedWithList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogAccessControlSharedWithList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogAccessControlSharedWithList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogAccessControlSharedWithList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogAccessControlSharedWithList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogAccessControlSharedWithListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

