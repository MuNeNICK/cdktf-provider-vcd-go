package datavcdedgegateway

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdedgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList interface {
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
	Get(index *float64) DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList
type jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList {
	_init_.Initialize()

	if err := validateNewDataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList{}

	_jsii_.Create(
		"vcd.dataVcdEdgegateway.DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList_Override(d DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdEdgegateway.DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) Get(index *float64) DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegatewayExternalNetworkSubnetSuballocatePoolList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

