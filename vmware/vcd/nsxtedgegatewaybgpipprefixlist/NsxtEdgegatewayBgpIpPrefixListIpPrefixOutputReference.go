package nsxtedgegatewaybgpipprefixlist

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtedgegatewaybgpipprefixlist/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	GreaterThanOrEqualTo() *float64
	SetGreaterThanOrEqualTo(val *float64)
	GreaterThanOrEqualToInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LessThanOrEqualTo() *float64
	SetLessThanOrEqualTo(val *float64)
	LessThanOrEqualToInput() *float64
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
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
	ResetGreaterThanOrEqualTo()
	ResetLessThanOrEqualTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference
type jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GreaterThanOrEqualTo() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanOrEqualTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GreaterThanOrEqualToInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanOrEqualToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) LessThanOrEqualTo() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanOrEqualTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) LessThanOrEqualToInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanOrEqualToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference{}

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpIpPrefixList.NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference_Override(n NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpIpPrefixList.NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetGreaterThanOrEqualTo(val *float64) {
	if err := j.validateSetGreaterThanOrEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThanOrEqualTo",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetLessThanOrEqualTo(val *float64) {
	if err := j.validateSetLessThanOrEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThanOrEqualTo",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ResetGreaterThanOrEqualTo() {
	_jsii_.InvokeVoid(
		n,
		"resetGreaterThanOrEqualTo",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ResetLessThanOrEqualTo() {
	_jsii_.InvokeVoid(
		n,
		"resetLessThanOrEqualTo",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

