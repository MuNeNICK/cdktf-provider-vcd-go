package providervdc

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/providervdc/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc vcd_provider_vdc}.
type ProviderVdc interface {
	cdktf.TerraformResource
	Capabilities() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeCapacity() ProviderVdcComputeCapacityList
	ComputeProviderScope() *string
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
	ExternalNetworkIds() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HighestSupportedHardwareVersion() *string
	SetHighestSupportedHardwareVersion(val *string)
	HighestSupportedHardwareVersionInput() *string
	HostIds() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkPoolIds() *[]*string
	SetNetworkPoolIds(val *[]*string)
	NetworkPoolIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	NsxtManagerId() *string
	SetNsxtManagerId(val *string)
	NsxtManagerIdInput() *string
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
	ResourcePoolIds() *[]*string
	SetResourcePoolIds(val *[]*string)
	ResourcePoolIdsInput() *[]*string
	Status() *float64
	StorageContainerIds() *[]*string
	StorageProfileIds() *[]*string
	StorageProfileNames() *[]*string
	SetStorageProfileNames(val *[]*string)
	StorageProfileNamesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UniversalNetworkPoolId() *string
	VcenterId() *string
	SetVcenterId(val *string)
	VcenterIdInput() *string
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
	ResetIsEnabled()
	ResetNetworkPoolIds()
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

// The jsii proxy struct for ProviderVdc
type jsiiProxy_ProviderVdc struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProviderVdc) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ComputeCapacity() ProviderVdcComputeCapacityList {
	var returns ProviderVdcComputeCapacityList
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ComputeProviderScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeProviderScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ExternalNetworkIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalNetworkIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) HighestSupportedHardwareVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highestSupportedHardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) HighestSupportedHardwareVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highestSupportedHardwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) HostIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) NetworkPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) NetworkPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) NsxtManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtManagerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) NsxtManagerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtManagerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ResourcePoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) ResourcePoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) StorageContainerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageContainerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) StorageProfileIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageProfileIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) StorageProfileNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageProfileNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) StorageProfileNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageProfileNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) UniversalNetworkPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universalNetworkPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) VcenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProviderVdc) VcenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc vcd_provider_vdc} Resource.
func NewProviderVdc(scope constructs.Construct, id *string, config *ProviderVdcConfig) ProviderVdc {
	_init_.Initialize()

	if err := validateNewProviderVdcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProviderVdc{}

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc vcd_provider_vdc} Resource.
func NewProviderVdc_Override(p ProviderVdc, scope constructs.Construct, id *string, config *ProviderVdcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.providerVdc.ProviderVdc",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProviderVdc)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetHighestSupportedHardwareVersion(val *string) {
	if err := j.validateSetHighestSupportedHardwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highestSupportedHardwareVersion",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetNetworkPoolIds(val *[]*string) {
	if err := j.validateSetNetworkPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPoolIds",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetNsxtManagerId(val *string) {
	if err := j.validateSetNsxtManagerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtManagerId",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetResourcePoolIds(val *[]*string) {
	if err := j.validateSetResourcePoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePoolIds",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetStorageProfileNames(val *[]*string) {
	if err := j.validateSetStorageProfileNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfileNames",
		val,
	)
}

func (j *jsiiProxy_ProviderVdc)SetVcenterId(val *string) {
	if err := j.validateSetVcenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcenterId",
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
func ProviderVdc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProviderVdc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.providerVdc.ProviderVdc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProviderVdc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProviderVdc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.providerVdc.ProviderVdc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProviderVdc_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProviderVdc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.providerVdc.ProviderVdc",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProviderVdc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.providerVdc.ProviderVdc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProviderVdc) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProviderVdc) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProviderVdc) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProviderVdc) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProviderVdc) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProviderVdc) ResetNetworkPoolIds() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkPoolIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProviderVdc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProviderVdc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProviderVdc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

