package catalog

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/catalog/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog vcd_catalog}.
type Catalog interface {
	cdktf.TerraformResource
	CacheEnabled() interface{}
	SetCacheEnabled(val interface{})
	CacheEnabledInput() interface{}
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
	MediaItemList() *[]*string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataEntry() CatalogMetadataEntryList
	MetadataEntryInput() interface{}
	MetadataInput() *map[string]*string
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
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PreserveIdentityInformation() interface{}
	SetPreserveIdentityInformation(val interface{})
	PreserveIdentityInformationInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublishEnabled() interface{}
	SetPublishEnabled(val interface{})
	PublishEnabledInput() interface{}
	PublishSubscriptionType() *string
	PublishSubscriptionUrl() *string
	// Experimental.
	RawOverrides() interface{}
	StorageProfileId() *string
	SetStorageProfileId(val *string)
	StorageProfileIdInput() *string
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
	PutMetadataEntry(value interface{})
	ResetCacheEnabled()
	ResetDescription()
	ResetId()
	ResetMetadata()
	ResetMetadataEntry()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPreserveIdentityInformation()
	ResetPublishEnabled()
	ResetStorageProfileId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Catalog
type jsiiProxy_Catalog struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Catalog) CacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CatalogVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"catalogVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DeleteForce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DeleteForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteForceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DeleteRecursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DeleteRecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRecursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Href() *string {
	var returns *string
	_jsii_.Get(
		j,
		"href",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IsLocal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IsPublished() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPublished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IsShared() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isShared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MediaItemList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mediaItemList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MetadataEntry() CatalogMetadataEntryList {
	var returns CatalogMetadataEntryList
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MetadataEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) NumberOfMedia() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfMedia",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) NumberOfVappTemplates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfVappTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) OwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PreserveIdentityInformation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveIdentityInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PreserveIdentityInformationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveIdentityInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PublishEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PublishEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PublishSubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishSubscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PublishSubscriptionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishSubscriptionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) StorageProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) StorageProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) VappTemplateList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vappTemplateList",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog vcd_catalog} Resource.
func NewCatalog(scope constructs.Construct, id *string, config *CatalogConfig) Catalog {
	_init_.Initialize()

	if err := validateNewCatalogParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Catalog{}

	_jsii_.Create(
		"vcd.catalog.Catalog",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog vcd_catalog} Resource.
func NewCatalog_Override(c Catalog, scope constructs.Construct, id *string, config *CatalogConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.catalog.Catalog",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Catalog)SetCacheEnabled(val interface{}) {
	if err := j.validateSetCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheEnabled",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetDeleteForce(val interface{}) {
	if err := j.validateSetDeleteForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteForce",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetDeleteRecursive(val interface{}) {
	if err := j.validateSetDeleteRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRecursive",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetPreserveIdentityInformation(val interface{}) {
	if err := j.validateSetPreserveIdentityInformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveIdentityInformation",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetPublishEnabled(val interface{}) {
	if err := j.validateSetPublishEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishEnabled",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetStorageProfileId(val *string) {
	if err := j.validateSetStorageProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfileId",
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
func Catalog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.catalog.Catalog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Catalog_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.catalog.Catalog",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Catalog_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.catalog.Catalog",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Catalog_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.catalog.Catalog",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Catalog) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Catalog) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Catalog) PutMetadataEntry(value interface{}) {
	if err := c.validatePutMetadataEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetadataEntry",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Catalog) ResetCacheEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetMetadataEntry() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadataEntry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetOrg() {
	_jsii_.InvokeVoid(
		c,
		"resetOrg",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetPreserveIdentityInformation() {
	_jsii_.InvokeVoid(
		c,
		"resetPreserveIdentityInformation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetPublishEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPublishEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetStorageProfileId() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageProfileId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

