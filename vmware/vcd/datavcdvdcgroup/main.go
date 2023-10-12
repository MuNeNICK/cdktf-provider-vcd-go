package datavcdvdcgroup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		reflect.TypeOf((*DataVcdVdcGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicyStatus", GoGetter: "DefaultPolicyStatus"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicyStatusInput", GoGetter: "DefaultPolicyStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dfwEnabled", GoGetter: "DfwEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "errorMessage", GoGetter: "ErrorMessage"},
			_jsii_.MemberProperty{JsiiProperty: "errorMessageInput", GoGetter: "ErrorMessageInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "localEgress", GoGetter: "LocalEgress"},
			_jsii_.MemberProperty{JsiiProperty: "localEgressInput", GoGetter: "LocalEgressInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkPoolId", GoGetter: "NetworkPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "networkPoolIdInput", GoGetter: "NetworkPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkPoolUniversalId", GoGetter: "NetworkPoolUniversalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkPoolUniversalIdInput", GoGetter: "NetworkPoolUniversalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkProviderType", GoGetter: "NetworkProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "networkProviderTypeInput", GoGetter: "NetworkProviderTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "participatingOrgVdcs", GoGetter: "ParticipatingOrgVdcs"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultPolicyStatus", GoMethod: "ResetDefaultPolicyStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorMessage", GoMethod: "ResetErrorMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalEgress", GoMethod: "ResetLocalEgress"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkPoolId", GoMethod: "ResetNetworkPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkPoolUniversalId", GoMethod: "ResetNetworkPoolUniversalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkProviderType", GoMethod: "ResetNetworkProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUniversalNetworkingEnabled", GoMethod: "ResetUniversalNetworkingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "universalNetworkingEnabled", GoGetter: "UniversalNetworkingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "universalNetworkingEnabledInput", GoGetter: "UniversalNetworkingEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataVcdVdcGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroupConfig",
		reflect.TypeOf((*DataVcdVdcGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroupParticipatingOrgVdcs",
		reflect.TypeOf((*DataVcdVdcGroupParticipatingOrgVdcs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroupParticipatingOrgVdcsList",
		reflect.TypeOf((*DataVcdVdcGroupParticipatingOrgVdcsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataVcdVdcGroupParticipatingOrgVdcsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroupParticipatingOrgVdcsOutputReference",
		reflect.TypeOf((*DataVcdVdcGroupParticipatingOrgVdcsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "faultDomainTag", GoGetter: "FaultDomainTag"},
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
			_jsii_.MemberProperty{JsiiProperty: "isRemoteOrg", GoGetter: "IsRemoteOrg"},
			_jsii_.MemberProperty{JsiiProperty: "networkProviderScope", GoGetter: "NetworkProviderScope"},
			_jsii_.MemberProperty{JsiiProperty: "orgId", GoGetter: "OrgId"},
			_jsii_.MemberProperty{JsiiProperty: "orgName", GoGetter: "OrgName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "siteId", GoGetter: "SiteId"},
			_jsii_.MemberProperty{JsiiProperty: "siteName", GoGetter: "SiteName"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vdcId", GoGetter: "VdcId"},
			_jsii_.MemberProperty{JsiiProperty: "vdcName", GoGetter: "VdcName"},
		},
		func() interface{} {
			j := jsiiProxy_DataVcdVdcGroupParticipatingOrgVdcsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
