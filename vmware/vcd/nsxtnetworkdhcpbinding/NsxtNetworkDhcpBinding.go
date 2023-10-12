package nsxtnetworkdhcpbinding

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/nsxtnetworkdhcpbinding/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding vcd_nsxt_network_dhcp_binding}.
type NsxtNetworkDhcpBinding interface {
	cdktf.TerraformResource
	BindingType() *string
	SetBindingType(val *string)
	BindingTypeInput() *string
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
	DhcpV4Config() NsxtNetworkDhcpBindingDhcpV4ConfigOutputReference
	DhcpV4ConfigInput() *NsxtNetworkDhcpBindingDhcpV4Config
	DhcpV6Config() NsxtNetworkDhcpBindingDhcpV6ConfigOutputReference
	DhcpV6ConfigInput() *NsxtNetworkDhcpBindingDhcpV6Config
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
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	LeaseTime() *float64
	SetLeaseTime(val *float64)
	LeaseTimeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MacAddress() *string
	SetMacAddress(val *string)
	MacAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OrgNetworkId() *string
	SetOrgNetworkId(val *string)
	OrgNetworkIdInput() *string
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
	PutDhcpV4Config(value *NsxtNetworkDhcpBindingDhcpV4Config)
	PutDhcpV6Config(value *NsxtNetworkDhcpBindingDhcpV6Config)
	ResetDescription()
	ResetDhcpV4Config()
	ResetDhcpV6Config()
	ResetDnsServers()
	ResetId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxtNetworkDhcpBinding
type jsiiProxy_NsxtNetworkDhcpBinding struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) BindingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) BindingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DhcpV4Config() NsxtNetworkDhcpBindingDhcpV4ConfigOutputReference {
	var returns NsxtNetworkDhcpBindingDhcpV4ConfigOutputReference
	_jsii_.Get(
		j,
		"dhcpV4Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DhcpV4ConfigInput() *NsxtNetworkDhcpBindingDhcpV4Config {
	var returns *NsxtNetworkDhcpBindingDhcpV4Config
	_jsii_.Get(
		j,
		"dhcpV4ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DhcpV6Config() NsxtNetworkDhcpBindingDhcpV6ConfigOutputReference {
	var returns NsxtNetworkDhcpBindingDhcpV6ConfigOutputReference
	_jsii_.Get(
		j,
		"dhcpV6Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DhcpV6ConfigInput() *NsxtNetworkDhcpBindingDhcpV6Config {
	var returns *NsxtNetworkDhcpBindingDhcpV6Config
	_jsii_.Get(
		j,
		"dhcpV6ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) LeaseTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) LeaseTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) OrgNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) OrgNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding vcd_nsxt_network_dhcp_binding} Resource.
func NewNsxtNetworkDhcpBinding(scope constructs.Construct, id *string, config *NsxtNetworkDhcpBindingConfig) NsxtNetworkDhcpBinding {
	_init_.Initialize()

	if err := validateNewNsxtNetworkDhcpBindingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtNetworkDhcpBinding{}

	_jsii_.Create(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_dhcp_binding vcd_nsxt_network_dhcp_binding} Resource.
func NewNsxtNetworkDhcpBinding_Override(n NsxtNetworkDhcpBinding, scope constructs.Construct, id *string, config *NsxtNetworkDhcpBindingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetBindingType(val *string) {
	if err := j.validateSetBindingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindingType",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetLeaseTime(val *float64) {
	if err := j.validateSetLeaseTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leaseTime",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetOrgNetworkId(val *string) {
	if err := j.validateSetOrgNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgNetworkId",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkDhcpBinding)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func NsxtNetworkDhcpBinding_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcpBinding_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkDhcpBinding_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcpBinding_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkDhcpBinding_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkDhcpBinding_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtNetworkDhcpBinding_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtNetworkDhcpBinding.NsxtNetworkDhcpBinding",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkDhcpBinding) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) PutDhcpV4Config(value *NsxtNetworkDhcpBindingDhcpV4Config) {
	if err := n.validatePutDhcpV4ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDhcpV4Config",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) PutDhcpV6Config(value *NsxtNetworkDhcpBindingDhcpV6Config) {
	if err := n.validatePutDhcpV6ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDhcpV6Config",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetDhcpV4Config() {
	_jsii_.InvokeVoid(
		n,
		"resetDhcpV4Config",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetDhcpV6Config() {
	_jsii_.InvokeVoid(
		n,
		"resetDhcpV6Config",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetDnsServers() {
	_jsii_.InvokeVoid(
		n,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkDhcpBinding) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

