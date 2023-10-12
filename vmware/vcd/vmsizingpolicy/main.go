package vmsizingpolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vcd.vmSizingPolicy.VmSizingPolicy",
		reflect.TypeOf((*VmSizingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "cpuInput", GoGetter: "CpuInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "memoryInput", GoGetter: "MemoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "org", GoGetter: "Org"},
			_jsii_.MemberProperty{JsiiProperty: "orgInput", GoGetter: "OrgInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCpu", GoMethod: "PutCpu"},
			_jsii_.MemberMethod{JsiiMethod: "putMemory", GoMethod: "PutMemory"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpu", GoMethod: "ResetCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemory", GoMethod: "ResetMemory"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrg", GoMethod: "ResetOrg"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_VmSizingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.vmSizingPolicy.VmSizingPolicyConfig",
		reflect.TypeOf((*VmSizingPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vcd.vmSizingPolicy.VmSizingPolicyCpu",
		reflect.TypeOf((*VmSizingPolicyCpu)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.vmSizingPolicy.VmSizingPolicyCpuOutputReference",
		reflect.TypeOf((*VmSizingPolicyCpuOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "coresPerSocket", GoGetter: "CoresPerSocket"},
			_jsii_.MemberProperty{JsiiProperty: "coresPerSocketInput", GoGetter: "CoresPerSocketInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "countInput", GoGetter: "CountInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "limitInMhz", GoGetter: "LimitInMhz"},
			_jsii_.MemberProperty{JsiiProperty: "limitInMhzInput", GoGetter: "LimitInMhzInput"},
			_jsii_.MemberProperty{JsiiProperty: "reservationGuarantee", GoGetter: "ReservationGuarantee"},
			_jsii_.MemberProperty{JsiiProperty: "reservationGuaranteeInput", GoGetter: "ReservationGuaranteeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoresPerSocket", GoMethod: "ResetCoresPerSocket"},
			_jsii_.MemberMethod{JsiiMethod: "resetCount", GoMethod: "ResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimitInMhz", GoMethod: "ResetLimitInMhz"},
			_jsii_.MemberMethod{JsiiMethod: "resetReservationGuarantee", GoMethod: "ResetReservationGuarantee"},
			_jsii_.MemberMethod{JsiiMethod: "resetShares", GoMethod: "ResetShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpeedInMhz", GoMethod: "ResetSpeedInMhz"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shares", GoGetter: "Shares"},
			_jsii_.MemberProperty{JsiiProperty: "sharesInput", GoGetter: "SharesInput"},
			_jsii_.MemberProperty{JsiiProperty: "speedInMhz", GoGetter: "SpeedInMhz"},
			_jsii_.MemberProperty{JsiiProperty: "speedInMhzInput", GoGetter: "SpeedInMhzInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VmSizingPolicyCpuOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vcd.vmSizingPolicy.VmSizingPolicyMemory",
		reflect.TypeOf((*VmSizingPolicyMemory)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vcd.vmSizingPolicy.VmSizingPolicyMemoryOutputReference",
		reflect.TypeOf((*VmSizingPolicyMemoryOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "limitInMb", GoGetter: "LimitInMb"},
			_jsii_.MemberProperty{JsiiProperty: "limitInMbInput", GoGetter: "LimitInMbInput"},
			_jsii_.MemberProperty{JsiiProperty: "reservationGuarantee", GoGetter: "ReservationGuarantee"},
			_jsii_.MemberProperty{JsiiProperty: "reservationGuaranteeInput", GoGetter: "ReservationGuaranteeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimitInMb", GoMethod: "ResetLimitInMb"},
			_jsii_.MemberMethod{JsiiMethod: "resetReservationGuarantee", GoMethod: "ResetReservationGuarantee"},
			_jsii_.MemberMethod{JsiiMethod: "resetShares", GoMethod: "ResetShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeInMb", GoMethod: "ResetSizeInMb"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shares", GoGetter: "Shares"},
			_jsii_.MemberProperty{JsiiProperty: "sharesInput", GoGetter: "SharesInput"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMb", GoGetter: "SizeInMb"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInMbInput", GoGetter: "SizeInMbInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VmSizingPolicyMemoryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
