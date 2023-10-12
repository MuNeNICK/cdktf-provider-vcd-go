package nsxtedgegateway


type NsxtEdgegatewaySubnet struct {
	// Gateway address for a subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#gateway NsxtEdgegateway#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Prefix length for a subnet (e.g. 24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#prefix_length NsxtEdgegateway#prefix_length}
	PrefixLength *float64 `field:"required" json:"prefixLength" yaml:"prefixLength"`
	// allocated_ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#allocated_ips NsxtEdgegateway#allocated_ips}
	AllocatedIps interface{} `field:"optional" json:"allocatedIps" yaml:"allocatedIps"`
	// Primary IP address for the edge gateway - will be auto-assigned if not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#primary_ip NsxtEdgegateway#primary_ip}
	PrimaryIp *string `field:"optional" json:"primaryIp" yaml:"primaryIp"`
}

