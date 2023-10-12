package nsxvdistributedfirewall


type NsxvDistributedFirewallRuleSource struct {
	// Name of the source entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#name NsxvDistributedFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the source entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#type NsxvDistributedFirewall#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the source entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#value NsxvDistributedFirewall#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

