package nsxtedgegatewaybgpconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayBgpConfigurationConfig struct {
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
	// Edge gateway ID for BGP Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#edge_gateway_id NsxtEdgegatewayBgpConfiguration#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Defines if BGP service is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#enabled NsxtEdgegatewayBgpConfiguration#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Defines if ECMP (Equal-cost multi-path routing) is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#ecmp_enabled NsxtEdgegatewayBgpConfiguration#ecmp_enabled}
	EcmpEnabled interface{} `field:"optional" json:"ecmpEnabled" yaml:"ecmpEnabled"`
	// Graceful restart configuration on Edge Gateway. One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#graceful_restart_mode NsxtEdgegatewayBgpConfiguration#graceful_restart_mode}
	GracefulRestartMode *string `field:"optional" json:"gracefulRestartMode" yaml:"gracefulRestartMode"`
	// Maximum time taken (in seconds) for a BGP session to be established after a restart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#graceful_restart_timer NsxtEdgegatewayBgpConfiguration#graceful_restart_timer}
	GracefulRestartTimer *float64 `field:"optional" json:"gracefulRestartTimer" yaml:"gracefulRestartTimer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#id NsxtEdgegatewayBgpConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Autonomous system number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#local_as_number NsxtEdgegatewayBgpConfiguration#local_as_number}
	LocalAsNumber *string `field:"optional" json:"localAsNumber" yaml:"localAsNumber"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#org NsxtEdgegatewayBgpConfiguration#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Maximum time (in seconds) before stale routes are removed when BGP restarts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration#stale_route_timer NsxtEdgegatewayBgpConfiguration#stale_route_timer}
	StaleRouteTimer *float64 `field:"optional" json:"staleRouteTimer" yaml:"staleRouteTimer"`
}

