//go:build no_runtime_type_checking

package networkdirect

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkDirectMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkDirectMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkDirectMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkDirectMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkDirectMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkDirectMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkDirectMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

