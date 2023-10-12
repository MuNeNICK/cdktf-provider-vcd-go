package nsxvdistributedfirewall


type NsxvDistributedFirewallRule struct {
	// Action of the rule (allow, deny).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#action NsxvDistributedFirewall#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// applied_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#applied_to NsxvDistributedFirewall#applied_to}
	AppliedTo interface{} `field:"required" json:"appliedTo" yaml:"appliedTo"`
	// Direction of the rule (in, out, inout).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#direction NsxvDistributedFirewall#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#application NsxvDistributedFirewall#application}
	Application interface{} `field:"optional" json:"application" yaml:"application"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#destination NsxvDistributedFirewall#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Whether the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#enabled NsxvDistributedFirewall#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If true, the content of the destination elements is reversed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#exclude_destination NsxvDistributedFirewall#exclude_destination}
	ExcludeDestination interface{} `field:"optional" json:"excludeDestination" yaml:"excludeDestination"`
	// If true, the content of the source elements is reversed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#exclude_source NsxvDistributedFirewall#exclude_source}
	ExcludeSource interface{} `field:"optional" json:"excludeSource" yaml:"excludeSource"`
	// Whether the rule traffic is logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#logged NsxvDistributedFirewall#logged}
	Logged interface{} `field:"optional" json:"logged" yaml:"logged"`
	// Firewall Rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#name NsxvDistributedFirewall#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Packet type of the rule (any, ipv4, ipv6).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#packet_type NsxvDistributedFirewall#packet_type}
	PacketType *string `field:"optional" json:"packetType" yaml:"packetType"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_distributed_firewall#source NsxvDistributedFirewall#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

