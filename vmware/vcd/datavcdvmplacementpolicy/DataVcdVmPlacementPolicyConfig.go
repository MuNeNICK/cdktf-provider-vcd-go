package datavcdvmplacementpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdVmPlacementPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm_placement_policy#name DataVcdVmPlacementPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm_placement_policy#id DataVcdVmPlacementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of the Provider VDC to which the VM Placement Policy belongs.
	//
	// To be used by System administrators instead of `vdc_id`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm_placement_policy#provider_vdc_id DataVcdVmPlacementPolicy#provider_vdc_id}
	ProviderVdcId *string `field:"optional" json:"providerVdcId" yaml:"providerVdcId"`
	// ID of the VDC to which the VM Placement Policy is assigned.
	//
	// To be used by tenants instead of `provider_vdc_id`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm_placement_policy#vdc_id DataVcdVmPlacementPolicy#vdc_id}
	VdcId *string `field:"optional" json:"vdcId" yaml:"vdcId"`
}

