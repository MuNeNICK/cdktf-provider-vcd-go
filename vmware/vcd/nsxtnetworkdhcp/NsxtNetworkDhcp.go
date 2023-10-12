package nsxtnetworkdhcp

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtnetworkdhcp/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp vcd_nsxt_network_dhcp}.
type NsxtNetworkDhcp interface {
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
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
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
	LeaseTime() *float64
	SetLeaseTime(val *float64)
	LeaseTimeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenerIpAddress() *string
	SetListenerIpAddress(val *string)
	ListenerIpAddressInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OrgNetworkId() *string
	SetOrgNetworkId(val *string)
	OrgNetworkIdInput() *string
	Pool() NsxtNetworkDhcpPoolList
	PoolInput() interface{}
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
	PutPool(value interface{})
	ResetDnsServers()
	ResetId()
	ResetLeaseTime()
	ResetListenerIpAddress()
	ResetMode()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPool()
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

// The jsii proxy struct for NsxtNetworkDhcp
type jsiiProxy_NsxtNetworkDhcp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtNetworkDhcp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) LeaseTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) LeaseTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) ListenerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) ListenerIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) OrgNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) OrgNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Pool() NsxtNetworkDhcpPoolList {
	var returns NsxtNetworkDhcpPoolList
	_jsii_.Get(
		j,
		"pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) PoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"poolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcp) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp vcd_nsxt_network_dhcp} Resource.
func NewNsxtNetworkDhcp(scope constructs.Construct, id *string, config *NsxtNetworkDhcpConfig) NsxtNetworkDhcp {
	_init_.Initialize()

	if err := validateNewNsxtNetworkDhcpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtNetworkDhcp{}

	_jsii_.Create(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp vcd_nsxt_network_dhcp} Resource.
func NewNsxtNetworkDhcp_Override(n NsxtNetworkDhcp, scope constructs.Construct, id *string, config *NsxtNetworkDhcpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetLeaseTime(val *float64) {
	if err := j.validateSetLeaseTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leaseTime",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetListenerIpAddress(val *string) {
	if err := j.validateSetListenerIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenerIpAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetOrgNetworkId(val *string) {
	if err := j.validateSetOrgNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgNetworkId",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcp)SetVdc(val *string) {
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
func NsxtNetworkDhcp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkDhcp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkDhcp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtNetworkDhcp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtNetworkDhcp.NsxtNetworkDhcp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcp) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtNetworkDhcp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtNetworkDhcp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkDhcp) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) PutPool(value interface{}) {
	if err := n.validatePutPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPool",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetDnsServers() {
	_jsii_.InvokeVoid(
		n,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetLeaseTime() {
	_jsii_.InvokeVoid(
		n,
		"resetLeaseTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetListenerIpAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetListenerIpAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetMode() {
	_jsii_.InvokeVoid(
		n,
		"resetMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetPool() {
	_jsii_.InvokeVoid(
		n,
		"resetPool",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

