package nsxvsnat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvSnatConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#edge_gateway NsxvSnat#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Org or external network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#network_name NsxvSnat#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// Network type. One of 'ext', 'org'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#network_type NsxvSnat#network_type}
	NetworkType *string `field:"required" json:"networkType" yaml:"networkType"`
	// Original address or address range. This is the the source address for SNAT rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#original_address NsxvSnat#original_address}
	OriginalAddress *string `field:"required" json:"originalAddress" yaml:"originalAddress"`
	// NAT rule description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#description NsxvSnat#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule should be enabled. Default 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#enabled NsxvSnat#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#id NsxvSnat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether logging should be enabled for this rule. Default 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#logging_enabled NsxvSnat#logging_enabled}
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#org NsxvSnat#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Optional. Allows to set custom rule tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#rule_tag NsxvSnat#rule_tag}
	RuleTag *float64 `field:"optional" json:"ruleTag" yaml:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#rule_type NsxvSnat#rule_type}
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// Translated address or address range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#translated_address NsxvSnat#translated_address}
	TranslatedAddress *string `field:"optional" json:"translatedAddress" yaml:"translatedAddress"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_snat#vdc NsxvSnat#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

