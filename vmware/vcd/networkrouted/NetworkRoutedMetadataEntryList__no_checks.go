//go:build no_runtime_type_checking

package networkrouted

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkRoutedMetadataEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkRoutedMetadataEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedMetadataEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedMetadataEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedMetadataEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkRoutedMetadataEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkRoutedMetadataEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

