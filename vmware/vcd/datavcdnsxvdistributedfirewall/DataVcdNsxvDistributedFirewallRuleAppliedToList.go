package datavcdnsxvdistributedfirewall

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdnsxvdistributedfirewall/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxvDistributedFirewallRuleAppliedToList interface {
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
	Get(index *float64) DataVcdNsxvDistributedFirewallRuleAppliedToOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataVcdNsxvDistributedFirewallRuleAppliedToList
type jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataVcdNsxvDistributedFirewallRuleAppliedToList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataVcdNsxvDistributedFirewallRuleAppliedToList {
	_init_.Initialize()

	if err := validateNewDataVcdNsxvDistributedFirewallRuleAppliedToListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList{}

	_jsii_.Create(
		"vcd.dataVcdNsxvDistributedFirewall.DataVcdNsxvDistributedFirewallRuleAppliedToList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataVcdNsxvDistributedFirewallRuleAppliedToList_Override(d DataVcdNsxvDistributedFirewallRuleAppliedToList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdNsxvDistributedFirewall.DataVcdNsxvDistributedFirewallRuleAppliedToList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) Get(index *float64) DataVcdNsxvDistributedFirewallRuleAppliedToOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataVcdNsxvDistributedFirewallRuleAppliedToOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataVcdNsxvDistributedFirewallRuleAppliedToList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

