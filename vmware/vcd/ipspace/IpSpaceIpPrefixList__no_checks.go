//go:build no_runtime_type_checking

package ipspace

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpSpaceIpPrefixList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IpSpaceIpPrefixList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IpSpaceIpPrefixList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIpSpaceIpPrefixListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

