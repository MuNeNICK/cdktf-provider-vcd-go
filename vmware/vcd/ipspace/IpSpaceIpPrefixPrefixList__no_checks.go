//go:build no_runtime_type_checking

package ipspace

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpSpaceIpPrefixPrefixList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IpSpaceIpPrefixPrefixList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixPrefixList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixPrefixList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixPrefixList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixPrefixList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIpSpaceIpPrefixPrefixListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

