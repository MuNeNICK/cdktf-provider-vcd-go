package nsxtdynamicsecuritygroup


type NsxtDynamicSecurityGroupCriteriaRule struct {
	// Operator can be one of 'EQUALS', 'CONTAINS', 'STARTS_WITH', 'ENDS_WITH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_dynamic_security_group#operator NsxtDynamicSecurityGroup#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Type of object matching 'VM_TAG' or 'VM_NAME'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_dynamic_security_group#type NsxtDynamicSecurityGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_dynamic_security_group#value NsxtDynamicSecurityGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

