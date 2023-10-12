package nsxtalbvirtualservice

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/nsxtalbvirtualservice/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service vcd_nsxt_alb_virtual_service}.
type NsxtAlbVirtualService interface {
	cdktf.TerraformResource
	ApplicationProfileType() *string
	SetApplicationProfileType(val *string)
	ApplicationProfileTypeInput() *string
	CaCertificateId() *string
	SetCaCertificateId(val *string)
	CaCertificateIdInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ipv6VirtualIpAddress() *string
	SetIpv6VirtualIpAddress(val *string)
	Ipv6VirtualIpAddressInput() *string
	IsTransparentModeEnabled() interface{}
	SetIsTransparentModeEnabled(val interface{})
	IsTransparentModeEnabledInput() interface{}
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
	PoolId() *string
	SetPoolId(val *string)
	PoolIdInput() *string
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
	ServiceEngineGroupId() *string
	SetServiceEngineGroupId(val *string)
	ServiceEngineGroupIdInput() *string
	ServicePort() NsxtAlbVirtualServiceServicePortList
	ServicePortInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Vdc() *string
	SetVdc(val *string)
	VdcInput() *string
	VirtualIpAddress() *string
	SetVirtualIpAddress(val *string)
	VirtualIpAddressInput() *string
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
	PutServicePort(value interface{})
	ResetCaCertificateId()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetIpv6VirtualIpAddress()
	ResetIsTransparentModeEnabled()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServicePort()
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

// The jsii proxy struct for NsxtAlbVirtualService
type jsiiProxy_NsxtAlbVirtualService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtAlbVirtualService) ApplicationProfileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationProfileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ApplicationProfileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationProfileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) CaCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) CaCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Ipv6VirtualIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6VirtualIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Ipv6VirtualIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6VirtualIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) IsTransparentModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransparentModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) IsTransparentModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransparentModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) PoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) PoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ServiceEngineGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ServiceEngineGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEngineGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ServicePort() NsxtAlbVirtualServiceServicePortList {
	var returns NsxtAlbVirtualServiceServicePortList
	_jsii_.Get(
		j,
		"servicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) ServicePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servicePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) VirtualIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbVirtualService) VirtualIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualIpAddressInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service vcd_nsxt_alb_virtual_service} Resource.
func NewNsxtAlbVirtualService(scope constructs.Construct, id *string, config *NsxtAlbVirtualServiceConfig) NsxtAlbVirtualService {
	_init_.Initialize()

	if err := validateNewNsxtAlbVirtualServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtAlbVirtualService{}

	_jsii_.Create(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_virtual_service vcd_nsxt_alb_virtual_service} Resource.
func NewNsxtAlbVirtualService_Override(n NsxtAlbVirtualService, scope constructs.Construct, id *string, config *NsxtAlbVirtualServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetApplicationProfileType(val *string) {
	if err := j.validateSetApplicationProfileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationProfileType",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetCaCertificateId(val *string) {
	if err := j.validateSetCaCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetIpv6VirtualIpAddress(val *string) {
	if err := j.validateSetIpv6VirtualIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6VirtualIpAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetIsTransparentModeEnabled(val interface{}) {
	if err := j.validateSetIsTransparentModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTransparentModeEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetPoolId(val *string) {
	if err := j.validateSetPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetServiceEngineGroupId(val *string) {
	if err := j.validateSetServiceEngineGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEngineGroupId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetVdc(val *string) {
	if err := j.validateSetVdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdc",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbVirtualService)SetVirtualIpAddress(val *string) {
	if err := j.validateSetVirtualIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualIpAddress",
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
func NsxtAlbVirtualService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbVirtualService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbVirtualService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbVirtualService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbVirtualService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbVirtualService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtAlbVirtualService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtAlbVirtualService.NsxtAlbVirtualService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtAlbVirtualService) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtAlbVirtualService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtAlbVirtualService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbVirtualService) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) PutServicePort(value interface{}) {
	if err := n.validatePutServicePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putServicePort",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetCaCertificateId() {
	_jsii_.InvokeVoid(
		n,
		"resetCaCertificateId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetIpv6VirtualIpAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv6VirtualIpAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetIsTransparentModeEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetIsTransparentModeEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetServicePort() {
	_jsii_.InvokeVoid(
		n,
		"resetServicePort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbVirtualService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbVirtualService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbVirtualService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbVirtualService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

