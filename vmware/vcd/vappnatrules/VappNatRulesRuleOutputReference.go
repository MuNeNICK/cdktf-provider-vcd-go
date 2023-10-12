package vappnatrules

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vappnatrules/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappNatRulesRuleOutputReference interface {
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
	ExternalIp() *string
	SetExternalIp(val *string)
	ExternalIpInput() *string
	ExternalPort() *float64
	SetExternalPort(val *float64)
	ExternalPortInput() *float64
	ForwardToPort() *float64
	SetForwardToPort(val *float64)
	ForwardToPortInput() *float64
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MappingMode() *string
	SetMappingMode(val *string)
	MappingModeInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmId() *string
	SetVmId(val *string)
	VmIdInput() *string
	VmNicId() *float64
	SetVmNicId(val *float64)
	VmNicIdInput() *float64
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
	ResetExternalIp()
	ResetExternalPort()
	ResetForwardToPort()
	ResetMappingMode()
	ResetProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VappNatRulesRuleOutputReference
type jsiiProxy_VappNatRulesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ExternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ExternalIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ExternalPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ExternalPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ForwardToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forwardToPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ForwardToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forwardToPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) MappingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) MappingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) VmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) VmIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) VmNicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmNicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference) VmNicIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmNicIdInput",
		&returns,
	)
	return returns
}


func NewVappNatRulesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VappNatRulesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewVappNatRulesRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappNatRulesRuleOutputReference{}

	_jsii_.Create(
		"vcd.vappNatRules.VappNatRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVappNatRulesRuleOutputReference_Override(v VappNatRulesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappNatRules.VappNatRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetExternalIp(val *string) {
	if err := j.validateSetExternalIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIp",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetExternalPort(val *float64) {
	if err := j.validateSetExternalPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalPort",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetForwardToPort(val *float64) {
	if err := j.validateSetForwardToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardToPort",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetMappingMode(val *string) {
	if err := j.validateSetMappingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingMode",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetVmId(val *string) {
	if err := j.validateSetVmIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmId",
		val,
	)
}

func (j *jsiiProxy_VappNatRulesRuleOutputReference)SetVmNicId(val *float64) {
	if err := j.validateSetVmNicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmNicId",
		val,
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ResetExternalIp() {
	_jsii_.InvokeVoid(
		v,
		"resetExternalIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ResetExternalPort() {
	_jsii_.InvokeVoid(
		v,
		"resetExternalPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ResetForwardToPort() {
	_jsii_.InvokeVoid(
		v,
		"resetForwardToPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ResetMappingMode() {
	_jsii_.InvokeVoid(
		v,
		"resetMappingMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNatRulesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VappNatRulesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

