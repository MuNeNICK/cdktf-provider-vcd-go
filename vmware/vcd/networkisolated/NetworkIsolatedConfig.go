package networkisolated

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkIsolatedConfig struct {
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
	// A unique name for this network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#name NetworkIsolated#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description for the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#description NetworkIsolated#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dhcp_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#dhcp_pool NetworkIsolated#dhcp_pool}
	DhcpPool interface{} `field:"optional" json:"dhcpPool" yaml:"dhcpPool"`
	// First DNS server to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#dns1 NetworkIsolated#dns1}
	Dns1 *string `field:"optional" json:"dns1" yaml:"dns1"`
	// Second DNS server to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#dns2 NetworkIsolated#dns2}
	Dns2 *string `field:"optional" json:"dns2" yaml:"dns2"`
	// A FQDN for the virtual machines on this network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#dns_suffix NetworkIsolated#dns_suffix}
	DnsSuffix *string `field:"optional" json:"dnsSuffix" yaml:"dnsSuffix"`
	// The gateway for this network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#gateway NetworkIsolated#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#id NetworkIsolated#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key value map of metadata to assign to this network. Key and value can be any string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#metadata NetworkIsolated#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#metadata_entry NetworkIsolated#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The netmask for the new network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#netmask NetworkIsolated#netmask}
	Netmask *string `field:"optional" json:"netmask" yaml:"netmask"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#org NetworkIsolated#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Defines if this network is shared between multiple VDCs in the Org.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#shared NetworkIsolated#shared}
	Shared interface{} `field:"optional" json:"shared" yaml:"shared"`
	// static_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#static_ip_pool NetworkIsolated#static_ip_pool}
	StaticIpPool interface{} `field:"optional" json:"staticIpPool" yaml:"staticIpPool"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#vdc NetworkIsolated#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

