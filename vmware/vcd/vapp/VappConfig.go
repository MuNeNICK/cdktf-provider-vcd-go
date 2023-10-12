package vapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappConfig struct {
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
	// A name for the vApp, unique withing the VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#name Vapp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description of the vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#description Vapp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key/value settings for guest properties. Will be picked up by new VMs when created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#guest_properties Vapp#guest_properties}
	GuestProperties *map[string]*string `field:"optional" json:"guestProperties" yaml:"guestProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#id Vapp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lease block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#lease Vapp#lease}
	Lease *VappLease `field:"optional" json:"lease" yaml:"lease"`
	// Key value map of metadata to assign to this vApp. Key and value can be any string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#metadata Vapp#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#metadata_entry Vapp#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#org Vapp#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// A boolean value stating if this vApp should be powered on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#power_on Vapp#power_on}
	PowerOn interface{} `field:"optional" json:"powerOn" yaml:"powerOn"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#vdc Vapp#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

