package edgegateway

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.edgegateway.Edgegateway",
		reflect.TypeOf((*Edgegateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "configurationInput", GoGetter: "ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultExternalNetworkIp", GoGetter: "DefaultExternalNetworkIp"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "distributedRouting", GoGetter: "DistributedRouting"},
			_jsii_.MemberProperty{JsiiProperty: "distributedRoutingInput", GoGetter: "DistributedRoutingInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalNetwork", GoGetter: "ExternalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "externalNetworkInput", GoGetter: "ExternalNetworkInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalNetworkIps", GoGetter: "ExternalNetworkIps"},
			_jsii_.MemberProperty{JsiiProperty: "fipsModeEnabled", GoGetter: "FipsModeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "fipsModeEnabledInput", GoGetter: "FipsModeEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fwDefaultRuleAction", GoGetter: "FwDefaultRuleAction"},
			_jsii_.MemberProperty{JsiiProperty: "fwDefaultRuleActionInput", GoGetter: "FwDefaultRuleActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fwDefaultRuleLoggingEnabled", GoGetter: "FwDefaultRuleLoggingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "fwDefaultRuleLoggingEnabledInput", GoGetter: "FwDefaultRuleLoggingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fwEnabled", GoGetter: "FwEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "fwEnabledInput", GoGetter: "FwEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "haEnabled", GoGetter: "HaEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "haEnabledInput", GoGetter: "HaEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lbAccelerationEnabled", GoGetter: "LbAccelerationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lbAccelerationEnabledInput", GoGetter: "LbAccelerationEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lbEnabled", GoGetter: "LbEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lbEnabledInput", GoGetter: "LbEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lbLoggingEnabled", GoGetter: "LbLoggingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lbLoggingEnabledInput", GoGetter: "LbLoggingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lbLoglevel", GoGetter: "LbLoglevel"},
			_jsii_.MemberProperty{JsiiProperty: "lbLoglevelInput", GoGetter: "LbLoglevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalNetwork", GoMethod: "PutExternalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDistributedRouting", GoMethod: "ResetDistributedRouting"},
			_jsii_.MemberMethod{JsiiMethod: "resetFipsModeEnabled", GoMethod: "ResetFipsModeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwDefaultRuleAction", GoMethod: "ResetFwDefaultRuleAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwDefaultRuleLoggingEnabled", GoMethod: "ResetFwDefaultRuleLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwEnabled", GoMethod: "ResetFwEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaEnabled", GoMethod: "ResetHaEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbAccelerationEnabled", GoMethod: "ResetLbAccelerationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbEnabled", GoMethod: "ResetLbEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbLoggingEnabled", GoMethod: "ResetLbLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbLoglevel", GoMethod: "ResetLbLoglevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseDefaultRouteForDnsRelay", GoMethod: "ResetUseDefaultRouteForDnsRelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultRouteForDnsRelay", GoGetter: "UseDefaultRouteForDnsRelay"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultRouteForDnsRelayInput", GoGetter: "UseDefaultRouteForDnsRelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
		},
		func() interface{} {
			j := jsiiProxy_Edgegateway{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegateway.EdgegatewayConfig",
		reflect.TypeOf((*EdgegatewayConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.edgegateway.EdgegatewayExternalNetwork",
		reflect.TypeOf((*EdgegatewayExternalNetwork)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkList",
		reflect.TypeOf((*EdgegatewayExternalNetworkList)(nil)).Elem(),
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
			j := jsiiProxy_EdgegatewayExternalNetworkList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkOutputReference",
		reflect.TypeOf((*EdgegatewayExternalNetworkOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableRateLimit", GoGetter: "EnableRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "enableRateLimitInput", GoGetter: "EnableRateLimitInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "incomingRateLimit", GoGetter: "IncomingRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "incomingRateLimitInput", GoGetter: "IncomingRateLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "outgoingRateLimit", GoGetter: "OutgoingRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "outgoingRateLimitInput", GoGetter: "OutgoingRateLimitInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSubnet", GoMethod: "PutSubnet"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableRateLimit", GoMethod: "ResetEnableRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncomingRateLimit", GoMethod: "ResetIncomingRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutgoingRateLimit", GoMethod: "ResetOutgoingRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnet", GoMethod: "ResetSubnet"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "subnet", GoGetter: "Subnet"},
			_jsii_.MemberProperty{JsiiProperty: "subnetInput", GoGetter: "SubnetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayExternalNetworkOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnet",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetList",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnetList)(nil)).Elem(),
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
			j := jsiiProxy_EdgegatewayExternalNetworkSubnetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetOutputReference",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressInput", GoGetter: "IpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "netmaskInput", GoGetter: "NetmaskInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSuballocatePool", GoMethod: "PutSuballocatePool"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpAddress", GoMethod: "ResetIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuballocatePool", GoMethod: "ResetSuballocatePool"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseForDefaultRoute", GoMethod: "ResetUseForDefaultRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "suballocatePool", GoGetter: "SuballocatePool"},
			_jsii_.MemberProperty{JsiiProperty: "suballocatePoolInput", GoGetter: "SuballocatePoolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useForDefaultRoute", GoGetter: "UseForDefaultRoute"},
			_jsii_.MemberProperty{JsiiProperty: "useForDefaultRouteInput", GoGetter: "UseForDefaultRouteInput"},
		},
		func() interface{} {
			j := jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetSuballocatePool",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnetSuballocatePool)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetSuballocatePoolList",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnetSuballocatePoolList)(nil)).Elem(),
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
			j := jsiiProxy_EdgegatewayExternalNetworkSubnetSuballocatePoolList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference",
		reflect.TypeOf((*EdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_EdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
