//go:build no_runtime_type_checking

package providervdc

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProviderVdcComputeCapacityCpuListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

