package datavcdcatalogvapptemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdCatalogVappTemplateConfig struct {
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
	// ID of the catalog containing the vApp Template. Can't be used if a specific VDC identifier is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#catalog_id DataVcdCatalogVappTemplate#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#filter DataVcdCatalogVappTemplate#filter}
	Filter *DataVcdCatalogVappTemplateFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#id DataVcdCatalogVappTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the vApp Template. It is optional when a filter is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#name DataVcdCatalogVappTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#org DataVcdCatalogVappTemplate#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// ID of the VDC to which the vApp Template belongs.
	//
	// Can't be used if a specific Catalog identifier is set
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#vdc_id DataVcdCatalogVappTemplate#vdc_id}
	VdcId *string `field:"optional" json:"vdcId" yaml:"vdcId"`
}

