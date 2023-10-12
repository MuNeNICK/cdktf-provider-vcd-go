package vm

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/vm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm vcd_vm}.
type Vm interface {
	cdktf.TerraformResource
	AcceptAllEulas() interface{}
	SetAcceptAllEulas(val interface{})
	AcceptAllEulasInput() interface{}
	BootImage() *string
	SetBootImage(val *string)
	BootImageId() *string
	SetBootImageId(val *string)
	BootImageIdInput() *string
	BootImageInput() *string
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputerName() *string
	SetComputerName(val *string)
	ComputerNameInput() *string
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
	CpuCores() *float64
	SetCpuCores(val *float64)
	CpuCoresInput() *float64
	CpuHotAddEnabled() interface{}
	SetCpuHotAddEnabled(val interface{})
	CpuHotAddEnabledInput() interface{}
	CpuLimit() *float64
	SetCpuLimit(val *float64)
	CpuLimitInput() *float64
	CpuPriority() *string
	SetCpuPriority(val *string)
	CpuPriorityInput() *string
	CpuReservation() *float64
	SetCpuReservation(val *float64)
	CpuReservationInput() *float64
	Cpus() *float64
	SetCpus(val *float64)
	CpuShares() *float64
	SetCpuShares(val *float64)
	CpuSharesInput() *float64
	CpusInput() *float64
	Customization() VmCustomizationOutputReference
	CustomizationInput() *VmCustomization
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disk() VmDiskList
	DiskInput() interface{}
	ExposeHardwareVirtualization() interface{}
	SetExposeHardwareVirtualization(val interface{})
	ExposeHardwareVirtualizationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestProperties() *map[string]*string
	SetGuestProperties(val *map[string]*string)
	GuestPropertiesInput() *map[string]*string
	HardwareVersion() *string
	SetHardwareVersion(val *string)
	HardwareVersionInput() *string
	Href() *string
	SetHref(val *string)
	HrefInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalDisk() VmInternalDiskList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *float64
	SetMemory(val *float64)
	MemoryHotAddEnabled() interface{}
	SetMemoryHotAddEnabled(val interface{})
	MemoryHotAddEnabledInput() interface{}
	MemoryInput() *float64
	MemoryLimit() *float64
	SetMemoryLimit(val *float64)
	MemoryLimitInput() *float64
	MemoryPriority() *string
	SetMemoryPriority(val *string)
	MemoryPriorityInput() *string
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MemoryShares() *float64
	SetMemoryShares(val *float64)
	MemorySharesInput() *float64
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataEntry() VmMetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() VmNetworkList
	NetworkDhcpWaitSeconds() *float64
	SetNetworkDhcpWaitSeconds(val *float64)
	NetworkDhcpWaitSecondsInput() *float64
	NetworkInput() interface{}
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	OverrideTemplateDisk() VmOverrideTemplateDiskList
	OverrideTemplateDiskInput() interface{}
	PlacementPolicyId() *string
	SetPlacementPolicyId(val *string)
	PlacementPolicyIdInput() *string
	PowerOn() interface{}
	SetPowerOn(val interface{})
	PowerOnInput() interface{}
	PreventUpdatePowerOff() interface{}
	SetPreventUpdatePowerOff(val interface{})
	PreventUpdatePowerOffInput() interface{}
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
	SecurityTags() *[]*string
	SetSecurityTags(val *[]*string)
	SecurityTagsInput() *[]*string
	SizingPolicyId() *string
	SetSizingPolicyId(val *string)
	SizingPolicyIdInput() *string
	Status() *float64
	StatusText() *string
	StorageProfile() *string
	SetStorageProfile(val *string)
	StorageProfileInput() *string
	TemplateName() *string
	SetTemplateName(val *string)
	TemplateNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VappName() *string
	SetVappName(val *string)
	VappNameInput() *string
	VappTemplateId() *string
	SetVappTemplateId(val *string)
	VappTemplateIdInput() *string
	Vdc() *string
	SetVdc(val *string)
	VdcInput() *string
	VmNameInTemplate() *string
	SetVmNameInTemplate(val *string)
	VmNameInTemplateInput() *string
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
	PutCustomization(value *VmCustomization)
	PutDisk(value interface{})
	PutMetadataEntry(value interface{})
	PutNetwork(value interface{})
	PutOverrideTemplateDisk(value interface{})
	ResetAcceptAllEulas()
	ResetBootImage()
	ResetBootImageId()
	ResetCatalogName()
	ResetComputerName()
	ResetCpuCores()
	ResetCpuHotAddEnabled()
	ResetCpuLimit()
	ResetCpuPriority()
	ResetCpuReservation()
	ResetCpus()
	ResetCpuShares()
	ResetCustomization()
	ResetDescription()
	ResetDisk()
	ResetExposeHardwareVirtualization()
	ResetGuestProperties()
	ResetHardwareVersion()
	ResetHref()
	ResetId()
	ResetMemory()
	ResetMemoryHotAddEnabled()
	ResetMemoryLimit()
	ResetMemoryPriority()
	ResetMemoryReservation()
	ResetMemoryShares()
	ResetMetadata()
	ResetMetadataEntry()
	ResetNetwork()
	ResetNetworkDhcpWaitSeconds()
	ResetOrg()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrideTemplateDisk()
	ResetPlacementPolicyId()
	ResetPowerOn()
	ResetPreventUpdatePowerOff()
	ResetSecurityTags()
	ResetSizingPolicyId()
	ResetStorageProfile()
	ResetTemplateName()
	ResetVappName()
	ResetVappTemplateId()
	ResetVdc()
	ResetVmNameInTemplate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Vm
type jsiiProxy_Vm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Vm) AcceptAllEulas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAllEulas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) AcceptAllEulasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAllEulasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) BootImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) BootImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) BootImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) BootImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuCoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Cpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpuSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Customization() VmCustomizationOutputReference {
	var returns VmCustomizationOutputReference
	_jsii_.Get(
		j,
		"customization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) CustomizationInput() *VmCustomization {
	var returns *VmCustomization
	_jsii_.Get(
		j,
		"customizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Disk() VmDiskList {
	var returns VmDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ExposeHardwareVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeHardwareVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ExposeHardwareVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeHardwareVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) GuestProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"guestProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) GuestPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"guestPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) HardwareVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) HardwareVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) HrefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hrefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) InternalDisk() VmInternalDiskList {
	var returns VmInternalDiskList
	_jsii_.Get(
		j,
		"internalDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemoryShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MemorySharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MetadataEntry() VmMetadataEntryList {
	var returns VmMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Network() VmNetworkList {
	var returns VmNetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) NetworkDhcpWaitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) NetworkDhcpWaitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) OverrideTemplateDisk() VmOverrideTemplateDiskList {
	var returns VmOverrideTemplateDiskList
	_jsii_.Get(
		j,
		"overrideTemplateDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) OverrideTemplateDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideTemplateDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PlacementPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PlacementPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PowerOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PowerOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PreventUpdatePowerOff() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUpdatePowerOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) PreventUpdatePowerOffInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUpdatePowerOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) SecurityTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) SecurityTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) SizingPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) SizingPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) StatusText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) StorageProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) StorageProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) TemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VappName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VappNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VappTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VappTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VmNameInTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNameInTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VmNameInTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNameInTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vm) VmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm vcd_vm} Resource.
func NewVm(scope constructs.Construct, id *string, config *VmConfig) Vm {
	_init_.Initialize()

	if err := validateNewVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Vm{}

	_jsii_.Create(
		"vcd.vm.Vm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm vcd_vm} Resource.
func NewVm_Override(v Vm, scope constructs.Construct, id *string, config *VmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vm.Vm",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_Vm)SetAcceptAllEulas(val interface{}) {
	if err := j.validateSetAcceptAllEulasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptAllEulas",
		val,
	)
}

func (j *jsiiProxy_Vm)SetBootImage(val *string) {
	if err := j.validateSetBootImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImage",
		val,
	)
}

func (j *jsiiProxy_Vm)SetBootImageId(val *string) {
	if err := j.validateSetBootImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImageId",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_Vm)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_Vm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuCores(val *float64) {
	if err := j.validateSetCpuCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCores",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuHotAddEnabled(val interface{}) {
	if err := j.validateSetCpuHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuPriority(val *string) {
	if err := j.validateSetCpuPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPriority",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpus(val *float64) {
	if err := j.validateSetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpus",
		val,
	)
}

func (j *jsiiProxy_Vm)SetCpuShares(val *float64) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_Vm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Vm)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Vm)SetExposeHardwareVirtualization(val interface{}) {
	if err := j.validateSetExposeHardwareVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeHardwareVirtualization",
		val,
	)
}

func (j *jsiiProxy_Vm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Vm)SetGuestProperties(val *map[string]*string) {
	if err := j.validateSetGuestPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestProperties",
		val,
	)
}

func (j *jsiiProxy_Vm)SetHardwareVersion(val *string) {
	if err := j.validateSetHardwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardwareVersion",
		val,
	)
}

func (j *jsiiProxy_Vm)SetHref(val *string) {
	if err := j.validateSetHrefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"href",
		val,
	)
}

func (j *jsiiProxy_Vm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Vm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemoryHotAddEnabled(val interface{}) {
	if err := j.validateSetMemoryHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemoryPriority(val *string) {
	if err := j.validateSetMemoryPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPriority",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMemoryShares(val *float64) {
	if err := j.validateSetMemorySharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShares",
		val,
	)
}

func (j *jsiiProxy_Vm)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Vm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Vm)SetNetworkDhcpWaitSeconds(val *float64) {
	if err := j.validateSetNetworkDhcpWaitSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkDhcpWaitSeconds",
		val,
	)
}

func (j *jsiiProxy_Vm)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_Vm)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_Vm)SetPlacementPolicyId(val *string) {
	if err := j.validateSetPlacementPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPolicyId",
		val,
	)
}

func (j *jsiiProxy_Vm)SetPowerOn(val interface{}) {
	if err := j.validateSetPowerOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerOn",
		val,
	)
}

func (j *jsiiProxy_Vm)SetPreventUpdatePowerOff(val interface{}) {
	if err := j.validateSetPreventUpdatePowerOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUpdatePowerOff",
		val,
	)
}

func (j *jsiiProxy_Vm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Vm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Vm)SetSecurityTags(val *[]*string) {
	if err := j.validateSetSecurityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityTags",
		val,
	)
}

func (j *jsiiProxy_Vm)SetSizingPolicyId(val *string) {
	if err := j.validateSetSizingPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicyId",
		val,
	)
}

func (j *jsiiProxy_Vm)SetStorageProfile(val *string) {
	if err := j.validateSetStorageProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfile",
		val,
	)
}

func (j *jsiiProxy_Vm)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_Vm)SetVappName(val *string) {
	if err := j.validateSetVappNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappName",
		val,
	)
}

func (j *jsiiProxy_Vm)SetVappTemplateId(val *string) {
	if err := j.validateSetVappTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappTemplateId",
		val,
	)
}

func (j *jsiiProxy_Vm)SetVdc(val *string) {
	if err := j.validateSetVdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdc",
		val,
	)
}

func (j *jsiiProxy_Vm)SetVmNameInTemplate(val *string) {
	if err := j.validateSetVmNameInTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmNameInTemplate",
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
func Vm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vm.Vm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Vm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vm.Vm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Vm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vm.Vm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Vm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.vm.Vm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_Vm) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_Vm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_Vm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_Vm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_Vm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_Vm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_Vm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_Vm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_Vm) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_Vm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_Vm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_Vm) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_Vm) PutCustomization(value *VmCustomization) {
	if err := v.validatePutCustomizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomization",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Vm) PutDisk(value interface{}) {
	if err := v.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Vm) PutMetadataEntry(value interface{}) {
	if err := v.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Vm) PutNetwork(value interface{}) {
	if err := v.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetwork",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Vm) PutOverrideTemplateDisk(value interface{}) {
	if err := v.validatePutOverrideTemplateDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOverrideTemplateDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_Vm) ResetAcceptAllEulas() {
	_jsii_.InvokeVoid(
		v,
		"resetAcceptAllEulas",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetBootImage() {
	_jsii_.InvokeVoid(
		v,
		"resetBootImage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetBootImageId() {
	_jsii_.InvokeVoid(
		v,
		"resetBootImageId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCatalogName() {
	_jsii_.InvokeVoid(
		v,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetComputerName() {
	_jsii_.InvokeVoid(
		v,
		"resetComputerName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuCores() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuCores",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuPriority() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuPriority",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpus() {
	_jsii_.InvokeVoid(
		v,
		"resetCpus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCpuShares() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetCustomization() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomization",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetExposeHardwareVirtualization() {
	_jsii_.InvokeVoid(
		v,
		"resetExposeHardwareVirtualization",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetGuestProperties() {
	_jsii_.InvokeVoid(
		v,
		"resetGuestProperties",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetHardwareVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetHardwareVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetHref() {
	_jsii_.InvokeVoid(
		v,
		"resetHref",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemory() {
	_jsii_.InvokeVoid(
		v,
		"resetMemory",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemoryHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemoryPriority() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryPriority",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMemoryShares() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMetadata() {
	_jsii_.InvokeVoid(
		v,
		"resetMetadata",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		v,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetNetworkDhcpWaitSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkDhcpWaitSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetOrg() {
	_jsii_.InvokeVoid(
		v,
		"resetOrg",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetOsType() {
	_jsii_.InvokeVoid(
		v,
		"resetOsType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetOverrideTemplateDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideTemplateDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetPlacementPolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetPlacementPolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetPowerOn() {
	_jsii_.InvokeVoid(
		v,
		"resetPowerOn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetPreventUpdatePowerOff() {
	_jsii_.InvokeVoid(
		v,
		"resetPreventUpdatePowerOff",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetSecurityTags() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetSizingPolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetSizingPolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetStorageProfile() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageProfile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetTemplateName() {
	_jsii_.InvokeVoid(
		v,
		"resetTemplateName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetVappName() {
	_jsii_.InvokeVoid(
		v,
		"resetVappName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetVappTemplateId() {
	_jsii_.InvokeVoid(
		v,
		"resetVappTemplateId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetVdc() {
	_jsii_.InvokeVoid(
		v,
		"resetVdc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) ResetVmNameInTemplate() {
	_jsii_.InvokeVoid(
		v,
		"resetVmNameInTemplate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Vm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

