//go:build no_runtime_type_checking

package vappvm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappVmOverrideTemplateDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappVmOverrideTemplateDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappVmOverrideTemplateDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappVmOverrideTemplateDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappVmOverrideTemplateDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappVmOverrideTemplateDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappVmOverrideTemplateDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

