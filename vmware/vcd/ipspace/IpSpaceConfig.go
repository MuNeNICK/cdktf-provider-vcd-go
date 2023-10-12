package ipspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IpSpaceConfig struct {
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
	// A set of internal scope IPs in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#internal_scope IpSpace#internal_scope}
	InternalScope *[]*string `field:"required" json:"internalScope" yaml:"internalScope"`
	// Name of IP space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#name IpSpace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of IP space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#type IpSpace#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Description of IP space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#description IpSpace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// External scope in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#external_scope IpSpace#external_scope}
	ExternalScope *string `field:"optional" json:"externalScope" yaml:"externalScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#id IpSpace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#ip_prefix IpSpace#ip_prefix}
	IpPrefix interface{} `field:"optional" json:"ipPrefix" yaml:"ipPrefix"`
	// ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#ip_range IpSpace#ip_range}
	IpRange interface{} `field:"optional" json:"ipRange" yaml:"ipRange"`
	// IP ranges quota. '-1' - unlimited, '0' - no quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#ip_range_quota IpSpace#ip_range_quota}
	IpRangeQuota *string `field:"optional" json:"ipRangeQuota" yaml:"ipRangeQuota"`
	// Org ID for 'SHARED' IP spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#org_id IpSpace#org_id}
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// Flag whether route advertisement should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ip_space#route_advertisement_enabled IpSpace#route_advertisement_enabled}
	RouteAdvertisementEnabled interface{} `field:"optional" json:"routeAdvertisementEnabled" yaml:"routeAdvertisementEnabled"`
}

