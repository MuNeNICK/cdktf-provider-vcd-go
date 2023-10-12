package nsxvdistributedfirewall


type NsxvDistributedFirewallRuleDestination struct {
	// Name of the destination entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#name NsxvDistributedFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the destination entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#type NsxvDistributedFirewall#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the destination entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#value NsxvDistributedFirewall#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

