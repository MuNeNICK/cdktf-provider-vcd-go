//go:build no_runtime_type_checking

package vm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmOverrideTemplateDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmOverrideTemplateDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmOverrideTemplateDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VmOverrideTemplateDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmOverrideTemplateDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmOverrideTemplateDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmOverrideTemplateDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

