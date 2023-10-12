package edgegatewayvpn

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		reflect.TypeOf((*EdgegatewayVpn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGateway", GoGetter: "EdgeGateway"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayInput", GoGetter: "EdgeGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionProtocol", GoGetter: "EncryptionProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionProtocolInput", GoGetter: "EncryptionProtocolInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localId", GoGetter: "LocalId"},
			_jsii_.MemberProperty{JsiiProperty: "localIdInput", GoGetter: "LocalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "localIpAddress", GoGetter: "LocalIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "localIpAddressInput", GoGetter: "LocalIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnets", GoGetter: "LocalSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetsInput", GoGetter: "LocalSubnetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "mtu", GoGetter: "Mtu"},
			_jsii_.MemberProperty{JsiiProperty: "mtuInput", GoGetter: "MtuInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerId", GoGetter: "PeerId"},
			_jsii_.MemberProperty{JsiiProperty: "peerIdInput", GoGetter: "PeerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerIpAddress", GoGetter: "PeerIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "peerIpAddressInput", GoGetter: "PeerIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnets", GoGetter: "PeerSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetsInput", GoGetter: "PeerSubnetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putLocalSubnets", GoMethod: "PutLocalSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "putPeerSubnets", GoMethod: "PutPeerSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalSubnets", GoMethod: "ResetLocalSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerSubnets", GoMethod: "ResetPeerSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberProperty{JsiiProperty: "sharedSecret", GoGetter: "SharedSecret"},
			_jsii_.MemberProperty{JsiiProperty: "sharedSecretInput", GoGetter: "SharedSecretInput"},
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
			j := jsiiProxy_EdgegatewayVpn{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegatewayVpn.EdgegatewayVpnConfig",
		reflect.TypeOf((*EdgegatewayVpnConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.edgegatewayVpn.EdgegatewayVpnLocalSubnets",
		reflect.TypeOf((*EdgegatewayVpnLocalSubnets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.edgegatewayVpn.EdgegatewayVpnLocalSubnetsList",
		reflect.TypeOf((*EdgegatewayVpnLocalSubnetsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayVpnLocalSubnetsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.edgegatewayVpn.EdgegatewayVpnLocalSubnetsOutputReference",
		reflect.TypeOf((*EdgegatewayVpnLocalSubnetsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetGateway", GoGetter: "LocalSubnetGateway"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetGatewayInput", GoGetter: "LocalSubnetGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetMask", GoGetter: "LocalSubnetMask"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetMaskInput", GoGetter: "LocalSubnetMaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetName", GoGetter: "LocalSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "localSubnetNameInput", GoGetter: "LocalSubnetNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegatewayVpn.EdgegatewayVpnPeerSubnets",
		reflect.TypeOf((*EdgegatewayVpnPeerSubnets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.edgegatewayVpn.EdgegatewayVpnPeerSubnetsList",
		reflect.TypeOf((*EdgegatewayVpnPeerSubnetsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayVpnPeerSubnetsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.edgegatewayVpn.EdgegatewayVpnPeerSubnetsOutputReference",
		reflect.TypeOf((*EdgegatewayVpnPeerSubnetsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetGateway", GoGetter: "PeerSubnetGateway"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetGatewayInput", GoGetter: "PeerSubnetGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetMask", GoGetter: "PeerSubnetMask"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetMaskInput", GoGetter: "PeerSubnetMaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetName", GoGetter: "PeerSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "peerSubnetNameInput", GoGetter: "PeerSubnetNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayVpnPeerSubnetsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
