package nsxtalbvirtualservice


type NsxtAlbVirtualServiceServicePort struct {
	// Starting port in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#start_port NsxtAlbVirtualService#start_port}
	StartPort *float64 `field:"required" json:"startPort" yaml:"startPort"`
	// One of 'TCP_PROXY', 'TCP_FAST_PATH', 'UDP_FAST_PATH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#type NsxtAlbVirtualService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Last port in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#end_port NsxtAlbVirtualService#end_port}
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// Defines if certificate should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#ssl_enabled NsxtAlbVirtualService#ssl_enabled}
	SslEnabled interface{} `field:"optional" json:"sslEnabled" yaml:"sslEnabled"`
}

