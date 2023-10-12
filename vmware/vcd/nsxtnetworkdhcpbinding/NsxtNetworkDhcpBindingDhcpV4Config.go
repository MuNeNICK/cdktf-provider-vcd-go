package nsxtnetworkdhcpbinding


type NsxtNetworkDhcpBindingDhcpV4Config struct {
	// IPv4 gateway address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#gateway_ip_address NsxtNetworkDhcpBinding#gateway_ip_address}
	GatewayIpAddress *string `field:"optional" json:"gatewayIpAddress" yaml:"gatewayIpAddress"`
	// Hostname for the DHCP client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#hostname NsxtNetworkDhcpBinding#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

