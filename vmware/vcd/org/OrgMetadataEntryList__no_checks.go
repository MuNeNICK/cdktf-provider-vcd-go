//go:build no_runtime_type_checking

package org

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrgMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OrgMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OrgMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OrgMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OrgMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OrgMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOrgMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

