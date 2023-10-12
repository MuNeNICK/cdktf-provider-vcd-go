package ipspaceipallocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IpSpaceIpAllocationConfig struct {
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
	// Org ID for IP Allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#org_id IpSpaceIpAllocation#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Type of allocation. One of `FLOATING_IP``, `IP_PREFIX`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#type IpSpaceIpAllocation#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Custom description can only be set when usage_state is set to 'USED_MANUAL'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#description IpSpaceIpAllocation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#id IpSpaceIpAllocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP Space ID for IP Allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#ip_space_id IpSpaceIpAllocation#ip_space_id}
	IpSpaceId *string `field:"optional" json:"ipSpaceId" yaml:"ipSpaceId"`
	// Required if 'type' is IP_PREFIX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#prefix_length IpSpaceIpAllocation#prefix_length}
	PrefixLength *string `field:"optional" json:"prefixLength" yaml:"prefixLength"`
	// Can be set to 'USED_MANUAL' to mark the IP Allocation for manual use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_ip_allocation#usage_state IpSpaceIpAllocation#usage_state}
	UsageState *string `field:"optional" json:"usageState" yaml:"usageState"`
}

