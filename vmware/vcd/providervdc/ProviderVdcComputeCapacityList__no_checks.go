//go:build no_runtime_type_checking

package providervdc

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProviderVdcComputeCapacityList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProviderVdcComputeCapacityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProviderVdcComputeCapacityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

