package vappnatrules

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vappnatrules/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappNatRulesRuleList interface {
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
	Get(index *float64) VappNatRulesRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VappNatRulesRuleList
type jsiiProxy_VappNatRulesRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VappNatRulesRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVappNatRulesRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VappNatRulesRuleList {
	_init_.Initialize()

	if err := validateNewVappNatRulesRuleListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappNatRulesRuleList{}

	_jsii_.Create(
		"vcd.vappNatRules.VappNatRulesRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVappNatRulesRuleList_Override(v VappNatRulesRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappNatRules.VappNatRulesRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VappNatRulesRuleList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VappNatRulesRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleList) Get(index *float64) VappNatRulesRuleOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VappNatRulesRuleOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VappNatRulesRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

