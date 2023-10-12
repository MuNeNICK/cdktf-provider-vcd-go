package externalnetworkv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalNetworkV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#name ExternalNetworkV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dedicate this External Network to an Org ID (only with IP Spaces, VCD 10.4.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#dedicated_org_id ExternalNetworkV2#dedicated_org_id}
	DedicatedOrgId *string `field:"optional" json:"dedicatedOrgId" yaml:"dedicatedOrgId"`
	// Network description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#description ExternalNetworkV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#id ExternalNetworkV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#ip_scope ExternalNetworkV2#ip_scope}
	IpScope interface{} `field:"optional" json:"ipScope" yaml:"ipScope"`
	// nsxt_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#nsxt_network ExternalNetworkV2#nsxt_network}
	NsxtNetwork *ExternalNetworkV2NsxtNetwork `field:"optional" json:"nsxtNetwork" yaml:"nsxtNetwork"`
	// Enables IP Spaces for this network (default 'false'). VCD 10.4.1+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#use_ip_spaces ExternalNetworkV2#use_ip_spaces}
	UseIpSpaces interface{} `field:"optional" json:"useIpSpaces" yaml:"useIpSpaces"`
	// vsphere_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/external_network_v2#vsphere_network ExternalNetworkV2#vsphere_network}
	VsphereNetwork interface{} `field:"optional" json:"vsphereNetwork" yaml:"vsphereNetwork"`
}

