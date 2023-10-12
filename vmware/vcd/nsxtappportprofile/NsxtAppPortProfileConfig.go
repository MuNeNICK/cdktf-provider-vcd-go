package nsxtappportprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAppPortProfileConfig struct {
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
	// app_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#app_port NsxtAppPortProfile#app_port}
	AppPort interface{} `field:"required" json:"appPort" yaml:"appPort"`
	// Application Port Profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#name NsxtAppPortProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Scope - 'PROVIDER' or 'TENANT'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#scope NsxtAppPortProfile#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// ID of VDC, VDC Group, or NSX-T Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#context_id NsxtAppPortProfile#context_id}
	ContextId *string `field:"optional" json:"contextId" yaml:"contextId"`
	// Application Port Profile description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#description NsxtAppPortProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#id NsxtAppPortProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of NSX-T manager. Only required for 'PROVIDER' scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#nsxt_manager_id NsxtAppPortProfile#nsxt_manager_id}
	NsxtManagerId *string `field:"optional" json:"nsxtManagerId" yaml:"nsxtManagerId"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#org NsxtAppPortProfile#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#vdc NsxtAppPortProfile#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

