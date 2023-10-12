package datavcdcatalogmedia

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdCatalogMediaConfig struct {
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
	// catalog name where upload the Media file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#catalog DataVcdCatalogMedia#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// ID of the catalog from where to retrieve the Media item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#catalog_id DataVcdCatalogMedia#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#filter DataVcdCatalogMedia#filter}
	Filter *DataVcdCatalogMediaFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#id DataVcdCatalogMedia#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// media name (Optional when 'filter' is used).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#name DataVcdCatalogMedia#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#org DataVcdCatalogMedia#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
}

