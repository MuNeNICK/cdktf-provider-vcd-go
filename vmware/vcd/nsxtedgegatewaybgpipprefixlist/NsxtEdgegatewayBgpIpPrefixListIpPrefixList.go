package nsxtedgegatewaybgpipprefixlist

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtedgegatewaybgpipprefixlist/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayBgpIpPrefixListIpPrefixList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxtEdgegatewayBgpIpPrefixListIpPrefixList
type jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNsxtEdgegatewayBgpIpPrefixListIpPrefixList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NsxtEdgegatewayBgpIpPrefixListIpPrefixList {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewayBgpIpPrefixListIpPrefixListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList{}

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpIpPrefixList.NsxtEdgegatewayBgpIpPrefixListIpPrefixList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNsxtEdgegatewayBgpIpPrefixListIpPrefixList_Override(n NsxtEdgegatewayBgpIpPrefixListIpPrefixList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpIpPrefixList.NsxtEdgegatewayBgpIpPrefixListIpPrefixList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) Get(index *float64) NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NsxtEdgegatewayBgpIpPrefixListIpPrefixOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpIpPrefixListIpPrefixList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

