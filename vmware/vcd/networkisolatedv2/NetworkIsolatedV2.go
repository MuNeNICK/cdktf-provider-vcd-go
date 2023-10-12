package networkisolatedv2

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/networkisolatedv2/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2 vcd_network_isolated_v2}.
type NetworkIsolatedV2 interface {
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
	IsShared() interface{}
	SetIsShared(val interface{})
	IsSharedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataEntry() NetworkIsolatedV2MetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	SecondaryStaticIpPool() NetworkIsolatedV2SecondaryStaticIpPoolList
	SecondaryStaticIpPoolInput() interface{}
	StaticIpPool() NetworkIsolatedV2StaticIpPoolList
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
	PutMetadataEntry(value interface{})
	PutSecondaryStaticIpPool(value interface{})
	PutStaticIpPool(value interface{})
	ResetDescription()
	ResetDns1()
	ResetDns2()
	ResetDnsSuffix()
	ResetDualStackEnabled()
	ResetId()
	ResetIsShared()
	ResetMetadata()
	ResetMetadataEntry()
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

// The jsii proxy struct for NetworkIsolatedV2
type jsiiProxy_NetworkIsolatedV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkIsolatedV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Dns1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Dns1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Dns2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Dns2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DnsSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DnsSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DualStackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dualStackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) DualStackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dualStackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) IsShared() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isShared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) IsSharedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSharedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) MetadataEntry() NetworkIsolatedV2MetadataEntryList {
	var returns NetworkIsolatedV2MetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) OwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryPrefixLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryPrefixLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryPrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryStaticIpPool() NetworkIsolatedV2SecondaryStaticIpPoolList {
	var returns NetworkIsolatedV2SecondaryStaticIpPoolList
	_jsii_.Get(
		j,
		"secondaryStaticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) SecondaryStaticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryStaticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) StaticIpPool() NetworkIsolatedV2StaticIpPoolList {
	var returns NetworkIsolatedV2StaticIpPoolList
	_jsii_.Get(
		j,
		"staticIpPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) StaticIpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIsolatedV2) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2 vcd_network_isolated_v2} Resource.
func NewNetworkIsolatedV2(scope constructs.Construct, id *string, config *NetworkIsolatedV2Config) NetworkIsolatedV2 {
	_init_.Initialize()

	if err := validateNewNetworkIsolatedV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkIsolatedV2{}

	_jsii_.Create(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated_v2 vcd_network_isolated_v2} Resource.
func NewNetworkIsolatedV2_Override(n NetworkIsolatedV2, scope constructs.Construct, id *string, config *NetworkIsolatedV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDns1(val *string) {
	if err := j.validateSetDns1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns1",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDns2(val *string) {
	if err := j.validateSetDns2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns2",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDnsSuffix(val *string) {
	if err := j.validateSetDnsSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSuffix",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetDualStackEnabled(val interface{}) {
	if err := j.validateSetDualStackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dualStackEnabled",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetIsShared(val interface{}) {
	if err := j.validateSetIsSharedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isShared",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetOwnerId(val *string) {
	if err := j.validateSetOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetPrefixLength(val *float64) {
	if err := j.validateSetPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixLength",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetSecondaryGateway(val *string) {
	if err := j.validateSetSecondaryGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryGateway",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetSecondaryPrefixLength(val *string) {
	if err := j.validateSetSecondaryPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrefixLength",
		val,
	)
}

func (j *jsiiProxy_NetworkIsolatedV2)SetVdc(val *string) {
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
func NetworkIsolatedV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkIsolatedV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkIsolatedV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkIsolatedV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkIsolatedV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.networkIsolatedV2.NetworkIsolatedV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkIsolatedV2) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkIsolatedV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkIsolatedV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkIsolatedV2) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) PutMetadataEntry(value interface{}) {
	if err := n.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) PutSecondaryStaticIpPool(value interface{}) {
	if err := n.validatePutSecondaryStaticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSecondaryStaticIpPool",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) PutStaticIpPool(value interface{}) {
	if err := n.validatePutStaticIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStaticIpPool",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetDns1() {
	_jsii_.InvokeVoid(
		n,
		"resetDns1",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetDns2() {
	_jsii_.InvokeVoid(
		n,
		"resetDns2",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetDnsSuffix() {
	_jsii_.InvokeVoid(
		n,
		"resetDnsSuffix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetDualStackEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetDualStackEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetIsShared() {
	_jsii_.InvokeVoid(
		n,
		"resetIsShared",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetOwnerId() {
	_jsii_.InvokeVoid(
		n,
		"resetOwnerId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetSecondaryGateway() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryGateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetSecondaryPrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryPrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetSecondaryStaticIpPool() {
	_jsii_.InvokeVoid(
		n,
		"resetSecondaryStaticIpPool",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetStaticIpPool() {
	_jsii_.InvokeVoid(
		n,
		"resetStaticIpPool",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIsolatedV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIsolatedV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

