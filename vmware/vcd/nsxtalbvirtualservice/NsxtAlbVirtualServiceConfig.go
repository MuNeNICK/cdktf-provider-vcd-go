package nsxtalbvirtualservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbVirtualServiceConfig struct {
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
	// HTTP, HTTPS, L4, L4_TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#application_profile_type NsxtAlbVirtualService#application_profile_type}
	ApplicationProfileType *string `field:"required" json:"applicationProfileType" yaml:"applicationProfileType"`
	// Edge gateway ID in which ALB Pool should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#edge_gateway_id NsxtAlbVirtualService#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Name of ALB Virtual Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#name NsxtAlbVirtualService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Pool ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#pool_id NsxtAlbVirtualService#pool_id}
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
	// Service Engine Group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#service_engine_group_id NsxtAlbVirtualService#service_engine_group_id}
	ServiceEngineGroupId *string `field:"required" json:"serviceEngineGroupId" yaml:"serviceEngineGroupId"`
	// Virtual IP address (VIP) for Virtual Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#virtual_ip_address NsxtAlbVirtualService#virtual_ip_address}
	VirtualIpAddress *string `field:"required" json:"virtualIpAddress" yaml:"virtualIpAddress"`
	// Optional certificate ID to use for exposing service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#ca_certificate_id NsxtAlbVirtualService#ca_certificate_id}
	CaCertificateId *string `field:"optional" json:"caCertificateId" yaml:"caCertificateId"`
	// Description of ALB Virtual Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#description NsxtAlbVirtualService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Virtual Service is enabled or disabled (default true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#enabled NsxtAlbVirtualService#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#id NsxtAlbVirtualService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IPv6 Virtual IP address (VIP) for Virtual Service (VCD 10.4.0+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#ipv6_virtual_ip_address NsxtAlbVirtualService#ipv6_virtual_ip_address}
	Ipv6VirtualIpAddress *string `field:"optional" json:"ipv6VirtualIpAddress" yaml:"ipv6VirtualIpAddress"`
	// Preserves Client IP on a Virtual Service (VCD 10.4.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#is_transparent_mode_enabled NsxtAlbVirtualService#is_transparent_mode_enabled}
	IsTransparentModeEnabled interface{} `field:"optional" json:"isTransparentModeEnabled" yaml:"isTransparentModeEnabled"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#org NsxtAlbVirtualService#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// service_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#service_port NsxtAlbVirtualService#service_port}
	ServicePort interface{} `field:"optional" json:"servicePort" yaml:"servicePort"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service#vdc NsxtAlbVirtualService#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

