package provider

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/provider/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs vcd}.
type VcdProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AllowApiTokenFile() interface{}
	SetAllowApiTokenFile(val interface{})
	AllowApiTokenFileInput() interface{}
	AllowServiceAccountTokenFile() interface{}
	SetAllowServiceAccountTokenFile(val interface{})
	AllowServiceAccountTokenFileInput() interface{}
	AllowUnverifiedSsl() interface{}
	SetAllowUnverifiedSsl(val interface{})
	AllowUnverifiedSslInput() interface{}
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenFile() *string
	SetApiTokenFile(val *string)
	ApiTokenFileInput() *string
	ApiTokenInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IgnoreMetadataChanges() interface{}
	SetIgnoreMetadataChanges(val interface{})
	IgnoreMetadataChangesInput() interface{}
	ImportSeparator() *string
	SetImportSeparator(val *string)
	ImportSeparatorInput() *string
	Logging() interface{}
	SetLogging(val interface{})
	LoggingFile() *string
	SetLoggingFile(val *string)
	LoggingFileInput() *string
	LoggingInput() interface{}
	MaxRetryTimeout() *float64
	SetMaxRetryTimeout(val *float64)
	MaxRetryTimeoutInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	SamlAdfsRptId() *string
	SetSamlAdfsRptId(val *string)
	SamlAdfsRptIdInput() *string
	ServiceAccountTokenFile() *string
	SetServiceAccountTokenFile(val *string)
	ServiceAccountTokenFileInput() *string
	Sysorg() *string
	SetSysorg(val *string)
	SysorgInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
	Vdc() *string
	SetVdc(val *string)
	VdcInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAllowApiTokenFile()
	ResetAllowServiceAccountTokenFile()
	ResetAllowUnverifiedSsl()
	ResetApiToken()
	ResetApiTokenFile()
	ResetAuthType()
	ResetIgnoreMetadataChanges()
	ResetImportSeparator()
	ResetLogging()
	ResetLoggingFile()
	ResetMaxRetryTimeout()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetSamlAdfsRptId()
	ResetServiceAccountTokenFile()
	ResetSysorg()
	ResetToken()
	ResetUser()
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

// The jsii proxy struct for VcdProvider
type jsiiProxy_VcdProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_VcdProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowApiTokenFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowApiTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowApiTokenFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowApiTokenFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowServiceAccountTokenFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowServiceAccountTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowServiceAccountTokenFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowServiceAccountTokenFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowUnverifiedSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AllowUnverifiedSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ApiTokenFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ApiTokenFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) IgnoreMetadataChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMetadataChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) IgnoreMetadataChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMetadataChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ImportSeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importSeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ImportSeparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importSeparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) LoggingFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) LoggingFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) MaxRetryTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) MaxRetryTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) SamlAdfsRptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAdfsRptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) SamlAdfsRptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAdfsRptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ServiceAccountTokenFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) ServiceAccountTokenFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountTokenFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Sysorg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sysorg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) SysorgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sysorgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VcdProvider) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs vcd} Resource.
func NewVcdProvider(scope constructs.Construct, id *string, config *VcdProviderConfig) VcdProvider {
	_init_.Initialize()

	if err := validateNewVcdProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VcdProvider{}

	_jsii_.Create(
		"vcd.provider.VcdProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs vcd} Resource.
func NewVcdProvider_Override(v VcdProvider, scope constructs.Construct, id *string, config *VcdProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.provider.VcdProvider",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VcdProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetAllowApiTokenFile(val interface{}) {
	if err := j.validateSetAllowApiTokenFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowApiTokenFile",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetAllowServiceAccountTokenFile(val interface{}) {
	if err := j.validateSetAllowServiceAccountTokenFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowServiceAccountTokenFile",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetAllowUnverifiedSsl(val interface{}) {
	if err := j.validateSetAllowUnverifiedSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnverifiedSsl",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetApiTokenFile(val *string) {
	_jsii_.Set(
		j,
		"apiTokenFile",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetIgnoreMetadataChanges(val interface{}) {
	if err := j.validateSetIgnoreMetadataChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMetadataChanges",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetImportSeparator(val *string) {
	_jsii_.Set(
		j,
		"importSeparator",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetLogging(val interface{}) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetLoggingFile(val *string) {
	_jsii_.Set(
		j,
		"loggingFile",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetMaxRetryTimeout(val *float64) {
	_jsii_.Set(
		j,
		"maxRetryTimeout",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetOrg(val *string) {
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetSamlAdfsRptId(val *string) {
	_jsii_.Set(
		j,
		"samlAdfsRptId",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetServiceAccountTokenFile(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountTokenFile",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetSysorg(val *string) {
	_jsii_.Set(
		j,
		"sysorg",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_VcdProvider)SetVdc(val *string) {
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
func VcdProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcdProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.provider.VcdProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VcdProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcdProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.provider.VcdProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VcdProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVcdProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.provider.VcdProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VcdProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.provider.VcdProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VcdProvider) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VcdProvider) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VcdProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		v,
		"resetAlias",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetAllowApiTokenFile() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowApiTokenFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetAllowServiceAccountTokenFile() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowServiceAccountTokenFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetAllowUnverifiedSsl() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowUnverifiedSsl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetApiToken() {
	_jsii_.InvokeVoid(
		v,
		"resetApiToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetApiTokenFile() {
	_jsii_.InvokeVoid(
		v,
		"resetApiTokenFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetAuthType() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetIgnoreMetadataChanges() {
	_jsii_.InvokeVoid(
		v,
		"resetIgnoreMetadataChanges",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetImportSeparator() {
	_jsii_.InvokeVoid(
		v,
		"resetImportSeparator",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetLogging() {
	_jsii_.InvokeVoid(
		v,
		"resetLogging",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetLoggingFile() {
	_jsii_.InvokeVoid(
		v,
		"resetLoggingFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetMaxRetryTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxRetryTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetSamlAdfsRptId() {
	_jsii_.InvokeVoid(
		v,
		"resetSamlAdfsRptId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetServiceAccountTokenFile() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceAccountTokenFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetSysorg() {
	_jsii_.InvokeVoid(
		v,
		"resetSysorg",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetToken() {
	_jsii_.InvokeVoid(
		v,
		"resetToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetUser() {
	_jsii_.InvokeVoid(
		v,
		"resetUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) ResetVdc() {
	_jsii_.InvokeVoid(
		v,
		"resetVdc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VcdProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcdProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcdProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VcdProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

