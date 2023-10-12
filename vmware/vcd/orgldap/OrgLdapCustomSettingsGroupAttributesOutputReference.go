package orgldap

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/orgldap/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgLdapCustomSettingsGroupAttributesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GroupBackLinkIdentifier() *string
	SetGroupBackLinkIdentifier(val *string)
	GroupBackLinkIdentifierInput() *string
	GroupMembershipIdentifier() *string
	SetGroupMembershipIdentifier(val *string)
	GroupMembershipIdentifierInput() *string
	InternalValue() *OrgLdapCustomSettingsGroupAttributes
	SetInternalValue(val *OrgLdapCustomSettingsGroupAttributes)
	Membership() *string
	SetMembership(val *string)
	MembershipInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ObjectClass() *string
	SetObjectClass(val *string)
	ObjectClassInput() *string
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

// The jsii proxy struct for OrgLdapCustomSettingsGroupAttributesOutputReference
type jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GroupBackLinkIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupBackLinkIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GroupBackLinkIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupBackLinkIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GroupMembershipIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMembershipIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GroupMembershipIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMembershipIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) InternalValue() *OrgLdapCustomSettingsGroupAttributes {
	var returns *OrgLdapCustomSettingsGroupAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) Membership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) MembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ObjectClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ObjectClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) UniqueIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) UniqueIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueIdentifierInput",
		&returns,
	)
	return returns
}


func NewOrgLdapCustomSettingsGroupAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrgLdapCustomSettingsGroupAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewOrgLdapCustomSettingsGroupAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference{}

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsGroupAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrgLdapCustomSettingsGroupAttributesOutputReference_Override(o OrgLdapCustomSettingsGroupAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.orgLdap.OrgLdapCustomSettingsGroupAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetGroupBackLinkIdentifier(val *string) {
	if err := j.validateSetGroupBackLinkIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBackLinkIdentifier",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetGroupMembershipIdentifier(val *string) {
	if err := j.validateSetGroupMembershipIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMembershipIdentifier",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetInternalValue(val *OrgLdapCustomSettingsGroupAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetMembership(val *string) {
	if err := j.validateSetMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membership",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetObjectClass(val *string) {
	if err := j.validateSetObjectClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectClass",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference)SetUniqueIdentifier(val *string) {
	if err := j.validateSetUniqueIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueIdentifier",
		val,
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ResetGroupBackLinkIdentifier() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupBackLinkIdentifier",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrgLdapCustomSettingsGroupAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

