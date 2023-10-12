package nsxvdnat

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxvdnat/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat vcd_nsxv_dnat}.
type NsxvDnat interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IcmpType() *string
	SetIcmpType(val *string)
	IcmpTypeInput() *string
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
	NetworkName() *string
	SetNetworkName(val *string)
	NetworkNameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OriginalAddress() *string
	SetOriginalAddress(val *string)
	OriginalAddressInput() *string
	OriginalPort() *string
	SetOriginalPort(val *string)
	OriginalPortInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TranslatedAddress() *string
	SetTranslatedAddress(val *string)
	TranslatedAddressInput() *string
	TranslatedPort() *string
	SetTranslatedPort(val *string)
	TranslatedPortInput() *string
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
	ResetDescription()
	ResetEnabled()
	ResetIcmpType()
	ResetId()
	ResetLoggingEnabled()
	ResetOrg()
	ResetOriginalPort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetRuleTag()
	ResetRuleType()
	ResetTranslatedAddress()
	ResetTranslatedPort()
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

// The jsii proxy struct for NsxvDnat
type jsiiProxy_NsxvDnat struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxvDnat) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) EdgeGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) EdgeGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) LoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) LoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) NetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) NetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) OriginalAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) OriginalAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) OriginalPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) OriginalPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) RuleTag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) RuleTagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) RuleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) RuleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TranslatedAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TranslatedAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TranslatedPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) TranslatedPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvDnat) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat vcd_nsxv_dnat} Resource.
func NewNsxvDnat(scope constructs.Construct, id *string, config *NsxvDnatConfig) NsxvDnat {
	_init_.Initialize()

	if err := validateNewNsxvDnatParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxvDnat{}

	_jsii_.Create(
		"vcd.nsxvDnat.NsxvDnat",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_dnat vcd_nsxv_dnat} Resource.
func NewNsxvDnat_Override(n NsxvDnat, scope constructs.Construct, id *string, config *NsxvDnatConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxvDnat.NsxvDnat",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxvDnat)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetEdgeGateway(val *string) {
	if err := j.validateSetEdgeGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGateway",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetLoggingEnabled(val interface{}) {
	if err := j.validateSetLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetNetworkName(val *string) {
	if err := j.validateSetNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkName",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetOriginalAddress(val *string) {
	if err := j.validateSetOriginalAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalAddress",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetOriginalPort(val *string) {
	if err := j.validateSetOriginalPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalPort",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetRuleTag(val *float64) {
	if err := j.validateSetRuleTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleTag",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetRuleType(val *string) {
	if err := j.validateSetRuleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleType",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetTranslatedAddress(val *string) {
	if err := j.validateSetTranslatedAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"translatedAddress",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetTranslatedPort(val *string) {
	if err := j.validateSetTranslatedPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"translatedPort",
		val,
	)
}

func (j *jsiiProxy_NsxvDnat)SetVdc(val *string) {
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
func NsxvDnat_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvDnat_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvDnat.NsxvDnat",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxvDnat_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvDnat_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvDnat.NsxvDnat",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxvDnat_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxvDnat_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxvDnat.NsxvDnat",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxvDnat_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxvDnat.NsxvDnat",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxvDnat) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxvDnat) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxvDnat) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvDnat) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxvDnat) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxvDnat) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxvDnat) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxvDnat) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxvDnat) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxvDnat) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxvDnat) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvDnat) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxvDnat) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetIcmpType() {
	_jsii_.InvokeVoid(
		n,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetLoggingEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetLoggingEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetOriginalPort() {
	_jsii_.InvokeVoid(
		n,
		"resetOriginalPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetRuleTag() {
	_jsii_.InvokeVoid(
		n,
		"resetRuleTag",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetRuleType() {
	_jsii_.InvokeVoid(
		n,
		"resetRuleType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetTranslatedAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetTranslatedAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetTranslatedPort() {
	_jsii_.InvokeVoid(
		n,
		"resetTranslatedPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvDnat) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDnat) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDnat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvDnat) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

