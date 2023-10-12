package externalnetworkv2

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/externalnetworkv2/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalNetworkV2NsxtNetworkOutputReference interface {
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
	InternalValue() *ExternalNetworkV2NsxtNetwork
	SetInternalValue(val *ExternalNetworkV2NsxtNetwork)
	NsxtManagerId() *string
	SetNsxtManagerId(val *string)
	NsxtManagerIdInput() *string
	NsxtSegmentName() *string
	SetNsxtSegmentName(val *string)
	NsxtSegmentNameInput() *string
	NsxtTier0RouterId() *string
	SetNsxtTier0RouterId(val *string)
	NsxtTier0RouterIdInput() *string
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
	ResetNsxtSegmentName()
	ResetNsxtTier0RouterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExternalNetworkV2NsxtNetworkOutputReference
type jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) InternalValue() *ExternalNetworkV2NsxtNetwork {
	var returns *ExternalNetworkV2NsxtNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtManagerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtManagerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtManagerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtSegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtSegmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtSegmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtSegmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtTier0RouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtTier0RouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) NsxtTier0RouterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtTier0RouterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExternalNetworkV2NsxtNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExternalNetworkV2NsxtNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewExternalNetworkV2NsxtNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference{}

	_jsii_.Create(
		"vcd.externalNetworkV2.ExternalNetworkV2NsxtNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExternalNetworkV2NsxtNetworkOutputReference_Override(e ExternalNetworkV2NsxtNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.externalNetworkV2.ExternalNetworkV2NsxtNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetInternalValue(val *ExternalNetworkV2NsxtNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetNsxtManagerId(val *string) {
	if err := j.validateSetNsxtManagerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtManagerId",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetNsxtSegmentName(val *string) {
	if err := j.validateSetNsxtSegmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtSegmentName",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetNsxtTier0RouterId(val *string) {
	if err := j.validateSetNsxtTier0RouterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtTier0RouterId",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ResetNsxtSegmentName() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxtSegmentName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ResetNsxtTier0RouterId() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxtTier0RouterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalNetworkV2NsxtNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

