package vmsizingpolicy

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/vmsizingpolicy/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmSizingPolicyCpuOutputReference interface {
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
	CoresPerSocket() *string
	SetCoresPerSocket(val *string)
	CoresPerSocketInput() *string
	Count() *string
	SetCount(val *string)
	CountInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VmSizingPolicyCpu
	SetInternalValue(val *VmSizingPolicyCpu)
	LimitInMhz() *string
	SetLimitInMhz(val *string)
	LimitInMhzInput() *string
	ReservationGuarantee() *string
	SetReservationGuarantee(val *string)
	ReservationGuaranteeInput() *string
	Shares() *string
	SetShares(val *string)
	SharesInput() *string
	SpeedInMhz() *string
	SetSpeedInMhz(val *string)
	SpeedInMhzInput() *string
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
	ResetCoresPerSocket()
	ResetCount()
	ResetLimitInMhz()
	ResetReservationGuarantee()
	ResetShares()
	ResetSpeedInMhz()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmSizingPolicyCpuOutputReference
type jsiiProxy_VmSizingPolicyCpuOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) CoresPerSocket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coresPerSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) CoresPerSocketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coresPerSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) Count() *string {
	var returns *string
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) CountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) InternalValue() *VmSizingPolicyCpu {
	var returns *VmSizingPolicyCpu
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) LimitInMhz() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitInMhz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) LimitInMhzInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitInMhzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) ReservationGuarantee() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationGuarantee",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) ReservationGuaranteeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationGuaranteeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) Shares() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) SharesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) SpeedInMhz() *string {
	var returns *string
	_jsii_.Get(
		j,
		"speedInMhz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) SpeedInMhzInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"speedInMhzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmSizingPolicyCpuOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmSizingPolicyCpuOutputReference {
	_init_.Initialize()

	if err := validateNewVmSizingPolicyCpuOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmSizingPolicyCpuOutputReference{}

	_jsii_.Create(
		"vcd.vmSizingPolicy.VmSizingPolicyCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmSizingPolicyCpuOutputReference_Override(v VmSizingPolicyCpuOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vmSizingPolicy.VmSizingPolicyCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetCoresPerSocket(val *string) {
	if err := j.validateSetCoresPerSocketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coresPerSocket",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetCount(val *string) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetInternalValue(val *VmSizingPolicyCpu) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetLimitInMhz(val *string) {
	if err := j.validateSetLimitInMhzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitInMhz",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetReservationGuarantee(val *string) {
	if err := j.validateSetReservationGuaranteeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationGuarantee",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetShares(val *string) {
	if err := j.validateSetSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shares",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetSpeedInMhz(val *string) {
	if err := j.validateSetSpeedInMhzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"speedInMhz",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmSizingPolicyCpuOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetCoresPerSocket() {
	_jsii_.InvokeVoid(
		v,
		"resetCoresPerSocket",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		v,
		"resetCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetLimitInMhz() {
	_jsii_.InvokeVoid(
		v,
		"resetLimitInMhz",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetReservationGuarantee() {
	_jsii_.InvokeVoid(
		v,
		"resetReservationGuarantee",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetShares() {
	_jsii_.InvokeVoid(
		v,
		"resetShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ResetSpeedInMhz() {
	_jsii_.InvokeVoid(
		v,
		"resetSpeedInMhz",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VmSizingPolicyCpuOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

