package nsxvdnat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvDnatConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#edge_gateway NsxvDnat#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Org or external network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#network_name NsxvDnat#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// Network type. One of 'ext', 'org'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#network_type NsxvDnat#network_type}
	NetworkType *string `field:"required" json:"networkType" yaml:"networkType"`
	// Original address or address range. This is the the destination address for DNAT rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#original_address NsxvDnat#original_address}
	OriginalAddress *string `field:"required" json:"originalAddress" yaml:"originalAddress"`
	// NAT rule description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#description NsxvDnat#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule should be enabled. Default 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#enabled NsxvDnat#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// ICMP type.
	//
	// Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`, `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`, `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#icmp_type NsxvDnat#icmp_type}
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#id NsxvDnat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether logging should be enabled for this rule. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#logging_enabled NsxvDnat#logging_enabled}
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#org NsxvDnat#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Original port. This is the destination port for DNAT rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#original_port NsxvDnat#original_port}
	OriginalPort *string `field:"optional" json:"originalPort" yaml:"originalPort"`
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#protocol NsxvDnat#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Optional. Allows to set custom rule tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#rule_tag NsxvDnat#rule_tag}
	RuleTag *float64 `field:"optional" json:"ruleTag" yaml:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#rule_type NsxvDnat#rule_type}
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// Translated address or address range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#translated_address NsxvDnat#translated_address}
	TranslatedAddress *string `field:"optional" json:"translatedAddress" yaml:"translatedAddress"`
	// Translated port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#translated_port NsxvDnat#translated_port}
	TranslatedPort *string `field:"optional" json:"translatedPort" yaml:"translatedPort"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat#vdc NsxvDnat#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

