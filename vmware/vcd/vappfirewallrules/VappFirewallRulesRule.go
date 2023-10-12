package vappfirewallrules


type VappFirewallRulesRule struct {
	// Destination IP address to which the rule applies. A value of `Any` matches any IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#destination_ip VappFirewallRules#destination_ip}
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// Destination port to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#destination_port VappFirewallRules#destination_port}
	DestinationPort *string `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Destination VM identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#destination_vm_id VappFirewallRules#destination_vm_id}
	DestinationVmId *string `field:"optional" json:"destinationVmId" yaml:"destinationVmId"`
	// The value can be one of: `assigned` - assigned internal IP will be automatically chosen.
	//
	// `NAT`: NATed external IP will be automatically chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#destination_vm_ip_type VappFirewallRules#destination_vm_ip_type}
	DestinationVmIpType *string `field:"optional" json:"destinationVmIpType" yaml:"destinationVmIpType"`
	// Destination VM NIC ID to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#destination_vm_nic_id VappFirewallRules#destination_vm_nic_id}
	DestinationVmNicId *float64 `field:"optional" json:"destinationVmNicId" yaml:"destinationVmNicId"`
	// 'true' value will enable firewall rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#enabled VappFirewallRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// 'true' value will enable rule logging. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#enable_logging VappFirewallRules#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#name VappFirewallRules#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// One of: `drop` (drop packets that match the rule), `allow` (allow packets that match the rule to pass through the firewall).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#policy VappFirewallRules#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Specify the protocols to which the rule should be applied. One of: `any`, `icmp`, `tcp`, `udp`, `tcp&udp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#protocol VappFirewallRules#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Source IP address to which the rule applies. A value of `Any` matches any IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#source_ip VappFirewallRules#source_ip}
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Source port to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#source_port VappFirewallRules#source_port}
	SourcePort *string `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// Source VM identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#source_vm_id VappFirewallRules#source_vm_id}
	SourceVmId *string `field:"optional" json:"sourceVmId" yaml:"sourceVmId"`
	// The value can be one of: `assigned` - assigned internal IP will be automatically chosen.
	//
	// `NAT`: NATed external IP will be automatically chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#source_vm_ip_type VappFirewallRules#source_vm_ip_type}
	SourceVmIpType *string `field:"optional" json:"sourceVmIpType" yaml:"sourceVmIpType"`
	// Source VM NIC ID to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#source_vm_nic_id VappFirewallRules#source_vm_nic_id}
	SourceVmNicId *float64 `field:"optional" json:"sourceVmNicId" yaml:"sourceVmNicId"`
}

