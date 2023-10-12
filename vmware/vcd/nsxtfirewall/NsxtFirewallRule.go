package nsxtfirewall


type NsxtFirewallRule struct {
	// Defines if the rule should 'ALLOW' or 'DROP' matching traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#action NsxtFirewall#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Direction on which Firewall Rule applies (One of 'IN', 'OUT', 'IN_OUT').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#direction NsxtFirewall#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Firewall Rule Protocol (One of 'IPV4', 'IPV6', 'IPV4_IPV6').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#ip_protocol NsxtFirewall#ip_protocol}
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// Firewall Rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#name NsxtFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of Application Port Profile IDs. Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#app_port_profile_ids NsxtFirewall#app_port_profile_ids}
	AppPortProfileIds *[]*string `field:"optional" json:"appPortProfileIds" yaml:"appPortProfileIds"`
	// A set of Destination Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#destination_ids NsxtFirewall#destination_ids}
	DestinationIds *[]*string `field:"optional" json:"destinationIds" yaml:"destinationIds"`
	// Defined if Firewall Rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#enabled NsxtFirewall#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Defines if matching traffic should be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#logging NsxtFirewall#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A set of Source Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_firewall#source_ids NsxtFirewall#source_ids}
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
}

