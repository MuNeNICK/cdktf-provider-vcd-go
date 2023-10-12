package orgldap

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/orgldap/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgLdapCustomSettingsUserAttributesOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	Fqn() *string
	GivenName() *string
	SetGivenName(val *string)
	GivenNameInput() *string
	GroupBackLinkIdentifier() *string
	SetGroupBackLinkIdentifier(val *string)
	GroupBackLinkIdentifierInput() *string
	GroupMembershipIdentifier() *string
	SetGroupMembershipIdentifier(val *string)
	GroupMembershipIdentifierInput() *string
	InternalValue() *OrgLdapCustomSettingsUserAttributes
	SetInternalValue(val *OrgLdapCustomSettingsUserAttributes)
	ObjectClass() *string
	SetObjectClass(val *string)
	ObjectClassInput() *string
	Surname() *string
	SetSurname(val *string)
	SurnameInput() *string
	Telephone() *string
	SetTelephone(val *string)
	TelephoneInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UniqueIdentifier() *string
	SetUniqueIdentifier(val *string)
	UniqueIdentifierInput() *string
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
	ResetGroupBackLinkIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrgLdapCustomSettingsUserAttributesOutputReference
type jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GroupBackLinkIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupBackLinkIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GroupBackLinkIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupBackLinkIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GroupMembershipIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMembershipIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GroupMembershipIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMembershipIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) InternalValue() *OrgLdapCustomSettingsUserAttributes {
	var returns *OrgLdapCustomSettingsUserAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ObjectClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ObjectClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) SurnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Telephone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) TelephoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) UniqueIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) UniqueIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewOrgLdapCustomSettingsUserAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrgLdapCustomSettingsUserAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewOrgLdapCustomSettingsUserAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference{}

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsUserAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrgLdapCustomSettingsUserAttributesOutputReference_Override(o OrgLdapCustomSettingsUserAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsUserAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetGivenName(val *string) {
	if err := j.validateSetGivenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetGroupBackLinkIdentifier(val *string) {
	if err := j.validateSetGroupBackLinkIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBackLinkIdentifier",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetGroupMembershipIdentifier(val *string) {
	if err := j.validateSetGroupMembershipIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMembershipIdentifier",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetInternalValue(val *OrgLdapCustomSettingsUserAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetObjectClass(val *string) {
	if err := j.validateSetObjectClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectClass",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetSurname(val *string) {
	if err := j.validateSetSurnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surname",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetTelephone(val *string) {
	if err := j.validateSetTelephoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telephone",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetUniqueIdentifier(val *string) {
	if err := j.validateSetUniqueIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueIdentifier",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ResetGroupBackLinkIdentifier() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupBackLinkIdentifier",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrgLdapCustomSettingsUserAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

