package vappvm

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/vappvm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappVmNetworkOutputReference interface {
	cdktf.ComplexObject
	AdapterType() *string
	SetAdapterType(val *string)
	AdapterTypeInput() *string
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
	Connected() interface{}
	SetConnected(val interface{})
	ConnectedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() *string
	SetIp(val *string)
	IpAllocationMode() *string
	SetIpAllocationMode(val *string)
	IpAllocationModeInput() *string
	IpInput() *string
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	Mac() *string
	SetMac(val *string)
	MacInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAdapterType()
	ResetConnected()
	ResetIp()
	ResetIpAllocationMode()
	ResetIsPrimary()
	ResetMac()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VappVmNetworkOutputReference
type jsiiProxy_VappVmNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VappVmNetworkOutputReference) AdapterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adapterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) AdapterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adapterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Connected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) ConnectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) IpAllocationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) IpAllocationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Mac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) MacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmNetworkOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewVappVmNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VappVmNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewVappVmNetworkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappVmNetworkOutputReference{}

	_jsii_.Create(
		"vcd.vappVm.VappVmNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVappVmNetworkOutputReference_Override(v VappVmNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappVm.VappVmNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetAdapterType(val *string) {
	if err := j.validateSetAdapterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adapterType",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetConnected(val interface{}) {
	if err := j.validateSetConnectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connected",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetIp(val *string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetIpAllocationMode(val *string) {
	if err := j.validateSetIpAllocationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAllocationMode",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetIsPrimary(val interface{}) {
	if err := j.validateSetIsPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetMac(val *string) {
	if err := j.validateSetMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mac",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VappVmNetworkOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetAdapterType() {
	_jsii_.InvokeVoid(
		v,
		"resetAdapterType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetConnected() {
	_jsii_.InvokeVoid(
		v,
		"resetConnected",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		v,
		"resetIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetIpAllocationMode() {
	_jsii_.InvokeVoid(
		v,
		"resetIpAllocationMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetIsPrimary() {
	_jsii_.InvokeVoid(
		v,
		"resetIsPrimary",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetMac() {
	_jsii_.InvokeVoid(
		v,
		"resetMac",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VappVmNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

