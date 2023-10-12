package datavcdresourcelist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdResourceListConfig struct {
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
	// Unique name of the Info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#name DataVcdResourceList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Which resource we should list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#resource_type DataVcdResourceList#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#id DataVcdResourceList#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// How the list should be built.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#list_mode DataVcdResourceList#list_mode}
	ListMode *string `field:"optional" json:"listMode" yaml:"listMode"`
	// Separator for name_id combination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#name_id_separator DataVcdResourceList#name_id_separator}
	NameIdSeparator *string `field:"optional" json:"nameIdSeparator" yaml:"nameIdSeparator"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#org DataVcdResourceList#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of the parent to the resources being retrieved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#parent DataVcdResourceList#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/resource_list#vdc DataVcdResourceList#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

