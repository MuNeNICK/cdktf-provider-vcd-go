package nsxtdynamicsecuritygroup


type NsxtDynamicSecurityGroupCriteria struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_dynamic_security_group#rule NsxtDynamicSecurityGroup#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

