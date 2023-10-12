package edgegatewayvpn


type EdgegatewayVpnLocalSubnets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_subnet_gateway EdgegatewayVpn#local_subnet_gateway}.
	LocalSubnetGateway *string `field:"required" json:"localSubnetGateway" yaml:"localSubnetGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_subnet_mask EdgegatewayVpn#local_subnet_mask}.
	LocalSubnetMask *string `field:"required" json:"localSubnetMask" yaml:"localSubnetMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn#local_subnet_name EdgegatewayVpn#local_subnet_name}.
	LocalSubnetName *string `field:"required" json:"localSubnetName" yaml:"localSubnetName"`
}

