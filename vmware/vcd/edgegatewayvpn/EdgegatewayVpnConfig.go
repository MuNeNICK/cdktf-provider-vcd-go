package edgegatewayvpn

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgegatewayVpnConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#edge_gateway EdgegatewayVpn#edge_gateway}.
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#encryption_protocol EdgegatewayVpn#encryption_protocol}.
	EncryptionProtocol *string `field:"required" json:"encryptionProtocol" yaml:"encryptionProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_id EdgegatewayVpn#local_id}.
	LocalId *string `field:"required" json:"localId" yaml:"localId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_ip_address EdgegatewayVpn#local_ip_address}.
	LocalIpAddress *string `field:"required" json:"localIpAddress" yaml:"localIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#mtu EdgegatewayVpn#mtu}.
	Mtu *float64 `field:"required" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#name EdgegatewayVpn#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_id EdgegatewayVpn#peer_id}.
	PeerId *string `field:"required" json:"peerId" yaml:"peerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_ip_address EdgegatewayVpn#peer_ip_address}.
	PeerIpAddress *string `field:"required" json:"peerIpAddress" yaml:"peerIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#shared_secret EdgegatewayVpn#shared_secret}.
	SharedSecret *string `field:"required" json:"sharedSecret" yaml:"sharedSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#description EdgegatewayVpn#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#id EdgegatewayVpn#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// local_subnets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_subnets EdgegatewayVpn#local_subnets}
	LocalSubnets interface{} `field:"optional" json:"localSubnets" yaml:"localSubnets"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#org EdgegatewayVpn#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// peer_subnets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_subnets EdgegatewayVpn#peer_subnets}
	PeerSubnets interface{} `field:"optional" json:"peerSubnets" yaml:"peerSubnets"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#vdc EdgegatewayVpn#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

