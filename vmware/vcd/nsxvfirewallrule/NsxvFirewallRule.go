package nsxvfirewallrule

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxvfirewallrule/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule vcd_nsxv_firewall_rule}.
type NsxvFirewallRule interface {
	cdktf.TerraformResource
	AboveRuleId() *string
	SetAboveRuleId(val *string)
	AboveRuleIdInput() *string
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destination() NsxvFirewallRuleDestinationOutputReference
	DestinationInput() *NsxvFirewallRuleDestination
	EdgeGateway() *string
	SetEdgeGateway(val *string)
	EdgeGatewayInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingEnabled() interface{}
	SetLoggingEnabled(val interface{})
	LoggingEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RuleTag() *float64
	SetRuleTag(val *float64)
	RuleTagInput() *float64
	RuleType() *string
	SetRuleType(val *string)
	RuleTypeInput() *string
	Service() NsxvFirewallRuleServiceList
	ServiceInput() interface{}
	Source() NsxvFirewallRuleSourceOutputReference
	SourceInput() *NsxvFirewallRuleSource
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Vdc() *string
	SetVdc(val *string)
	VdcInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDestination(value *NsxvFirewallRuleDestination)
	PutService(value interface{})
	PutSource(value *NsxvFirewallRuleSource)
	ResetAboveRuleId()
	ResetAction()
	ResetEnabled()
	ResetId()
	ResetLoggingEnabled()
	ResetName()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuleTag()
	ResetRuleType()
	ResetVdc()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxvFirewallRule
type jsiiProxy_NsxvFirewallRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxvFirewallRule) AboveRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboveRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) AboveRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboveRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Destination() NsxvFirewallRuleDestinationOutputReference {
	var returns NsxvFirewallRuleDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) DestinationInput() *NsxvFirewallRuleDestination {
	var returns *NsxvFirewallRuleDestination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) EdgeGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) EdgeGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) LoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) LoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) RuleTag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) RuleTagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) RuleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) RuleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Service() NsxvFirewallRuleServiceList {
	var returns NsxvFirewallRuleServiceList
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) ServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Source() NsxvFirewallRuleSourceOutputReference {
	var returns NsxvFirewallRuleSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) SourceInput() *NsxvFirewallRuleSource {
	var returns *NsxvFirewallRuleSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRule) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule vcd_nsxv_firewall_rule} Resource.
func NewNsxvFirewallRule(scope constructs.Construct, id *string, config *NsxvFirewallRuleConfig) NsxvFirewallRule {
	_init_.Initialize()

	if err := validateNewNsxvFirewallRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxvFirewallRule{}

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule vcd_nsxv_firewall_rule} Resource.
func NewNsxvFirewallRule_Override(n NsxvFirewallRule, scope constructs.Construct, id *string, config *NsxvFirewallRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetAboveRuleId(val *string) {
	if err := j.validateSetAboveRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aboveRuleId",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetEdgeGateway(val *string) {
	if err := j.validateSetEdgeGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGateway",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetLoggingEnabled(val interface{}) {
	if err := j.validateSetLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetRuleTag(val *float64) {
	if err := j.validateSetRuleTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleTag",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetRuleType(val *string) {
	if err := j.validateSetRuleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleType",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRule)SetVdc(val *string) {
	if err := j.validateSetVdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdc",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NsxvFirewallRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvFirewallRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxvFirewallRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvFirewallRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxvFirewallRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvFirewallRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxvFirewallRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxvFirewallRule.NsxvFirewallRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxvFirewallRule) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxvFirewallRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxvFirewallRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvFirewallRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxvFirewallRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxvFirewallRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxvFirewallRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxvFirewallRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxvFirewallRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRule) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxvFirewallRule) PutDestination(value *NsxvFirewallRuleDestination) {
	if err := n.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestination",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvFirewallRule) PutService(value interface{}) {
	if err := n.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putService",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvFirewallRule) PutSource(value *NsxvFirewallRuleSource) {
	if err := n.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSource",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetAboveRuleId() {
	_jsii_.InvokeVoid(
		n,
		"resetAboveRuleId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetAction() {
	_jsii_.InvokeVoid(
		n,
		"resetAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetLoggingEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetLoggingEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetName() {
	_jsii_.InvokeVoid(
		n,
		"resetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetRuleTag() {
	_jsii_.InvokeVoid(
		n,
		"resetRuleTag",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetRuleType() {
	_jsii_.InvokeVoid(
		n,
		"resetRuleType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

