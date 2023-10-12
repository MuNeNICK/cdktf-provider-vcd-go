package datavcdnsxtappportprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtAppPortProfileConfig struct {
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
	// Application Port Profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#name DataVcdNsxtAppPortProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Scope - 'SYSTEM', 'PROVIDER' or 'TENANT'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#scope DataVcdNsxtAppPortProfile#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// ID of VDC, VDC Group, or NSX-T Manager. Required if the VCD instance has more than one NSX-T manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#context_id DataVcdNsxtAppPortProfile#context_id}
	ContextId *string `field:"optional" json:"contextId" yaml:"contextId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#id DataVcdNsxtAppPortProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of NSX-T manager. Only required for 'PROVIDER' scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#nsxt_manager_id DataVcdNsxtAppPortProfile#nsxt_manager_id}
	NsxtManagerId *string `field:"optional" json:"nsxtManagerId" yaml:"nsxtManagerId"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#org DataVcdNsxtAppPortProfile#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_app_port_profile#vdc DataVcdNsxtAppPortProfile#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

