package orguser

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/orguser/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user vcd_org_user}.
type OrgUser interface {
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
	DeployedVmQuota() *float64
	SetDeployedVmQuota(val *float64)
	DeployedVmQuotaInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	GroupNames() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstantMessaging() *string
	SetInstantMessaging(val *string)
	InstantMessagingInput() *string
	IsExternal() interface{}
	SetIsExternal(val interface{})
	IsExternalInput() interface{}
	IsGroupRole() interface{}
	SetIsGroupRole(val interface{})
	IsGroupRoleInput() interface{}
	IsLocked() interface{}
	SetIsLocked(val interface{})
	IsLockedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordFile() *string
	SetPasswordFile(val *string)
	PasswordFileInput() *string
	PasswordInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	StoredVmQuota() *float64
	SetStoredVmQuota(val *float64)
	StoredVmQuotaInput() *float64
	TakeOwnership() interface{}
	SetTakeOwnership(val interface{})
	TakeOwnershipInput() interface{}
	Telephone() *string
	SetTelephone(val *string)
	TelephoneInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetDeployedVmQuota()
	ResetDescription()
	ResetEmailAddress()
	ResetEnabled()
	ResetFullName()
	ResetId()
	ResetInstantMessaging()
	ResetIsExternal()
	ResetIsGroupRole()
	ResetIsLocked()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPasswordFile()
	ResetProviderType()
	ResetStoredVmQuota()
	ResetTakeOwnership()
	ResetTelephone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OrgUser
type jsiiProxy_OrgUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrgUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) DeployedVmQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVmQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) DeployedVmQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deployedVmQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) GroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) InstantMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instantMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) InstantMessagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instantMessagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsExternal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExternal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsExternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExternalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsGroupRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isGroupRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsGroupRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isGroupRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsLocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) IsLockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) PasswordFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) PasswordFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) StoredVmQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storedVmQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) StoredVmQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storedVmQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TakeOwnership() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"takeOwnership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TakeOwnershipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"takeOwnershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) Telephone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TelephoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user vcd_org_user} Resource.
func NewOrgUser(scope constructs.Construct, id *string, config *OrgUserConfig) OrgUser {
	_init_.Initialize()

	if err := validateNewOrgUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgUser{}

	_jsii_.Create(
		"vcd.orgUser.OrgUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user vcd_org_user} Resource.
func NewOrgUser_Override(o OrgUser, scope constructs.Construct, id *string, config *OrgUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgUser.OrgUser",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrgUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetDeployedVmQuota(val *float64) {
	if err := j.validateSetDeployedVmQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployedVmQuota",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetInstantMessaging(val *string) {
	if err := j.validateSetInstantMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instantMessaging",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetIsExternal(val interface{}) {
	if err := j.validateSetIsExternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isExternal",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetIsGroupRole(val interface{}) {
	if err := j.validateSetIsGroupRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isGroupRole",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetIsLocked(val interface{}) {
	if err := j.validateSetIsLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLocked",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetPasswordFile(val *string) {
	if err := j.validateSetPasswordFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordFile",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetProviderType(val *string) {
	if err := j.validateSetProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetStoredVmQuota(val *float64) {
	if err := j.validateSetStoredVmQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedVmQuota",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetTakeOwnership(val interface{}) {
	if err := j.validateSetTakeOwnershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"takeOwnership",
		val,
	)
}

func (j *jsiiProxy_OrgUser)SetTelephone(val *string) {
	if err := j.validateSetTelephoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telephone",
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
func OrgUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgUser.OrgUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrgUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgUser.OrgUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrgUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.orgUser.OrgUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrgUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.orgUser.OrgUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrgUser) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OrgUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgUser) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrgUser) ResetDeployedVmQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetDeployedVmQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		o,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetFullName() {
	_jsii_.InvokeVoid(
		o,
		"resetFullName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetInstantMessaging() {
	_jsii_.InvokeVoid(
		o,
		"resetInstantMessaging",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetIsExternal() {
	_jsii_.InvokeVoid(
		o,
		"resetIsExternal",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetIsGroupRole() {
	_jsii_.InvokeVoid(
		o,
		"resetIsGroupRole",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetIsLocked() {
	_jsii_.InvokeVoid(
		o,
		"resetIsLocked",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetOrg() {
	_jsii_.InvokeVoid(
		o,
		"resetOrg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetPasswordFile() {
	_jsii_.InvokeVoid(
		o,
		"resetPasswordFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetProviderType() {
	_jsii_.InvokeVoid(
		o,
		"resetProviderType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetStoredVmQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetStoredVmQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetTakeOwnership() {
	_jsii_.InvokeVoid(
		o,
		"resetTakeOwnership",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) ResetTelephone() {
	_jsii_.InvokeVoid(
		o,
		"resetTelephone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

