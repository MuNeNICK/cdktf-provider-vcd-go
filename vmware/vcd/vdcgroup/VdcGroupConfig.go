package vdcgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VdcGroupConfig struct {
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
	// Name of VDC group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#name VdcGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Participating VDC IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#participating_vdc_ids VdcGroup#participating_vdc_ids}
	ParticipatingVdcIds *[]*string `field:"required" json:"participatingVdcIds" yaml:"participatingVdcIds"`
	// Starting VDC ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#starting_vdc_id VdcGroup#starting_vdc_id}
	StartingVdcId *string `field:"required" json:"startingVdcId" yaml:"startingVdcId"`
	// Default Policy Status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#default_policy_status VdcGroup#default_policy_status}
	DefaultPolicyStatus interface{} `field:"optional" json:"defaultPolicyStatus" yaml:"defaultPolicyStatus"`
	// VDC group description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#description VdcGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Distributed firewall status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#dfw_enabled VdcGroup#dfw_enabled}
	DfwEnabled interface{} `field:"optional" json:"dfwEnabled" yaml:"dfwEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#id VdcGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#org VdcGroup#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// A flag to remove default firewall rule when DFW and Default Policy are both enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group#remove_default_firewall_rule VdcGroup#remove_default_firewall_rule}
	RemoveDefaultFirewallRule interface{} `field:"optional" json:"removeDefaultFirewallRule" yaml:"removeDefaultFirewallRule"`
}

