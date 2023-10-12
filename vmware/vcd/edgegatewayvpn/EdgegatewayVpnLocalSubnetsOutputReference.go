package edgegatewayvpn

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/edgegatewayvpn/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgegatewayVpnLocalSubnetsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalSubnetGateway() *string
	SetLocalSubnetGateway(val *string)
	LocalSubnetGatewayInput() *string
	LocalSubnetMask() *string
	SetLocalSubnetMask(val *string)
	LocalSubnetMaskInput() *string
	LocalSubnetName() *string
	SetLocalSubnetName(val *string)
	LocalSubnetNameInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgegatewayVpnLocalSubnetsOutputReference
type jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetMaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetMaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) LocalSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEdgegatewayVpnLocalSubnetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EdgegatewayVpnLocalSubnetsOutputReference {
	_init_.Initialize()

	if err := validateNewEdgegatewayVpnLocalSubnetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference{}

	_jsii_.Create(
		"vcd.edgegatewayVpn.EdgegatewayVpnLocalSubnetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEdgegatewayVpnLocalSubnetsOutputReference_Override(e EdgegatewayVpnLocalSubnetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.edgegatewayVpn.EdgegatewayVpnLocalSubnetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetLocalSubnetGateway(val *string) {
	if err := j.validateSetLocalSubnetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSubnetGateway",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetLocalSubnetMask(val *string) {
	if err := j.validateSetLocalSubnetMaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSubnetMask",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetLocalSubnetName(val *string) {
	if err := j.validateSetLocalSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSubnetName",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EdgegatewayVpnLocalSubnetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

