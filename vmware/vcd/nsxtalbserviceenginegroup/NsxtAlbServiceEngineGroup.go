package nsxtalbserviceenginegroup

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtalbserviceenginegroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group vcd_nsxt_alb_service_engine_group}.
type NsxtAlbServiceEngineGroup interface {
	cdktf.TerraformResource
	AlbCloudId() *string
	SetAlbCloudId(val *string)
	AlbCloudIdInput() *string
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
	DeployedVirtualServices() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaMode() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImportableServiceEngineGroupName() *string
	SetImportableServiceEngineGroupName(val *string)
	ImportableServiceEngineGroupNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxVirtualServices() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Overallocated() interface{}
	SetOverallocated(val interface{})
	OverallocatedInput() interface{}
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
	ReservationModel() *string
	SetReservationModel(val *string)
	ReservationModelInput() *string
	ReservedVirtualServices() *float64
	SupportedFeatureSet() *string
	SetSupportedFeatureSet(val *string)
	SupportedFeatureSetInput() *string
	SyncOnRefresh() interface{}
	SetSyncOnRefresh(val interface{})
	SyncOnRefreshInput() interface{}
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
	ResetDescription()
	ResetId()
	ResetOverallocated()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSupportedFeatureSet()
	ResetSyncOnRefresh()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxtAlbServiceEngineGroup
type jsiiProxy_NsxtAlbServiceEngineGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) AlbCloudId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albCloudId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) AlbCloudIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"albCloudIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) DeployedVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) HaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ImportableServiceEngineGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importableServiceEngineGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ImportableServiceEngineGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importableServiceEngineGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) MaxVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Overallocated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overallocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) OverallocatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overallocatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ReservationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ReservationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) ReservedVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) SupportedFeatureSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedFeatureSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) SupportedFeatureSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedFeatureSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) SyncOnRefresh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) SyncOnRefreshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group vcd_nsxt_alb_service_engine_group} Resource.
func NewNsxtAlbServiceEngineGroup(scope constructs.Construct, id *string, config *NsxtAlbServiceEngineGroupConfig) NsxtAlbServiceEngineGroup {
	_init_.Initialize()

	if err := validateNewNsxtAlbServiceEngineGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtAlbServiceEngineGroup{}

	_jsii_.Create(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group vcd_nsxt_alb_service_engine_group} Resource.
func NewNsxtAlbServiceEngineGroup_Override(n NsxtAlbServiceEngineGroup, scope constructs.Construct, id *string, config *NsxtAlbServiceEngineGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetAlbCloudId(val *string) {
	if err := j.validateSetAlbCloudIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"albCloudId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetImportableServiceEngineGroupName(val *string) {
	if err := j.validateSetImportableServiceEngineGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importableServiceEngineGroupName",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetOverallocated(val interface{}) {
	if err := j.validateSetOverallocatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overallocated",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetReservationModel(val *string) {
	if err := j.validateSetReservationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationModel",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetSupportedFeatureSet(val *string) {
	if err := j.validateSetSupportedFeatureSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedFeatureSet",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbServiceEngineGroup)SetSyncOnRefresh(val interface{}) {
	if err := j.validateSetSyncOnRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncOnRefresh",
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
func NsxtAlbServiceEngineGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbServiceEngineGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbServiceEngineGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbServiceEngineGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbServiceEngineGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbServiceEngineGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtAlbServiceEngineGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtAlbServiceEngineGroup.NsxtAlbServiceEngineGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetOverallocated() {
	_jsii_.InvokeVoid(
		n,
		"resetOverallocated",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetSupportedFeatureSet() {
	_jsii_.InvokeVoid(
		n,
		"resetSupportedFeatureSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ResetSyncOnRefresh() {
	_jsii_.InvokeVoid(
		n,
		"resetSyncOnRefresh",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbServiceEngineGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

