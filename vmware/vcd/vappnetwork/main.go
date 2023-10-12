package vappnetwork

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.vappNetwork.VappNetwork",
		reflect.TypeOf((*VappNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpPool", GoGetter: "DhcpPool"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpPoolInput", GoGetter: "DhcpPoolInput"},
			_jsii_.MemberProperty{JsiiProperty: "dns1", GoGetter: "Dns1"},
			_jsii_.MemberProperty{JsiiProperty: "dns1Input", GoGetter: "Dns1Input"},
			_jsii_.MemberProperty{JsiiProperty: "dns2", GoGetter: "Dns2"},
			_jsii_.MemberProperty{JsiiProperty: "dns2Input", GoGetter: "Dns2Input"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSuffix", GoGetter: "DnsSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSuffixInput", GoGetter: "DnsSuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayInput", GoGetter: "GatewayInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "guestVlanAllowed", GoGetter: "GuestVlanAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "guestVlanAllowedInput", GoGetter: "GuestVlanAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "netmaskInput", GoGetter: "NetmaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberProperty{JsiiProperty: "orgNetworkName", GoGetter: "OrgNetworkName"},
			_jsii_.MemberProperty{JsiiProperty: "orgNetworkNameInput", GoGetter: "OrgNetworkNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "prefixLength", GoGetter: "PrefixLength"},
			_jsii_.MemberProperty{JsiiProperty: "prefixLengthInput", GoGetter: "PrefixLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDhcpPool", GoMethod: "PutDhcpPool"},
			_jsii_.MemberMethod{JsiiMethod: "putStaticIpPool", GoMethod: "PutStaticIpPool"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "rebootVappOnRemoval", GoGetter: "RebootVappOnRemoval"},
			_jsii_.MemberProperty{JsiiProperty: "rebootVappOnRemovalInput", GoGetter: "RebootVappOnRemovalInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDhcpPool", GoMethod: "ResetDhcpPool"},
			_jsii_.MemberMethod{JsiiMethod: "resetDns1", GoMethod: "ResetDns1"},
			_jsii_.MemberMethod{JsiiMethod: "resetDns2", GoMethod: "ResetDns2"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsSuffix", GoMethod: "ResetDnsSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "resetGuestVlanAllowed", GoMethod: "ResetGuestVlanAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetmask", GoMethod: "ResetNetmask"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrgNetworkName", GoMethod: "ResetOrgNetworkName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixLength", GoMethod: "ResetPrefixLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetRebootVappOnRemoval", GoMethod: "ResetRebootVappOnRemoval"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetainIpMacEnabled", GoMethod: "ResetRetainIpMacEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetStaticIpPool", GoMethod: "ResetStaticIpPool"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberProperty{JsiiProperty: "retainIpMacEnabled", GoGetter: "RetainIpMacEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "retainIpMacEnabledInput", GoGetter: "RetainIpMacEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "staticIpPool", GoGetter: "StaticIpPool"},
			_jsii_.MemberProperty{JsiiProperty: "staticIpPoolInput", GoGetter: "StaticIpPoolInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vappName", GoGetter: "VappName"},
			_jsii_.MemberProperty{JsiiProperty: "vappNameInput", GoGetter: "VappNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
		},
		func() interface{} {
			j := jsiiProxy_VappNetwork{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.vappNetwork.VappNetworkConfig",
		reflect.TypeOf((*VappNetworkConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.vappNetwork.VappNetworkDhcpPool",
		reflect.TypeOf((*VappNetworkDhcpPool)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.vappNetwork.VappNetworkDhcpPoolList",
		reflect.TypeOf((*VappNetworkDhcpPoolList)(nil)).Elem(),
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
			j := jsiiProxy_VappNetworkDhcpPoolList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.vappNetwork.VappNetworkDhcpPoolOutputReference",
		reflect.TypeOf((*VappNetworkDhcpPoolOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTime", GoGetter: "DefaultLeaseTime"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTimeInput", GoGetter: "DefaultLeaseTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "endAddress", GoGetter: "EndAddress"},
			_jsii_.MemberProperty{JsiiProperty: "endAddressInput", GoGetter: "EndAddressInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTime", GoGetter: "MaxLeaseTime"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTimeInput", GoGetter: "MaxLeaseTimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLeaseTime", GoMethod: "ResetDefaultLeaseTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndAddress", GoMethod: "ResetEndAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTime", GoMethod: "ResetMaxLeaseTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startAddress", GoGetter: "StartAddress"},
			_jsii_.MemberProperty{JsiiProperty: "startAddressInput", GoGetter: "StartAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VappNetworkDhcpPoolOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.vappNetwork.VappNetworkStaticIpPool",
		reflect.TypeOf((*VappNetworkStaticIpPool)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.vappNetwork.VappNetworkStaticIpPoolList",
		reflect.TypeOf((*VappNetworkStaticIpPoolList)(nil)).Elem(),
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
			j := jsiiProxy_VappNetworkStaticIpPoolList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.vappNetwork.VappNetworkStaticIpPoolOutputReference",
		reflect.TypeOf((*VappNetworkStaticIpPoolOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endAddress", GoGetter: "EndAddress"},
			_jsii_.MemberProperty{JsiiProperty: "endAddressInput", GoGetter: "EndAddressInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startAddress", GoGetter: "StartAddress"},
			_jsii_.MemberProperty{JsiiProperty: "startAddressInput", GoGetter: "StartAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VappNetworkStaticIpPoolOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
