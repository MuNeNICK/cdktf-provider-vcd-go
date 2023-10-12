package vappstaticrouting


type VappStaticRoutingRule struct {
	// Name for the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_static_routing#name VappStaticRouting#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network specification in CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_static_routing#network_cidr VappStaticRouting#network_cidr}
	NetworkCidr *string `field:"required" json:"networkCidr" yaml:"networkCidr"`
	// IP Address of Next Hop router/gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_static_routing#next_hop_ip VappStaticRouting#next_hop_ip}
	NextHopIp *string `field:"required" json:"nextHopIp" yaml:"nextHopIp"`
}

