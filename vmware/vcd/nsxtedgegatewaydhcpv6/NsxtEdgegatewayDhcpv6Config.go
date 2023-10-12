package nsxtedgegatewaydhcpv6

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayDhcpv6Config struct {
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
	// Edge gateway ID for Rate limiting (DHCPv6) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#edge_gateway_id NsxtEdgegatewayDhcpv6#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// DHCPv6 configuration mode. One of 'SLAAC', 'DHCPv6', 'DISABLED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#mode NsxtEdgegatewayDhcpv6#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// A set of DNS Servers (only applicable for 'SLAAC' mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#dns_servers NsxtEdgegatewayDhcpv6#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A set of domain names (only applicable for 'SLAAC' mode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#domain_names NsxtEdgegatewayDhcpv6#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#id NsxtEdgegatewayDhcpv6#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_dhcpv6#org NsxtEdgegatewayDhcpv6#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
}

