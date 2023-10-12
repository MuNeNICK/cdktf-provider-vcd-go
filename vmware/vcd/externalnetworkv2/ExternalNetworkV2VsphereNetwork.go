package externalnetworkv2


type ExternalNetworkV2VsphereNetwork struct {
	// The name of the port group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#portgroup_id ExternalNetworkV2#portgroup_id}
	PortgroupId *string `field:"required" json:"portgroupId" yaml:"portgroupId"`
	// The vCenter server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#vcenter_id ExternalNetworkV2#vcenter_id}
	VcenterId *string `field:"required" json:"vcenterId" yaml:"vcenterId"`
}

