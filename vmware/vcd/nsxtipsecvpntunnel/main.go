package nsxtipsecvpntunnel

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		reflect.TypeOf((*NsxtIpsecVpnTunnel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationMode", GoGetter: "AuthenticationMode"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationModeInput", GoGetter: "AuthenticationModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificateId", GoGetter: "CaCertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificateIdInput", GoGetter: "CaCertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificateId", GoGetter: "CertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "certificateIdInput", GoGetter: "CertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayId", GoGetter: "EdgeGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayIdInput", GoGetter: "EdgeGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeFailReason", GoGetter: "IkeFailReason"},
			_jsii_.MemberProperty{JsiiProperty: "ikeServiceStatus", GoGetter: "IkeServiceStatus"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localIpAddress", GoGetter: "LocalIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "localIpAddressInput", GoGetter: "LocalIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "localNetworks", GoGetter: "LocalNetworks"},
			_jsii_.MemberProperty{JsiiProperty: "localNetworksInput", GoGetter: "LocalNetworksInput"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInput", GoGetter: "LoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preSharedKey", GoGetter: "PreSharedKey"},
			_jsii_.MemberProperty{JsiiProperty: "preSharedKeyInput", GoGetter: "PreSharedKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putSecurityProfileCustomization", GoMethod: "PutSecurityProfileCustomization"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteId", GoGetter: "RemoteId"},
			_jsii_.MemberProperty{JsiiProperty: "remoteIdInput", GoGetter: "RemoteIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "remoteIpAddress", GoGetter: "RemoteIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "remoteIpAddressInput", GoGetter: "RemoteIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "remoteNetworks", GoGetter: "RemoteNetworks"},
			_jsii_.MemberProperty{JsiiProperty: "remoteNetworksInput", GoGetter: "RemoteNetworksInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticationMode", GoMethod: "ResetAuthenticationMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertificateId", GoMethod: "ResetCaCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateId", GoMethod: "ResetCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogging", GoMethod: "ResetLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteId", GoMethod: "ResetRemoteId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteNetworks", GoMethod: "ResetRemoteNetworks"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityProfileCustomization", GoMethod: "ResetSecurityProfileCustomization"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberProperty{JsiiProperty: "securityProfile", GoGetter: "SecurityProfile"},
			_jsii_.MemberProperty{JsiiProperty: "securityProfileCustomization", GoGetter: "SecurityProfileCustomization"},
			_jsii_.MemberProperty{JsiiProperty: "securityProfileCustomizationInput", GoGetter: "SecurityProfileCustomizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
		},
		func() interface{} {
			j := jsiiProxy_NsxtIpsecVpnTunnel{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnelConfig",
		reflect.TypeOf((*NsxtIpsecVpnTunnelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnelSecurityProfileCustomization",
		reflect.TypeOf((*NsxtIpsecVpnTunnelSecurityProfileCustomization)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference",
		reflect.TypeOf((*NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dpdProbeInternal", GoGetter: "DpdProbeInternal"},
			_jsii_.MemberProperty{JsiiProperty: "dpdProbeInternalInput", GoGetter: "DpdProbeInternalInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ikeDhGroups", GoGetter: "IkeDhGroups"},
			_jsii_.MemberProperty{JsiiProperty: "ikeDhGroupsInput", GoGetter: "IkeDhGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeDigestAlgorithms", GoGetter: "IkeDigestAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "ikeDigestAlgorithmsInput", GoGetter: "IkeDigestAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeEncryptionAlgorithms", GoGetter: "IkeEncryptionAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "ikeEncryptionAlgorithmsInput", GoGetter: "IkeEncryptionAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeSaLifetime", GoGetter: "IkeSaLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "ikeSaLifetimeInput", GoGetter: "IkeSaLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeVersion", GoGetter: "IkeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "ikeVersionInput", GoGetter: "IkeVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDpdProbeInternal", GoMethod: "ResetDpdProbeInternal"},
			_jsii_.MemberMethod{JsiiMethod: "resetIkeDigestAlgorithms", GoMethod: "ResetIkeDigestAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetIkeSaLifetime", GoMethod: "ResetIkeSaLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelDfPolicy", GoMethod: "ResetTunnelDfPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelDigestAlgorithms", GoMethod: "ResetTunnelDigestAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelPfsEnabled", GoMethod: "ResetTunnelPfsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelSaLifetime", GoMethod: "ResetTunnelSaLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDfPolicy", GoGetter: "TunnelDfPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDfPolicyInput", GoGetter: "TunnelDfPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDhGroups", GoGetter: "TunnelDhGroups"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDhGroupsInput", GoGetter: "TunnelDhGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDigestAlgorithms", GoGetter: "TunnelDigestAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelDigestAlgorithmsInput", GoGetter: "TunnelDigestAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelEncryptionAlgorithms", GoGetter: "TunnelEncryptionAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelEncryptionAlgorithmsInput", GoGetter: "TunnelEncryptionAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelPfsEnabled", GoGetter: "TunnelPfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelPfsEnabledInput", GoGetter: "TunnelPfsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelSaLifetime", GoGetter: "TunnelSaLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelSaLifetimeInput", GoGetter: "TunnelSaLifetimeInput"},
		},
		func() interface{} {
			j := jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
