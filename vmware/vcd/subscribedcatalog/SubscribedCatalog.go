package subscribedcatalog

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/subscribedcatalog/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog vcd_subscribed_catalog}.
type SubscribedCatalog interface {
	cdktf.TerraformResource
	CancelFailedTasks() interface{}
	SetCancelFailedTasks(val interface{})
	CancelFailedTasksInput() interface{}
	CatalogVersion() *float64
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
	Created() *string
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
	FailedTasks() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Href() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsLocal() cdktf.IResolvable
	IsPublished() cdktf.IResolvable
	IsShared() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MakeLocalCopy() interface{}
	SetMakeLocalCopy(val interface{})
	MakeLocalCopyInput() interface{}
	MediaItemList() *[]*string
	Metadata() cdktf.StringMap
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumberOfMedia() *float64
	NumberOfVappTemplates() *float64
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OwnerName() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublishSubscriptionType() *string
	// Experimental.
	RawOverrides() interface{}
	RunningTasks() *[]*string
	StorageProfileId() *string
	SetStorageProfileId(val *string)
	StorageProfileIdInput() *string
	StoreTasks() interface{}
	SetStoreTasks(val interface{})
	StoreTasksInput() interface{}
	SubscriptionPassword() *string
	SetSubscriptionPassword(val *string)
	SubscriptionPasswordInput() *string
	SubscriptionUrl() *string
	SetSubscriptionUrl(val *string)
	SubscriptionUrlInput() *string
	SyncAll() interface{}
	SetSyncAll(val interface{})
	SyncAllInput() interface{}
	SyncAllMediaItems() interface{}
	SetSyncAllMediaItems(val interface{})
	SyncAllMediaItemsInput() interface{}
	SyncAllVappTemplates() interface{}
	SetSyncAllVappTemplates(val interface{})
	SyncAllVappTemplatesInput() interface{}
	SyncCatalog() interface{}
	SetSyncCatalog(val interface{})
	SyncCatalogInput() interface{}
	SyncMediaItems() *[]*string
	SetSyncMediaItems(val *[]*string)
	SyncMediaItemsInput() *[]*string
	SyncOnRefresh() interface{}
	SetSyncOnRefresh(val interface{})
	SyncOnRefreshInput() interface{}
	SyncVappTemplates() *[]*string
	SetSyncVappTemplates(val *[]*string)
	SyncVappTemplatesInput() *[]*string
	TasksFileName() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VappTemplateList() *[]*string
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
	ResetCancelFailedTasks()
	ResetDeleteForce()
	ResetDeleteRecursive()
	ResetId()
	ResetMakeLocalCopy()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageProfileId()
	ResetStoreTasks()
	ResetSubscriptionPassword()
	ResetSyncAll()
	ResetSyncAllMediaItems()
	ResetSyncAllVappTemplates()
	ResetSyncCatalog()
	ResetSyncMediaItems()
	ResetSyncOnRefresh()
	ResetSyncVappTemplates()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SubscribedCatalog
type jsiiProxy_SubscribedCatalog struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SubscribedCatalog) CancelFailedTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cancelFailedTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) CancelFailedTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cancelFailedTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) CatalogVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"catalogVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) DeleteForce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) DeleteForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) DeleteRecursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) DeleteRecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) FailedTasks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failedTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) IsLocal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) IsPublished() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPublished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) IsShared() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isShared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) MakeLocalCopy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"makeLocalCopy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) MakeLocalCopyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"makeLocalCopyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) MediaItemList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mediaItemList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) NumberOfMedia() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfMedia",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) NumberOfVappTemplates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfVappTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) OwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) PublishSubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishSubscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) RunningTasks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runningTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) StorageProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) StorageProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) StoreTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) StoreTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SubscriptionPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SubscriptionPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SubscriptionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SubscriptionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAllMediaItems() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAllMediaItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAllMediaItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAllMediaItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAllVappTemplates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAllVappTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncAllVappTemplatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncAllVappTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncCatalog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncCatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncMediaItems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syncMediaItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncMediaItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syncMediaItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncOnRefresh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncOnRefreshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncVappTemplates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syncVappTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) SyncVappTemplatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syncVappTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) TasksFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tasksFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscribedCatalog) VappTemplateList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vappTemplateList",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog vcd_subscribed_catalog} Resource.
func NewSubscribedCatalog(scope constructs.Construct, id *string, config *SubscribedCatalogConfig) SubscribedCatalog {
	_init_.Initialize()

	if err := validateNewSubscribedCatalogParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SubscribedCatalog{}

	_jsii_.Create(
		"vcd.subscribedCatalog.SubscribedCatalog",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog vcd_subscribed_catalog} Resource.
func NewSubscribedCatalog_Override(s SubscribedCatalog, scope constructs.Construct, id *string, config *SubscribedCatalogConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.subscribedCatalog.SubscribedCatalog",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetCancelFailedTasks(val interface{}) {
	if err := j.validateSetCancelFailedTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cancelFailedTasks",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetDeleteForce(val interface{}) {
	if err := j.validateSetDeleteForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteForce",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetDeleteRecursive(val interface{}) {
	if err := j.validateSetDeleteRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRecursive",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetMakeLocalCopy(val interface{}) {
	if err := j.validateSetMakeLocalCopyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"makeLocalCopy",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetStorageProfileId(val *string) {
	if err := j.validateSetStorageProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfileId",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetStoreTasks(val interface{}) {
	if err := j.validateSetStoreTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeTasks",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSubscriptionPassword(val *string) {
	if err := j.validateSetSubscriptionPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionPassword",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSubscriptionUrl(val *string) {
	if err := j.validateSetSubscriptionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionUrl",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncAll(val interface{}) {
	if err := j.validateSetSyncAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncAll",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncAllMediaItems(val interface{}) {
	if err := j.validateSetSyncAllMediaItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncAllMediaItems",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncAllVappTemplates(val interface{}) {
	if err := j.validateSetSyncAllVappTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncAllVappTemplates",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncCatalog(val interface{}) {
	if err := j.validateSetSyncCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncCatalog",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncMediaItems(val *[]*string) {
	if err := j.validateSetSyncMediaItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMediaItems",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncOnRefresh(val interface{}) {
	if err := j.validateSetSyncOnRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncOnRefresh",
		val,
	)
}

func (j *jsiiProxy_SubscribedCatalog)SetSyncVappTemplates(val *[]*string) {
	if err := j.validateSetSyncVappTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncVappTemplates",
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
func SubscribedCatalog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscribedCatalog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.subscribedCatalog.SubscribedCatalog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SubscribedCatalog_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscribedCatalog_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.subscribedCatalog.SubscribedCatalog",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SubscribedCatalog_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubscribedCatalog_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.subscribedCatalog.SubscribedCatalog",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SubscribedCatalog_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.subscribedCatalog.SubscribedCatalog",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SubscribedCatalog) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SubscribedCatalog) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetCancelFailedTasks() {
	_jsii_.InvokeVoid(
		s,
		"resetCancelFailedTasks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetDeleteForce() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteForce",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetDeleteRecursive() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteRecursive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetMakeLocalCopy() {
	_jsii_.InvokeVoid(
		s,
		"resetMakeLocalCopy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetOrg() {
	_jsii_.InvokeVoid(
		s,
		"resetOrg",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetStorageProfileId() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageProfileId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetStoreTasks() {
	_jsii_.InvokeVoid(
		s,
		"resetStoreTasks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSubscriptionPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetSubscriptionPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncAll() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncAllMediaItems() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncAllMediaItems",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncAllVappTemplates() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncAllVappTemplates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncCatalog() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncCatalog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncMediaItems() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncMediaItems",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncOnRefresh() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncOnRefresh",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) ResetSyncVappTemplates() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncVappTemplates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscribedCatalog) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscribedCatalog) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

