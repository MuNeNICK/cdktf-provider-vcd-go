package independentdisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IndependentDiskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#name IndependentDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// size in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#size_in_mb IndependentDisk#size_in_mb}
	SizeInMb *float64 `field:"required" json:"sizeInMb" yaml:"sizeInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#bus_sub_type IndependentDisk#bus_sub_type}.
	BusSubType *string `field:"optional" json:"busSubType" yaml:"busSubType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#bus_type IndependentDisk#bus_type}.
	BusType *string `field:"optional" json:"busType" yaml:"busType"`
	// independent disk description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#description IndependentDisk#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#id IndependentDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key value map of metadata to assign to this disk. Key and value can be any string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#metadata IndependentDisk#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#metadata_entry IndependentDisk#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#org IndependentDisk#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// This is the sharing type. This attribute can only have values defined one of: `DiskSharing`,`ControllerSharing`, `None`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#sharing_type IndependentDisk#sharing_type}
	SharingType *string `field:"optional" json:"sharingType" yaml:"sharingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#storage_profile IndependentDisk#storage_profile}.
	StorageProfile *string `field:"optional" json:"storageProfile" yaml:"storageProfile"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/independent_disk#vdc IndependentDisk#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

