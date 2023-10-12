package nsxvdhcprelay

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvDhcpRelayConfig struct {
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
	// Edge gateway name for DHCP relay settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#edge_gateway NsxvDhcpRelay#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// relay_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#relay_agent NsxvDhcpRelay#relay_agent}
	RelayAgent interface{} `field:"required" json:"relayAgent" yaml:"relayAgent"`
	// A set of IP domain names of DHCP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#domain_names NsxvDhcpRelay#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#id NsxvDhcpRelay#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of IP address of DHCP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#ip_addresses NsxvDhcpRelay#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// A set of IP set names which consist DHCP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#ip_sets NsxvDhcpRelay#ip_sets}
	IpSets *[]*string `field:"optional" json:"ipSets" yaml:"ipSets"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#org NsxvDhcpRelay#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#vdc NsxvDhcpRelay#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

