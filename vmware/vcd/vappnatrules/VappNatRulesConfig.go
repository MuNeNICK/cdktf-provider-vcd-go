package vappnatrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappNatRulesConfig struct {
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
	// One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#nat_type VappNatRules#nat_type}
	NatType *string `field:"required" json:"natType" yaml:"natType"`
	// vApp network identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#network_id VappNatRules#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// vApp identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#vapp_id VappNatRules#vapp_id}
	VappId *string `field:"required" json:"vappId" yaml:"vappId"`
	// Enable or disable NAT service. Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#enabled VappNatRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#enable_ip_masquerade VappNatRules#enable_ip_masquerade}
	EnableIpMasquerade interface{} `field:"optional" json:"enableIpMasquerade" yaml:"enableIpMasquerade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#id VappNatRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#org VappNatRules#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#rule VappNatRules#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_nat_rules#vdc VappNatRules#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

