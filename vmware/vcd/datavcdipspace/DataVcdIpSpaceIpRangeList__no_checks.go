//go:build no_runtime_type_checking

package datavcdipspace

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdIpSpaceIpRangeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdIpSpaceIpRangeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpRangeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpRangeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdIpSpaceIpRangeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdIpSpaceIpRangeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

