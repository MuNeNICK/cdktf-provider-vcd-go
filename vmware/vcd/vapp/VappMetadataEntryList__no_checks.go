//go:build no_runtime_type_checking

package vapp

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

