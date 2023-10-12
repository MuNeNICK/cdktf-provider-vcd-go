package vappfirewallrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappFirewallRulesConfig struct {
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
	// Specifies what to do should none of the rules match. Either `allow` or `drop`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#default_action VappFirewallRules#default_action}
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// vApp network identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#network_id VappFirewallRules#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// vApp identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#vapp_id VappFirewallRules#vapp_id}
	VappId *string `field:"required" json:"vappId" yaml:"vappId"`
	// Enable or disable firewall service. Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#enabled VappFirewallRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#id VappFirewallRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Flag to enable logging for default action. Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#log_default_action VappFirewallRules#log_default_action}
	LogDefaultAction interface{} `field:"optional" json:"logDefaultAction" yaml:"logDefaultAction"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#org VappFirewallRules#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#rule VappFirewallRules#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_firewall_rules#vdc VappFirewallRules#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

