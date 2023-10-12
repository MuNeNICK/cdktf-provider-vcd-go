package edgegatewaysettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		reflect.TypeOf((*EdgegatewaySettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayId", GoGetter: "EdgeGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayIdInput", GoGetter: "EdgeGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayName", GoGetter: "EdgeGatewayName"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayNameInput", GoGetter: "EdgeGatewayNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetEdgeGatewayId", GoMethod: "ResetEdgeGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEdgeGatewayName", GoMethod: "ResetEdgeGatewayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwDefaultRuleAction", GoMethod: "ResetFwDefaultRuleAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwDefaultRuleLoggingEnabled", GoMethod: "ResetFwDefaultRuleLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFwEnabled", GoMethod: "ResetFwEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbAccelerationEnabled", GoMethod: "ResetLbAccelerationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbEnabled", GoMethod: "ResetLbEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbLoggingEnabled", GoMethod: "ResetLbLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLbLoglevel", GoMethod: "ResetLbLoglevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
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
			j := jsiiProxy_EdgegatewaySettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.edgegatewaySettings.EdgegatewaySettingsConfig",
		reflect.TypeOf((*EdgegatewaySettingsConfig)(nil)).Elem(),
	)
}
