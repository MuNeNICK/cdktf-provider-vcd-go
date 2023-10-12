package vdcgroup

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vdcgroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group vcd_vdc_group}.
type VdcGroup interface {
	cdktf.TerraformResource
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
	DefaultPolicyStatus() interface{}
	SetDefaultPolicyStatus(val interface{})
	DefaultPolicyStatusInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DfwEnabled() interface{}
	SetDfwEnabled(val interface{})
	DfwEnabledInput() interface{}
	ErrorMessage() *string
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
	LocalEgress() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkPoolId() *string
	NetworkPoolUniversalId() *string
	NetworkProviderType() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	ParticipatingOrgVdcs() VdcGroupParticipatingOrgVdcsList
	ParticipatingVdcIds() *[]*string
	SetParticipatingVdcIds(val *[]*string)
	ParticipatingVdcIdsInput() *[]*string
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
	RemoveDefaultFirewallRule() interface{}
	SetRemoveDefaultFirewallRule(val interface{})
	RemoveDefaultFirewallRuleInput() interface{}
	StartingVdcId() *string
	SetStartingVdcId(val *string)
	StartingVdcIdInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	UniversalNetworkingEnabled() cdktf.IResolvable
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
	ResetDefaultPolicyStatus()
	ResetDescription()
	ResetDfwEnabled()
	ResetId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoveDefaultFirewallRule()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VdcGroup
type jsiiProxy_VdcGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VdcGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DefaultPolicyStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPolicyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DefaultPolicyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPolicyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DfwEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dfwEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) DfwEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dfwEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) LocalEgress() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"localEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) NetworkPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) NetworkPoolUniversalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolUniversalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) NetworkProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ParticipatingOrgVdcs() VdcGroupParticipatingOrgVdcsList {
	var returns VdcGroupParticipatingOrgVdcsList
	_jsii_.Get(
		j,
		"participatingOrgVdcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ParticipatingVdcIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"participatingVdcIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) ParticipatingVdcIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"participatingVdcIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) RemoveDefaultFirewallRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultFirewallRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) RemoveDefaultFirewallRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultFirewallRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) StartingVdcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingVdcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) StartingVdcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingVdcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VdcGroup) UniversalNetworkingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"universalNetworkingEnabled",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group vcd_vdc_group} Resource.
func NewVdcGroup(scope constructs.Construct, id *string, config *VdcGroupConfig) VdcGroup {
	_init_.Initialize()

	if err := validateNewVdcGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VdcGroup{}

	_jsii_.Create(
		"vcd.vdcGroup.VdcGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vdc_group vcd_vdc_group} Resource.
func NewVdcGroup_Override(v VdcGroup, scope constructs.Construct, id *string, config *VdcGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vdcGroup.VdcGroup",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VdcGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetDefaultPolicyStatus(val interface{}) {
	if err := j.validateSetDefaultPolicyStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPolicyStatus",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetDfwEnabled(val interface{}) {
	if err := j.validateSetDfwEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dfwEnabled",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetParticipatingVdcIds(val *[]*string) {
	if err := j.validateSetParticipatingVdcIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"participatingVdcIds",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetRemoveDefaultFirewallRule(val interface{}) {
	if err := j.validateSetRemoveDefaultFirewallRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeDefaultFirewallRule",
		val,
	)
}

func (j *jsiiProxy_VdcGroup)SetStartingVdcId(val *string) {
	if err := j.validateSetStartingVdcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingVdcId",
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
func VdcGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVdcGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vdcGroup.VdcGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VdcGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVdcGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vdcGroup.VdcGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VdcGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVdcGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vdcGroup.VdcGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VdcGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.vdcGroup.VdcGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VdcGroup) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VdcGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VdcGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VdcGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VdcGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VdcGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VdcGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VdcGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VdcGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VdcGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VdcGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroup) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VdcGroup) ResetDefaultPolicyStatus() {
	_jsii_.InvokeVoid(
		v,
		"resetDefaultPolicyStatus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetDfwEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetDfwEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetOrg() {
	_jsii_.InvokeVoid(
		v,
		"resetOrg",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) ResetRemoveDefaultFirewallRule() {
	_jsii_.InvokeVoid(
		v,
		"resetRemoveDefaultFirewallRule",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VdcGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VdcGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

