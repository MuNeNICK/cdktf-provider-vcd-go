package lbappprofile

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/lbappprofile/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile vcd_lb_app_profile}.
type LbAppProfile interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CookieMode() *string
	SetCookieMode(val *string)
	CookieModeInput() *string
	CookieName() *string
	SetCookieName(val *string)
	CookieNameInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EdgeGateway() *string
	SetEdgeGateway(val *string)
	EdgeGatewayInput() *string
	EnablePoolSideSsl() interface{}
	SetEnablePoolSideSsl(val interface{})
	EnablePoolSideSslInput() interface{}
	EnableSslPassthrough() interface{}
	SetEnableSslPassthrough(val interface{})
	EnableSslPassthroughInput() interface{}
	Expiration() *float64
	SetExpiration(val *float64)
	ExpirationInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpRedirectUrl() *string
	SetHttpRedirectUrl(val *string)
	HttpRedirectUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsertXForwardedHttpHeader() interface{}
	SetInsertXForwardedHttpHeader(val interface{})
	InsertXForwardedHttpHeaderInput() interface{}
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
	PersistenceMechanism() *string
	SetPersistenceMechanism(val *string)
	PersistenceMechanismInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetCookieMode()
	ResetCookieName()
	ResetEnablePoolSideSsl()
	ResetEnableSslPassthrough()
	ResetExpiration()
	ResetHttpRedirectUrl()
	ResetId()
	ResetInsertXForwardedHttpHeader()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistenceMechanism()
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

// The jsii proxy struct for LbAppProfile
type jsiiProxy_LbAppProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbAppProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) CookieMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) CookieModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) CookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) CookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EdgeGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EdgeGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EnablePoolSideSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePoolSideSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EnablePoolSideSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePoolSideSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EnableSslPassthrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) EnableSslPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslPassthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Expiration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) ExpirationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) HttpRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) HttpRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) InsertXForwardedHttpHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insertXForwardedHttpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) InsertXForwardedHttpHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insertXForwardedHttpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) PersistenceMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) PersistenceMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbAppProfile) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile vcd_lb_app_profile} Resource.
func NewLbAppProfile(scope constructs.Construct, id *string, config *LbAppProfileConfig) LbAppProfile {
	_init_.Initialize()

	if err := validateNewLbAppProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbAppProfile{}

	_jsii_.Create(
		"vcd.lbAppProfile.LbAppProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile vcd_lb_app_profile} Resource.
func NewLbAppProfile_Override(l LbAppProfile, scope constructs.Construct, id *string, config *LbAppProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.lbAppProfile.LbAppProfile",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbAppProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetCookieMode(val *string) {
	if err := j.validateSetCookieModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieMode",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetCookieName(val *string) {
	if err := j.validateSetCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieName",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetEdgeGateway(val *string) {
	if err := j.validateSetEdgeGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGateway",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetEnablePoolSideSsl(val interface{}) {
	if err := j.validateSetEnablePoolSideSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePoolSideSsl",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetEnableSslPassthrough(val interface{}) {
	if err := j.validateSetEnableSslPassthroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSslPassthrough",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetExpiration(val *float64) {
	if err := j.validateSetExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiration",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetHttpRedirectUrl(val *string) {
	if err := j.validateSetHttpRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetInsertXForwardedHttpHeader(val interface{}) {
	if err := j.validateSetInsertXForwardedHttpHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insertXForwardedHttpHeader",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetPersistenceMechanism(val *string) {
	if err := j.validateSetPersistenceMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistenceMechanism",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_LbAppProfile)SetVdc(val *string) {
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
func LbAppProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbAppProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.lbAppProfile.LbAppProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbAppProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbAppProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.lbAppProfile.LbAppProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbAppProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbAppProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.lbAppProfile.LbAppProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbAppProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.lbAppProfile.LbAppProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbAppProfile) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbAppProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbAppProfile) ResetCookieMode() {
	_jsii_.InvokeVoid(
		l,
		"resetCookieMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetCookieName() {
	_jsii_.InvokeVoid(
		l,
		"resetCookieName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetEnablePoolSideSsl() {
	_jsii_.InvokeVoid(
		l,
		"resetEnablePoolSideSsl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetEnableSslPassthrough() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableSslPassthrough",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetExpiration() {
	_jsii_.InvokeVoid(
		l,
		"resetExpiration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetHttpRedirectUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpRedirectUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetInsertXForwardedHttpHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetInsertXForwardedHttpHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetOrg() {
	_jsii_.InvokeVoid(
		l,
		"resetOrg",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetPersistenceMechanism() {
	_jsii_.InvokeVoid(
		l,
		"resetPersistenceMechanism",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) ResetVdc() {
	_jsii_.InvokeVoid(
		l,
		"resetVdc",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbAppProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbAppProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

