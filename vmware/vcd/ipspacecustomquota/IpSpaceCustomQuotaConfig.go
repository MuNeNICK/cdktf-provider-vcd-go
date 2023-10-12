package ipspacecustomquota

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IpSpaceCustomQuotaConfig struct {
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
	// ID of IP Space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#ip_space_id IpSpaceCustomQuota#ip_space_id}
	IpSpaceId *string `field:"required" json:"ipSpaceId" yaml:"ipSpaceId"`
	// Organization ID for setting custom quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#org_id IpSpaceCustomQuota#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#id IpSpaceCustomQuota#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_prefix_quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#ip_prefix_quota IpSpaceCustomQuota#ip_prefix_quota}
	IpPrefixQuota interface{} `field:"optional" json:"ipPrefixQuota" yaml:"ipPrefixQuota"`
	// IP range quota. '-1' - unlimited, '0' - no quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space_custom_quota#ip_range_quota IpSpaceCustomQuota#ip_range_quota}
	IpRangeQuota *string `field:"optional" json:"ipRangeQuota" yaml:"ipRangeQuota"`
}

