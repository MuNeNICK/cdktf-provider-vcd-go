package datavcdipspaceipallocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdIpSpaceIpAllocationConfig struct {
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
	// IP Address or Prefix of the allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#ip_address DataVcdIpSpaceIpAllocation#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// IP Space ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#ip_space_id DataVcdIpSpaceIpAllocation#ip_space_id}
	IpSpaceId *string `field:"required" json:"ipSpaceId" yaml:"ipSpaceId"`
	// Org ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#org_id DataVcdIpSpaceIpAllocation#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Type of IP Allocation. One of 'FLOATING_IP' or 'IP_PREFIX'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#type DataVcdIpSpaceIpAllocation#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// IP Allocation Description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#description DataVcdIpSpaceIpAllocation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#id DataVcdIpSpaceIpAllocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// One of 'UNUSED', 'USED', 'USED_MANUAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/ip_space_ip_allocation#usage_state DataVcdIpSpaceIpAllocation#usage_state}
	UsageState *string `field:"optional" json:"usageState" yaml:"usageState"`
}

