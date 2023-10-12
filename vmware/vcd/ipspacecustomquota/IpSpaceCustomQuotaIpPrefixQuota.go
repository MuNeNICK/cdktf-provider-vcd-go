package ipspacecustomquota


type IpSpaceCustomQuotaIpPrefixQuota struct {
	// Prefix length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#prefix_length IpSpaceCustomQuota#prefix_length}
	PrefixLength *string `field:"required" json:"prefixLength" yaml:"prefixLength"`
	// IP Prefix Quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#quota IpSpaceCustomQuota#quota}
	Quota *string `field:"required" json:"quota" yaml:"quota"`
}

