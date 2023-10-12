package nsxtnetworkdhcp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtNetworkDhcpConfig struct {
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
	// Parent Org VDC network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#org_network_id NsxtNetworkDhcp#org_network_id}
	OrgNetworkId *string `field:"required" json:"orgNetworkId" yaml:"orgNetworkId"`
	// The DNS server IPs to be assigned by this DHCP service. 2 values maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#dns_servers NsxtNetworkDhcp#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#id NsxtNetworkDhcp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Lease time in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#lease_time NsxtNetworkDhcp#lease_time}
	LeaseTime *float64 `field:"optional" json:"leaseTime" yaml:"leaseTime"`
	// IP Address of DHCP server in network. Only applicable when mode=NETWORK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#listener_ip_address NsxtNetworkDhcp#listener_ip_address}
	ListenerIpAddress *string `field:"optional" json:"listenerIpAddress" yaml:"listenerIpAddress"`
	// DHCP mode. One of 'EDGE' (default), 'NETWORK', 'RELAY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#mode NsxtNetworkDhcp#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#org NsxtNetworkDhcp#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#pool NsxtNetworkDhcp#pool}
	Pool interface{} `field:"optional" json:"pool" yaml:"pool"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp#vdc NsxtNetworkDhcp#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

