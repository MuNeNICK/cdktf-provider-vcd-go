package nsxtnetworkimported

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtNetworkImportedConfig struct {
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
	// Gateway IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#gateway NsxtNetworkImported#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#name NsxtNetworkImported#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#prefix_length NsxtNetworkImported#prefix_length}
	PrefixLength *float64 `field:"required" json:"prefixLength" yaml:"prefixLength"`
	// Network description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#description NsxtNetworkImported#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// DNS server 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#dns1 NsxtNetworkImported#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// DNS server 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#dns2 NsxtNetworkImported#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// DNS suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#dns_suffix NsxtNetworkImported#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// Boolean value if Dual-Stack mode should be enabled (default `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#dual_stack_enabled NsxtNetworkImported#dual_stack_enabled}
	DualStackEnabled interface{} `field:"optional" json:"dualStackEnabled" yaml:"dualStackEnabled"`
	// Name of existing Distributed Virtual Port Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#dvpg_name NsxtNetworkImported#dvpg_name}
	DvpgName *string `field:"optional" json:"dvpgName" yaml:"dvpgName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#id NsxtNetworkImported#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of existing NSX-T Logical Switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#nsxt_logical_switch_name NsxtNetworkImported#nsxt_logical_switch_name}
	NsxtLogicalSwitchName *string `field:"optional" json:"nsxtLogicalSwitchName" yaml:"nsxtLogicalSwitchName"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#org NsxtNetworkImported#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// ID of VDC or VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#owner_id NsxtNetworkImported#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Secondary gateway (can only be IPv6 and requires enabled Dual Stack mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#secondary_gateway NsxtNetworkImported#secondary_gateway}
	SecondaryGateway *string `field:"optional" json:"secondaryGateway" yaml:"secondaryGateway"`
	// Secondary prefix (can only be IPv6 and requires enabled Dual Stack mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#secondary_prefix_length NsxtNetworkImported#secondary_prefix_length}
	SecondaryPrefixLength *string `field:"optional" json:"secondaryPrefixLength" yaml:"secondaryPrefixLength"`
	// secondary_static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#secondary_static_ip_pool NsxtNetworkImported#secondary_static_ip_pool}
	SecondaryStaticIpPool interface{} `field:"optional" json:"secondaryStaticIpPool" yaml:"secondaryStaticIpPool"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#static_ip_pool NsxtNetworkImported#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported#vdc NsxtNetworkImported#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

