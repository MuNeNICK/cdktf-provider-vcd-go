package vmplacementpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmPlacementPolicyConfig struct {
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
	// Name of the VM Placement Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#name VmPlacementPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the Provider VDC to which the VM Placement Policy belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#provider_vdc_id VmPlacementPolicy#provider_vdc_id}
	ProviderVdcId *string `field:"required" json:"providerVdcId" yaml:"providerVdcId"`
	// Description of the VM Placement Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#description VmPlacementPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#id VmPlacementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IDs of one or more Logical VM Groups to define this VM Placement Policy.
	//
	// There is an AND relationship among all the entries set in this attribute
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#logical_vm_group_ids VmPlacementPolicy#logical_vm_group_ids}
	LogicalVmGroupIds *[]*string `field:"optional" json:"logicalVmGroupIds" yaml:"logicalVmGroupIds"`
	// IDs of the collection of VMs with similar host requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_placement_policy#vm_group_ids VmPlacementPolicy#vm_group_ids}
	VmGroupIds *[]*string `field:"optional" json:"vmGroupIds" yaml:"vmGroupIds"`
}

