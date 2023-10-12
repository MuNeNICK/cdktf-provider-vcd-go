package providervdc

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/providervdc/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProviderVdcComputeCapacityMemoryList interface {
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
	Get(index *float64) ProviderVdcComputeCapacityMemoryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProviderVdcComputeCapacityMemoryList
type jsiiProxy_ProviderVdcComputeCapacityMemoryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewProviderVdcComputeCapacityMemoryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ProviderVdcComputeCapacityMemoryList {
	_init_.Initialize()

	if err := validateNewProviderVdcComputeCapacityMemoryListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProviderVdcComputeCapacityMemoryList{}

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdcComputeCapacityMemoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewProviderVdcComputeCapacityMemoryList_Override(p ProviderVdcComputeCapacityMemoryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdcComputeCapacityMemoryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityMemoryList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) Get(index *float64) ProviderVdcComputeCapacityMemoryOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ProviderVdcComputeCapacityMemoryOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdcComputeCapacityMemoryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

