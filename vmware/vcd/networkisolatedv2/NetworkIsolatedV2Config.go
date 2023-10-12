package networkisolatedv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkIsolatedV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#gateway NetworkIsolatedV2#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// Network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#name NetworkIsolatedV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#prefix_length NetworkIsolatedV2#prefix_length}
	PrefixLength *float64 `field:"required" json:"prefixLength" yaml:"prefixLength"`
	// Network description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#description NetworkIsolatedV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// DNS server 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#dns1 NetworkIsolatedV2#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// DNS server 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#dns2 NetworkIsolatedV2#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// DNS suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#dns_suffix NetworkIsolatedV2#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// Boolean value if Dual-Stack mode should be enabled (default `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#dual_stack_enabled NetworkIsolatedV2#dual_stack_enabled}
	DualStackEnabled interface{} `field:"optional" json:"dualStackEnabled" yaml:"dualStackEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#id NetworkIsolatedV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// NSX-V only - share this network with other VDCs in this organization. Default - false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#is_shared NetworkIsolatedV2#is_shared}
	IsShared interface{} `field:"optional" json:"isShared" yaml:"isShared"`
	// Key value map of metadata to assign to this network. Key and value can be any string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#metadata NetworkIsolatedV2#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#metadata_entry NetworkIsolatedV2#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#org NetworkIsolatedV2#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// ID of VDC or VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#owner_id NetworkIsolatedV2#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Secondary gateway (can only be IPv6 and requires enabled Dual Stack mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#secondary_gateway NetworkIsolatedV2#secondary_gateway}
	SecondaryGateway *string `field:"optional" json:"secondaryGateway" yaml:"secondaryGateway"`
	// Secondary prefix (can only be IPv6 and requires enabled Dual Stack mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#secondary_prefix_length NetworkIsolatedV2#secondary_prefix_length}
	SecondaryPrefixLength *string `field:"optional" json:"secondaryPrefixLength" yaml:"secondaryPrefixLength"`
	// secondary_static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#secondary_static_ip_pool NetworkIsolatedV2#secondary_static_ip_pool}
	SecondaryStaticIpPool interface{} `field:"optional" json:"secondaryStaticIpPool" yaml:"secondaryStaticIpPool"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#static_ip_pool NetworkIsolatedV2#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2#vdc NetworkIsolatedV2#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

