//go:build no_runtime_type_checking

package vappaccesscontrol

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappAccessControlSharedWithList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappAccessControlSharedWithList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappAccessControlSharedWithList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappAccessControlSharedWithList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappAccessControlSharedWithList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappAccessControlSharedWithList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappAccessControlSharedWithListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

