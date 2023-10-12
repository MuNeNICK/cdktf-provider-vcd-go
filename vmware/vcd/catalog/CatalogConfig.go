package catalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogConfig struct {
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
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains, regardless of their state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#delete_force Catalog#delete_force}
	DeleteForce interface{} `field:"required" json:"deleteForce" yaml:"deleteForce"`
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#delete_recursive Catalog#delete_recursive}
	DeleteRecursive interface{} `field:"required" json:"deleteRecursive" yaml:"deleteRecursive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#name Catalog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// True enables early catalog export to optimize synchronization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#cache_enabled Catalog#cache_enabled}
	CacheEnabled interface{} `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#description Catalog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#id Catalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key and value pairs for catalog metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#metadata Catalog#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#metadata_entry Catalog#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#org Catalog#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#password Catalog#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Include BIOS UUIDs and MAC addresses in the downloaded OVF package.
	//
	// Preserving the identity information limits the portability of the package and you should use it only when necessary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#preserve_identity_information Catalog#preserve_identity_information}
	PreserveIdentityInformation interface{} `field:"optional" json:"preserveIdentityInformation" yaml:"preserveIdentityInformation"`
	// True allows to publish a catalog externally to make its vApp templates and media files available for subscription by organizations outside the Cloud Director installation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#publish_enabled Catalog#publish_enabled}
	PublishEnabled interface{} `field:"optional" json:"publishEnabled" yaml:"publishEnabled"`
	// Optional storage profile ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog#storage_profile_id Catalog#storage_profile_id}
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
}

