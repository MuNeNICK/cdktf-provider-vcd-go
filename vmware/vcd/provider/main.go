package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.provider.VcdProvider",
		reflect.TypeOf((*VcdProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowApiTokenFile", GoGetter: "AllowApiTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "allowApiTokenFileInput", GoGetter: "AllowApiTokenFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowServiceAccountTokenFile", GoGetter: "AllowServiceAccountTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "allowServiceAccountTokenFileInput", GoGetter: "AllowServiceAccountTokenFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSsl", GoGetter: "AllowUnverifiedSsl"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSslInput", GoGetter: "AllowUnverifiedSslInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiToken", GoGetter: "ApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenFile", GoGetter: "ApiTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenFileInput", GoGetter: "ApiTokenFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenInput", GoGetter: "ApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "authType", GoGetter: "AuthType"},
			_jsii_.MemberProperty{JsiiProperty: "authTypeInput", GoGetter: "AuthTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreMetadataChanges", GoGetter: "IgnoreMetadataChanges"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreMetadataChangesInput", GoGetter: "IgnoreMetadataChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "importSeparator", GoGetter: "ImportSeparator"},
			_jsii_.MemberProperty{JsiiProperty: "importSeparatorInput", GoGetter: "ImportSeparatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "loggingFile", GoGetter: "LoggingFile"},
			_jsii_.MemberProperty{JsiiProperty: "loggingFileInput", GoGetter: "LoggingFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInput", GoGetter: "LoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetryTimeout", GoGetter: "MaxRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetryTimeoutInput", GoGetter: "MaxRetryTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowApiTokenFile", GoMethod: "ResetAllowApiTokenFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowServiceAccountTokenFile", GoMethod: "ResetAllowServiceAccountTokenFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUnverifiedSsl", GoMethod: "ResetAllowUnverifiedSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiToken", GoMethod: "ResetApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiTokenFile", GoMethod: "ResetApiTokenFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthType", GoMethod: "ResetAuthType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreMetadataChanges", GoMethod: "ResetIgnoreMetadataChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetImportSeparator", GoMethod: "ResetImportSeparator"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogging", GoMethod: "ResetLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingFile", GoMethod: "ResetLoggingFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetryTimeout", GoMethod: "ResetMaxRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetSamlAdfsRptId", GoMethod: "ResetSamlAdfsRptId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountTokenFile", GoMethod: "ResetServiceAccountTokenFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSysorg", GoMethod: "ResetSysorg"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUser", GoMethod: "ResetUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberProperty{JsiiProperty: "samlAdfsRptId", GoGetter: "SamlAdfsRptId"},
			_jsii_.MemberProperty{JsiiProperty: "samlAdfsRptIdInput", GoGetter: "SamlAdfsRptIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountTokenFile", GoGetter: "ServiceAccountTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountTokenFileInput", GoGetter: "ServiceAccountTokenFileInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "sysorg", GoGetter: "Sysorg"},
			_jsii_.MemberProperty{JsiiProperty: "sysorgInput", GoGetter: "SysorgInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
		},
		func() interface{} {
			j := jsiiProxy_VcdProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.provider.VcdProviderConfig",
		reflect.TypeOf((*VcdProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.provider.VcdProviderIgnoreMetadataChanges",
		reflect.TypeOf((*VcdProviderIgnoreMetadataChanges)(nil)).Elem(),
	)
}
