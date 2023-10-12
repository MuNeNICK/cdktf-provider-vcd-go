package rde

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdeConfig struct {
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
	// The name of the Runtime Defined Entity. It can be non-unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#name Rde#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Runtime Defined Entity Type ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#rde_type_id Rde#rde_type_id}
	RdeTypeId *string `field:"required" json:"rdeTypeId" yaml:"rdeTypeId"`
	// If `true`, the Runtime Defined Entity will be resolved by this provider.
	//
	// If `false`, it won't beresolved and must be done either by an external component action or by an update. The Runtime Defined Entity can't bedeleted until the entity is resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#resolve Rde#resolve}
	Resolve interface{} `field:"required" json:"resolve" yaml:"resolve"`
	// An external entity's ID that this Runtime Defined Entity may have a relation to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#external_id Rde#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#id Rde#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A JSON representation of the Runtime Defined Entity that is defined by the user and is used to initialize/override its contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#input_entity Rde#input_entity}
	InputEntity *string `field:"optional" json:"inputEntity" yaml:"inputEntity"`
	// URL that should point to a JSON representation of the Runtime Defined Entity and is used to initialize/override its contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#input_entity_url Rde#input_entity_url}
	InputEntityUrl *string `field:"optional" json:"inputEntityUrl" yaml:"inputEntityUrl"`
	// The name of organization that will own this Runtime Defined Entity, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#org Rde#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// If `true`, the Runtime Defined Entity will be resolved before it gets deleted, to ensure forced deletion.Destroy will fail if it is not resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde#resolve_on_removal Rde#resolve_on_removal}
	ResolveOnRemoval interface{} `field:"optional" json:"resolveOnRemoval" yaml:"resolveOnRemoval"`
}

