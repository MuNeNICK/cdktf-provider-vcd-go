package nsxtdistributedfirewall


type NsxtDistributedFirewallRule struct {
	// Defines if the rule should 'ALLOW', 'DROP', 'REJECT' matching traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#action NsxtDistributedFirewall#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Firewall Rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#name NsxtDistributedFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of Application Port Profile IDs. Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#app_port_profile_ids NsxtDistributedFirewall#app_port_profile_ids}
	AppPortProfileIds *[]*string `field:"optional" json:"appPortProfileIds" yaml:"appPortProfileIds"`
	// Comment that is shown next to rule in UI (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#comment NsxtDistributedFirewall#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Description is not shown in UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#description NsxtDistributedFirewall#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Reverses firewall matching for to match all except Destinations Groups specified in 'destination_ids' (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#destination_groups_excluded NsxtDistributedFirewall#destination_groups_excluded}
	DestinationGroupsExcluded interface{} `field:"optional" json:"destinationGroupsExcluded" yaml:"destinationGroupsExcluded"`
	// A set of Destination Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#destination_ids NsxtDistributedFirewall#destination_ids}
	DestinationIds *[]*string `field:"optional" json:"destinationIds" yaml:"destinationIds"`
	// Direction on which Firewall Rule applies (One of 'IN', 'OUT', 'IN_OUT').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#direction NsxtDistributedFirewall#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Defined if Firewall Rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#enabled NsxtDistributedFirewall#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Firewall Rule Protocol (One of 'IPV4', 'IPV6', 'IPV4_IPV6').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#ip_protocol NsxtDistributedFirewall#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// Defines if matching traffic should be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#logging NsxtDistributedFirewall#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A set of Network Context Profile IDs. Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#network_context_profile_ids NsxtDistributedFirewall#network_context_profile_ids}
	NetworkContextProfileIds *[]*string `field:"optional" json:"networkContextProfileIds" yaml:"networkContextProfileIds"`
	// Reverses firewall matching for to match all except Source Groups specified in 'source_ids' (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#source_groups_excluded NsxtDistributedFirewall#source_groups_excluded}
	SourceGroupsExcluded interface{} `field:"optional" json:"sourceGroupsExcluded" yaml:"sourceGroupsExcluded"`
	// A set of Source Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall#source_ids NsxtDistributedFirewall#source_ids}
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
}

