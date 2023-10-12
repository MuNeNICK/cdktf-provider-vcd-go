package vapporgnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappOrgNetworkConfig struct {
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
	// Organization network name to which vApp network is connected to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#org_network_name VappOrgNetwork#org_network_name}
	OrgNetworkName *string `field:"required" json:"orgNetworkName" yaml:"orgNetworkName"`
	// vApp network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#vapp_name VappOrgNetwork#vapp_name}
	VappName *string `field:"required" json:"vappName" yaml:"vappName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#id VappOrgNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are accessed in this vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#is_fenced VappOrgNetwork#is_fenced}
	IsFenced interface{} `field:"optional" json:"isFenced" yaml:"isFenced"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#org VappOrgNetwork#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#reboot_vapp_on_removal VappOrgNetwork#reboot_vapp_on_removal}
	RebootVappOnRemoval interface{} `field:"optional" json:"rebootVappOnRemoval" yaml:"rebootVappOnRemoval"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#retain_ip_mac_enabled VappOrgNetwork#retain_ip_mac_enabled}
	RetainIpMacEnabled interface{} `field:"optional" json:"retainIpMacEnabled" yaml:"retainIpMacEnabled"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_org_network#vdc VappOrgNetwork#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

