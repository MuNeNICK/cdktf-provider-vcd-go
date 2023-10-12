package rdetypebehavior

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdeTypeBehaviorConfig struct {
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
	// The ID of the original RDE Interface Behavior to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior#rde_interface_behavior_id RdeTypeBehavior#rde_interface_behavior_id}
	RdeInterfaceBehaviorId *string `field:"required" json:"rdeInterfaceBehaviorId" yaml:"rdeInterfaceBehaviorId"`
	// The ID of the RDE Type that owns the Behavior override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior#rde_type_id RdeTypeBehavior#rde_type_id}
	RdeTypeId *string `field:"required" json:"rdeTypeId" yaml:"rdeTypeId"`
	// The description of the contract of the overridden Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior#description RdeTypeBehavior#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Execution map of the Behavior that overrides the original.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior#execution RdeTypeBehavior#execution}
	Execution *map[string]*string `field:"optional" json:"execution" yaml:"execution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior#id RdeTypeBehavior#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

