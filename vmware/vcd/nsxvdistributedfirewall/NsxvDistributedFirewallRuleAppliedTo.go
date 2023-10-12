package nsxvdistributedfirewall


type NsxvDistributedFirewallRuleAppliedTo struct {
	// Name of the applied-to entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#name NsxvDistributedFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the applied-to entity (one of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#type NsxvDistributedFirewall#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the applied-to entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#value NsxvDistributedFirewall#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

