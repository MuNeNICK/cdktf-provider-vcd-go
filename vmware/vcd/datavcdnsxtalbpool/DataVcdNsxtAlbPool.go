package datavcdnsxtalbpool

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdnsxtalbpool/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool vcd_nsxt_alb_pool}.
type DataVcdNsxtAlbPool interface {
	cdktf.TerraformDataSource
	Algorithm() *string
	AssociatedVirtualServiceIds() *[]*string
	SetAssociatedVirtualServiceIds(val *[]*string)
	AssociatedVirtualServiceIdsInput() *[]*string
	AssociatedVirtualServices() *[]*string
	CaCertificateIds() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CnCheckEnabled() cdktf.IResolvable
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultPort() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DomainNames() *[]*string
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
	Enabled() cdktf.IResolvable
	EnabledMemberCount() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GracefulTimeoutPeriod() *float64
	HealthMessage() *string
	HealthMonitor() DataVcdNsxtAlbPoolHealthMonitorList
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Member() DataVcdNsxtAlbPoolMemberList
	MemberCount() *float64
	MemberGroupId() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	PassiveMonitoringEnabled() interface{}
	SetPassiveMonitoringEnabled(val interface{})
	PassiveMonitoringEnabledInput() interface{}
	PersistenceProfile() DataVcdNsxtAlbPoolPersistenceProfileList
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
	UpMemberCount() *float64
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
	ResetAssociatedVirtualServiceIds()
	ResetId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassiveMonitoringEnabled()
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

// The jsii proxy struct for DataVcdNsxtAlbPool
type jsiiProxy_DataVcdNsxtAlbPool struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) AssociatedVirtualServiceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedVirtualServiceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) AssociatedVirtualServiceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedVirtualServiceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) AssociatedVirtualServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) CaCertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"caCertificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) CnCheckEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cnCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) EnabledMemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledMemberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) GracefulTimeoutPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracefulTimeoutPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) HealthMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) HealthMonitor() DataVcdNsxtAlbPoolHealthMonitorList {
	var returns DataVcdNsxtAlbPoolHealthMonitorList
	_jsii_.Get(
		j,
		"healthMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Member() DataVcdNsxtAlbPoolMemberList {
	var returns DataVcdNsxtAlbPoolMemberList
	_jsii_.Get(
		j,
		"member",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) MemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) MemberGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) PassiveMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passiveMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) PassiveMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passiveMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) PersistenceProfile() DataVcdNsxtAlbPoolPersistenceProfileList {
	var returns DataVcdNsxtAlbPoolPersistenceProfileList
	_jsii_.Get(
		j,
		"persistenceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) UpMemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upMemberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtAlbPool) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool vcd_nsxt_alb_pool} Data Source.
func NewDataVcdNsxtAlbPool(scope constructs.Construct, id *string, config *DataVcdNsxtAlbPoolConfig) DataVcdNsxtAlbPool {
	_init_.Initialize()

	if err := validateNewDataVcdNsxtAlbPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdNsxtAlbPool{}

	_jsii_.Create(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool vcd_nsxt_alb_pool} Data Source.
func NewDataVcdNsxtAlbPool_Override(d DataVcdNsxtAlbPool, scope constructs.Construct, id *string, config *DataVcdNsxtAlbPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetAssociatedVirtualServiceIds(val *[]*string) {
	if err := j.validateSetAssociatedVirtualServiceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedVirtualServiceIds",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetPassiveMonitoringEnabled(val interface{}) {
	if err := j.validateSetPassiveMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtAlbPool)SetVdc(val *string) {
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
func DataVcdNsxtAlbPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNsxtAlbPool_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbPool_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdNsxtAlbPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdNsxtAlbPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdNsxtAlbPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdNsxtAlbPool.DataVcdNsxtAlbPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtAlbPool) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetAssociatedVirtualServiceIds() {
	_jsii_.InvokeVoid(
		d,
		"resetAssociatedVirtualServiceIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetPassiveMonitoringEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPassiveMonitoringEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ResetVdc() {
	_jsii_.InvokeVoid(
		d,
		"resetVdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtAlbPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

