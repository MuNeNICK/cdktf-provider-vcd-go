package externalnetwork


type ExternalNetworkVsphereNetwork struct {
	// The name of the port group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#name ExternalNetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The vSphere port group type. One of: DV_PORTGROUP (distributed virtual port group), NETWORK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#type ExternalNetwork#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The vCenter server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network#vcenter ExternalNetwork#vcenter}
	Vcenter *string `field:"required" json:"vcenter" yaml:"vcenter"`
}

