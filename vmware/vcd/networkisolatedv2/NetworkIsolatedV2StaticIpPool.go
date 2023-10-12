package networkisolatedv2


type NetworkIsolatedV2StaticIpPool struct {
	// End address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#end_address NetworkIsolatedV2#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#start_address NetworkIsolatedV2#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

