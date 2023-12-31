package datavcdnsxtalbedgegatewayserviceenginegroup

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/datavcdnsxtalbedgegatewayserviceenginegroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group}.
type DataVcdNsxtAlbEdgegatewayServiceEngineGroup interface {
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
	DeployedVirtualServices() *float64
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
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
	MaxVirtualServices() *float64
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
	ReservedVirtualServices() *string
	ServiceEngineGroupId() *string
	SetServiceEngineGroupId(val *string)
	ServiceEngineGroupIdInput() *string
	ServiceEngineGroupName() *string
	SetServiceEngineGroupName(val *string)
	ServiceEngineGroupNameInput() *string
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
	ResetId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceEngineGroupId()
	ResetServiceEngineGroupName()
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

// The jsii proxy struct for DataVcdNsxtAlbEdgegatewayServiceEngineGroup
type jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) DeployedVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) MaxVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ReservedVirtualServices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group} Data Source.
func NewDataVcdNsxtAlbEdgegatewayServiceEngineGroup(scope constructs.Construct, id *string, config *DataVcdNsxtAlbEdgegatewayServiceEngineGroupConfig) DataVcdNsxtAlbEdgegatewayServiceEngineGroup {
	_init_.Initialize()

	if err := validateNewDataVcdNsxtAlbEdgegatewayServiceEngineGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup{}

	_jsii_.Create(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group} Data Source.
func NewDataVcdNsxtAlbEdgegatewayServiceEngineGroup_Override(d DataVcdNsxtAlbEdgegatewayServiceEngineGroup, scope constructs.Construct, id *string, config *DataVcdNsxtAlbEdgegatewayServiceEngineGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetServiceEngineGroupId(val *string) {
	if err := j.validateSetServiceEngineGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEngineGroupId",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetServiceEngineGroupName(val *string) {
	if err := j.validateSetServiceEngineGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEngineGroupName",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup)SetVdc(val *string) {
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
func DataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdNsxtAlbEdgegatewayServiceEngineGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdNsxtAlbEdgegatewayServiceEngineGroup.DataVcdNsxtAlbEdgegatewayServiceEngineGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetServiceEngineGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceEngineGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetServiceEngineGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceEngineGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ResetVdc() {
	_jsii_.InvokeVoid(
		d,
		"resetVdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbEdgegatewayServiceEngineGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

