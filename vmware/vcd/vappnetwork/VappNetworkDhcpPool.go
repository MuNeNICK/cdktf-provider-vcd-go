package vappnetwork


type VappNetworkDhcpPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#start_address VappNetwork#start_address}.
	StartAddress *string `field:"required" json:"startAddress" yaml:"startAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#default_lease_time VappNetwork#default_lease_time}.
	DefaultLeaseTime *float64 `field:"optional" json:"defaultLeaseTime" yaml:"defaultLeaseTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#enabled VappNetwork#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#end_address VappNetwork#end_address}.
	EndAddress *string `field:"optional" json:"endAddress" yaml:"endAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#max_lease_time VappNetwork#max_lease_time}.
	MaxLeaseTime *float64 `field:"optional" json:"maxLeaseTime" yaml:"maxLeaseTime"`
}

