package vmsizingpolicy

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/vmsizingpolicy/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmSizingPolicyMemoryOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *VmSizingPolicyMemory
	SetInternalValue(val *VmSizingPolicyMemory)
	LimitInMb() *string
	SetLimitInMb(val *string)
	LimitInMbInput() *string
	ReservationGuarantee() *string
	SetReservationGuarantee(val *string)
	ReservationGuaranteeInput() *string
	Shares() *string
	SetShares(val *string)
	SharesInput() *string
	SizeInMb() *string
	SetSizeInMb(val *string)
	SizeInMbInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetLimitInMb()
	ResetReservationGuarantee()
	ResetShares()
	ResetSizeInMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmSizingPolicyMemoryOutputReference
type jsiiProxy_VmSizingPolicyMemoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) InternalValue() *VmSizingPolicyMemory {
	var returns *VmSizingPolicyMemory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) LimitInMb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) LimitInMbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) ReservationGuarantee() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationGuarantee",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) ReservationGuaranteeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationGuaranteeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) Shares() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) SharesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) SizeInMb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) SizeInMbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmSizingPolicyMemoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmSizingPolicyMemoryOutputReference {
	_init_.Initialize()

	if err := validateNewVmSizingPolicyMemoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmSizingPolicyMemoryOutputReference{}

	_jsii_.Create(
		"vcd.vmSizingPolicy.VmSizingPolicyMemoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmSizingPolicyMemoryOutputReference_Override(v VmSizingPolicyMemoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vmSizingPolicy.VmSizingPolicyMemoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetInternalValue(val *VmSizingPolicyMemory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetLimitInMb(val *string) {
	if err := j.validateSetLimitInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitInMb",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetReservationGuarantee(val *string) {
	if err := j.validateSetReservationGuaranteeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationGuarantee",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetShares(val *string) {
	if err := j.validateSetSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shares",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetSizeInMb(val *string) {
	if err := j.validateSetSizeInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInMb",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyMemoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ResetLimitInMb() {
	_jsii_.InvokeVoid(
		v,
		"resetLimitInMb",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ResetReservationGuarantee() {
	_jsii_.InvokeVoid(
		v,
		"resetReservationGuarantee",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ResetShares() {
	_jsii_.InvokeVoid(
		v,
		"resetShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ResetSizeInMb() {
	_jsii_.InvokeVoid(
		v,
		"resetSizeInMb",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VmSizingPolicyMemoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

