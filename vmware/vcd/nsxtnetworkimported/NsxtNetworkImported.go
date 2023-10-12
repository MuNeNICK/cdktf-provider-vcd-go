package nsxtnetworkimported

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/nsxtnetworkimported/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported vcd_nsxt_network_imported}.
type NsxtNetworkImported interface {
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
	Dns1() *string
	SetDns1(val *string)
	Dns1Input() *string
	Dns2() *string
	SetDns2(val *string)
	Dns2Input() *string
	DnsSuffix() *string
	SetDnsSuffix(val *string)
	DnsSuffixInput() *string
	DualStackEnabled() interface{}
	SetDualStackEnabled(val interface{})
	DualStackEnabledInput() interface{}
	DvpgId() *string
	DvpgName() *string
	SetDvpgName(val *string)
	DvpgNameInput() *string
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
	// The tree node.
	Node() constructs.Node
	NsxtLogicalSwitchId() *string
	NsxtLogicalSwitchName() *string
	SetNsxtLogicalSwitchName(val *string)
	NsxtLogicalSwitchNameInput() *string
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OwnerId() *string
	SetOwnerId(val *string)
	OwnerIdInput() *string
	PrefixLength() *float64
	SetPrefixLength(val *float64)
	PrefixLengthInput() *float64
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
	SecondaryGateway() *string
	SetSecondaryGateway(val *string)
	SecondaryGatewayInput() *string
	SecondaryPrefixLength() *string
	SetSecondaryPrefixLength(val *string)
	SecondaryPrefixLengthInput() *string
	SecondaryStaticIpPool() NsxtNetworkImportedSecondaryStaticIpPoolList
	SecondaryStaticIpPoolInput() interface{}
	StaticIpPool() NsxtNetworkImportedStaticIpPoolList
	StaticIpPoolInput() interface{}
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
	PutSecondaryStaticIpPool(value interface{})
	PutStaticIpPool(value interface{})
	ResetDescription()
	ResetDns1()
	ResetDns2()
	ResetDnsSuffix()
	ResetDualStackEnabled()
	ResetDvpgName()
	ResetId()
	ResetNsxtLogicalSwitchName()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerId()
	ResetSecondaryGateway()
	ResetSecondaryPrefixLength()
	ResetSecondaryStaticIpPool()
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

// The jsii proxy struct for NsxtNetworkImported
type jsiiProxy_NsxtNetworkImported struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtNetworkImported) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Dns1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Dns1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Dns2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Dns2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DnsSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DnsSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DualStackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dualStackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DualStackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dualStackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DvpgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvpgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DvpgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvpgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) DvpgNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvpgNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) NsxtLogicalSwitchId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtLogicalSwitchId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) NsxtLogicalSwitchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtLogicalSwitchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) NsxtLogicalSwitchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtLogicalSwitchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) OwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryPrefixLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryPrefixLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryStaticIpPool() NsxtNetworkImportedSecondaryStaticIpPoolList {
	var returns NsxtNetworkImportedSecondaryStaticIpPoolList
	_jsii_.Get(
		j,
		"secondaryStaticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) SecondaryStaticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryStaticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) StaticIpPool() NsxtNetworkImportedStaticIpPoolList {
	var returns NsxtNetworkImportedStaticIpPoolList
	_jsii_.Get(
		j,
		"staticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) StaticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtNetworkImported) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported vcd_nsxt_network_imported} Resource.
func NewNsxtNetworkImported(scope constructs.Construct, id *string, config *NsxtNetworkImportedConfig) NsxtNetworkImported {
	_init_.Initialize()

	if err := validateNewNsxtNetworkImportedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtNetworkImported{}

	_jsii_.Create(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_network_imported vcd_nsxt_network_imported} Resource.
func NewNsxtNetworkImported_Override(n NsxtNetworkImported, scope constructs.Construct, id *string, config *NsxtNetworkImportedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDns1(val *string) {
	if err := j.validateSetDns1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns1",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDns2(val *string) {
	if err := j.validateSetDns2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns2",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDnsSuffix(val *string) {
	if err := j.validateSetDnsSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSuffix",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDualStackEnabled(val interface{}) {
	if err := j.validateSetDualStackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dualStackEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetDvpgName(val *string) {
	if err := j.validateSetDvpgNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvpgName",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetNsxtLogicalSwitchName(val *string) {
	if err := j.validateSetNsxtLogicalSwitchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtLogicalSwitchName",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetOwnerId(val *string) {
	if err := j.validateSetOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetPrefixLength(val *float64) {
	if err := j.validateSetPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixLength",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetSecondaryGateway(val *string) {
	if err := j.validateSetSecondaryGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryGateway",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetSecondaryPrefixLength(val *string) {
	if err := j.validateSetSecondaryPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrefixLength",
		val,
	)
}

func (j *jsiiProxy_NsxtNetworkImported)SetVdc(val *string) {
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
func NsxtNetworkImported_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkImported_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkImported_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkImported_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtNetworkImported_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtNetworkImported_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtNetworkImported_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtNetworkImported.NsxtNetworkImported",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtNetworkImported) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtNetworkImported) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtNetworkImported) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkImported) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtNetworkImported) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtNetworkImported) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtNetworkImported) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtNetworkImported) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtNetworkImported) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtNetworkImported) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtNetworkImported) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtNetworkImported) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtNetworkImported) PutSecondaryStaticIpPool(value interface{}) {
	if err := n.validatePutSecondaryStaticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSecondaryStaticIpPool",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtNetworkImported) PutStaticIpPool(value interface{}) {
	if err := n.validatePutStaticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStaticIpPool",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDns1() {
	_jsii_.InvokeVoid(
		n,
		"resetDns1",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDns2() {
	_jsii_.InvokeVoid(
		n,
		"resetDns2",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDnsSuffix() {
	_jsii_.InvokeVoid(
		n,
		"resetDnsSuffix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDualStackEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetDualStackEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetDvpgName() {
	_jsii_.InvokeVoid(
		n,
		"resetDvpgName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetNsxtLogicalSwitchName() {
	_jsii_.InvokeVoid(
		n,
		"resetNsxtLogicalSwitchName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetOwnerId() {
	_jsii_.InvokeVoid(
		n,
		"resetOwnerId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetSecondaryGateway() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryGateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetSecondaryPrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryPrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetSecondaryStaticIpPool() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryStaticIpPool",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetStaticIpPool() {
	_jsii_.InvokeVoid(
		n,
		"resetStaticIpPool",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtNetworkImported) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkImported) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkImported) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtNetworkImported) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

