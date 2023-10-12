package vmaffinityrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmAffinityRuleConfig struct {
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
	// VM affinity rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#name VmAffinityRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One of 'Affinity', 'Anti-Affinity'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#polarity VmAffinityRule#polarity}
	Polarity *string `field:"required" json:"polarity" yaml:"polarity"`
	// Set of VM IDs assigned to this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#vm_ids VmAffinityRule#vm_ids}
	VmIds *[]*string `field:"required" json:"vmIds" yaml:"vmIds"`
	// True if this affinity rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#enabled VmAffinityRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#id VmAffinityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#org VmAffinityRule#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// True if this affinity rule is required.
	//
	// When a rule is mandatory, a host failover will not power on the VM if doing so would violate the rule
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#required VmAffinityRule#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_affinity_rule#vdc VmAffinityRule#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

