package catalogaccesscontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogAccessControlConfig struct {
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
	// The ID of Catalog to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#catalog_id CatalogAccessControl#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Whether the Catalog is shared with everyone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#shared_with_everyone CatalogAccessControl#shared_with_everyone}
	SharedWithEveryone interface{} `field:"required" json:"sharedWithEveryone" yaml:"sharedWithEveryone"`
	// Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#everyone_access_level CatalogAccessControl#everyone_access_level}
	EveryoneAccessLevel *string `field:"optional" json:"everyoneAccessLevel" yaml:"everyoneAccessLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#id CatalogAccessControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#org CatalogAccessControl#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// shared_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#shared_with CatalogAccessControl#shared_with}
	SharedWith interface{} `field:"optional" json:"sharedWith" yaml:"sharedWith"`
}

