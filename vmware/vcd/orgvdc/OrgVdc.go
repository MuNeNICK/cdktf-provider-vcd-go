package orgvdc

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/orgvdc/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc vcd_org_vdc}.
type OrgVdc interface {
	cdktf.TerraformResource
	AllocationModel() *string
	SetAllocationModel(val *string)
	AllocationModelInput() *string
	AllowOverCommit() interface{}
	SetAllowOverCommit(val interface{})
	AllowOverCommitInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeCapacity() OrgVdcComputeCapacityOutputReference
	ComputeCapacityInput() *OrgVdcComputeCapacity
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
	CpuGuaranteed() *float64
	SetCpuGuaranteed(val *float64)
	CpuGuaranteedInput() *float64
	CpuSpeed() *float64
	SetCpuSpeed(val *float64)
	CpuSpeedInput() *float64
	DefaultComputePolicyId() *string
	SetDefaultComputePolicyId(val *string)
	DefaultComputePolicyIdInput() *string
	DefaultVmSizingPolicyId() *string
	SetDefaultVmSizingPolicyId(val *string)
	DefaultVmSizingPolicyIdInput() *string
	DeleteForce() interface{}
	SetDeleteForce(val interface{})
	DeleteForceInput() interface{}
	DeleteRecursive() interface{}
	SetDeleteRecursive(val interface{})
	DeleteRecursiveInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeClusterId() *string
	SetEdgeClusterId(val *string)
	EdgeClusterIdInput() *string
	Elasticity() interface{}
	SetElasticity(val interface{})
	ElasticityInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnableFastProvisioning() interface{}
	SetEnableFastProvisioning(val interface{})
	EnableFastProvisioningInput() interface{}
	EnableNsxvDistributedFirewall() interface{}
	SetEnableNsxvDistributedFirewall(val interface{})
	EnableNsxvDistributedFirewallInput() interface{}
	EnableThinProvisioning() interface{}
	SetEnableThinProvisioning(val interface{})
	EnableThinProvisioningInput() interface{}
	EnableVmDiscovery() interface{}
	SetEnableVmDiscovery(val interface{})
	EnableVmDiscoveryInput() interface{}
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
	IncludeVmMemoryOverhead() interface{}
	SetIncludeVmMemoryOverhead(val interface{})
	IncludeVmMemoryOverheadInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MemoryGuaranteed() *float64
	SetMemoryGuaranteed(val *float64)
	MemoryGuaranteedInput() *float64
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataEntry() OrgVdcMetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkPoolName() *string
	SetNetworkPoolName(val *string)
	NetworkPoolNameInput() *string
	NetworkQuota() *float64
	SetNetworkQuota(val *float64)
	NetworkQuotaInput() *float64
	NicQuota() *float64
	SetNicQuota(val *float64)
	NicQuotaInput() *float64
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderVdcName() *string
	SetProviderVdcName(val *string)
	ProviderVdcNameInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	StorageProfile() OrgVdcStorageProfileList
	StorageProfileInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VmPlacementPolicyIds() *[]*string
	SetVmPlacementPolicyIds(val *[]*string)
	VmPlacementPolicyIdsInput() *[]*string
	VmQuota() *float64
	SetVmQuota(val *float64)
	VmQuotaInput() *float64
	VmSizingPolicyIds() *[]*string
	SetVmSizingPolicyIds(val *[]*string)
	VmSizingPolicyIdsInput() *[]*string
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
	PutComputeCapacity(value *OrgVdcComputeCapacity)
	PutMetadataEntry(value interface{})
	PutStorageProfile(value interface{})
	ResetAllowOverCommit()
	ResetCpuGuaranteed()
	ResetCpuSpeed()
	ResetDefaultComputePolicyId()
	ResetDefaultVmSizingPolicyId()
	ResetDescription()
	ResetEdgeClusterId()
	ResetElasticity()
	ResetEnabled()
	ResetEnableFastProvisioning()
	ResetEnableNsxvDistributedFirewall()
	ResetEnableThinProvisioning()
	ResetEnableVmDiscovery()
	ResetId()
	ResetIncludeVmMemoryOverhead()
	ResetMemoryGuaranteed()
	ResetMetadata()
	ResetMetadataEntry()
	ResetNetworkPoolName()
	ResetNetworkQuota()
	ResetNicQuota()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetVmPlacementPolicyIds()
	ResetVmQuota()
	ResetVmSizingPolicyIds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OrgVdc
type jsiiProxy_OrgVdc struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrgVdc) AllocationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) AllocationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) AllowOverCommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverCommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) AllowOverCommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverCommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ComputeCapacity() OrgVdcComputeCapacityOutputReference {
	var returns OrgVdcComputeCapacityOutputReference
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ComputeCapacityInput() *OrgVdcComputeCapacity {
	var returns *OrgVdcComputeCapacity
	_jsii_.Get(
		j,
		"computeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) CpuGuaranteed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuGuaranteed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) CpuGuaranteedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuGuaranteedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) CpuSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) CpuSpeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSpeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DefaultComputePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultComputePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DefaultComputePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultComputePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DefaultVmSizingPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVmSizingPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DefaultVmSizingPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVmSizingPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DeleteForce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DeleteForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DeleteRecursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DeleteRecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EdgeClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EdgeClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Elasticity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ElasticityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableFastProvisioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFastProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableFastProvisioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFastProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableNsxvDistributedFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNsxvDistributedFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableNsxvDistributedFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNsxvDistributedFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableThinProvisioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableThinProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableThinProvisioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableThinProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableVmDiscovery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVmDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) EnableVmDiscoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVmDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) IncludeVmMemoryOverhead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVmMemoryOverhead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) IncludeVmMemoryOverheadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVmMemoryOverheadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) MemoryGuaranteed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryGuaranteed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) MemoryGuaranteedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryGuaranteedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) MetadataEntry() OrgVdcMetadataEntryList {
	var returns OrgVdcMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NetworkPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NetworkPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NetworkQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NetworkQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NicQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nicQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) NicQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nicQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ProviderVdcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVdcName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) ProviderVdcNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVdcNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) StorageProfile() OrgVdcStorageProfileList {
	var returns OrgVdcStorageProfileList
	_jsii_.Get(
		j,
		"storageProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) StorageProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmPlacementPolicyIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmPlacementPolicyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmPlacementPolicyIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmPlacementPolicyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmSizingPolicyIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmSizingPolicyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgVdc) VmSizingPolicyIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmSizingPolicyIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc vcd_org_vdc} Resource.
func NewOrgVdc(scope constructs.Construct, id *string, config *OrgVdcConfig) OrgVdc {
	_init_.Initialize()

	if err := validateNewOrgVdcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgVdc{}

	_jsii_.Create(
		"vcd.orgVdc.OrgVdc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc vcd_org_vdc} Resource.
func NewOrgVdc_Override(o OrgVdc, scope constructs.Construct, id *string, config *OrgVdcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgVdc.OrgVdc",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrgVdc)SetAllocationModel(val *string) {
	if err := j.validateSetAllocationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationModel",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetAllowOverCommit(val interface{}) {
	if err := j.validateSetAllowOverCommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverCommit",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetCpuGuaranteed(val *float64) {
	if err := j.validateSetCpuGuaranteedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuGuaranteed",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetCpuSpeed(val *float64) {
	if err := j.validateSetCpuSpeedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuSpeed",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDefaultComputePolicyId(val *string) {
	if err := j.validateSetDefaultComputePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultComputePolicyId",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDefaultVmSizingPolicyId(val *string) {
	if err := j.validateSetDefaultVmSizingPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultVmSizingPolicyId",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDeleteForce(val interface{}) {
	if err := j.validateSetDeleteForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteForce",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDeleteRecursive(val interface{}) {
	if err := j.validateSetDeleteRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRecursive",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEdgeClusterId(val *string) {
	if err := j.validateSetEdgeClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeClusterId",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetElasticity(val interface{}) {
	if err := j.validateSetElasticityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticity",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEnableFastProvisioning(val interface{}) {
	if err := j.validateSetEnableFastProvisioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFastProvisioning",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEnableNsxvDistributedFirewall(val interface{}) {
	if err := j.validateSetEnableNsxvDistributedFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNsxvDistributedFirewall",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEnableThinProvisioning(val interface{}) {
	if err := j.validateSetEnableThinProvisioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableThinProvisioning",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetEnableVmDiscovery(val interface{}) {
	if err := j.validateSetEnableVmDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVmDiscovery",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetIncludeVmMemoryOverhead(val interface{}) {
	if err := j.validateSetIncludeVmMemoryOverheadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVmMemoryOverhead",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetMemoryGuaranteed(val *float64) {
	if err := j.validateSetMemoryGuaranteedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryGuaranteed",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetNetworkPoolName(val *string) {
	if err := j.validateSetNetworkPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPoolName",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetNetworkQuota(val *float64) {
	if err := j.validateSetNetworkQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkQuota",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetNicQuota(val *float64) {
	if err := j.validateSetNicQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicQuota",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetProviderVdcName(val *string) {
	if err := j.validateSetProviderVdcNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerVdcName",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetVmPlacementPolicyIds(val *[]*string) {
	if err := j.validateSetVmPlacementPolicyIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmPlacementPolicyIds",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetVmQuota(val *float64) {
	if err := j.validateSetVmQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmQuota",
		val,
	)
}

func (j *jsiiProxy_OrgVdc)SetVmSizingPolicyIds(val *[]*string) {
	if err := j.validateSetVmSizingPolicyIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSizingPolicyIds",
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
func OrgVdc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgVdc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgVdc.OrgVdc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrgVdc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgVdc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgVdc.OrgVdc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrgVdc_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgVdc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgVdc.OrgVdc",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrgVdc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.orgVdc.OrgVdc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrgVdc) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OrgVdc) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrgVdc) PutComputeCapacity(value *OrgVdcComputeCapacity) {
	if err := o.validatePutComputeCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putComputeCapacity",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgVdc) PutMetadataEntry(value interface{}) {
	if err := o.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgVdc) PutStorageProfile(value interface{}) {
	if err := o.validatePutStorageProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putStorageProfile",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgVdc) ResetAllowOverCommit() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowOverCommit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetCpuGuaranteed() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuGuaranteed",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetCpuSpeed() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuSpeed",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetDefaultComputePolicyId() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultComputePolicyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetDefaultVmSizingPolicyId() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultVmSizingPolicyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEdgeClusterId() {
	_jsii_.InvokeVoid(
		o,
		"resetEdgeClusterId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetElasticity() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEnableFastProvisioning() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableFastProvisioning",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEnableNsxvDistributedFirewall() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableNsxvDistributedFirewall",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEnableThinProvisioning() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableThinProvisioning",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetEnableVmDiscovery() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableVmDiscovery",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetIncludeVmMemoryOverhead() {
	_jsii_.InvokeVoid(
		o,
		"resetIncludeVmMemoryOverhead",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetMemoryGuaranteed() {
	_jsii_.InvokeVoid(
		o,
		"resetMemoryGuaranteed",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetMetadata() {
	_jsii_.InvokeVoid(
		o,
		"resetMetadata",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		o,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetNetworkPoolName() {
	_jsii_.InvokeVoid(
		o,
		"resetNetworkPoolName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetNetworkQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetNetworkQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetNicQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetNicQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetOrg() {
	_jsii_.InvokeVoid(
		o,
		"resetOrg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetVmPlacementPolicyIds() {
	_jsii_.InvokeVoid(
		o,
		"resetVmPlacementPolicyIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetVmQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetVmQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) ResetVmSizingPolicyIds() {
	_jsii_.InvokeVoid(
		o,
		"resetVmSizingPolicyIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgVdc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgVdc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

