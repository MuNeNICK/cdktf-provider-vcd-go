package catalogitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogItemConfig struct {
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
	// catalog name where upload the OVA file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#catalog CatalogItem#catalog}
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// catalog item name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#name CatalogItem#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Key and value pairs for catalog item metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#catalog_item_metadata CatalogItem#catalog_item_metadata}
	CatalogItemMetadata *map[string]*string `field:"optional" json:"catalogItemMetadata" yaml:"catalogItemMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#description CatalogItem#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#id CatalogItem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key and value pairs for the metadata of the vApp template associated to this catalog item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#metadata CatalogItem#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#metadata_entry CatalogItem#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#org CatalogItem#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Absolute or relative path to OVA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#ova_path CatalogItem#ova_path}
	OvaPath *string `field:"optional" json:"ovaPath" yaml:"ovaPath"`
	// URL of OVF file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#ovf_url CatalogItem#ovf_url}
	OvfUrl *string `field:"optional" json:"ovfUrl" yaml:"ovfUrl"`
	// shows upload progress in stdout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#show_upload_progress CatalogItem#show_upload_progress}
	ShowUploadProgress interface{} `field:"optional" json:"showUploadProgress" yaml:"showUploadProgress"`
	// size of upload file piece size in mega bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_item#upload_piece_size CatalogItem#upload_piece_size}
	UploadPieceSize *float64 `field:"optional" json:"uploadPieceSize" yaml:"uploadPieceSize"`
}

