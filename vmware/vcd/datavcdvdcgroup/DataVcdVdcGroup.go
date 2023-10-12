package datavcdvdcgroup

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdvdcgroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group vcd_vdc_group}.
type DataVcdVdcGroup interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DfwEnabled() cdktf.IResolvable
	ErrorMessage() *string
	SetErrorMessage(val *string)
	ErrorMessageInput() *string
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
	LocalEgress() interface{}
	SetLocalEgress(val interface{})
	LocalEgressInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkPoolId() *string
	SetNetworkPoolId(val *string)
	NetworkPoolIdInput() *string
	NetworkPoolUniversalId() *string
	SetNetworkPoolUniversalId(val *string)
	NetworkPoolUniversalIdInput() *string
	NetworkProviderType() *string
	SetNetworkProviderType(val *string)
	NetworkProviderTypeInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	ParticipatingOrgVdcs() DataVcdVdcGroupParticipatingOrgVdcsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UniversalNetworkingEnabled() interface{}
	SetUniversalNetworkingEnabled(val interface{})
	UniversalNetworkingEnabledInput() interface{}
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
	ResetErrorMessage()
	ResetId()
	ResetLocalEgress()
	ResetName()
	ResetNetworkPoolId()
	ResetNetworkPoolUniversalId()
	ResetNetworkProviderType()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStatus()
	ResetType()
	ResetUniversalNetworkingEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataVcdVdcGroup
type jsiiProxy_DataVcdVdcGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdVdcGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) DefaultPolicyStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPolicyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) DefaultPolicyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPolicyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) DfwEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dfwEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) ErrorMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) LocalEgress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) LocalEgressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localEgressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkPoolUniversalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolUniversalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkPoolUniversalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolUniversalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) NetworkProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) ParticipatingOrgVdcs() DataVcdVdcGroupParticipatingOrgVdcsList {
	var returns DataVcdVdcGroupParticipatingOrgVdcsList
	_jsii_.Get(
		j,
		"participatingOrgVdcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) UniversalNetworkingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"universalNetworkingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVdcGroup) UniversalNetworkingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"universalNetworkingEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group vcd_vdc_group} Data Source.
func NewDataVcdVdcGroup(scope constructs.Construct, id *string, config *DataVcdVdcGroupConfig) DataVcdVdcGroup {
	_init_.Initialize()

	if err := validateNewDataVcdVdcGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdVdcGroup{}

	_jsii_.Create(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vdc_group vcd_vdc_group} Data Source.
func NewDataVcdVdcGroup_Override(d DataVcdVdcGroup, scope constructs.Construct, id *string, config *DataVcdVdcGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetDefaultPolicyStatus(val interface{}) {
	if err := j.validateSetDefaultPolicyStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPolicyStatus",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetErrorMessage(val *string) {
	if err := j.validateSetErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorMessage",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetLocalEgress(val interface{}) {
	if err := j.validateSetLocalEgressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localEgress",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetNetworkPoolId(val *string) {
	if err := j.validateSetNetworkPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPoolId",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetNetworkPoolUniversalId(val *string) {
	if err := j.validateSetNetworkPoolUniversalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPoolUniversalId",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetNetworkProviderType(val *string) {
	if err := j.validateSetNetworkProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkProviderType",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataVcdVdcGroup)SetUniversalNetworkingEnabled(val interface{}) {
	if err := j.validateSetUniversalNetworkingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"universalNetworkingEnabled",
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
func DataVcdVdcGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVdcGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdVdcGroup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVdcGroup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdVdcGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVdcGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdVdcGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdVdcGroup.DataVcdVdcGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetDefaultPolicyStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultPolicyStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetErrorMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetErrorMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetLocalEgress() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalEgress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetNetworkPoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkPoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetNetworkPoolUniversalId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkPoolUniversalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetNetworkProviderType() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkProviderType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) ResetUniversalNetworkingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetUniversalNetworkingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVdcGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVdcGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

