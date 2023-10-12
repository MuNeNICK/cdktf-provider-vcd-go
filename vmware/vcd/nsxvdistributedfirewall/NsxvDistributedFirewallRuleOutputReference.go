package nsxvdistributedfirewall

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/nsxvdistributedfirewall/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvDistributedFirewallRuleOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	Application() NsxvDistributedFirewallRuleApplicationList
	ApplicationInput() interface{}
	AppliedTo() NsxvDistributedFirewallRuleAppliedToList
	AppliedToInput() interface{}
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
	Destination() NsxvDistributedFirewallRuleDestinationList
	DestinationInput() interface{}
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludeDestination() interface{}
	SetExcludeDestination(val interface{})
	ExcludeDestinationInput() interface{}
	ExcludeSource() interface{}
	SetExcludeSource(val interface{})
	ExcludeSourceInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Logged() interface{}
	SetLogged(val interface{})
	LoggedInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PacketType() *string
	SetPacketType(val *string)
	PacketTypeInput() *string
	Source() NsxvDistributedFirewallRuleSourceList
	SourceInput() interface{}
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
	PutApplication(value interface{})
	PutAppliedTo(value interface{})
	PutDestination(value interface{})
	PutSource(value interface{})
	ResetApplication()
	ResetDestination()
	ResetEnabled()
	ResetExcludeDestination()
	ResetExcludeSource()
	ResetLogged()
	ResetName()
	ResetPacketType()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxvDistributedFirewallRuleOutputReference
type jsiiProxy_NsxvDistributedFirewallRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Application() NsxvDistributedFirewallRuleApplicationList {
	var returns NsxvDistributedFirewallRuleApplicationList
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) AppliedTo() NsxvDistributedFirewallRuleAppliedToList {
	var returns NsxvDistributedFirewallRuleAppliedToList
	_jsii_.Get(
		j,
		"appliedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) AppliedToInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appliedToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Destination() NsxvDistributedFirewallRuleDestinationList {
	var returns NsxvDistributedFirewallRuleDestinationList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ExcludeDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ExcludeDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ExcludeSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ExcludeSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Logged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) LoggedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PacketType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PacketTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Source() NsxvDistributedFirewallRuleSourceList {
	var returns NsxvDistributedFirewallRuleSourceList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNsxvDistributedFirewallRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NsxvDistributedFirewallRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNsxvDistributedFirewallRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxvDistributedFirewallRuleOutputReference{}

	_jsii_.Create(
		"vcd.nsxvDistributedFirewall.NsxvDistributedFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNsxvDistributedFirewallRuleOutputReference_Override(n NsxvDistributedFirewallRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxvDistributedFirewall.NsxvDistributedFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetExcludeDestination(val interface{}) {
	if err := j.validateSetExcludeDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeDestination",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetExcludeSource(val interface{}) {
	if err := j.validateSetExcludeSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeSource",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetLogged(val interface{}) {
	if err := j.validateSetLoggedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logged",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetPacketType(val *string) {
	if err := j.validateSetPacketTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packetType",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxvDistributedFirewallRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PutApplication(value interface{}) {
	if err := n.validatePutApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putApplication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PutAppliedTo(value interface{}) {
	if err := n.validatePutAppliedToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAppliedTo",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PutDestination(value interface{}) {
	if err := n.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestination",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) PutSource(value interface{}) {
	if err := n.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSource",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetApplication() {
	_jsii_.InvokeVoid(
		n,
		"resetApplication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		n,
		"resetDestination",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetExcludeDestination() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeDestination",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetExcludeSource() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetLogged() {
	_jsii_.InvokeVoid(
		n,
		"resetLogged",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		n,
		"resetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetPacketType() {
	_jsii_.InvokeVoid(
		n,
		"resetPacketType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		n,
		"resetSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDistributedFirewallRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

