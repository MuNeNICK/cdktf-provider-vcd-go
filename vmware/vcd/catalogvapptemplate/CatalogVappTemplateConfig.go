package catalogvapptemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogVappTemplateConfig struct {
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
	// ID of the Catalog where to upload the OVA file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#catalog_id CatalogVappTemplate#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// vApp Template name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#name CatalogVappTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the vApp Template. Not to be used with `ovf_url` when target OVA has a description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#description CatalogVappTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#id CatalogVappTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key and value pairs for the metadata of this vApp Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#metadata CatalogVappTemplate#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#metadata_entry CatalogVappTemplate#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#org CatalogVappTemplate#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Absolute or relative path to OVA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#ova_path CatalogVappTemplate#ova_path}
	OvaPath *string `field:"optional" json:"ovaPath" yaml:"ovaPath"`
	// URL of OVF file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#ovf_url CatalogVappTemplate#ovf_url}
	OvfUrl *string `field:"optional" json:"ovfUrl" yaml:"ovfUrl"`
	// Size of upload file piece size in megabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_vapp_template#upload_piece_size CatalogVappTemplate#upload_piece_size}
	UploadPieceSize *float64 `field:"optional" json:"uploadPieceSize" yaml:"uploadPieceSize"`
}

