package externalnetwork


type ExternalNetworkIpScopeStaticIpPool struct {
	// End address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#end_address ExternalNetwork#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#start_address ExternalNetwork#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

