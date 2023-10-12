package edgegatewaysettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgegatewaySettingsConfig struct {
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
	// ID of the edge gateway. Required when 'edge_gateway_name' is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#edge_gateway_id EdgegatewaySettings#edge_gateway_id}
	EdgeGatewayId *string `field:"optional" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Name of the edge gateway. Required when 'edge_gateway_id' is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#edge_gateway_name EdgegatewaySettings#edge_gateway_name}
	EdgeGatewayName *string `field:"optional" json:"edgeGatewayName" yaml:"edgeGatewayName"`
	// 'accept' or 'deny'. Default 'deny'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#fw_default_rule_action EdgegatewaySettings#fw_default_rule_action}
	FwDefaultRuleAction *string `field:"optional" json:"fwDefaultRuleAction" yaml:"fwDefaultRuleAction"`
	// Enable logging for default rule. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#fw_default_rule_logging_enabled EdgegatewaySettings#fw_default_rule_logging_enabled}
	FwDefaultRuleLoggingEnabled interface{} `field:"optional" json:"fwDefaultRuleLoggingEnabled" yaml:"fwDefaultRuleLoggingEnabled"`
	// Enable firewall. Default 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#fw_enabled EdgegatewaySettings#fw_enabled}
	FwEnabled interface{} `field:"optional" json:"fwEnabled" yaml:"fwEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#id EdgegatewaySettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable load balancer acceleration. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#lb_acceleration_enabled EdgegatewaySettings#lb_acceleration_enabled}
	LbAccelerationEnabled interface{} `field:"optional" json:"lbAccelerationEnabled" yaml:"lbAccelerationEnabled"`
	// Enable load balancing. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#lb_enabled EdgegatewaySettings#lb_enabled}
	LbEnabled interface{} `field:"optional" json:"lbEnabled" yaml:"lbEnabled"`
	// Enable load balancer logging. (Disabled by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#lb_logging_enabled EdgegatewaySettings#lb_logging_enabled}
	LbLoggingEnabled interface{} `field:"optional" json:"lbLoggingEnabled" yaml:"lbLoggingEnabled"`
	// Log level. One of 'emergency', 'alert', 'critical', 'error', 'warning', 'notice', 'info', 'debug'. ('info' by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#lb_loglevel EdgegatewaySettings#lb_loglevel}
	LbLoglevel *string `field:"optional" json:"lbLoglevel" yaml:"lbLoglevel"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#org EdgegatewaySettings#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings#vdc EdgegatewaySettings#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

