package provider


type VcdProviderIgnoreMetadataChanges struct {
	// One of 'error', 'warn' or 'none'.
	//
	// Configures whether a conflict between this ignored metadata block and the metadata entries set in Terraform should fail, warn or do nothing. Defaults to 'error'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#conflict_action VcdProvider#conflict_action}
	ConflictAction *string `field:"optional" json:"conflictAction" yaml:"conflictAction"`
	// Regular expression of the metadata entry keys to ignore. Either `key_regex` or `value_regex` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#key_regex VcdProvider#key_regex}
	KeyRegex *string `field:"optional" json:"keyRegex" yaml:"keyRegex"`
	// Ignores metadata from the specific entity in VCD named like this argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#resource_name VcdProvider#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// Ignores metadata from the specific resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#resource_type VcdProvider#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Regular expression of the metadata entry values to ignore. Either `key_regex` or `value_regex` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#value_regex VcdProvider#value_regex}
	ValueRegex *string `field:"optional" json:"valueRegex" yaml:"valueRegex"`
}

