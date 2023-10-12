package org

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/org/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org vcd_org}.
type Org interface {
	cdktf.TerraformResource
	CanPublishCatalogs() interface{}
	SetCanPublishCatalogs(val interface{})
	CanPublishCatalogsInput() interface{}
	CanPublishExternalCatalogs() interface{}
	SetCanPublishExternalCatalogs(val interface{})
	CanPublishExternalCatalogsInput() interface{}
	CanSubscribeExternalCatalogs() interface{}
	SetCanSubscribeExternalCatalogs(val interface{})
	CanSubscribeExternalCatalogsInput() interface{}
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
	DelayAfterPowerOnSeconds() *float64
	SetDelayAfterPowerOnSeconds(val *float64)
	DelayAfterPowerOnSecondsInput() *float64
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
	DeployedVmQuota() *float64
	SetDeployedVmQuota(val *float64)
	DeployedVmQuotaInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
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
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataEntry() OrgMetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	StoredVmQuota() *float64
	SetStoredVmQuota(val *float64)
	StoredVmQuotaInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VappLease() OrgVappLeaseOutputReference
	VappLeaseInput() *OrgVappLease
	VappTemplateLease() OrgVappTemplateLeaseOutputReference
	VappTemplateLeaseInput() *OrgVappTemplateLease
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
	PutVappLease(value *OrgVappLease)
	PutVappTemplateLease(value *OrgVappTemplateLease)
	ResetCanPublishCatalogs()
	ResetCanPublishExternalCatalogs()
	ResetCanSubscribeExternalCatalogs()
	ResetDelayAfterPowerOnSeconds()
	ResetDeployedVmQuota()
	ResetDescription()
	ResetId()
	ResetIsEnabled()
	ResetMetadata()
	ResetMetadataEntry()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStoredVmQuota()
	ResetVappLease()
	ResetVappTemplateLease()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Org
type jsiiProxy_Org struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Org) CanPublishCatalogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canPublishCatalogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CanPublishCatalogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canPublishCatalogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CanPublishExternalCatalogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canPublishExternalCatalogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CanPublishExternalCatalogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canPublishExternalCatalogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CanSubscribeExternalCatalogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canSubscribeExternalCatalogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CanSubscribeExternalCatalogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canSubscribeExternalCatalogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DelayAfterPowerOnSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayAfterPowerOnSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DelayAfterPowerOnSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayAfterPowerOnSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeleteForce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeleteForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeleteRecursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeleteRecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeployedVmQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVmQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DeployedVmQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVmQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) MetadataEntry() OrgMetadataEntryList {
	var returns OrgMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) StoredVmQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storedVmQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) StoredVmQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storedVmQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) VappLease() OrgVappLeaseOutputReference {
	var returns OrgVappLeaseOutputReference
	_jsii_.Get(
		j,
		"vappLease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) VappLeaseInput() *OrgVappLease {
	var returns *OrgVappLease
	_jsii_.Get(
		j,
		"vappLeaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) VappTemplateLease() OrgVappTemplateLeaseOutputReference {
	var returns OrgVappTemplateLeaseOutputReference
	_jsii_.Get(
		j,
		"vappTemplateLease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Org) VappTemplateLeaseInput() *OrgVappTemplateLease {
	var returns *OrgVappTemplateLease
	_jsii_.Get(
		j,
		"vappTemplateLeaseInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org vcd_org} Resource.
func NewOrg(scope constructs.Construct, id *string, config *OrgConfig) Org {
	_init_.Initialize()

	if err := validateNewOrgParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Org{}

	_jsii_.Create(
		"vcd.org.Org",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org vcd_org} Resource.
func NewOrg_Override(o Org, scope constructs.Construct, id *string, config *OrgConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.org.Org",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_Org)SetCanPublishCatalogs(val interface{}) {
	if err := j.validateSetCanPublishCatalogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canPublishCatalogs",
		val,
	)
}

func (j *jsiiProxy_Org)SetCanPublishExternalCatalogs(val interface{}) {
	if err := j.validateSetCanPublishExternalCatalogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canPublishExternalCatalogs",
		val,
	)
}

func (j *jsiiProxy_Org)SetCanSubscribeExternalCatalogs(val interface{}) {
	if err := j.validateSetCanSubscribeExternalCatalogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canSubscribeExternalCatalogs",
		val,
	)
}

func (j *jsiiProxy_Org)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Org)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Org)SetDelayAfterPowerOnSeconds(val *float64) {
	if err := j.validateSetDelayAfterPowerOnSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delayAfterPowerOnSeconds",
		val,
	)
}

func (j *jsiiProxy_Org)SetDeleteForce(val interface{}) {
	if err := j.validateSetDeleteForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteForce",
		val,
	)
}

func (j *jsiiProxy_Org)SetDeleteRecursive(val interface{}) {
	if err := j.validateSetDeleteRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRecursive",
		val,
	)
}

func (j *jsiiProxy_Org)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Org)SetDeployedVmQuota(val *float64) {
	if err := j.validateSetDeployedVmQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployedVmQuota",
		val,
	)
}

func (j *jsiiProxy_Org)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Org)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Org)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_Org)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Org)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_Org)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Org)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Org)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Org)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Org)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Org)SetStoredVmQuota(val *float64) {
	if err := j.validateSetStoredVmQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedVmQuota",
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
func Org_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrg_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.org.Org",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Org_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrg_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.org.Org",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Org_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrg_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.org.Org",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Org_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.org.Org",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_Org) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_Org) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_Org) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_Org) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_Org) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_Org) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_Org) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_Org) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_Org) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_Org) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_Org) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_Org) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_Org) PutMetadataEntry(value interface{}) {
	if err := o.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_Org) PutVappLease(value *OrgVappLease) {
	if err := o.validatePutVappLeaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVappLease",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_Org) PutVappTemplateLease(value *OrgVappTemplateLease) {
	if err := o.validatePutVappTemplateLeaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVappTemplateLease",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_Org) ResetCanPublishCatalogs() {
	_jsii_.InvokeVoid(
		o,
		"resetCanPublishCatalogs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetCanPublishExternalCatalogs() {
	_jsii_.InvokeVoid(
		o,
		"resetCanPublishExternalCatalogs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetCanSubscribeExternalCatalogs() {
	_jsii_.InvokeVoid(
		o,
		"resetCanSubscribeExternalCatalogs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetDelayAfterPowerOnSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetDelayAfterPowerOnSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetDeployedVmQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetDeployedVmQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetMetadata() {
	_jsii_.InvokeVoid(
		o,
		"resetMetadata",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		o,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetStoredVmQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetStoredVmQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetVappLease() {
	_jsii_.InvokeVoid(
		o,
		"resetVappLease",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) ResetVappTemplateLease() {
	_jsii_.InvokeVoid(
		o,
		"resetVappTemplateLease",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Org) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Org) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Org) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Org) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

