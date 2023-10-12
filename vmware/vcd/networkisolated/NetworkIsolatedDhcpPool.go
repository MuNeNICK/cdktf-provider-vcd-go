package networkisolated


type NetworkIsolatedDhcpPool struct {
	// The final address in the IP Range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#end_address NetworkIsolated#end_address}
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// The first address in the IP Range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#start_address NetworkIsolated#start_address}
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
	// The default DHCP lease time to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#default_lease_time NetworkIsolated#default_lease_time}
	DefaultLeaseTime *float64 `field:"optional" json:"defaultLeaseTime" yaml:"defaultLeaseTime"`
	// The maximum DHCP lease time to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#max_lease_time NetworkIsolated#max_lease_time}
	MaxLeaseTime *float64 `field:"optional" json:"maxLeaseTime" yaml:"maxLeaseTime"`
}

