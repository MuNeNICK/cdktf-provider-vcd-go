package nsxtnatrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtNatRuleConfig struct {
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
	// Edge gateway name in which NAT Rule is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#edge_gateway_id NsxtNatRule#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Name of NAT rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#name NsxtNatRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#rule_type NsxtNatRule#rule_type}
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Application Port Profile to apply for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#app_port_profile_id NsxtNatRule#app_port_profile_id}
	AppPortProfileId *string `field:"optional" json:"appPortProfileId" yaml:"appPortProfileId"`
	// Description of NAT rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#description NsxtNatRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// For DNAT only.
	//
	// Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#dnat_external_port NsxtNatRule#dnat_external_port}
	DnatExternalPort *string `field:"optional" json:"dnatExternalPort" yaml:"dnatExternalPort"`
	// Enables or disables this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#enabled NsxtNatRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// IP address or CIDR of external network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#external_address NsxtNatRule#external_address}
	ExternalAddress *string `field:"optional" json:"externalAddress" yaml:"externalAddress"`
	// VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of 'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#firewall_match NsxtNatRule#firewall_match}
	FirewallMatch *string `field:"optional" json:"firewallMatch" yaml:"firewallMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#id NsxtNatRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP address or CIDR of the virtual machines for which you are configuring NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#internal_address NsxtNatRule#internal_address}
	InternalAddress *string `field:"optional" json:"internalAddress" yaml:"internalAddress"`
	// Enable logging when this rule is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#logging NsxtNatRule#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#org NsxtNatRule#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a higher precedence for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#priority NsxtNatRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// For SNAT only.
	//
	// If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain or an IP address range in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#snat_destination_address NsxtNatRule#snat_destination_address}
	SnatDestinationAddress *string `field:"optional" json:"snatDestinationAddress" yaml:"snatDestinationAddress"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_nat_rule#vdc NsxtNatRule#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

