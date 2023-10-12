package orgvdcaccesscontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgVdcAccessControlConfig struct {
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
	// Whether the VDC is shared with everyone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#shared_with_everyone OrgVdcAccessControl#shared_with_everyone}
	SharedWithEveryone interface{} `field:"required" json:"sharedWithEveryone" yaml:"sharedWithEveryone"`
	// Access level when the VDC is shared with everyone (only ReadOnly is available). Required when shared_with_everyone is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#everyone_access_level OrgVdcAccessControl#everyone_access_level}
	EveryoneAccessLevel *string `field:"optional" json:"everyoneAccessLevel" yaml:"everyoneAccessLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#id OrgVdcAccessControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#org OrgVdcAccessControl#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// shared_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#shared_with OrgVdcAccessControl#shared_with}
	SharedWith interface{} `field:"optional" json:"sharedWith" yaml:"sharedWith"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#vdc OrgVdcAccessControl#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

