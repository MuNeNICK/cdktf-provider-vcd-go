//go:build no_runtime_type_checking

package vm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VmDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

