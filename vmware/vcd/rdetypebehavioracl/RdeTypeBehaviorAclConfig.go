package rdetypebehavioracl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdeTypeBehaviorAclConfig struct {
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
	// Set of Access Level IDs to associate to the Behavior defined in `behavior_id` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior_acl#access_level_ids RdeTypeBehaviorAcl#access_level_ids}
	AccessLevelIds *[]*string `field:"required" json:"accessLevelIds" yaml:"accessLevelIds"`
	// The ID of either a RDE Interface Behavior or RDE Type Behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior_acl#behavior_id RdeTypeBehaviorAcl#behavior_id}
	BehaviorId *string `field:"required" json:"behaviorId" yaml:"behaviorId"`
	// The ID of the RDE Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior_acl#rde_type_id RdeTypeBehaviorAcl#rde_type_id}
	RdeTypeId *string `field:"required" json:"rdeTypeId" yaml:"rdeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rde_type_behavior_acl#id RdeTypeBehaviorAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

