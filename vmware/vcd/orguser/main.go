package orguser

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.orgUser.OrgUser",
		reflect.TypeOf((*OrgUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deployedVmQuota", GoGetter: "DeployedVmQuota"},
			_jsii_.MemberProperty{JsiiProperty: "deployedVmQuotaInput", GoGetter: "DeployedVmQuotaInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddress", GoGetter: "EmailAddress"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddressInput", GoGetter: "EmailAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fullName", GoGetter: "FullName"},
			_jsii_.MemberProperty{JsiiProperty: "fullNameInput", GoGetter: "FullNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupNames", GoGetter: "GroupNames"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "instantMessaging", GoGetter: "InstantMessaging"},
			_jsii_.MemberProperty{JsiiProperty: "instantMessagingInput", GoGetter: "InstantMessagingInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isExternal", GoGetter: "IsExternal"},
			_jsii_.MemberProperty{JsiiProperty: "isExternalInput", GoGetter: "IsExternalInput"},
			_jsii_.MemberProperty{JsiiProperty: "isGroupRole", GoGetter: "IsGroupRole"},
			_jsii_.MemberProperty{JsiiProperty: "isGroupRoleInput", GoGetter: "IsGroupRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "isLocked", GoGetter: "IsLocked"},
			_jsii_.MemberProperty{JsiiProperty: "isLockedInput", GoGetter: "IsLockedInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordFile", GoGetter: "PasswordFile"},
			_jsii_.MemberProperty{JsiiProperty: "passwordFileInput", GoGetter: "PasswordFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerType", GoGetter: "ProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "providerTypeInput", GoGetter: "ProviderTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeployedVmQuota", GoMethod: "ResetDeployedVmQuota"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailAddress", GoMethod: "ResetEmailAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFullName", GoMethod: "ResetFullName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstantMessaging", GoMethod: "ResetInstantMessaging"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsExternal", GoMethod: "ResetIsExternal"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsGroupRole", GoMethod: "ResetIsGroupRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsLocked", GoMethod: "ResetIsLocked"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordFile", GoMethod: "ResetPasswordFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderType", GoMethod: "ResetProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "resetStoredVmQuota", GoMethod: "ResetStoredVmQuota"},
			_jsii_.MemberMethod{JsiiMethod: "resetTakeOwnership", GoMethod: "ResetTakeOwnership"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephone", GoMethod: "ResetTelephone"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleInput", GoGetter: "RoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "storedVmQuota", GoGetter: "StoredVmQuota"},
			_jsii_.MemberProperty{JsiiProperty: "storedVmQuotaInput", GoGetter: "StoredVmQuotaInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "takeOwnership", GoGetter: "TakeOwnership"},
			_jsii_.MemberProperty{JsiiProperty: "takeOwnershipInput", GoGetter: "TakeOwnershipInput"},
			_jsii_.MemberProperty{JsiiProperty: "telephone", GoGetter: "Telephone"},
			_jsii_.MemberProperty{JsiiProperty: "telephoneInput", GoGetter: "TelephoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_OrgUser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.orgUser.OrgUserConfig",
		reflect.TypeOf((*OrgUserConfig)(nil)).Elem(),
	)
}