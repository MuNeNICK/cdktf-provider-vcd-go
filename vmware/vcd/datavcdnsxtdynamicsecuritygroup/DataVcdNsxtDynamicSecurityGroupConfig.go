package datavcdnsxtdynamicsecuritygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtDynamicSecurityGroupConfig struct {
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
	// Dynamic Security Group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_dynamic_security_group#name DataVcdNsxtDynamicSecurityGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// VDC Group ID in which Dynamic Security Group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_dynamic_security_group#vdc_group_id DataVcdNsxtDynamicSecurityGroup#vdc_group_id}
	VdcGroupId *string `field:"required" json:"vdcGroupId" yaml:"vdcGroupId"`
	// Dynamic Security Group description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_dynamic_security_group#description DataVcdNsxtDynamicSecurityGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_dynamic_security_group#id DataVcdNsxtDynamicSecurityGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_dynamic_security_group#org DataVcdNsxtDynamicSecurityGroup#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
}

