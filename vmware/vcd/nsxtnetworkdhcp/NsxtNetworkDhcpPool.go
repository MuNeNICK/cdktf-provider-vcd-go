package nsxtnetworkdhcp


type NsxtNetworkDhcpPool struct {
	// End address of DHCP pool IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#end_address NsxtNetworkDhcp#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of DHCP pool IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#start_address NsxtNetworkDhcp#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

