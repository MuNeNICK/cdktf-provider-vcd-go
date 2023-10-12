package vappnetwork


type VappNetworkStaticIpPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#end_address VappNetwork#end_address}.
	EndAddress *string `field:"required" json:"endAddress" yaml:"endAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#start_address VappNetwork#start_address}.
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
}

