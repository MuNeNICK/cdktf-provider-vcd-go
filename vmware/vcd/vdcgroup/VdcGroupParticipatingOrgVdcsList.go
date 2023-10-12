package vdcgroup

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vdcgroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VdcGroupParticipatingOrgVdcsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) VdcGroupParticipatingOrgVdcsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VdcGroupParticipatingOrgVdcsList
type jsiiProxy_VdcGroupParticipatingOrgVdcsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVdcGroupParticipatingOrgVdcsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VdcGroupParticipatingOrgVdcsList {
	_init_.Initialize()

	if err := validateNewVdcGroupParticipatingOrgVdcsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VdcGroupParticipatingOrgVdcsList{}

	_jsii_.Create(
		"vcd.vdcGroup.VdcGroupParticipatingOrgVdcsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVdcGroupParticipatingOrgVdcsList_Override(v VdcGroupParticipatingOrgVdcsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vdcGroup.VdcGroupParticipatingOrgVdcsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VdcGroupParticipatingOrgVdcsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VdcGroupParticipatingOrgVdcsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroupParticipatingOrgVdcsList) Get(index *float64) VdcGroupParticipatingOrgVdcsOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VdcGroupParticipatingOrgVdcsOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroupParticipatingOrgVdcsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VdcGroupParticipatingOrgVdcsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

