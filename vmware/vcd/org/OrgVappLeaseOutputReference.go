package org

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/org/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgVappLeaseOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteOnStorageLeaseExpiration() interface{}
	SetDeleteOnStorageLeaseExpiration(val interface{})
	DeleteOnStorageLeaseExpirationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OrgVappLease
	SetInternalValue(val *OrgVappLease)
	MaximumRuntimeLeaseInSec() *float64
	SetMaximumRuntimeLeaseInSec(val *float64)
	MaximumRuntimeLeaseInSecInput() *float64
	MaximumStorageLeaseInSec() *float64
	SetMaximumStorageLeaseInSec(val *float64)
	MaximumStorageLeaseInSecInput() *float64
	PowerOffOnRuntimeLeaseExpiration() interface{}
	SetPowerOffOnRuntimeLeaseExpiration(val interface{})
	PowerOffOnRuntimeLeaseExpirationInput() interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrgVappLeaseOutputReference
type jsiiProxy_OrgVappLeaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) DeleteOnStorageLeaseExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnStorageLeaseExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) DeleteOnStorageLeaseExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnStorageLeaseExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) InternalValue() *OrgVappLease {
	var returns *OrgVappLease
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) MaximumRuntimeLeaseInSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRuntimeLeaseInSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) MaximumRuntimeLeaseInSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRuntimeLeaseInSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) MaximumStorageLeaseInSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumStorageLeaseInSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) MaximumStorageLeaseInSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumStorageLeaseInSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) PowerOffOnRuntimeLeaseExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOffOnRuntimeLeaseExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) PowerOffOnRuntimeLeaseExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOffOnRuntimeLeaseExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVappLeaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrgVappLeaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrgVappLeaseOutputReference {
	_init_.Initialize()

	if err := validateNewOrgVappLeaseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgVappLeaseOutputReference{}

	_jsii_.Create(
		"vcd.org.OrgVappLeaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrgVappLeaseOutputReference_Override(o OrgVappLeaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.org.OrgVappLeaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetDeleteOnStorageLeaseExpiration(val interface{}) {
	if err := j.validateSetDeleteOnStorageLeaseExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnStorageLeaseExpiration",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetInternalValue(val *OrgVappLease) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetMaximumRuntimeLeaseInSec(val *float64) {
	if err := j.validateSetMaximumRuntimeLeaseInSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRuntimeLeaseInSec",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetMaximumStorageLeaseInSec(val *float64) {
	if err := j.validateSetMaximumStorageLeaseInSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumStorageLeaseInSec",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetPowerOffOnRuntimeLeaseExpiration(val interface{}) {
	if err := j.validateSetPowerOffOnRuntimeLeaseExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerOffOnRuntimeLeaseExpiration",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgVappLeaseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrgVappLeaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVappLeaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrgVappLeaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

