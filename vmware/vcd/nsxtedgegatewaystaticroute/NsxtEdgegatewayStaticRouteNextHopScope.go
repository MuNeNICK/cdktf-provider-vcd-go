package nsxtedgegatewaystaticroute


type NsxtEdgegatewayStaticRouteNextHopScope struct {
	// ID of Scope element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_static_route#id NsxtEdgegatewayStaticRoute#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Scope type - One of 'NETWORK', 'SYSTEM_OWNED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_static_route#type NsxtEdgegatewayStaticRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

