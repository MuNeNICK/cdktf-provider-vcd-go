package datavcdnetworkdirect


type DataVcdNetworkDirectFilter struct {
	// Search by IP. The value can be a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct#ip DataVcdNetworkDirect#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct#metadata DataVcdNetworkDirect#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct#name_regex DataVcdNetworkDirect#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

