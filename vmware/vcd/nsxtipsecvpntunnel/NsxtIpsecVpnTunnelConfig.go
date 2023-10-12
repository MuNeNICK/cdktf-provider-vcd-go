package nsxtipsecvpntunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtIpsecVpnTunnelConfig struct {
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
	// Edge gateway name in which IP Sec VPN configuration is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#edge_gateway_id NsxtIpsecVpnTunnel#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// IPv4 Address for the endpoint. This has to be a sub-allocated IP on the Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#local_ip_address NsxtIpsecVpnTunnel#local_ip_address}
	LocalIpAddress *string `field:"required" json:"localIpAddress" yaml:"localIpAddress"`
	// Set of local networks in CIDR format. At least one value is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#local_networks NsxtIpsecVpnTunnel#local_networks}
	LocalNetworks *[]*string `field:"required" json:"localNetworks" yaml:"localNetworks"`
	// Name of IP Sec VPN Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#name NsxtIpsecVpnTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Pre-Shared Key (PSK).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#pre_shared_key NsxtIpsecVpnTunnel#pre_shared_key}
	PreSharedKey *string `field:"required" json:"preSharedKey" yaml:"preSharedKey"`
	// Public IPv4 Address of the remote device terminating the VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#remote_ip_address NsxtIpsecVpnTunnel#remote_ip_address}
	RemoteIpAddress *string `field:"required" json:"remoteIpAddress" yaml:"remoteIpAddress"`
	// One of 'PSK' (default), 'CERTIFICATE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#authentication_mode NsxtIpsecVpnTunnel#authentication_mode}
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Optional CA certificate ID to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ca_certificate_id NsxtIpsecVpnTunnel#ca_certificate_id}
	CaCertificateId *string `field:"optional" json:"caCertificateId" yaml:"caCertificateId"`
	// Optional certificate ID to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#certificate_id NsxtIpsecVpnTunnel#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// Description IP Sec VPN Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#description NsxtIpsecVpnTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables or disables this configuration (default true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#enabled NsxtIpsecVpnTunnel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#id NsxtIpsecVpnTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets whether logging for the tunnel is enabled or not. (default - false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#logging NsxtIpsecVpnTunnel#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#org NsxtIpsecVpnTunnel#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Custom remote ID of the peer site. 'remote_ip_address' is used by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#remote_id NsxtIpsecVpnTunnel#remote_id}
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#remote_networks NsxtIpsecVpnTunnel#remote_networks}
	RemoteNetworks *[]*string `field:"optional" json:"remoteNetworks" yaml:"remoteNetworks"`
	// security_profile_customization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#security_profile_customization NsxtIpsecVpnTunnel#security_profile_customization}
	SecurityProfileCustomization *NsxtIpsecVpnTunnelSecurityProfileCustomization `field:"optional" json:"securityProfileCustomization" yaml:"securityProfileCustomization"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#vdc NsxtIpsecVpnTunnel#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

