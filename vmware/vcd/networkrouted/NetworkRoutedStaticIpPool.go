package networkrouted


type NetworkRoutedStaticIpPool struct {
	// The final address in the IP Range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed#end_address NetworkRouted#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// The first address in the IP Range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed#start_address NetworkRouted#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

