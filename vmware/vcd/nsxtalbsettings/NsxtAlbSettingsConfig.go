package nsxtalbsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbSettingsConfig struct {
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
	// Edge gateway ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#edge_gateway_id NsxtAlbSettings#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Defines if ALB is enabled on Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#is_active NsxtAlbSettings#is_active}
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#id NsxtAlbSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IPv6 network definition in Gateway CIDR format which will be used by Load Balancer service on Edge (VCD 10.4.0+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#ipv6_service_network_specification NsxtAlbSettings#ipv6_service_network_specification}
	Ipv6ServiceNetworkSpecification *string `field:"optional" json:"ipv6ServiceNetworkSpecification" yaml:"ipv6ServiceNetworkSpecification"`
	// Enabling transparent mode allows to configure Preserve Client IP on a Virtual Service (VCD 10.4.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#is_transparent_mode_enabled NsxtAlbSettings#is_transparent_mode_enabled}
	IsTransparentModeEnabled interface{} `field:"optional" json:"isTransparentModeEnabled" yaml:"isTransparentModeEnabled"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#org NsxtAlbSettings#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#service_network_specification NsxtAlbSettings#service_network_specification}
	ServiceNetworkSpecification *string `field:"optional" json:"serviceNetworkSpecification" yaml:"serviceNetworkSpecification"`
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#supported_feature_set NsxtAlbSettings#supported_feature_set}
	SupportedFeatureSet *string `field:"optional" json:"supportedFeatureSet" yaml:"supportedFeatureSet"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_settings#vdc NsxtAlbSettings#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

