package nsxvdhcprelay


type NsxvDhcpRelayRelayAgent struct {
	// Org network which is to be used for relaying DHCP message to specified servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#network_name NsxvDhcpRelay#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// Optional gateway IP address of org network which is to be used for relaying DHCP message to specified servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dhcp_relay#gateway_ip_address NsxvDhcpRelay#gateway_ip_address}
	GatewayIpAddress *string `field:"optional" json:"gatewayIpAddress" yaml:"gatewayIpAddress"`
}

