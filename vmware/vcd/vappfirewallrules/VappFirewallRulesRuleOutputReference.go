package vappfirewallrules

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/vappfirewallrules/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappFirewallRulesRuleOutputReference interface {
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
	DestinationIp() *string
	SetDestinationIp(val *string)
	DestinationIpInput() *string
	DestinationPort() *string
	SetDestinationPort(val *string)
	DestinationPortInput() *string
	DestinationVmId() *string
	SetDestinationVmId(val *string)
	DestinationVmIdInput() *string
	DestinationVmIpType() *string
	SetDestinationVmIpType(val *string)
	DestinationVmIpTypeInput() *string
	DestinationVmNicId() *float64
	SetDestinationVmNicId(val *float64)
	DestinationVmNicIdInput() *float64
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SourceIp() *string
	SetSourceIp(val *string)
	SourceIpInput() *string
	SourcePort() *string
	SetSourcePort(val *string)
	SourcePortInput() *string
	SourceVmId() *string
	SetSourceVmId(val *string)
	SourceVmIdInput() *string
	SourceVmIpType() *string
	SetSourceVmIpType(val *string)
	SourceVmIpTypeInput() *string
	SourceVmNicId() *float64
	SetSourceVmNicId(val *float64)
	SourceVmNicIdInput() *float64
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
	ResetDestinationIp()
	ResetDestinationPort()
	ResetDestinationVmId()
	ResetDestinationVmIpType()
	ResetDestinationVmNicId()
	ResetEnabled()
	ResetEnableLogging()
	ResetName()
	ResetPolicy()
	ResetProtocol()
	ResetSourceIp()
	ResetSourcePort()
	ResetSourceVmId()
	ResetSourceVmIpType()
	ResetSourceVmNicId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VappFirewallRulesRuleOutputReference
type jsiiProxy_VappFirewallRulesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationVmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationVmIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmIpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationVmIpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmIpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationVmIpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmNicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationVmNicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) DestinationVmNicIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationVmNicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourcePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourcePortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmIpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmIpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmIpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVmIpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmNicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceVmNicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) SourceVmNicIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceVmNicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVappFirewallRulesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VappFirewallRulesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewVappFirewallRulesRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappFirewallRulesRuleOutputReference{}

	_jsii_.Create(
		"vcd.vappFirewallRules.VappFirewallRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVappFirewallRulesRuleOutputReference_Override(v VappFirewallRulesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappFirewallRules.VappFirewallRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetDestinationIp(val *string) {
	if err := j.validateSetDestinationIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIp",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetDestinationPort(val *string) {
	if err := j.validateSetDestinationPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPort",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetDestinationVmId(val *string) {
	if err := j.validateSetDestinationVmIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationVmId",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetDestinationVmIpType(val *string) {
	if err := j.validateSetDestinationVmIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationVmIpType",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetDestinationVmNicId(val *float64) {
	if err := j.validateSetDestinationVmNicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationVmNicId",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetSourceIp(val *string) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetSourcePort(val *string) {
	if err := j.validateSetSourcePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePort",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetSourceVmId(val *string) {
	if err := j.validateSetSourceVmIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVmId",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetSourceVmIpType(val *string) {
	if err := j.validateSetSourceVmIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVmIpType",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetSourceVmNicId(val *float64) {
	if err := j.validateSetSourceVmNicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVmNicId",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VappFirewallRulesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetDestinationIp() {
	_jsii_.InvokeVoid(
		v,
		"resetDestinationIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		v,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetDestinationVmId() {
	_jsii_.InvokeVoid(
		v,
		"resetDestinationVmId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetDestinationVmIpType() {
	_jsii_.InvokeVoid(
		v,
		"resetDestinationVmIpType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetDestinationVmNicId() {
	_jsii_.InvokeVoid(
		v,
		"resetDestinationVmNicId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetSourcePort() {
	_jsii_.InvokeVoid(
		v,
		"resetSourcePort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetSourceVmId() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceVmId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetSourceVmIpType() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceVmIpType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ResetSourceVmNicId() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceVmNicId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VappFirewallRulesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

