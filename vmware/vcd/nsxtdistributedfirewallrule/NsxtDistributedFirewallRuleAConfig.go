package nsxtdistributedfirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtDistributedFirewallRuleAConfig struct {
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
	// Defines if the rule should 'ALLOW', 'DROP', 'REJECT' matching traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#action NsxtDistributedFirewallRuleA#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Firewall Rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#name NsxtDistributedFirewallRuleA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of VDC Group for Distributed Firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#vdc_group_id NsxtDistributedFirewallRuleA#vdc_group_id}
	VdcGroupId *string `field:"required" json:"vdcGroupId" yaml:"vdcGroupId"`
	// An optional firewall rule ID, to put new rule above during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#above_rule_id NsxtDistributedFirewallRuleA#above_rule_id}
	AboveRuleId *string `field:"optional" json:"aboveRuleId" yaml:"aboveRuleId"`
	// A set of Application Port Profile IDs. Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#app_port_profile_ids NsxtDistributedFirewallRuleA#app_port_profile_ids}
	AppPortProfileIds *[]*string `field:"optional" json:"appPortProfileIds" yaml:"appPortProfileIds"`
	// Comment that is shown next to rule in UI (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#comment NsxtDistributedFirewallRuleA#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Description is not shown in UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#description NsxtDistributedFirewallRuleA#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Reverses firewall matching for to match all except Destinations Groups specified in 'destination_ids' (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#destination_groups_excluded NsxtDistributedFirewallRuleA#destination_groups_excluded}
	DestinationGroupsExcluded interface{} `field:"optional" json:"destinationGroupsExcluded" yaml:"destinationGroupsExcluded"`
	// A set of Destination Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#destination_ids NsxtDistributedFirewallRuleA#destination_ids}
	DestinationIds *[]*string `field:"optional" json:"destinationIds" yaml:"destinationIds"`
	// Direction on which Firewall Rule applies (one of 'IN', 'OUT', 'IN_OUT').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#direction NsxtDistributedFirewallRuleA#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Defined if Firewall Rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#enabled NsxtDistributedFirewallRuleA#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#id NsxtDistributedFirewallRuleA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Firewall Rule Protocol (one of 'IPV4', 'IPV6', 'IPV4_IPV6').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#ip_protocol NsxtDistributedFirewallRuleA#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// Defines if matching traffic should be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#logging NsxtDistributedFirewallRuleA#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// A set of Network Context Profile IDs. Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#network_context_profile_ids NsxtDistributedFirewallRuleA#network_context_profile_ids}
	NetworkContextProfileIds *[]*string `field:"optional" json:"networkContextProfileIds" yaml:"networkContextProfileIds"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#org NsxtDistributedFirewallRuleA#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Reverses firewall matching for to match all except Source Groups specified in 'source_ids' (VCD 10.3.2+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#source_groups_excluded NsxtDistributedFirewallRuleA#source_groups_excluded}
	SourceGroupsExcluded interface{} `field:"optional" json:"sourceGroupsExcluded" yaml:"sourceGroupsExcluded"`
	// A set of Source Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule#source_ids NsxtDistributedFirewallRuleA#source_ids}
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
}

