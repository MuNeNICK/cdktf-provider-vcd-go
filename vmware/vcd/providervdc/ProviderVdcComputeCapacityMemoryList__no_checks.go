//go:build no_runtime_type_checking

package providervdc

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProviderVdcComputeCapacityMemoryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

