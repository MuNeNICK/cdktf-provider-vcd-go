package nsxtalbpool


type NsxtAlbPoolMember struct {
	// IP address of pool member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#ip_address NsxtAlbPool#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Defines if pool member is accepts traffic (default 'true').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#enabled NsxtAlbPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Member port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#port NsxtAlbPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Ratio of selecting eligible servers in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#ratio NsxtAlbPool#ratio}
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

