//go:build no_runtime_type_checking

package datavcdlbserverpool

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdLbServerPoolMemberList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdLbServerPoolMemberList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdLbServerPoolMemberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdLbServerPoolMemberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdLbServerPoolMemberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdLbServerPoolMemberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

