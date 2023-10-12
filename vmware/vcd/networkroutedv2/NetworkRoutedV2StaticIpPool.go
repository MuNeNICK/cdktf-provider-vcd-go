package networkroutedv2


type NetworkRoutedV2StaticIpPool struct {
	// End address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#end_address NetworkRoutedV2#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#start_address NetworkRoutedV2#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

