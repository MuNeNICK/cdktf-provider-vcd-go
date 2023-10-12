package nsxtedgegatewaystaticroute


type NsxtEdgegatewayStaticRouteNextHop struct {
	// Admin distance of next hop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_static_route#admin_distance NsxtEdgegatewayStaticRoute#admin_distance}
	AdminDistance *float64 `field:"required" json:"adminDistance" yaml:"adminDistance"`
	// IP Address of next hop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_static_route#ip_address NsxtEdgegatewayStaticRoute#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_static_route#scope NsxtEdgegatewayStaticRoute#scope}
	Scope *NsxtEdgegatewayStaticRouteNextHopScope `field:"optional" json:"scope" yaml:"scope"`
}

