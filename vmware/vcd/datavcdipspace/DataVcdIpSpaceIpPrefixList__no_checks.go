//go:build no_runtime_type_checking

package datavcdipspace

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdIpSpaceIpPrefixList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdIpSpaceIpPrefixList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpPrefixList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpPrefixList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpPrefixList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdIpSpaceIpPrefixListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

