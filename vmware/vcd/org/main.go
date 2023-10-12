package org

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.org.Org",
		reflect.TypeOf((*Org)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "canPublishCatalogs", GoGetter: "CanPublishCatalogs"},
			_jsii_.MemberProperty{JsiiProperty: "canPublishCatalogsInput", GoGetter: "CanPublishCatalogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "canPublishExternalCatalogs", GoGetter: "CanPublishExternalCatalogs"},
			_jsii_.MemberProperty{JsiiProperty: "canPublishExternalCatalogsInput", GoGetter: "CanPublishExternalCatalogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "canSubscribeExternalCatalogs", GoGetter: "CanSubscribeExternalCatalogs"},
			_jsii_.MemberProperty{JsiiProperty: "canSubscribeExternalCatalogsInput", GoGetter: "CanSubscribeExternalCatalogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "delayAfterPowerOnSeconds", GoGetter: "DelayAfterPowerOnSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "delayAfterPowerOnSecondsInput", GoGetter: "DelayAfterPowerOnSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteForce", GoGetter: "DeleteForce"},
			_jsii_.MemberProperty{JsiiProperty: "deleteForceInput", GoGetter: "DeleteForceInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteRecursive", GoGetter: "DeleteRecursive"},
			_jsii_.MemberProperty{JsiiProperty: "deleteRecursiveInput", GoGetter: "DeleteRecursiveInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deployedVmQuota", GoGetter: "DeployedVmQuota"},
			_jsii_.MemberProperty{JsiiProperty: "deployedVmQuotaInput", GoGetter: "DeployedVmQuotaInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabled", GoGetter: "IsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabledInput", GoGetter: "IsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataEntry", GoGetter: "MetadataEntry"},
			_jsii_.MemberProperty{JsiiProperty: "metadataEntryInput", GoGetter: "MetadataEntryInput"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putMetadataEntry", GoMethod: "PutMetadataEntry"},
			_jsii_.MemberMethod{JsiiMethod: "putVappLease", GoMethod: "PutVappLease"},
			_jsii_.MemberMethod{JsiiMethod: "putVappTemplateLease", GoMethod: "PutVappTemplateLease"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanPublishCatalogs", GoMethod: "ResetCanPublishCatalogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanPublishExternalCatalogs", GoMethod: "ResetCanPublishExternalCatalogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanSubscribeExternalCatalogs", GoMethod: "ResetCanSubscribeExternalCatalogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelayAfterPowerOnSeconds", GoMethod: "ResetDelayAfterPowerOnSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeployedVmQuota", GoMethod: "ResetDeployedVmQuota"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsEnabled", GoMethod: "ResetIsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataEntry", GoMethod: "ResetMetadataEntry"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStoredVmQuota", GoMethod: "ResetStoredVmQuota"},
			_jsii_.MemberMethod{JsiiMethod: "resetVappLease", GoMethod: "ResetVappLease"},
			_jsii_.MemberMethod{JsiiMethod: "resetVappTemplateLease", GoMethod: "ResetVappTemplateLease"},
			_jsii_.MemberProperty{JsiiProperty: "storedVmQuota", GoGetter: "StoredVmQuota"},
			_jsii_.MemberProperty{JsiiProperty: "storedVmQuotaInput", GoGetter: "StoredVmQuotaInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vappLease", GoGetter: "VappLease"},
			_jsii_.MemberProperty{JsiiProperty: "vappLeaseInput", GoGetter: "VappLeaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "vappTemplateLease", GoGetter: "VappTemplateLease"},
			_jsii_.MemberProperty{JsiiProperty: "vappTemplateLeaseInput", GoGetter: "VappTemplateLeaseInput"},
		},
		func() interface{} {
			j := jsiiProxy_Org{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.org.OrgConfig",
		reflect.TypeOf((*OrgConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.org.OrgMetadataEntry",
		reflect.TypeOf((*OrgMetadataEntry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.org.OrgMetadataEntryList",
		reflect.TypeOf((*OrgMetadataEntryList)(nil)).Elem(),
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
			j := jsiiProxy_OrgMetadataEntryList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"vcd.org.OrgMetadataEntryOutputReference",
		reflect.TypeOf((*OrgMetadataEntryOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "isSystem", GoGetter: "IsSystem"},
			_jsii_.MemberProperty{JsiiProperty: "isSystemInput", GoGetter: "IsSystemInput"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsSystem", GoMethod: "ResetIsSystem"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserAccess", GoMethod: "ResetUserAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "userAccess", GoGetter: "UserAccess"},
			_jsii_.MemberProperty{JsiiProperty: "userAccessInput", GoGetter: "UserAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrgMetadataEntryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.org.OrgVappLease",
		reflect.TypeOf((*OrgVappLease)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.org.OrgVappLeaseOutputReference",
		reflect.TypeOf((*OrgVappLeaseOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnStorageLeaseExpiration", GoGetter: "DeleteOnStorageLeaseExpiration"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnStorageLeaseExpirationInput", GoGetter: "DeleteOnStorageLeaseExpirationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maximumRuntimeLeaseInSec", GoGetter: "MaximumRuntimeLeaseInSec"},
			_jsii_.MemberProperty{JsiiProperty: "maximumRuntimeLeaseInSecInput", GoGetter: "MaximumRuntimeLeaseInSecInput"},
			_jsii_.MemberProperty{JsiiProperty: "maximumStorageLeaseInSec", GoGetter: "MaximumStorageLeaseInSec"},
			_jsii_.MemberProperty{JsiiProperty: "maximumStorageLeaseInSecInput", GoGetter: "MaximumStorageLeaseInSecInput"},
			_jsii_.MemberProperty{JsiiProperty: "powerOffOnRuntimeLeaseExpiration", GoGetter: "PowerOffOnRuntimeLeaseExpiration"},
			_jsii_.MemberProperty{JsiiProperty: "powerOffOnRuntimeLeaseExpirationInput", GoGetter: "PowerOffOnRuntimeLeaseExpirationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrgVappLeaseOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.org.OrgVappTemplateLease",
		reflect.TypeOf((*OrgVappTemplateLease)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.org.OrgVappTemplateLeaseOutputReference",
		reflect.TypeOf((*OrgVappTemplateLeaseOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnStorageLeaseExpiration", GoGetter: "DeleteOnStorageLeaseExpiration"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnStorageLeaseExpirationInput", GoGetter: "DeleteOnStorageLeaseExpirationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maximumStorageLeaseInSec", GoGetter: "MaximumStorageLeaseInSec"},
			_jsii_.MemberProperty{JsiiProperty: "maximumStorageLeaseInSecInput", GoGetter: "MaximumStorageLeaseInSecInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OrgVappTemplateLeaseOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
