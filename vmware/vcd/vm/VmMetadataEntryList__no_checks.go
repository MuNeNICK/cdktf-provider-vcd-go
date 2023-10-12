//go:build no_runtime_type_checking

package vm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VmMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

