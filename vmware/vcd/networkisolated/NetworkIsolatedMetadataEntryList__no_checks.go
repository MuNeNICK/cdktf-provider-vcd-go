//go:build no_runtime_type_checking

package networkisolated

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIsolatedMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkIsolatedMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkIsolatedMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkIsolatedMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

