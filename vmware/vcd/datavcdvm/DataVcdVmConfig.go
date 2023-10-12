package datavcdvm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdVmConfig struct {
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
	// A name for the VM, unique within the vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#name DataVcdVm#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#id DataVcdVm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional number of seconds to try and wait for DHCP IP (valid for 'network' block only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#network_dhcp_wait_seconds DataVcdVm#network_dhcp_wait_seconds}
	NetworkDhcpWaitSeconds *float64 `field:"optional" json:"networkDhcpWaitSeconds" yaml:"networkDhcpWaitSeconds"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#org DataVcdVm#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// VM placement policy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#placement_policy_id DataVcdVm#placement_policy_id}
	PlacementPolicyId *string `field:"optional" json:"placementPolicyId" yaml:"placementPolicyId"`
	// VM sizing policy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#sizing_policy_id DataVcdVm#sizing_policy_id}
	SizingPolicyId *string `field:"optional" json:"sizingPolicyId" yaml:"sizingPolicyId"`
	// The vApp this VM belongs to - Required, unless it is a standalone VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#vapp_name DataVcdVm#vapp_name}
	VappName *string `field:"optional" json:"vappName" yaml:"vappName"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm#vdc DataVcdVm#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

