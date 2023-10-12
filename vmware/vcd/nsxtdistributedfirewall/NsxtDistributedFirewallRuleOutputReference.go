package nsxtdistributedfirewall

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtdistributedfirewall/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtDistributedFirewallRuleOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	AppPortProfileIds() *[]*string
	SetAppPortProfileIds(val *[]*string)
	AppPortProfileIdsInput() *[]*string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationGroupsExcluded() interface{}
	SetDestinationGroupsExcluded(val interface{})
	DestinationGroupsExcludedInput() interface{}
	DestinationIds() *[]*string
	SetDestinationIds(val *[]*string)
	DestinationIdsInput() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	Logging() interface{}
	SetLogging(val interface{})
	LoggingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkContextProfileIds() *[]*string
	SetNetworkContextProfileIds(val *[]*string)
	NetworkContextProfileIdsInput() *[]*string
	SourceGroupsExcluded() interface{}
	SetSourceGroupsExcluded(val interface{})
	SourceGroupsExcludedInput() interface{}
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	SourceIdsInput() *[]*string
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
	ResetAppPortProfileIds()
	ResetComment()
	ResetDescription()
	ResetDestinationGroupsExcluded()
	ResetDestinationIds()
	ResetDirection()
	ResetEnabled()
	ResetIpProtocol()
	ResetLogging()
	ResetNetworkContextProfileIds()
	ResetSourceGroupsExcluded()
	ResetSourceIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxtDistributedFirewallRuleOutputReference
type jsiiProxy_NsxtDistributedFirewallRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) AppPortProfileIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appPortProfileIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) AppPortProfileIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appPortProfileIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DestinationGroupsExcluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationGroupsExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DestinationGroupsExcludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationGroupsExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DestinationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DestinationIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) NetworkContextProfileIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkContextProfileIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) NetworkContextProfileIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkContextProfileIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) SourceGroupsExcluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceGroupsExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) SourceGroupsExcludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceGroupsExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) SourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNsxtDistributedFirewallRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NsxtDistributedFirewallRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNsxtDistributedFirewallRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtDistributedFirewallRuleOutputReference{}

	_jsii_.Create(
		"vcd.nsxtDistributedFirewall.NsxtDistributedFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNsxtDistributedFirewallRuleOutputReference_Override(n NsxtDistributedFirewallRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtDistributedFirewall.NsxtDistributedFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetAppPortProfileIds(val *[]*string) {
	if err := j.validateSetAppPortProfileIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appPortProfileIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetDestinationGroupsExcluded(val interface{}) {
	if err := j.validateSetDestinationGroupsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationGroupsExcluded",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetDestinationIds(val *[]*string) {
	if err := j.validateSetDestinationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetLogging(val interface{}) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetNetworkContextProfileIds(val *[]*string) {
	if err := j.validateSetNetworkContextProfileIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkContextProfileIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetSourceGroupsExcluded(val interface{}) {
	if err := j.validateSetSourceGroupsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceGroupsExcluded",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetSourceIds(val *[]*string) {
	if err := j.validateSetSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetAppPortProfileIds() {
	_jsii_.InvokeVoid(
		n,
		"resetAppPortProfileIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		n,
		"resetComment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetDestinationGroupsExcluded() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationGroupsExcluded",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetDestinationIds() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		n,
		"resetDirection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		n,
		"resetLogging",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetNetworkContextProfileIds() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkContextProfileIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetSourceGroupsExcluded() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceGroupsExcluded",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ResetSourceIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

