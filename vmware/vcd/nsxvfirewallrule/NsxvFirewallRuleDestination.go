package nsxvfirewallrule


type NsxvFirewallRuleDestination struct {
	// Rule is applied to traffic going to any destinations except for the excluded destination. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#exclude NsxvFirewallRule#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// 'vse', 'internal', 'external' or network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#gateway_interfaces NsxvFirewallRule#gateway_interfaces}
	GatewayInterfaces *[]*string `field:"optional" json:"gatewayInterfaces" yaml:"gatewayInterfaces"`
	// IP address, CIDR, an IP range, or the keyword 'any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#ip_addresses NsxvFirewallRule#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Set of IP set names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#ip_sets NsxvFirewallRule#ip_sets}
	IpSets *[]*string `field:"optional" json:"ipSets" yaml:"ipSets"`
	// Set of org network names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#org_networks NsxvFirewallRule#org_networks}
	OrgNetworks *[]*string `field:"optional" json:"orgNetworks" yaml:"orgNetworks"`
	// Set of VM IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#vm_ids NsxvFirewallRule#vm_ids}
	VmIds *[]*string `field:"optional" json:"vmIds" yaml:"vmIds"`
}

