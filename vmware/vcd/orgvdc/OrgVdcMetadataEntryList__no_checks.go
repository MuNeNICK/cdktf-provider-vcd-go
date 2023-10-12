//go:build no_runtime_type_checking

package orgvdc

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrgVdcMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OrgVdcMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OrgVdcMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OrgVdcMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OrgVdcMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OrgVdcMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOrgVdcMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

