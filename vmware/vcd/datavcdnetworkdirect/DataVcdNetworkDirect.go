package datavcdnetworkdirect

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdnetworkdirect/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct vcd_network_direct}.
type DataVcdNetworkDirect interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	ExternalNetwork() *string
	ExternalNetworkDns1() *string
	ExternalNetworkDns2() *string
	ExternalNetworkDnsSuffix() *string
	ExternalNetworkGateway() *string
	ExternalNetworkNetmask() *string
	Filter() DataVcdNetworkDirectFilterOutputReference
	FilterInput() *DataVcdNetworkDirectFilter
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Href() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() cdktf.StringMap
	MetadataEntry() DataVcdNetworkDirectMetadataEntryList
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
	Shared() cdktf.IResolvable
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
	PutFilter(value *DataVcdNetworkDirectFilter)
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

// The jsii proxy struct for DataVcdNetworkDirect
type jsiiProxy_DataVcdNetworkDirect struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdNetworkDirect) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetworkDns1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkDns1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetworkDns2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkDns2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetworkDnsSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkDnsSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetworkGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ExternalNetworkNetmask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkNetmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Filter() DataVcdNetworkDirectFilterOutputReference {
	var returns DataVcdNetworkDirectFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) FilterInput() *DataVcdNetworkDirectFilter {
	var returns *DataVcdNetworkDirectFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) MetadataEntry() DataVcdNetworkDirectMetadataEntryList {
	var returns DataVcdNetworkDirectMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Shared() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"shared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNetworkDirect) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct vcd_network_direct} Data Source.
func NewDataVcdNetworkDirect(scope constructs.Construct, id *string, config *DataVcdNetworkDirectConfig) DataVcdNetworkDirect {
	_init_.Initialize()

	if err := validateNewDataVcdNetworkDirectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdNetworkDirect{}

	_jsii_.Create(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/network_direct vcd_network_direct} Data Source.
func NewDataVcdNetworkDirect_Override(d DataVcdNetworkDirect, scope constructs.Construct, id *string, config *DataVcdNetworkDirectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdNetworkDirect)SetVdc(val *string) {
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
func DataVcdNetworkDirect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNetworkDirect_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNetworkDirect_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNetworkDirect_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNetworkDirect_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNetworkDirect_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdNetworkDirect_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdNetworkDirect.DataVcdNetworkDirect",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdNetworkDirect) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdNetworkDirect) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdNetworkDirect) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNetworkDirect) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) PutFilter(value *DataVcdNetworkDirectFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) ResetVdc() {
	_jsii_.InvokeVoid(
		d,
		"resetVdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNetworkDirect) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNetworkDirect) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNetworkDirect) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNetworkDirect) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

