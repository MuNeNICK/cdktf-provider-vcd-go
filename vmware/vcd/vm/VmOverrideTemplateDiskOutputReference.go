package vm

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/vm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmOverrideTemplateDiskOutputReference interface {
	cdktf.ComplexObject
	BusNumber() *float64
	SetBusNumber(val *float64)
	BusNumberInput() *float64
	BusType() *string
	SetBusType(val *string)
	BusTypeInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	SizeInMb() *float64
	SetSizeInMb(val *float64)
	SizeInMbInput() *float64
	StorageProfile() *string
	SetStorageProfile(val *string)
	StorageProfileInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnitNumber() *float64
	SetUnitNumber(val *float64)
	UnitNumberInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIops()
	ResetStorageProfile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmOverrideTemplateDiskOutputReference
type jsiiProxy_VmOverrideTemplateDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) BusNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"busNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) BusNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"busNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) BusType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) BusTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) SizeInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) SizeInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) StorageProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) StorageProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) UnitNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference) UnitNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitNumberInput",
		&returns,
	)
	return returns
}


func NewVmOverrideTemplateDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VmOverrideTemplateDiskOutputReference {
	_init_.Initialize()

	if err := validateNewVmOverrideTemplateDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmOverrideTemplateDiskOutputReference{}

	_jsii_.Create(
		"vcd.vm.VmOverrideTemplateDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVmOverrideTemplateDiskOutputReference_Override(v VmOverrideTemplateDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vm.VmOverrideTemplateDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetBusNumber(val *float64) {
	if err := j.validateSetBusNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"busNumber",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetBusType(val *string) {
	if err := j.validateSetBusTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"busType",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetSizeInMb(val *float64) {
	if err := j.validateSetSizeInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInMb",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetStorageProfile(val *string) {
	if err := j.validateSetStorageProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfile",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VmOverrideTemplateDiskOutputReference)SetUnitNumber(val *float64) {
	if err := j.validateSetUnitNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitNumber",
		val,
	)
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		v,
		"resetIops",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) ResetStorageProfile() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageProfile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmOverrideTemplateDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

