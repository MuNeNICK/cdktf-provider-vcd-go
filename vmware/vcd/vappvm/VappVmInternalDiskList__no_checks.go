//go:build no_runtime_type_checking

package vappvm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappVmInternalDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappVmInternalDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappVmInternalDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappVmInternalDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappVmInternalDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappVmInternalDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

