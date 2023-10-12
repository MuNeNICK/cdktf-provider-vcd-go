package datavcdrdetypebehavior

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdRdeTypeBehaviorConfig struct {
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
	// The ID of either a RDE Interface Behavior or RDE Type Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type_behavior#behavior_id DataVcdRdeTypeBehavior#behavior_id}
	BehaviorId *string `field:"required" json:"behaviorId" yaml:"behaviorId"`
	// The ID of the RDE Type that owns the Behavior to fetch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type_behavior#rde_type_id DataVcdRdeTypeBehavior#rde_type_id}
	RdeTypeId *string `field:"required" json:"rdeTypeId" yaml:"rdeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type_behavior#id DataVcdRdeTypeBehavior#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

