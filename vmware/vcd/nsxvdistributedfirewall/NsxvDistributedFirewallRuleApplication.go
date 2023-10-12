package nsxvdistributedfirewall


type NsxvDistributedFirewallRuleApplication struct {
	// Destination port for this application. Leaving it empty means 'any' port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#destination_port NsxvDistributedFirewall#destination_port}
	DestinationPort *string `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Name of application (Application, ApplicationGroup).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#name NsxvDistributedFirewall#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol of the application (one of TCP, UDP, ICMP) (When not using name/value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#protocol NsxvDistributedFirewall#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Source port for this application. Leaving it empty means 'any' port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#source_port NsxvDistributedFirewall#source_port}
	SourcePort *string `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// Type of application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#type NsxvDistributedFirewall#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Value of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#value NsxvDistributedFirewall#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

