package vappnetwork

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/vappnetwork/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network vcd_vapp_network}.
type VappNetwork interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DhcpPool() VappNetworkDhcpPoolList
	DhcpPoolInput() interface{}
	Dns1() *string
	SetDns1(val *string)
	Dns1Input() *string
	Dns2() *string
	SetDns2(val *string)
	Dns2Input() *string
	DnsSuffix() *string
	SetDnsSuffix(val *string)
	DnsSuffixInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gateway() *string
	SetGateway(val *string)
	GatewayInput() *string
	GuestVlanAllowed() interface{}
	SetGuestVlanAllowed(val interface{})
	GuestVlanAllowedInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Netmask() *string
	SetNetmask(val *string)
	NetmaskInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OrgNetworkName() *string
	SetOrgNetworkName(val *string)
	OrgNetworkNameInput() *string
	PrefixLength() *string
	SetPrefixLength(val *string)
	PrefixLengthInput() *string
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
	RebootVappOnRemoval() interface{}
	SetRebootVappOnRemoval(val interface{})
	RebootVappOnRemovalInput() interface{}
	RetainIpMacEnabled() interface{}
	SetRetainIpMacEnabled(val interface{})
	RetainIpMacEnabledInput() interface{}
	StaticIpPool() VappNetworkStaticIpPoolList
	StaticIpPoolInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VappName() *string
	SetVappName(val *string)
	VappNameInput() *string
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
	PutDhcpPool(value interface{})
	PutStaticIpPool(value interface{})
	ResetDescription()
	ResetDhcpPool()
	ResetDns1()
	ResetDns2()
	ResetDnsSuffix()
	ResetGuestVlanAllowed()
	ResetId()
	ResetNetmask()
	ResetOrg()
	ResetOrgNetworkName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrefixLength()
	ResetRebootVappOnRemoval()
	ResetRetainIpMacEnabled()
	ResetStaticIpPool()
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

// The jsii proxy struct for VappNetwork
type jsiiProxy_VappNetwork struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VappNetwork) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DhcpPool() VappNetworkDhcpPoolList {
	var returns VappNetworkDhcpPoolList
	_jsii_.Get(
		j,
		"dhcpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DhcpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Dns1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Dns1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Dns2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Dns2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DnsSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) DnsSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) GuestVlanAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestVlanAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) GuestVlanAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestVlanAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Netmask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) NetmaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) OrgNetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) OrgNetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) PrefixLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) PrefixLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) RebootVappOnRemoval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebootVappOnRemoval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) RebootVappOnRemovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebootVappOnRemovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) RetainIpMacEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainIpMacEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) RetainIpMacEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainIpMacEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) StaticIpPool() VappNetworkStaticIpPoolList {
	var returns VappNetworkStaticIpPoolList
	_jsii_.Get(
		j,
		"staticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) StaticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) VappName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) VappNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappNetwork) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network vcd_vapp_network} Resource.
func NewVappNetwork(scope constructs.Construct, id *string, config *VappNetworkConfig) VappNetwork {
	_init_.Initialize()

	if err := validateNewVappNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappNetwork{}

	_jsii_.Create(
		"vcd.vappNetwork.VappNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_network vcd_vapp_network} Resource.
func NewVappNetwork_Override(v VappNetwork, scope constructs.Construct, id *string, config *VappNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappNetwork.VappNetwork",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VappNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetDns1(val *string) {
	if err := j.validateSetDns1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns1",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetDns2(val *string) {
	if err := j.validateSetDns2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns2",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetDnsSuffix(val *string) {
	if err := j.validateSetDnsSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSuffix",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetGuestVlanAllowed(val interface{}) {
	if err := j.validateSetGuestVlanAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestVlanAllowed",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetNetmask(val *string) {
	if err := j.validateSetNetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netmask",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetOrgNetworkName(val *string) {
	if err := j.validateSetOrgNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgNetworkName",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetPrefixLength(val *string) {
	if err := j.validateSetPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixLength",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetRebootVappOnRemoval(val interface{}) {
	if err := j.validateSetRebootVappOnRemovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebootVappOnRemoval",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetRetainIpMacEnabled(val interface{}) {
	if err := j.validateSetRetainIpMacEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainIpMacEnabled",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetVappName(val *string) {
	if err := j.validateSetVappNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappName",
		val,
	)
}

func (j *jsiiProxy_VappNetwork)SetVdc(val *string) {
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
func VappNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappNetwork.VappNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappNetwork.VappNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappNetwork.VappNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VappNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.vappNetwork.VappNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VappNetwork) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VappNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VappNetwork) PutDhcpPool(value interface{}) {
	if err := v.validatePutDhcpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDhcpPool",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappNetwork) PutStaticIpPool(value interface{}) {
	if err := v.validatePutStaticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStaticIpPool",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappNetwork) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetDhcpPool() {
	_jsii_.InvokeVoid(
		v,
		"resetDhcpPool",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetDns1() {
	_jsii_.InvokeVoid(
		v,
		"resetDns1",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetDns2() {
	_jsii_.InvokeVoid(
		v,
		"resetDns2",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetDnsSuffix() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsSuffix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetGuestVlanAllowed() {
	_jsii_.InvokeVoid(
		v,
		"resetGuestVlanAllowed",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetNetmask() {
	_jsii_.InvokeVoid(
		v,
		"resetNetmask",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetOrg() {
	_jsii_.InvokeVoid(
		v,
		"resetOrg",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetOrgNetworkName() {
	_jsii_.InvokeVoid(
		v,
		"resetOrgNetworkName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetPrefixLength() {
	_jsii_.InvokeVoid(
		v,
		"resetPrefixLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetRebootVappOnRemoval() {
	_jsii_.InvokeVoid(
		v,
		"resetRebootVappOnRemoval",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetRetainIpMacEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetRetainIpMacEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetStaticIpPool() {
	_jsii_.InvokeVoid(
		v,
		"resetStaticIpPool",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) ResetVdc() {
	_jsii_.InvokeVoid(
		v,
		"resetVdc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

