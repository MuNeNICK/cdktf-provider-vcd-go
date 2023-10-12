package rdeinterfacebehavior

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdeInterfaceBehaviorConfig struct {
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
	// Execution map of the Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_interface_behavior#execution RdeInterfaceBehavior#execution}
	Execution *map[string]*string `field:"required" json:"execution" yaml:"execution"`
	// Name of the Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_interface_behavior#name RdeInterfaceBehavior#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the RDE Interface that owns the Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_interface_behavior#rde_interface_id RdeInterfaceBehavior#rde_interface_id}
	RdeInterfaceId *string `field:"required" json:"rdeInterfaceId" yaml:"rdeInterfaceId"`
	// A description specifying the contract of the Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_interface_behavior#description RdeInterfaceBehavior#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_interface_behavior#id RdeInterfaceBehavior#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

