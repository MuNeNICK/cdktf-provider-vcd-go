package edgegateway

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/edgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgegatewayExternalNetworkSubnetOutputReference interface {
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
	Gateway() *string
	SetGateway(val *string)
	GatewayInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Netmask() *string
	SetNetmask(val *string)
	NetmaskInput() *string
	SuballocatePool() EdgegatewayExternalNetworkSubnetSuballocatePoolList
	SuballocatePoolInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseForDefaultRoute() interface{}
	SetUseForDefaultRoute(val interface{})
	UseForDefaultRouteInput() interface{}
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
	PutSuballocatePool(value interface{})
	ResetIpAddress()
	ResetSuballocatePool()
	ResetUseForDefaultRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgegatewayExternalNetworkSubnetOutputReference
type jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) Netmask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) NetmaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) SuballocatePool() EdgegatewayExternalNetworkSubnetSuballocatePoolList {
	var returns EdgegatewayExternalNetworkSubnetSuballocatePoolList
	_jsii_.Get(
		j,
		"suballocatePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) SuballocatePoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suballocatePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) UseForDefaultRoute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useForDefaultRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) UseForDefaultRouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useForDefaultRouteInput",
		&returns,
	)
	return returns
}


func NewEdgegatewayExternalNetworkSubnetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EdgegatewayExternalNetworkSubnetOutputReference {
	_init_.Initialize()

	if err := validateNewEdgegatewayExternalNetworkSubnetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference{}

	_jsii_.Create(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEdgegatewayExternalNetworkSubnetOutputReference_Override(e EdgegatewayExternalNetworkSubnetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.edgegateway.EdgegatewayExternalNetworkSubnetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetNetmask(val *string) {
	if err := j.validateSetNetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netmask",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference)SetUseForDefaultRoute(val interface{}) {
	if err := j.validateSetUseForDefaultRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useForDefaultRoute",
		val,
	)
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) PutSuballocatePool(value interface{}) {
	if err := e.validatePutSuballocatePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSuballocatePool",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ResetSuballocatePool() {
	_jsii_.InvokeVoid(
		e,
		"resetSuballocatePool",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ResetUseForDefaultRoute() {
	_jsii_.InvokeVoid(
		e,
		"resetUseForDefaultRoute",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EdgegatewayExternalNetworkSubnetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

