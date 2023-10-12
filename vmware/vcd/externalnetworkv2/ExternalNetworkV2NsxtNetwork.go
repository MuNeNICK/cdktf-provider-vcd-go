package externalnetworkv2


type ExternalNetworkV2NsxtNetwork struct {
	// ID of NSX-T manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#nsxt_manager_id ExternalNetworkV2#nsxt_manager_id}
	NsxtManagerId *string `field:"required" json:"nsxtManagerId" yaml:"nsxtManagerId"`
	// Name of NSX-T segment (for NSX-T segment backed external network).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#nsxt_segment_name ExternalNetworkV2#nsxt_segment_name}
	NsxtSegmentName *string `field:"optional" json:"nsxtSegmentName" yaml:"nsxtSegmentName"`
	// ID of NSX-T Tier-0 router (for T0 gateway backed external network).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#nsxt_tier0_router_id ExternalNetworkV2#nsxt_tier0_router_id}
	NsxtTier0RouterId *string `field:"optional" json:"nsxtTier0RouterId" yaml:"nsxtTier0RouterId"`
}

