//go:build no_runtime_type_checking

package lbserverpool

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbServerPoolMemberList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbServerPoolMemberList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbServerPoolMemberList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbServerPoolMemberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbServerPoolMemberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbServerPoolMemberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbServerPoolMemberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

