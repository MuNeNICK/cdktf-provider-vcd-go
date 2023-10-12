package edgegatewayvpn


type EdgegatewayVpnPeerSubnets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_subnet_gateway EdgegatewayVpn#peer_subnet_gateway}.
	PeerSubnetGateway *string `field:"required" json:"peerSubnetGateway" yaml:"peerSubnetGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_subnet_mask EdgegatewayVpn#peer_subnet_mask}.
	PeerSubnetMask *string `field:"required" json:"peerSubnetMask" yaml:"peerSubnetMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#peer_subnet_name EdgegatewayVpn#peer_subnet_name}.
	PeerSubnetName *string `field:"required" json:"peerSubnetName" yaml:"peerSubnetName"`
}

