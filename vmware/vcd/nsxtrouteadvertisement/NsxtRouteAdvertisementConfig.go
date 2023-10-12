package nsxtrouteadvertisement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtRouteAdvertisementConfig struct {
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
	// NSX-T Edge Gateway ID in which route advertisement is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_route_advertisement#edge_gateway_id NsxtRouteAdvertisement#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Defines if route advertisement is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_route_advertisement#enabled NsxtRouteAdvertisement#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_route_advertisement#id NsxtRouteAdvertisement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_route_advertisement#org NsxtRouteAdvertisement#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Set of subnets that will be advertised to Tier-0 gateway. Empty means none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_route_advertisement#subnets NsxtRouteAdvertisement#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

