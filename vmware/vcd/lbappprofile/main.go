package lbappprofile

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.lbAppProfile.LbAppProfile",
		reflect.TypeOf((*LbAppProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "cookieMode", GoGetter: "CookieMode"},
			_jsii_.MemberProperty{JsiiProperty: "cookieModeInput", GoGetter: "CookieModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cookieName", GoGetter: "CookieName"},
			_jsii_.MemberProperty{JsiiProperty: "cookieNameInput", GoGetter: "CookieNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGateway", GoGetter: "EdgeGateway"},
			_jsii_.MemberProperty{JsiiProperty: "edgeGatewayInput", GoGetter: "EdgeGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "enablePoolSideSsl", GoGetter: "EnablePoolSideSsl"},
			_jsii_.MemberProperty{JsiiProperty: "enablePoolSideSslInput", GoGetter: "EnablePoolSideSslInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSslPassthrough", GoGetter: "EnableSslPassthrough"},
			_jsii_.MemberProperty{JsiiProperty: "enableSslPassthroughInput", GoGetter: "EnableSslPassthroughInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiration", GoGetter: "Expiration"},
			_jsii_.MemberProperty{JsiiProperty: "expirationInput", GoGetter: "ExpirationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpRedirectUrl", GoGetter: "HttpRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "httpRedirectUrlInput", GoGetter: "HttpRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "insertXForwardedHttpHeader", GoGetter: "InsertXForwardedHttpHeader"},
			_jsii_.MemberProperty{JsiiProperty: "insertXForwardedHttpHeaderInput", GoGetter: "InsertXForwardedHttpHeaderInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "persistenceMechanism", GoGetter: "PersistenceMechanism"},
			_jsii_.MemberProperty{JsiiProperty: "persistenceMechanismInput", GoGetter: "PersistenceMechanismInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookieMode", GoMethod: "ResetCookieMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookieName", GoMethod: "ResetCookieName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePoolSideSsl", GoMethod: "ResetEnablePoolSideSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSslPassthrough", GoMethod: "ResetEnableSslPassthrough"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpiration", GoMethod: "ResetExpiration"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRedirectUrl", GoMethod: "ResetHttpRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsertXForwardedHttpHeader", GoMethod: "ResetInsertXForwardedHttpHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistenceMechanism", GoMethod: "ResetPersistenceMechanism"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
		},
		func() interface{} {
			j := jsiiProxy_LbAppProfile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.lbAppProfile.LbAppProfileConfig",
		reflect.TypeOf((*LbAppProfileConfig)(nil)).Elem(),
	)
}
