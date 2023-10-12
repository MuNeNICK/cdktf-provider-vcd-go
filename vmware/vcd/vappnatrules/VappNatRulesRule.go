package vappnatrules


type VappNatRulesRule struct {
	// VM to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#vm_id VappNatRules#vm_id}
	VmId *string `field:"required" json:"vmId" yaml:"vmId"`
	// VM NIC ID to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#vm_nic_id VappNatRules#vm_nic_id}
	VmNicId *float64 `field:"required" json:"vmNicId" yaml:"vmNicId"`
	// External IP address to forward to or External IP address to map to VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#external_ip VappNatRules#external_ip}
	ExternalIp *string `field:"optional" json:"externalIp" yaml:"externalIp"`
	// External port to forward.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#external_port VappNatRules#external_port}
	ExternalPort *float64 `field:"optional" json:"externalPort" yaml:"externalPort"`
	// Internal port to forward.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#forward_to_port VappNatRules#forward_to_port}
	ForwardToPort *float64 `field:"optional" json:"forwardToPort" yaml:"forwardToPort"`
	// Mapping mode. One of: `automatic`, `manual`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#mapping_mode VappNatRules#mapping_mode}
	MappingMode *string `field:"optional" json:"mappingMode" yaml:"mappingMode"`
	// Protocol to forward. One of: `TCP` (forward TCP packets), `UDP` (forward UDP packets), `TCP_UDP` (forward TCP and UDP packets).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#protocol VappNatRules#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

