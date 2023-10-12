package vminternaldisk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.vmInternalDisk.VmInternalDiskA",
		reflect.TypeOf((*VmInternalDiskA)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowVmReboot", GoGetter: "AllowVmReboot"},
			_jsii_.MemberProperty{JsiiProperty: "allowVmRebootInput", GoGetter: "AllowVmRebootInput"},
			_jsii_.MemberProperty{JsiiProperty: "busNumber", GoGetter: "BusNumber"},
			_jsii_.MemberProperty{JsiiProperty: "busNumberInput", GoGetter: "BusNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "busType", GoGetter: "BusType"},
			_jsii_.MemberProperty{JsiiProperty: "busTypeInput", GoGetter: "BusTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "iopsInput", GoGetter: "IopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowVmReboot", GoMethod: "ResetAllowVmReboot"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIops", GoMethod: "ResetIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageProfile", GoMethod: "ResetStorageProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetVdc", GoMethod: "ResetVdc"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMb", GoGetter: "SizeInMb"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMbInput", GoGetter: "SizeInMbInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageProfile", GoGetter: "StorageProfile"},
			_jsii_.MemberProperty{JsiiProperty: "storageProfileInput", GoGetter: "StorageProfileInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "thinProvisioned", GoGetter: "ThinProvisioned"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "unitNumber", GoGetter: "UnitNumber"},
			_jsii_.MemberProperty{JsiiProperty: "unitNumberInput", GoGetter: "UnitNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "vappName", GoGetter: "VappName"},
			_jsii_.MemberProperty{JsiiProperty: "vappNameInput", GoGetter: "VappNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vdc", GoGetter: "Vdc"},
			_jsii_.MemberProperty{JsiiProperty: "vdcInput", GoGetter: "VdcInput"},
			_jsii_.MemberProperty{JsiiProperty: "vmName", GoGetter: "VmName"},
			_jsii_.MemberProperty{JsiiProperty: "vmNameInput", GoGetter: "VmNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_VmInternalDiskA{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.vmInternalDisk.VmInternalDiskAConfig",
		reflect.TypeOf((*VmInternalDiskAConfig)(nil)).Elem(),
	)
}
