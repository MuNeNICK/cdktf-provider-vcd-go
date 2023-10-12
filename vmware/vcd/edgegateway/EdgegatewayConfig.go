package edgegateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgegatewayConfig struct {
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
	// Configuration of the vShield edge VM for this gateway. One of: compact, full ("Large"), full4 ("Quad Large"), x-large.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#configuration Edgegateway#configuration}
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// external_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#external_network Edgegateway#external_network}
	ExternalNetwork interface{} `field:"required" json:"externalNetwork" yaml:"externalNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#name Edgegateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#description Edgegateway#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable distributed routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#distributed_routing Edgegateway#distributed_routing}
	DistributedRouting interface{} `field:"optional" json:"distributedRouting" yaml:"distributedRouting"`
	// Enable FIPS mode. FIPS mode turns on the cipher suites that comply with FIPS. (False by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#fips_mode_enabled Edgegateway#fips_mode_enabled}
	FipsModeEnabled interface{} `field:"optional" json:"fipsModeEnabled" yaml:"fipsModeEnabled"`
	// 'accept' or 'deny'. Default 'deny'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#fw_default_rule_action Edgegateway#fw_default_rule_action}
	FwDefaultRuleAction *string `field:"optional" json:"fwDefaultRuleAction" yaml:"fwDefaultRuleAction"`
	// Enable logging for default rule. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#fw_default_rule_logging_enabled Edgegateway#fw_default_rule_logging_enabled}
	FwDefaultRuleLoggingEnabled interface{} `field:"optional" json:"fwDefaultRuleLoggingEnabled" yaml:"fwDefaultRuleLoggingEnabled"`
	// Enable firewall. Default 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#fw_enabled Edgegateway#fw_enabled}
	FwEnabled interface{} `field:"optional" json:"fwEnabled" yaml:"fwEnabled"`
	// Enable high availability on this edge gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#ha_enabled Edgegateway#ha_enabled}
	HaEnabled interface{} `field:"optional" json:"haEnabled" yaml:"haEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#id Edgegateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable load balancer acceleration. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#lb_acceleration_enabled Edgegateway#lb_acceleration_enabled}
	LbAccelerationEnabled interface{} `field:"optional" json:"lbAccelerationEnabled" yaml:"lbAccelerationEnabled"`
	// Enable load balancing. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#lb_enabled Edgegateway#lb_enabled}
	LbEnabled interface{} `field:"optional" json:"lbEnabled" yaml:"lbEnabled"`
	// Enable load balancer logging. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#lb_logging_enabled Edgegateway#lb_logging_enabled}
	LbLoggingEnabled interface{} `field:"optional" json:"lbLoggingEnabled" yaml:"lbLoggingEnabled"`
	// Log level. One of 'emergency', 'alert', 'critical', 'error', 'warning', 'notice', 'info', 'debug'. ('info' by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#lb_loglevel Edgegateway#lb_loglevel}
	LbLoglevel *string `field:"optional" json:"lbLoglevel" yaml:"lbLoglevel"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#org Edgegateway#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// If true, default gateway will be used for the edge gateways' default routing and DNS forwarding.(False by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#use_default_route_for_dns_relay Edgegateway#use_default_route_for_dns_relay}
	UseDefaultRouteForDnsRelay interface{} `field:"optional" json:"useDefaultRouteForDnsRelay" yaml:"useDefaultRouteForDnsRelay"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#vdc Edgegateway#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

