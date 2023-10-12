package datavcdvdcgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdVdcGroupConfig struct {
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
	// Default Policy Status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#default_policy_status DataVcdVdcGroup#default_policy_status}
	DefaultPolicyStatus interface{} `field:"optional" json:"defaultPolicyStatus" yaml:"defaultPolicyStatus"`
	// VDC group description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#description DataVcdVdcGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// More detailed error message when VDC group has error status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#error_message DataVcdVdcGroup#error_message}
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// VDC group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#id DataVcdVdcGroup#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#local_egress DataVcdVdcGroup#local_egress}
	LocalEgress interface{} `field:"optional" json:"localEgress" yaml:"localEgress"`
	// Name of VDC group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#name DataVcdVdcGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// ID of used network pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#network_pool_id DataVcdVdcGroup#network_pool_id}
	NetworkPoolId *string `field:"optional" json:"networkPoolId" yaml:"networkPoolId"`
	// The network provider’s universal id that is backing the universal network pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#network_pool_universal_id DataVcdVdcGroup#network_pool_universal_id}
	NetworkPoolUniversalId *string `field:"optional" json:"networkPoolUniversalId" yaml:"networkPoolUniversalId"`
	// Defines the networking provider backing the VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#network_provider_type DataVcdVdcGroup#network_provider_type}
	NetworkProviderType *string `field:"optional" json:"networkProviderType" yaml:"networkProviderType"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#org DataVcdVdcGroup#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#status DataVcdVdcGroup#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Defines the group as LOCAL or UNIVERSAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#type DataVcdVdcGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// True means that a VDC group router has been created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group#universal_networking_enabled DataVcdVdcGroup#universal_networking_enabled}
	UniversalNetworkingEnabled interface{} `field:"optional" json:"universalNetworkingEnabled" yaml:"universalNetworkingEnabled"`
}

