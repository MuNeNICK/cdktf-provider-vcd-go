package datavcdedgegateway

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdedgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/edgegateway vcd_edgegateway}.
type DataVcdEdgegateway interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Configuration() *string
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
	DistributedRouting() cdktf.IResolvable
	ExternalNetwork() DataVcdEdgegatewayExternalNetworkList
	ExternalNetworkIps() *[]*string
	Filter() DataVcdEdgegatewayFilterOutputReference
	FilterInput() *DataVcdEdgegatewayFilter
	FipsModeEnabled() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FwDefaultRuleAction() *string
	FwDefaultRuleLoggingEnabled() cdktf.IResolvable
	FwEnabled() cdktf.IResolvable
	HaEnabled() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	LbAccelerationEnabled() cdktf.IResolvable
	LbEnabled() cdktf.IResolvable
	LbLoggingEnabled() cdktf.IResolvable
	LbLoglevel() *string
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
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseDefaultRouteForDnsRelay() cdktf.IResolvable
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
	PutFilter(value *DataVcdEdgegatewayFilter)
	ResetFilter()
	ResetId()
	ResetName()
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

// The jsii proxy struct for DataVcdEdgegateway
type jsiiProxy_DataVcdEdgegateway struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdEdgegateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) DefaultExternalNetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExternalNetworkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) DistributedRouting() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"distributedRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) ExternalNetwork() DataVcdEdgegatewayExternalNetworkList {
	var returns DataVcdEdgegatewayExternalNetworkList
	_jsii_.Get(
		j,
		"externalNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) ExternalNetworkIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalNetworkIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Filter() DataVcdEdgegatewayFilterOutputReference {
	var returns DataVcdEdgegatewayFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FilterInput() *DataVcdEdgegatewayFilter {
	var returns *DataVcdEdgegatewayFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FipsModeEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fipsModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FwDefaultRuleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fwDefaultRuleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FwDefaultRuleLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fwDefaultRuleLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) FwEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fwEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) HaEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"haEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) LbAccelerationEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lbAccelerationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) LbEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) LbLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lbLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) LbLoglevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbLoglevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) UseDefaultRouteForDnsRelay() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useDefaultRouteForDnsRelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdEdgegateway) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/edgegateway vcd_edgegateway} Data Source.
func NewDataVcdEdgegateway(scope constructs.Construct, id *string, config *DataVcdEdgegatewayConfig) DataVcdEdgegateway {
	_init_.Initialize()

	if err := validateNewDataVcdEdgegatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdEdgegateway{}

	_jsii_.Create(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/edgegateway vcd_edgegateway} Data Source.
func NewDataVcdEdgegateway_Override(d DataVcdEdgegateway, scope constructs.Construct, id *string, config *DataVcdEdgegatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdEdgegateway)SetVdc(val *string) {
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
func DataVcdEdgegateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdEdgegateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdEdgegateway_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdEdgegateway_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdEdgegateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdEdgegateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdEdgegateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdEdgegateway.DataVcdEdgegateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdEdgegateway) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdEdgegateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdEdgegateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdEdgegateway) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) PutFilter(value *DataVcdEdgegatewayFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) ResetVdc() {
	_jsii_.InvokeVoid(
		d,
		"resetVdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdEdgegateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdEdgegateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

