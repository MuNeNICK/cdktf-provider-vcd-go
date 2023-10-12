package ipspace


type IpSpaceIpPrefix struct {
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#prefix IpSpace#prefix}
	Prefix interface{} `field:"required" json:"prefix" yaml:"prefix"`
	// Floating IP quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#default_quota IpSpace#default_quota}
	DefaultQuota *string `field:"optional" json:"defaultQuota" yaml:"defaultQuota"`
}

