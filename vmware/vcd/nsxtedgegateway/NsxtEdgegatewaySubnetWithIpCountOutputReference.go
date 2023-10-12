package nsxtedgegateway

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtedgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewaySubnetWithIpCountOutputReference interface {
	cdktf.ComplexObject
	AllocatedIpCount() *float64
	SetAllocatedIpCount(val *float64)
	AllocatedIpCountInput() *float64
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
	PrefixLength() *float64
	SetPrefixLength(val *float64)
	PrefixLengthInput() *float64
	PrimaryIp() *string
	SetPrimaryIp(val *string)
	PrimaryIpInput() *string
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
	ResetPrimaryIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxtEdgegatewaySubnetWithIpCountOutputReference
type jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) AllocatedIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) AllocatedIpCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) PrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) PrimaryIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNsxtEdgegatewaySubnetWithIpCountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NsxtEdgegatewaySubnetWithIpCountOutputReference {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewaySubnetWithIpCountOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference{}

	_jsii_.Create(
		"vcd.nsxtEdgegateway.NsxtEdgegatewaySubnetWithIpCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNsxtEdgegatewaySubnetWithIpCountOutputReference_Override(n NsxtEdgegatewaySubnetWithIpCountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegateway.NsxtEdgegatewaySubnetWithIpCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetAllocatedIpCount(val *float64) {
	if err := j.validateSetAllocatedIpCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedIpCount",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetPrefixLength(val *float64) {
	if err := j.validateSetPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixLength",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetPrimaryIp(val *string) {
	if err := j.validateSetPrimaryIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryIp",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) ResetPrimaryIp() {
	_jsii_.InvokeVoid(
		n,
		"resetPrimaryIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewaySubnetWithIpCountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

