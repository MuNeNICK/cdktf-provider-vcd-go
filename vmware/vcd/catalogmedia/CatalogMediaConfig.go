package catalogmedia

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogMediaConfig struct {
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
	// absolute or relative path to Media file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#media_path CatalogMedia#media_path}
	MediaPath *string `field:"required" json:"mediaPath" yaml:"mediaPath"`
	// media name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#name CatalogMedia#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// catalog name where to upload the Media file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#catalog CatalogMedia#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// ID of the catalog where to upload the Media file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#catalog_id CatalogMedia#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#description CatalogMedia#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#id CatalogMedia#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key and value pairs for catalog item metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#metadata CatalogMedia#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#metadata_entry CatalogMedia#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#org CatalogMedia#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// shows upload progress in stdout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#show_upload_progress CatalogMedia#show_upload_progress}
	ShowUploadProgress interface{} `field:"optional" json:"showUploadProgress" yaml:"showUploadProgress"`
	// size of upload file piece size in mega bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#upload_piece_size CatalogMedia#upload_piece_size}
	UploadPieceSize *float64 `field:"optional" json:"uploadPieceSize" yaml:"uploadPieceSize"`
}

