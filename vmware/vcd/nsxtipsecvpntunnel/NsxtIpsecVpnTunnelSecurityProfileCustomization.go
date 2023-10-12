package nsxtipsecvpntunnel


type NsxtIpsecVpnTunnelSecurityProfileCustomization struct {
	// Diffie-Hellman groups to be used if Perfect Forward Secrecy is enabled.
	//
	// One of GROUP2, GROUP5, GROUP14, GROUP15, GROUP16, GROUP19, GROUP20, GROUP21
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ike_dh_groups NsxtIpsecVpnTunnel#ike_dh_groups}
	IkeDhGroups *[]*string `field:"required" json:"ikeDhGroups" yaml:"ikeDhGroups"`
	// Encryption algorithms. One of SHA1, SHA2_256, SHA2_384, SHA2_512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ike_encryption_algorithms NsxtIpsecVpnTunnel#ike_encryption_algorithms}
	IkeEncryptionAlgorithms *[]*string `field:"required" json:"ikeEncryptionAlgorithms" yaml:"ikeEncryptionAlgorithms"`
	// IKE version one of IKE_V1, IKE_V2, IKE_FLEX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ike_version NsxtIpsecVpnTunnel#ike_version}
	IkeVersion *string `field:"required" json:"ikeVersion" yaml:"ikeVersion"`
	// Diffie-Hellman groups to be used is PFS is enabled. One of GROUP2, GROUP5, GROUP14, GROUP15, GROUP16, GROUP19, GROUP20, GROUP21.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_dh_groups NsxtIpsecVpnTunnel#tunnel_dh_groups}
	TunnelDhGroups *[]*string `field:"required" json:"tunnelDhGroups" yaml:"tunnelDhGroups"`
	// Encryption algorithms to use in IPSec tunnel establishment. One of AES_128, AES_256, AES_GCM_128, AES_GCM_192, AES_GCM_256, NO_ENCRYPTION_AUTH_AES_GMAC_128, NO_ENCRYPTION_AUTH_AES_GMAC_192, NO_ENCRYPTION_AUTH_AES_GMAC_256, NO_ENCRYPTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_encryption_algorithms NsxtIpsecVpnTunnel#tunnel_encryption_algorithms}
	TunnelEncryptionAlgorithms *[]*string `field:"required" json:"tunnelEncryptionAlgorithms" yaml:"tunnelEncryptionAlgorithms"`
	// Value in seconds of dead probe detection interval. Minimum is 3 seconds and the maximum is 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#dpd_probe_internal NsxtIpsecVpnTunnel#dpd_probe_internal}
	DpdProbeInternal *float64 `field:"optional" json:"dpdProbeInternal" yaml:"dpdProbeInternal"`
	// Secure hashing algorithms to use during the IKE negotiation. One of SHA1, SHA2_256, SHA2_384, SHA2_512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ike_digest_algorithms NsxtIpsecVpnTunnel#ike_digest_algorithms}
	IkeDigestAlgorithms *[]*string `field:"optional" json:"ikeDigestAlgorithms" yaml:"ikeDigestAlgorithms"`
	// Security Association life time (in seconds). It is number of seconds before the IPsec tunnel needs to reestablish.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#ike_sa_lifetime NsxtIpsecVpnTunnel#ike_sa_lifetime}
	IkeSaLifetime *float64 `field:"optional" json:"ikeSaLifetime" yaml:"ikeSaLifetime"`
	// Policy for handling defragmentation bit. One of COPY, CLEAR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_df_policy NsxtIpsecVpnTunnel#tunnel_df_policy}
	TunnelDfPolicy *string `field:"optional" json:"tunnelDfPolicy" yaml:"tunnelDfPolicy"`
	// Digest algorithms to be used for message digest. One of SHA1, SHA2_256, SHA2_384, SHA2_512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_digest_algorithms NsxtIpsecVpnTunnel#tunnel_digest_algorithms}
	TunnelDigestAlgorithms *[]*string `field:"optional" json:"tunnelDigestAlgorithms" yaml:"tunnelDigestAlgorithms"`
	// Perfect Forward Secrecy Enabled or Disabled. Default (enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_pfs_enabled NsxtIpsecVpnTunnel#tunnel_pfs_enabled}
	TunnelPfsEnabled interface{} `field:"optional" json:"tunnelPfsEnabled" yaml:"tunnelPfsEnabled"`
	// Security Association life time (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel#tunnel_sa_lifetime NsxtIpsecVpnTunnel#tunnel_sa_lifetime}
	TunnelSaLifetime *float64 `field:"optional" json:"tunnelSaLifetime" yaml:"tunnelSaLifetime"`
}

