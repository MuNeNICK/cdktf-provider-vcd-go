package datavcdnetworkisolated


type DataVcdNetworkIsolatedFilter struct {
	// Search by IP. The value can be a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_isolated#ip DataVcdNetworkIsolated#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_isolated#metadata DataVcdNetworkIsolated#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_isolated#name_regex DataVcdNetworkIsolated#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

