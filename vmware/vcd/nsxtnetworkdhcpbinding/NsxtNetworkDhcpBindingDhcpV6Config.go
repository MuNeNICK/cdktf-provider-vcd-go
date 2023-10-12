package nsxtnetworkdhcpbinding


type NsxtNetworkDhcpBindingDhcpV6Config struct {
	// Set of domain names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#domain_names NsxtNetworkDhcpBinding#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Set of SNTP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding#sntp_servers NsxtNetworkDhcpBinding#sntp_servers}
	SntpServers *[]*string `field:"optional" json:"sntpServers" yaml:"sntpServers"`
}

