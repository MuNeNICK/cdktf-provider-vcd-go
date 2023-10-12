package edgegatewaysettings

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/edgegatewaysettings/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings vcd_edgegateway_settings}.
type EdgegatewaySettings interface {
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
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
	EdgeGatewayName() *string
	SetEdgeGatewayName(val *string)
	EdgeGatewayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FwDefaultRuleAction() *string
	SetFwDefaultRuleAction(val *string)
	FwDefaultRuleActionInput() *string
	FwDefaultRuleLoggingEnabled() interface{}
	SetFwDefaultRuleLoggingEnabled(val interface{})
	FwDefaultRuleLoggingEnabledInput() interface{}
	FwEnabled() interface{}
	SetFwEnabled(val interface{})
	FwEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	LbAccelerationEnabled() interface{}
	SetLbAccelerationEnabled(val interface{})
	LbAccelerationEnabledInput() interface{}
	LbEnabled() interface{}
	SetLbEnabled(val interface{})
	LbEnabledInput() interface{}
	LbLoggingEnabled() interface{}
	SetLbLoggingEnabled(val interface{})
	LbLoggingEnabledInput() interface{}
	LbLoglevel() *string
	SetLbLoglevel(val *string)
	LbLoglevelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ResetEdgeGatewayId()
	ResetEdgeGatewayName()
	ResetFwDefaultRuleAction()
	ResetFwDefaultRuleLoggingEnabled()
	ResetFwEnabled()
	ResetId()
	ResetLbAccelerationEnabled()
	ResetLbEnabled()
	ResetLbLoggingEnabled()
	ResetLbLoglevel()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for EdgegatewaySettings
type jsiiProxy_EdgegatewaySettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EdgegatewaySettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) EdgeGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) EdgeGatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwDefaultRuleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fwDefaultRuleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwDefaultRuleActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fwDefaultRuleActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwDefaultRuleLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwDefaultRuleLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwDefaultRuleLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwDefaultRuleLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) FwEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbAccelerationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbAccelerationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbAccelerationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbAccelerationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbLoglevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbLoglevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) LbLoglevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbLoglevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewaySettings) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings vcd_edgegateway_settings} Resource.
func NewEdgegatewaySettings(scope constructs.Construct, id *string, config *EdgegatewaySettingsConfig) EdgegatewaySettings {
	_init_.Initialize()

	if err := validateNewEdgegatewaySettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgegatewaySettings{}

	_jsii_.Create(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_settings vcd_edgegateway_settings} Resource.
func NewEdgegatewaySettings_Override(e EdgegatewaySettings, scope constructs.Construct, id *string, config *EdgegatewaySettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetEdgeGatewayName(val *string) {
	if err := j.validateSetEdgeGatewayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayName",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetFwDefaultRuleAction(val *string) {
	if err := j.validateSetFwDefaultRuleActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwDefaultRuleAction",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetFwDefaultRuleLoggingEnabled(val interface{}) {
	if err := j.validateSetFwDefaultRuleLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwDefaultRuleLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetFwEnabled(val interface{}) {
	if err := j.validateSetFwEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwEnabled",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetLbAccelerationEnabled(val interface{}) {
	if err := j.validateSetLbAccelerationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbAccelerationEnabled",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetLbEnabled(val interface{}) {
	if err := j.validateSetLbEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbEnabled",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetLbLoggingEnabled(val interface{}) {
	if err := j.validateSetLbLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetLbLoglevel(val *string) {
	if err := j.validateSetLbLoglevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbLoglevel",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EdgegatewaySettings)SetVdc(val *string) {
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
func EdgegatewaySettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewaySettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgegatewaySettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewaySettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgegatewaySettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewaySettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EdgegatewaySettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.edgegatewaySettings.EdgegatewaySettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EdgegatewaySettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetEdgeGatewayId() {
	_jsii_.InvokeVoid(
		e,
		"resetEdgeGatewayId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetEdgeGatewayName() {
	_jsii_.InvokeVoid(
		e,
		"resetEdgeGatewayName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetFwDefaultRuleAction() {
	_jsii_.InvokeVoid(
		e,
		"resetFwDefaultRuleAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetFwDefaultRuleLoggingEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetFwDefaultRuleLoggingEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetFwEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetFwEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetLbAccelerationEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbAccelerationEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetLbEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetLbLoggingEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbLoggingEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetLbLoglevel() {
	_jsii_.InvokeVoid(
		e,
		"resetLbLoglevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetOrg() {
	_jsii_.InvokeVoid(
		e,
		"resetOrg",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) ResetVdc() {
	_jsii_.InvokeVoid(
		e,
		"resetVdc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewaySettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewaySettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

