//go:build no_runtime_type_checking

package vappnatrules

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VappNatRulesRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VappNatRulesRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VappNatRulesRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VappNatRulesRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VappNatRulesRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VappNatRulesRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVappNatRulesRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

