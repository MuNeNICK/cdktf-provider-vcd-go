package datavcdnetworkroutedv2


type DataVcdNetworkRoutedV2Filter struct {
	// Search by IP. The value can be a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_routed_v2#ip DataVcdNetworkRoutedV2#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_routed_v2#name_regex DataVcdNetworkRoutedV2#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

