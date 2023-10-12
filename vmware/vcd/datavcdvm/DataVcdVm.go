package datavcdvm

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/datavcdvm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm vcd_vm}.
type DataVcdVm interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputerName() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCores() *float64
	CpuHotAddEnabled() cdktf.IResolvable
	CpuLimit() *float64
	CpuPriority() *string
	CpuReservation() *float64
	Cpus() *float64
	CpuShares() *float64
	Customization() DataVcdVmCustomizationList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Disk() DataVcdVmDiskList
	ExposeHardwareVirtualization() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestProperties() cdktf.StringMap
	HardwareVersion() *string
	Href() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalDisk() DataVcdVmInternalDiskList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *float64
	MemoryHotAddEnabled() cdktf.IResolvable
	MemoryLimit() *float64
	MemoryPriority() *string
	MemoryReservation() *float64
	MemoryShares() *float64
	Metadata() cdktf.StringMap
	MetadataEntry() DataVcdVmMetadataEntryList
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() DataVcdVmNetworkList
	NetworkDhcpWaitSeconds() *float64
	SetNetworkDhcpWaitSeconds(val *float64)
	NetworkDhcpWaitSecondsInput() *float64
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OsType() *string
	PlacementPolicyId() *string
	SetPlacementPolicyId(val *string)
	PlacementPolicyIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SecurityTags() *[]*string
	SizingPolicyId() *string
	SetSizingPolicyId(val *string)
	SizingPolicyIdInput() *string
	Status() *float64
	StatusText() *string
	StorageProfile() *string
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
	VmType() *string
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
	ResetNetworkDhcpWaitSeconds()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementPolicyId()
	ResetSizingPolicyId()
	ResetVappName()
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

// The jsii proxy struct for DataVcdVm
type jsiiProxy_DataVcdVm struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVcdVm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuHotAddEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Cpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Customization() DataVcdVmCustomizationList {
	var returns DataVcdVmCustomizationList
	_jsii_.Get(
		j,
		"customization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Disk() DataVcdVmDiskList {
	var returns DataVcdVmDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) ExposeHardwareVirtualization() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"exposeHardwareVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) GuestProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"guestProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) HardwareVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) InternalDisk() DataVcdVmInternalDiskList {
	var returns DataVcdVmInternalDiskList
	_jsii_.Get(
		j,
		"internalDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MemoryHotAddEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MemoryPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MemoryShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) MetadataEntry() DataVcdVmMetadataEntryList {
	var returns DataVcdVmMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Network() DataVcdVmNetworkList {
	var returns DataVcdVmNetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) NetworkDhcpWaitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) NetworkDhcpWaitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) PlacementPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) PlacementPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) SecurityTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) SizingPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) SizingPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) StatusText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) StorageProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) VappName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) VappNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVm) VmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm vcd_vm} Data Source.
func NewDataVcdVm(scope constructs.Construct, id *string, config *DataVcdVmConfig) DataVcdVm {
	_init_.Initialize()

	if err := validateNewDataVcdVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdVm{}

	_jsii_.Create(
		"vcd.dataVcdVm.DataVcdVm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/vm vcd_vm} Data Source.
func NewDataVcdVm_Override(d DataVcdVm, scope constructs.Construct, id *string, config *DataVcdVmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdVm.DataVcdVm",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVcdVm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetNetworkDhcpWaitSeconds(val *float64) {
	if err := j.validateSetNetworkDhcpWaitSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkDhcpWaitSeconds",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetPlacementPolicyId(val *string) {
	if err := j.validateSetPlacementPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPolicyId",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetSizingPolicyId(val *string) {
	if err := j.validateSetSizingPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicyId",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetVappName(val *string) {
	if err := j.validateSetVappNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappName",
		val,
	)
}

func (j *jsiiProxy_DataVcdVm)SetVdc(val *string) {
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
func DataVcdVm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVm.DataVcdVm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdVm_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVm_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVm.DataVcdVm",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVcdVm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVcdVm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.dataVcdVm.DataVcdVm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVcdVm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.dataVcdVm.DataVcdVm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVcdVm) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVcdVm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdVm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdVm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdVm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdVm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdVm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdVm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdVm) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdVm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdVm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdVm) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVcdVm) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetNetworkDhcpWaitSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkDhcpWaitSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetOrg() {
	_jsii_.InvokeVoid(
		d,
		"resetOrg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetPlacementPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPlacementPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetSizingPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetSizingPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetVappName() {
	_jsii_.InvokeVoid(
		d,
		"resetVappName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) ResetVdc() {
	_jsii_.InvokeVoid(
		d,
		"resetVdc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdVm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

