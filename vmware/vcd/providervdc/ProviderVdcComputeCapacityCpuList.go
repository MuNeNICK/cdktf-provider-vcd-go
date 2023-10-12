package providervdc

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/providervdc/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProviderVdcComputeCapacityCpuList interface {
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
	Get(index *float64) ProviderVdcComputeCapacityCpuOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProviderVdcComputeCapacityCpuList
type jsiiProxy_ProviderVdcComputeCapacityCpuList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewProviderVdcComputeCapacityCpuList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ProviderVdcComputeCapacityCpuList {
	_init_.Initialize()

	if err := validateNewProviderVdcComputeCapacityCpuListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProviderVdcComputeCapacityCpuList{}

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdcComputeCapacityCpuList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewProviderVdcComputeCapacityCpuList_Override(p ProviderVdcComputeCapacityCpuList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdcComputeCapacityCpuList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ProviderVdcComputeCapacityCpuList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) Get(index *float64) ProviderVdcComputeCapacityCpuOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ProviderVdcComputeCapacityCpuOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_ProviderVdcComputeCapacityCpuList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

