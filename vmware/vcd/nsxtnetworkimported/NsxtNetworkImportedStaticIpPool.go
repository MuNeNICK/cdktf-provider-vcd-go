package nsxtnetworkimported


type NsxtNetworkImportedStaticIpPool struct {
	// End address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#end_address NsxtNetworkImported#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#start_address NsxtNetworkImported#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

