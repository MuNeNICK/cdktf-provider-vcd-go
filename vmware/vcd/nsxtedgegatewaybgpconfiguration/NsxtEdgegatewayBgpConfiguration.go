package nsxtedgegatewaybgpconfiguration

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtedgegatewaybgpconfiguration/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration vcd_nsxt_edgegateway_bgp_configuration}.
type NsxtEdgegatewayBgpConfiguration interface {
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
	EcmpEnabled() interface{}
	SetEcmpEnabled(val interface{})
	EcmpEnabledInput() interface{}
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
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
	GracefulRestartMode() *string
	SetGracefulRestartMode(val *string)
	GracefulRestartModeInput() *string
	GracefulRestartTimer() *float64
	SetGracefulRestartTimer(val *float64)
	GracefulRestartTimerInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAsNumber() *string
	SetLocalAsNumber(val *string)
	LocalAsNumberInput() *string
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
	StaleRouteTimer() *float64
	SetStaleRouteTimer(val *float64)
	StaleRouteTimerInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetEcmpEnabled()
	ResetGracefulRestartMode()
	ResetGracefulRestartTimer()
	ResetId()
	ResetLocalAsNumber()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStaleRouteTimer()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxtEdgegatewayBgpConfiguration
type jsiiProxy_NsxtEdgegatewayBgpConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) EcmpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecmpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) EcmpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecmpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GracefulRestartMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulRestartMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GracefulRestartModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulRestartModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GracefulRestartTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracefulRestartTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GracefulRestartTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracefulRestartTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) LocalAsNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAsNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) LocalAsNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAsNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) StaleRouteTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"staleRouteTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) StaleRouteTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"staleRouteTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration vcd_nsxt_edgegateway_bgp_configuration} Resource.
func NewNsxtEdgegatewayBgpConfiguration(scope constructs.Construct, id *string, config *NsxtEdgegatewayBgpConfigurationConfig) NsxtEdgegatewayBgpConfiguration {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewayBgpConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegatewayBgpConfiguration{}

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_configuration vcd_nsxt_edgegateway_bgp_configuration} Resource.
func NewNsxtEdgegatewayBgpConfiguration_Override(n NsxtEdgegatewayBgpConfiguration, scope constructs.Construct, id *string, config *NsxtEdgegatewayBgpConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetEcmpEnabled(val interface{}) {
	if err := j.validateSetEcmpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecmpEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetGracefulRestartMode(val *string) {
	if err := j.validateSetGracefulRestartModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulRestartMode",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetGracefulRestartTimer(val *float64) {
	if err := j.validateSetGracefulRestartTimerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulRestartTimer",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetLocalAsNumber(val *string) {
	if err := j.validateSetLocalAsNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAsNumber",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpConfiguration)SetStaleRouteTimer(val *float64) {
	if err := j.validateSetStaleRouteTimerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staleRouteTimer",
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
func NsxtEdgegatewayBgpConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegatewayBgpConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegatewayBgpConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtEdgegatewayBgpConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtEdgegatewayBgpConfiguration.NsxtEdgegatewayBgpConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetEcmpEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEcmpEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetGracefulRestartMode() {
	_jsii_.InvokeVoid(
		n,
		"resetGracefulRestartMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetGracefulRestartTimer() {
	_jsii_.InvokeVoid(
		n,
		"resetGracefulRestartTimer",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetLocalAsNumber() {
	_jsii_.InvokeVoid(
		n,
		"resetLocalAsNumber",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ResetStaleRouteTimer() {
	_jsii_.InvokeVoid(
		n,
		"resetStaleRouteTimer",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

