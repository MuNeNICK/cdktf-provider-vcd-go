package datavcdnsxtnetworkimported


type DataVcdNsxtNetworkImportedFilter struct {
	// Search by IP. The value can be a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_network_imported#ip DataVcdNsxtNetworkImported#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_network_imported#name_regex DataVcdNsxtNetworkImported#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

