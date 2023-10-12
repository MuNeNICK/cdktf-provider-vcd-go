//go:build no_runtime_type_checking

package vm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmInternalDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmInternalDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmInternalDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmInternalDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmInternalDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmInternalDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

