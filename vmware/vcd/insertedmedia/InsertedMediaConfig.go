package insertedmedia

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InsertedMediaConfig struct {
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
	// catalog name where to find media file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#catalog InsertedMedia#catalog}
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// media name to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#name InsertedMedia#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// vApp to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#vapp_name InsertedMedia#vapp_name}
	VappName *string `field:"required" json:"vappName" yaml:"vappName"`
	// VM in vApp in which media will be inserted or ejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#vm_name InsertedMedia#vm_name}
	VmName *string `field:"required" json:"vmName" yaml:"vmName"`
	// When ejecting answers automatically to question yes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#eject_force InsertedMedia#eject_force}
	EjectForce interface{} `field:"optional" json:"ejectForce" yaml:"ejectForce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#id InsertedMedia#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#org InsertedMedia#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/inserted_media#vdc InsertedMedia#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

