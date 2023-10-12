//go:build no_runtime_type_checking

package vappvm

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappVmDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappVmDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappVmDiskList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappVmDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappVmDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappVmDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappVmDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

