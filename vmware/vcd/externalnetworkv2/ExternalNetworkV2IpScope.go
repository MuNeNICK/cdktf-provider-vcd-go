package externalnetworkv2


type ExternalNetworkV2IpScope struct {
	// Gateway of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#gateway ExternalNetworkV2#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Network mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#prefix_length ExternalNetworkV2#prefix_length}
	PrefixLength *float64 `field:"required" json:"prefixLength" yaml:"prefixLength"`
	// Primary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#dns1 ExternalNetworkV2#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// Secondary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#dns2 ExternalNetworkV2#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// DNS suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#dns_suffix ExternalNetworkV2#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// If subnet is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#enabled ExternalNetworkV2#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#static_ip_pool ExternalNetworkV2#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
}

