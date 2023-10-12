package orgldap

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/orgldap/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgLdapCustomSettingsOutputReference interface {
	cdktf.ComplexObject
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
	BaseDistinguishedName() *string
	SetBaseDistinguishedName(val *string)
	BaseDistinguishedNameInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectorType() *string
	SetConnectorType(val *string)
	ConnectorTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupAttributes() OrgLdapCustomSettingsGroupAttributesOutputReference
	GroupAttributesInput() *OrgLdapCustomSettingsGroupAttributes
	InternalValue() *OrgLdapCustomSettings
	SetInternalValue(val *OrgLdapCustomSettings)
	IsSsl() interface{}
	SetIsSsl(val interface{})
	IsSslInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Server() *string
	SetServer(val *string)
	ServerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAttributes() OrgLdapCustomSettingsUserAttributesOutputReference
	UserAttributesInput() *OrgLdapCustomSettingsUserAttributes
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutGroupAttributes(value *OrgLdapCustomSettingsGroupAttributes)
	PutUserAttributes(value *OrgLdapCustomSettingsUserAttributes)
	ResetBaseDistinguishedName()
	ResetIsSsl()
	ResetPassword()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrgLdapCustomSettingsOutputReference
type jsiiProxy_OrgLdapCustomSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) BaseDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) BaseDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) ConnectorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) GroupAttributes() OrgLdapCustomSettingsGroupAttributesOutputReference {
	var returns OrgLdapCustomSettingsGroupAttributesOutputReference
	_jsii_.Get(
		j,
		"groupAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) GroupAttributesInput() *OrgLdapCustomSettingsGroupAttributes {
	var returns *OrgLdapCustomSettingsGroupAttributes
	_jsii_.Get(
		j,
		"groupAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) InternalValue() *OrgLdapCustomSettings {
	var returns *OrgLdapCustomSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) IsSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) IsSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) UserAttributes() OrgLdapCustomSettingsUserAttributesOutputReference {
	var returns OrgLdapCustomSettingsUserAttributesOutputReference
	_jsii_.Get(
		j,
		"userAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) UserAttributesInput() *OrgLdapCustomSettingsUserAttributes {
	var returns *OrgLdapCustomSettingsUserAttributes
	_jsii_.Get(
		j,
		"userAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewOrgLdapCustomSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrgLdapCustomSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewOrgLdapCustomSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgLdapCustomSettingsOutputReference{}

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrgLdapCustomSettingsOutputReference_Override(o OrgLdapCustomSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetBaseDistinguishedName(val *string) {
	if err := j.validateSetBaseDistinguishedNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetConnectorType(val *string) {
	if err := j.validateSetConnectorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetInternalValue(val *OrgLdapCustomSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetIsSsl(val interface{}) {
	if err := j.validateSetIsSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSsl",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetServer(val *string) {
	if err := j.validateSetServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) PutGroupAttributes(value *OrgLdapCustomSettingsGroupAttributes) {
	if err := o.validatePutGroupAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGroupAttributes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) PutUserAttributes(value *OrgLdapCustomSettingsUserAttributes) {
	if err := o.validatePutUserAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUserAttributes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ResetBaseDistinguishedName() {
	_jsii_.InvokeVoid(
		o,
		"resetBaseDistinguishedName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ResetIsSsl() {
	_jsii_.InvokeVoid(
		o,
		"resetIsSsl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

