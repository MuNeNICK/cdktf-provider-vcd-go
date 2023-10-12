package edgegateway


type EdgegatewayExternalNetworkSubnet struct {
	// Gateway address for a subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#gateway Edgegateway#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Netmask address for a subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#netmask Edgegateway#netmask}
	Netmask *string `field:"required" json:"netmask" yaml:"netmask"`
	// IP address on the edge gateway - will be auto-assigned if not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#ip_address Edgegateway#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// suballocate_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#suballocate_pool Edgegateway#suballocate_pool}
	SuballocatePool interface{} `field:"optional" json:"suballocatePool" yaml:"suballocatePool"`
	// Defines if this subnet should be used as default gateway for edge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#use_for_default_route Edgegateway#use_for_default_route}
	UseForDefaultRoute interface{} `field:"optional" json:"useForDefaultRoute" yaml:"useForDefaultRoute"`
}

