package ipspace


type IpSpaceIpRange struct {
	// End address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#end_address IpSpace#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Start address of the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#start_address IpSpace#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

