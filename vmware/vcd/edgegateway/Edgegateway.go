package edgegateway

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/edgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway vcd_edgegateway}.
type Edgegateway interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Configuration() *string
	SetConfiguration(val *string)
	ConfigurationInput() *string
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
	DefaultExternalNetworkIp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DistributedRouting() interface{}
	SetDistributedRouting(val interface{})
	DistributedRoutingInput() interface{}
	ExternalNetwork() EdgegatewayExternalNetworkList
	ExternalNetworkInput() interface{}
	ExternalNetworkIps() *[]*string
	FipsModeEnabled() interface{}
	SetFipsModeEnabled(val interface{})
	FipsModeEnabledInput() interface{}
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
	HaEnabled() interface{}
	SetHaEnabled(val interface{})
	HaEnabledInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseDefaultRouteForDnsRelay() interface{}
	SetUseDefaultRouteForDnsRelay(val interface{})
	UseDefaultRouteForDnsRelayInput() interface{}
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
	PutExternalNetwork(value interface{})
	ResetDescription()
	ResetDistributedRouting()
	ResetFipsModeEnabled()
	ResetFwDefaultRuleAction()
	ResetFwDefaultRuleLoggingEnabled()
	ResetFwEnabled()
	ResetHaEnabled()
	ResetId()
	ResetLbAccelerationEnabled()
	ResetLbEnabled()
	ResetLbLoggingEnabled()
	ResetLbLoglevel()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUseDefaultRouteForDnsRelay()
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

// The jsii proxy struct for Edgegateway
type jsiiProxy_Edgegateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Edgegateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) DefaultExternalNetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExternalNetworkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) DistributedRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributedRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) DistributedRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributedRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ExternalNetwork() EdgegatewayExternalNetworkList {
	var returns EdgegatewayExternalNetworkList
	_jsii_.Get(
		j,
		"externalNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ExternalNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ExternalNetworkIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalNetworkIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FipsModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FipsModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwDefaultRuleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fwDefaultRuleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwDefaultRuleActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fwDefaultRuleActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwDefaultRuleLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwDefaultRuleLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwDefaultRuleLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwDefaultRuleLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) FwEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fwEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) HaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) HaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbAccelerationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbAccelerationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbAccelerationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbAccelerationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbLoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbLoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbLoggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbLoglevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbLoglevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) LbLoglevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbLoglevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) UseDefaultRouteForDnsRelay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultRouteForDnsRelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) UseDefaultRouteForDnsRelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultRouteForDnsRelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Edgegateway) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway vcd_edgegateway} Resource.
func NewEdgegateway(scope constructs.Construct, id *string, config *EdgegatewayConfig) Edgegateway {
	_init_.Initialize()

	if err := validateNewEdgegatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Edgegateway{}

	_jsii_.Create(
		"vcd.edgegateway.Edgegateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway vcd_edgegateway} Resource.
func NewEdgegateway_Override(e Edgegateway, scope constructs.Construct, id *string, config *EdgegatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.edgegateway.Edgegateway",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Edgegateway)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetDistributedRouting(val interface{}) {
	if err := j.validateSetDistributedRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributedRouting",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetFipsModeEnabled(val interface{}) {
	if err := j.validateSetFipsModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fipsModeEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetFwDefaultRuleAction(val *string) {
	if err := j.validateSetFwDefaultRuleActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwDefaultRuleAction",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetFwDefaultRuleLoggingEnabled(val interface{}) {
	if err := j.validateSetFwDefaultRuleLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwDefaultRuleLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetFwEnabled(val interface{}) {
	if err := j.validateSetFwEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fwEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetHaEnabled(val interface{}) {
	if err := j.validateSetHaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetLbAccelerationEnabled(val interface{}) {
	if err := j.validateSetLbAccelerationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbAccelerationEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetLbEnabled(val interface{}) {
	if err := j.validateSetLbEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetLbLoggingEnabled(val interface{}) {
	if err := j.validateSetLbLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbLoggingEnabled",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetLbLoglevel(val *string) {
	if err := j.validateSetLbLoglevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbLoglevel",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetUseDefaultRouteForDnsRelay(val interface{}) {
	if err := j.validateSetUseDefaultRouteForDnsRelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultRouteForDnsRelay",
		val,
	)
}

func (j *jsiiProxy_Edgegateway)SetVdc(val *string) {
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
func Edgegateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegateway.Edgegateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Edgegateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegateway.Edgegateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Edgegateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegateway.Edgegateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Edgegateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.edgegateway.Edgegateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Edgegateway) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Edgegateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Edgegateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Edgegateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Edgegateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Edgegateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Edgegateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Edgegateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Edgegateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Edgegateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Edgegateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Edgegateway) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Edgegateway) PutExternalNetwork(value interface{}) {
	if err := e.validatePutExternalNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExternalNetwork",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Edgegateway) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetDistributedRouting() {
	_jsii_.InvokeVoid(
		e,
		"resetDistributedRouting",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetFipsModeEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetFipsModeEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetFwDefaultRuleAction() {
	_jsii_.InvokeVoid(
		e,
		"resetFwDefaultRuleAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetFwDefaultRuleLoggingEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetFwDefaultRuleLoggingEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetFwEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetFwEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetHaEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetHaEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetLbAccelerationEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbAccelerationEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetLbEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetLbLoggingEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLbLoggingEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetLbLoglevel() {
	_jsii_.InvokeVoid(
		e,
		"resetLbLoglevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetOrg() {
	_jsii_.InvokeVoid(
		e,
		"resetOrg",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetUseDefaultRouteForDnsRelay() {
	_jsii_.InvokeVoid(
		e,
		"resetUseDefaultRouteForDnsRelay",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) ResetVdc() {
	_jsii_.InvokeVoid(
		e,
		"resetVdc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Edgegateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Edgegateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Edgegateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Edgegateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

