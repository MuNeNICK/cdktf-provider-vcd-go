package orgvdc

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/orgvdc/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgVdcComputeCapacityOutputReference interface {
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
	Cpu() OrgVdcComputeCapacityCpuOutputReference
	CpuInput() *OrgVdcComputeCapacityCpu
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *OrgVdcComputeCapacity
	SetInternalValue(val *OrgVdcComputeCapacity)
	Memory() OrgVdcComputeCapacityMemoryOutputReference
	MemoryInput() *OrgVdcComputeCapacityMemory
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
	PutCpu(value *OrgVdcComputeCapacityCpu)
	PutMemory(value *OrgVdcComputeCapacityMemory)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrgVdcComputeCapacityOutputReference
type jsiiProxy_OrgVdcComputeCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) Cpu() OrgVdcComputeCapacityCpuOutputReference {
	var returns OrgVdcComputeCapacityCpuOutputReference
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) CpuInput() *OrgVdcComputeCapacityCpu {
	var returns *OrgVdcComputeCapacityCpu
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) InternalValue() *OrgVdcComputeCapacity {
	var returns *OrgVdcComputeCapacity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) Memory() OrgVdcComputeCapacityMemoryOutputReference {
	var returns OrgVdcComputeCapacityMemoryOutputReference
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) MemoryInput() *OrgVdcComputeCapacityMemory {
	var returns *OrgVdcComputeCapacityMemory
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrgVdcComputeCapacityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrgVdcComputeCapacityOutputReference {
	_init_.Initialize()

	if err := validateNewOrgVdcComputeCapacityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgVdcComputeCapacityOutputReference{}

	_jsii_.Create(
		"vcd.orgVdc.OrgVdcComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrgVdcComputeCapacityOutputReference_Override(o OrgVdcComputeCapacityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgVdc.OrgVdcComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference)SetInternalValue(val *OrgVdcComputeCapacity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgVdcComputeCapacityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) PutCpu(value *OrgVdcComputeCapacityCpu) {
	if err := o.validatePutCpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCpu",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) PutMemory(value *OrgVdcComputeCapacityMemory) {
	if err := o.validatePutMemoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMemory",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdcComputeCapacityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

