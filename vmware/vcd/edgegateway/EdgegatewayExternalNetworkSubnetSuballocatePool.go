package edgegateway


type EdgegatewayExternalNetworkSubnetSuballocatePool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#end_address Edgegateway#end_address}.
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#start_address Edgegateway#start_address}.
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

