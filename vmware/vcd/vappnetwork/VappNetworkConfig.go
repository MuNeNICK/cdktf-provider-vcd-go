package vappnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappNetworkConfig struct {
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
	// Gateway of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#gateway VappNetwork#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// vApp network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#name VappNetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// vApp to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#vapp_name VappNetwork#vapp_name}
	VappName *string `field:"required" json:"vappName" yaml:"vappName"`
	// Optional description for the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#description VappNetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dhcp_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#dhcp_pool VappNetwork#dhcp_pool}
	DhcpPool interface{} `field:"optional" json:"dhcpPool" yaml:"dhcpPool"`
	// Primary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#dns1 VappNetwork#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// Secondary DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#dns2 VappNetwork#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// DNS suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#dns_suffix VappNetwork#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// True if Network allows guest VLAN tagging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#guest_vlan_allowed VappNetwork#guest_vlan_allowed}
	GuestVlanAllowed interface{} `field:"optional" json:"guestVlanAllowed" yaml:"guestVlanAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#id VappNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Netmask address for a subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#netmask VappNetwork#netmask}
	Netmask *string `field:"optional" json:"netmask" yaml:"netmask"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#org VappNetwork#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// org network name to which vapp network is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#org_network_name VappNetwork#org_network_name}
	OrgNetworkName *string `field:"optional" json:"orgNetworkName" yaml:"orgNetworkName"`
	// Prefix length for a subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#prefix_length VappNetwork#prefix_length}
	PrefixLength *string `field:"optional" json:"prefixLength" yaml:"prefixLength"`
	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#reboot_vapp_on_removal VappNetwork#reboot_vapp_on_removal}
	RebootVappOnRemoval interface{} `field:"optional" json:"rebootVappOnRemoval" yaml:"rebootVappOnRemoval"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#retain_ip_mac_enabled VappNetwork#retain_ip_mac_enabled}
	RetainIpMacEnabled interface{} `field:"optional" json:"retainIpMacEnabled" yaml:"retainIpMacEnabled"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#static_ip_pool VappNetwork#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network#vdc VappNetwork#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

