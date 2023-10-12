package nsxtedgegateway


type NsxtEdgegatewaySubnetAllocatedIps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#end_address NsxtEdgegateway#end_address}.
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#start_address NsxtEdgegateway#start_address}.
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

