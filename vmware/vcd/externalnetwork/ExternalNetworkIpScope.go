package externalnetwork


type ExternalNetworkIpScope struct {
	// Gateway of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#gateway ExternalNetwork#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Network mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#netmask ExternalNetwork#netmask}
	Netmask *string `field:"required" json:"netmask" yaml:"netmask"`
	// Primary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#dns1 ExternalNetwork#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// Secondary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#dns2 ExternalNetwork#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// DNS suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#dns_suffix ExternalNetwork#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#static_ip_pool ExternalNetwork#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
}

