package nsxtalbedgegatewayserviceenginegroup

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtalbedgegatewayserviceenginegroup/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group}.
type NsxtAlbEdgegatewayServiceEngineGroup interface {
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
	SetMaxVirtualServices(val *float64)
	MaxVirtualServicesInput() *float64
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
	ReservedVirtualServices() *string
	SetReservedVirtualServices(val *string)
	ReservedVirtualServicesInput() *string
	ServiceEngineGroupId() *string
	SetServiceEngineGroupId(val *string)
	ServiceEngineGroupIdInput() *string
	ServiceEngineGroupName() *string
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
	ResetMaxVirtualServices()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReservedVirtualServices()
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

// The jsii proxy struct for NsxtAlbEdgegatewayServiceEngineGroup
type jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) DeployedVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) MaxVirtualServices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) MaxVirtualServicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVirtualServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ReservedVirtualServices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ReservedVirtualServicesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedVirtualServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ServiceEngineGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group} Resource.
func NewNsxtAlbEdgegatewayServiceEngineGroup(scope constructs.Construct, id *string, config *NsxtAlbEdgegatewayServiceEngineGroupConfig) NsxtAlbEdgegatewayServiceEngineGroup {
	_init_.Initialize()

	if err := validateNewNsxtAlbEdgegatewayServiceEngineGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup{}

	_jsii_.Create(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group vcd_nsxt_alb_edgegateway_service_engine_group} Resource.
func NewNsxtAlbEdgegatewayServiceEngineGroup_Override(n NsxtAlbEdgegatewayServiceEngineGroup, scope constructs.Construct, id *string, config *NsxtAlbEdgegatewayServiceEngineGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetMaxVirtualServices(val *float64) {
	if err := j.validateSetMaxVirtualServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVirtualServices",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetReservedVirtualServices(val *string) {
	if err := j.validateSetReservedVirtualServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedVirtualServices",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetServiceEngineGroupId(val *string) {
	if err := j.validateSetServiceEngineGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEngineGroupId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup)SetVdc(val *string) {
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
func NsxtAlbEdgegatewayServiceEngineGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbEdgegatewayServiceEngineGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbEdgegatewayServiceEngineGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbEdgegatewayServiceEngineGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbEdgegatewayServiceEngineGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtAlbEdgegatewayServiceEngineGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtAlbEdgegatewayServiceEngineGroup.NsxtAlbEdgegatewayServiceEngineGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetMaxVirtualServices() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxVirtualServices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetReservedVirtualServices() {
	_jsii_.InvokeVoid(
		n,
		"resetReservedVirtualServices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbEdgegatewayServiceEngineGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

