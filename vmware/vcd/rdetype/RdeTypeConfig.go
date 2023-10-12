package rdetype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdeTypeConfig struct {
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
	// The name of the Runtime Defined Entity Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#name RdeType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A unique namespace associated with the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#nss RdeType#nss}
	Nss *string `field:"required" json:"nss" yaml:"nss"`
	// The vendor name for the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#vendor RdeType#vendor}
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
	// The version of the Runtime Defined Entity Type.
	//
	// The version string must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#version RdeType#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The description of the Runtime Defined Entity Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#description RdeType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An external entity's ID that this definition may apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#external_id RdeType#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#id RdeType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// To be used when creating a new version of a Runtime Defined Entity Type.
	//
	// Specifies the version of the type that will be the template for the authorization configuration of the new version.The Type ACLs and the access requirements of the Type Behaviors of the new version will be copied from those of the inherited version.If not set, then the new type version will not inherit another version and will have the default authorization settings, just like the first version of a new type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#inherited_version RdeType#inherited_version}
	InheritedVersion *string `field:"optional" json:"inheritedVersion" yaml:"inheritedVersion"`
	// Set of Defined Interface URNs that this Runtime Defined Entity Type is referenced by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#interface_ids RdeType#interface_ids}
	InterfaceIds *[]*string `field:"optional" json:"interfaceIds" yaml:"interfaceIds"`
	// The JSON-Schema valid definition of the Runtime Defined Entity Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#schema RdeType#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// URL that should point to a JSON-Schema valid definition file of the Runtime Defined Entity Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type#schema_url RdeType#schema_url}
	SchemaUrl *string `field:"optional" json:"schemaUrl" yaml:"schemaUrl"`
}

