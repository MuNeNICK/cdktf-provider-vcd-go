package nsxtedgegatewaybgpipprefixlist


type NsxtEdgegatewayBgpIpPrefixListIpPrefix struct {
	// Action 'PERMIT' or 'DENY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list#action NsxtEdgegatewayBgpIpPrefixList#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Network in CIDR notation (e.g. '192.168.100.0/24', '2001:db8::/48').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list#network NsxtEdgegatewayBgpIpPrefixList#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Greater than or equal to subnet mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list#greater_than_or_equal_to NsxtEdgegatewayBgpIpPrefixList#greater_than_or_equal_to}
	GreaterThanOrEqualTo *float64 `field:"optional" json:"greaterThanOrEqualTo" yaml:"greaterThanOrEqualTo"`
	// Less than or equal to subnet mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list#less_than_or_equal_to NsxtEdgegatewayBgpIpPrefixList#less_than_or_equal_to}
	LessThanOrEqualTo *float64 `field:"optional" json:"lessThanOrEqualTo" yaml:"lessThanOrEqualTo"`
}

