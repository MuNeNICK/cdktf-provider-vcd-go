package nsxtedgegatewaybgpneighbor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayBgpNeighborConfig struct {
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
	// Edge gateway ID for BGP Neighbor Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#edge_gateway_id NsxtEdgegatewayBgpNeighbor#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// BGP Neighbor IP address (IPv4 or IPv6).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#ip_address NsxtEdgegatewayBgpNeighbor#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Remote Autonomous System (AS) number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#remote_as_number NsxtEdgegatewayBgpNeighbor#remote_as_number}
	RemoteAsNumber *string `field:"required" json:"remoteAsNumber" yaml:"remoteAsNumber"`
	// A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#allow_as_in NsxtEdgegatewayBgpNeighbor#allow_as_in}
	AllowAsIn interface{} `field:"optional" json:"allowAsIn" yaml:"allowAsIn"`
	// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#bfd_dead_multiple NsxtEdgegatewayBgpNeighbor#bfd_dead_multiple}
	BfdDeadMultiple *float64 `field:"optional" json:"bfdDeadMultiple" yaml:"bfdDeadMultiple"`
	// BFD configuration for failure detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#bfd_enabled NsxtEdgegatewayBgpNeighbor#bfd_enabled}
	BfdEnabled interface{} `field:"optional" json:"bfdEnabled" yaml:"bfdEnabled"`
	// Time interval (in milliseconds) between heartbeat packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#bfd_interval NsxtEdgegatewayBgpNeighbor#bfd_interval}
	BfdInterval *float64 `field:"optional" json:"bfdInterval" yaml:"bfdInterval"`
	// One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#graceful_restart_mode NsxtEdgegatewayBgpNeighbor#graceful_restart_mode}
	GracefulRestartMode *string `field:"optional" json:"gracefulRestartMode" yaml:"gracefulRestartMode"`
	// Time interval (in seconds) before declaring a peer dead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#hold_down_timer NsxtEdgegatewayBgpNeighbor#hold_down_timer}
	HoldDownTimer *float64 `field:"optional" json:"holdDownTimer" yaml:"holdDownTimer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#id NsxtEdgegatewayBgpNeighbor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An optional IP Prefix List ID for filtering 'IN' direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#in_filter_ip_prefix_list_id NsxtEdgegatewayBgpNeighbor#in_filter_ip_prefix_list_id}
	InFilterIpPrefixListId *string `field:"optional" json:"inFilterIpPrefixListId" yaml:"inFilterIpPrefixListId"`
	// Time interval (in seconds) between sending keep alive messages to a peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#keep_alive_timer NsxtEdgegatewayBgpNeighbor#keep_alive_timer}
	KeepAliveTimer *float64 `field:"optional" json:"keepAliveTimer" yaml:"keepAliveTimer"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#org NsxtEdgegatewayBgpNeighbor#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// An optional IP Prefix List ID for filtering 'OUT' direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#out_filter_ip_prefix_list_id NsxtEdgegatewayBgpNeighbor#out_filter_ip_prefix_list_id}
	OutFilterIpPrefixListId *string `field:"optional" json:"outFilterIpPrefixListId" yaml:"outFilterIpPrefixListId"`
	// Neighbor password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#password NsxtEdgegatewayBgpNeighbor#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// One of 'DISABLED', 'IPV4', 'IPV6'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor#route_filtering NsxtEdgegatewayBgpNeighbor#route_filtering}
	RouteFiltering *string `field:"optional" json:"routeFiltering" yaml:"routeFiltering"`
}

