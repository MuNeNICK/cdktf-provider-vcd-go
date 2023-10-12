package nsxvfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvFirewallRuleConfig struct {
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
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#destination NsxvFirewallRule#destination}
	Destination *NsxvFirewallRuleDestination `field:"required" json:"destination" yaml:"destination"`
	// Edge gateway name in which Firewall Rule is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#edge_gateway NsxvFirewallRule#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#service NsxvFirewallRule#service}
	Service interface{} `field:"required" json:"service" yaml:"service"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#source NsxvFirewallRule#source}
	Source *NsxvFirewallRuleSource `field:"required" json:"source" yaml:"source"`
	// This firewall rule will be inserted above the referred one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#above_rule_id NsxvFirewallRule#above_rule_id}
	AboveRuleId *string `field:"optional" json:"aboveRuleId" yaml:"aboveRuleId"`
	// 'accept' or 'deny'. Default 'accept'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#action NsxvFirewallRule#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Whether the rule should be enabled. Default 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#enabled NsxvFirewallRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#id NsxvFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether logging should be enabled for this rule. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#logging_enabled NsxvFirewallRule#logging_enabled}
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// Firewall rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#name NsxvFirewallRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#org NsxvFirewallRule#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Optional. Allows to set custom rule tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#rule_tag NsxvFirewallRule#rule_tag}
	RuleTag *float64 `field:"optional" json:"ruleTag" yaml:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#rule_type NsxvFirewallRule#rule_type}
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#vdc NsxvFirewallRule#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

