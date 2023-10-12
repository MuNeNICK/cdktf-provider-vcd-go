//go:build no_runtime_type_checking

package datavcdresourceschema

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVcdResourceSchemaAttributesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVcdResourceSchemaAttributesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVcdResourceSchemaAttributesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVcdResourceSchemaAttributesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVcdResourceSchemaAttributesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVcdResourceSchemaAttributesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

