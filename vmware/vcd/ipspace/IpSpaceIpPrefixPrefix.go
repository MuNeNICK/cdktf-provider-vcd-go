package ipspace


type IpSpaceIpPrefixPrefix struct {
	// First IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#first_ip IpSpace#first_ip}
	FirstIp *string `field:"required" json:"firstIp" yaml:"firstIp"`
	// Number of prefixes to define.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#prefix_count IpSpace#prefix_count}
	PrefixCount *string `field:"required" json:"prefixCount" yaml:"prefixCount"`
	// Prefix length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#prefix_length IpSpace#prefix_length}
	PrefixLength *string `field:"required" json:"prefixLength" yaml:"prefixLength"`
}

