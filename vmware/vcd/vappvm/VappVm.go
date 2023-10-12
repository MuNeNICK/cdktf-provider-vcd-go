package vappvm

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vappvm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm vcd_vapp_vm}.
type VappVm interface {
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
	Customization() VappVmCustomizationOutputReference
	CustomizationInput() *VappVmCustomization
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disk() VappVmDiskList
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
	InternalDisk() VappVmInternalDiskList
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
	MetadataEntry() VappVmMetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() VappVmNetworkList
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
	OverrideTemplateDisk() VappVmOverrideTemplateDiskList
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
	PutCustomization(value *VappVmCustomization)
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

// The jsii proxy struct for VappVm
type jsiiProxy_VappVm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VappVm) AcceptAllEulas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAllEulas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) AcceptAllEulasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptAllEulasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) BootImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) BootImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) BootImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) BootImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuCoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Cpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpuSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Customization() VappVmCustomizationOutputReference {
	var returns VappVmCustomizationOutputReference
	_jsii_.Get(
		j,
		"customization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) CustomizationInput() *VappVmCustomization {
	var returns *VappVmCustomization
	_jsii_.Get(
		j,
		"customizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Disk() VappVmDiskList {
	var returns VappVmDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ExposeHardwareVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeHardwareVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ExposeHardwareVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeHardwareVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) GuestProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"guestProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) GuestPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"guestPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) HardwareVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) HardwareVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) HrefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hrefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) InternalDisk() VappVmInternalDiskList {
	var returns VappVmInternalDiskList
	_jsii_.Get(
		j,
		"internalDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemoryShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MemorySharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MetadataEntry() VappVmMetadataEntryList {
	var returns VappVmMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Network() VappVmNetworkList {
	var returns VappVmNetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) NetworkDhcpWaitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) NetworkDhcpWaitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkDhcpWaitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) OverrideTemplateDisk() VappVmOverrideTemplateDiskList {
	var returns VappVmOverrideTemplateDiskList
	_jsii_.Get(
		j,
		"overrideTemplateDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) OverrideTemplateDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideTemplateDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PlacementPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PlacementPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PowerOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PowerOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powerOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PreventUpdatePowerOff() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUpdatePowerOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) PreventUpdatePowerOffInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUpdatePowerOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) SecurityTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) SecurityTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) SizingPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) SizingPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) StatusText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) StorageProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) StorageProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) TemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VappName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VappNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VappTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VappTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vappTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VmNameInTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNameInTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VmNameInTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmNameInTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVm) VmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm vcd_vapp_vm} Resource.
func NewVappVm(scope constructs.Construct, id *string, config *VappVmConfig) VappVm {
	_init_.Initialize()

	if err := validateNewVappVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappVm{}

	_jsii_.Create(
		"vcd.vappVm.VappVm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm vcd_vapp_vm} Resource.
func NewVappVm_Override(v VappVm, scope constructs.Construct, id *string, config *VappVmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappVm.VappVm",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VappVm)SetAcceptAllEulas(val interface{}) {
	if err := j.validateSetAcceptAllEulasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptAllEulas",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetBootImage(val *string) {
	if err := j.validateSetBootImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImage",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetBootImageId(val *string) {
	if err := j.validateSetBootImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootImageId",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuCores(val *float64) {
	if err := j.validateSetCpuCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCores",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuHotAddEnabled(val interface{}) {
	if err := j.validateSetCpuHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuPriority(val *string) {
	if err := j.validateSetCpuPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPriority",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpus(val *float64) {
	if err := j.validateSetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpus",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetCpuShares(val *float64) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetExposeHardwareVirtualization(val interface{}) {
	if err := j.validateSetExposeHardwareVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeHardwareVirtualization",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetGuestProperties(val *map[string]*string) {
	if err := j.validateSetGuestPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestProperties",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetHardwareVersion(val *string) {
	if err := j.validateSetHardwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardwareVersion",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetHref(val *string) {
	if err := j.validateSetHrefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"href",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemoryHotAddEnabled(val interface{}) {
	if err := j.validateSetMemoryHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemoryPriority(val *string) {
	if err := j.validateSetMemoryPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPriority",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMemoryShares(val *float64) {
	if err := j.validateSetMemorySharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShares",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetNetworkDhcpWaitSeconds(val *float64) {
	if err := j.validateSetNetworkDhcpWaitSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkDhcpWaitSeconds",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetPlacementPolicyId(val *string) {
	if err := j.validateSetPlacementPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPolicyId",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetPowerOn(val interface{}) {
	if err := j.validateSetPowerOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powerOn",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetPreventUpdatePowerOff(val interface{}) {
	if err := j.validateSetPreventUpdatePowerOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUpdatePowerOff",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetSecurityTags(val *[]*string) {
	if err := j.validateSetSecurityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityTags",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetSizingPolicyId(val *string) {
	if err := j.validateSetSizingPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicyId",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetStorageProfile(val *string) {
	if err := j.validateSetStorageProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfile",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetVappName(val *string) {
	if err := j.validateSetVappNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappName",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetVappTemplateId(val *string) {
	if err := j.validateSetVappTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vappTemplateId",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetVdc(val *string) {
	if err := j.validateSetVdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdc",
		val,
	)
}

func (j *jsiiProxy_VappVm)SetVmNameInTemplate(val *string) {
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
func VappVm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappVm.VappVm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappVm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappVm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappVm.VappVm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappVm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappVm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.vappVm.VappVm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VappVm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.vappVm.VappVm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VappVm) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VappVm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VappVm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappVm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VappVm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VappVm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VappVm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VappVm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VappVm) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VappVm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VappVm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappVm) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VappVm) PutCustomization(value *VappVmCustomization) {
	if err := v.validatePutCustomizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomization",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappVm) PutDisk(value interface{}) {
	if err := v.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappVm) PutMetadataEntry(value interface{}) {
	if err := v.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappVm) PutNetwork(value interface{}) {
	if err := v.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetwork",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappVm) PutOverrideTemplateDisk(value interface{}) {
	if err := v.validatePutOverrideTemplateDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOverrideTemplateDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VappVm) ResetAcceptAllEulas() {
	_jsii_.InvokeVoid(
		v,
		"resetAcceptAllEulas",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetBootImage() {
	_jsii_.InvokeVoid(
		v,
		"resetBootImage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetBootImageId() {
	_jsii_.InvokeVoid(
		v,
		"resetBootImageId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCatalogName() {
	_jsii_.InvokeVoid(
		v,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetComputerName() {
	_jsii_.InvokeVoid(
		v,
		"resetComputerName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuCores() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuCores",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuPriority() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuPriority",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpus() {
	_jsii_.InvokeVoid(
		v,
		"resetCpus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCpuShares() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetCustomization() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomization",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetExposeHardwareVirtualization() {
	_jsii_.InvokeVoid(
		v,
		"resetExposeHardwareVirtualization",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetGuestProperties() {
	_jsii_.InvokeVoid(
		v,
		"resetGuestProperties",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetHardwareVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetHardwareVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetHref() {
	_jsii_.InvokeVoid(
		v,
		"resetHref",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemory() {
	_jsii_.InvokeVoid(
		v,
		"resetMemory",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemoryHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemoryPriority() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryPriority",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMemoryShares() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMetadata() {
	_jsii_.InvokeVoid(
		v,
		"resetMetadata",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		v,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetNetworkDhcpWaitSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkDhcpWaitSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetOrg() {
	_jsii_.InvokeVoid(
		v,
		"resetOrg",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetOsType() {
	_jsii_.InvokeVoid(
		v,
		"resetOsType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetOverrideTemplateDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideTemplateDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetPlacementPolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetPlacementPolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetPowerOn() {
	_jsii_.InvokeVoid(
		v,
		"resetPowerOn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetPreventUpdatePowerOff() {
	_jsii_.InvokeVoid(
		v,
		"resetPreventUpdatePowerOff",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetSecurityTags() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetSizingPolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetSizingPolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetStorageProfile() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageProfile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetTemplateName() {
	_jsii_.InvokeVoid(
		v,
		"resetTemplateName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetVappTemplateId() {
	_jsii_.InvokeVoid(
		v,
		"resetVappTemplateId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetVdc() {
	_jsii_.InvokeVoid(
		v,
		"resetVdc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) ResetVmNameInTemplate() {
	_jsii_.InvokeVoid(
		v,
		"resetVmNameInTemplate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

