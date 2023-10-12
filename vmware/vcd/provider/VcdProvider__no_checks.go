//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VcdProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VcdProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVcdProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVcdProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVcdProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VcdProvider) validateSetAllowApiTokenFileParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcdProvider) validateSetAllowServiceAccountTokenFileParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcdProvider) validateSetAllowUnverifiedSslParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcdProvider) validateSetIgnoreMetadataChangesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VcdProvider) validateSetLoggingParameters(val interface{}) error {
	return nil
}

func validateNewVcdProviderParameters(scope constructs.Construct, id *string, config *VcdProviderConfig) error {
	return nil
}

