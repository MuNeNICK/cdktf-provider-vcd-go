package nsxtnetworkdhcpbinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtNetworkDhcpBindingConfig struct {
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
	// Binding type 'IPV4' or 'IPV6'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#binding_type NsxtNetworkDhcpBinding#binding_type}
	BindingType *string `field:"required" json:"bindingType" yaml:"bindingType"`
	// IP address of the DHCP binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#ip_address NsxtNetworkDhcpBinding#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Lease time in seconds. Minimum value is 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#lease_time NsxtNetworkDhcpBinding#lease_time}
	LeaseTime *float64 `field:"required" json:"leaseTime" yaml:"leaseTime"`
	// MAC address of the DHCP binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#mac_address NsxtNetworkDhcpBinding#mac_address}
	MacAddress *string `field:"required" json:"macAddress" yaml:"macAddress"`
	// Name of DHCP binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#name NsxtNetworkDhcpBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parent Org VDC network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#org_network_id NsxtNetworkDhcpBinding#org_network_id}
	OrgNetworkId *string `field:"required" json:"orgNetworkId" yaml:"orgNetworkId"`
	// Description of DHCP binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#description NsxtNetworkDhcpBinding#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dhcp_v4_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#dhcp_v4_config NsxtNetworkDhcpBinding#dhcp_v4_config}
	DhcpV4Config *NsxtNetworkDhcpBindingDhcpV4Config `field:"optional" json:"dhcpV4Config" yaml:"dhcpV4Config"`
	// dhcp_v6_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#dhcp_v6_config NsxtNetworkDhcpBinding#dhcp_v6_config}
	DhcpV6Config *NsxtNetworkDhcpBindingDhcpV6Config `field:"optional" json:"dhcpV6Config" yaml:"dhcpV6Config"`
	// The DNS server IPs to be assigned . 2 values maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#dns_servers NsxtNetworkDhcpBinding#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#id NsxtNetworkDhcpBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#org NsxtNetworkDhcpBinding#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
}

