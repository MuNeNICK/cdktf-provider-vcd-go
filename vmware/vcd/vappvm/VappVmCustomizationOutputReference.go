package vappvm

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/vappvm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappVmCustomizationOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AllowLocalAdminPassword() interface{}
	SetAllowLocalAdminPassword(val interface{})
	AllowLocalAdminPasswordInput() interface{}
	AutoGeneratePassword() interface{}
	SetAutoGeneratePassword(val interface{})
	AutoGeneratePasswordInput() interface{}
	ChangeSid() interface{}
	SetChangeSid(val interface{})
	ChangeSidInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Force() interface{}
	SetForce(val interface{})
	ForceInput() interface{}
	// Experimental.
	Fqn() *string
	Initscript() *string
	SetInitscript(val *string)
	InitscriptInput() *string
	InternalValue() *VappVmCustomization
	SetInternalValue(val *VappVmCustomization)
	JoinDomain() interface{}
	SetJoinDomain(val interface{})
	JoinDomainAccountOu() *string
	SetJoinDomainAccountOu(val *string)
	JoinDomainAccountOuInput() *string
	JoinDomainInput() interface{}
	JoinDomainName() *string
	SetJoinDomainName(val *string)
	JoinDomainNameInput() *string
	JoinDomainPassword() *string
	SetJoinDomainPassword(val *string)
	JoinDomainPasswordInput() *string
	JoinDomainUser() *string
	SetJoinDomainUser(val *string)
	JoinDomainUserInput() *string
	JoinOrgDomain() interface{}
	SetJoinOrgDomain(val interface{})
	JoinOrgDomainInput() interface{}
	MustChangePasswordOnFirstLogin() interface{}
	SetMustChangePasswordOnFirstLogin(val interface{})
	MustChangePasswordOnFirstLoginInput() interface{}
	NumberOfAutoLogons() *float64
	SetNumberOfAutoLogons(val *float64)
	NumberOfAutoLogonsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAdminPassword()
	ResetAllowLocalAdminPassword()
	ResetAutoGeneratePassword()
	ResetChangeSid()
	ResetEnabled()
	ResetForce()
	ResetInitscript()
	ResetJoinDomain()
	ResetJoinDomainAccountOu()
	ResetJoinDomainName()
	ResetJoinDomainPassword()
	ResetJoinDomainUser()
	ResetJoinOrgDomain()
	ResetMustChangePasswordOnFirstLogin()
	ResetNumberOfAutoLogons()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VappVmCustomizationOutputReference
type jsiiProxy_VappVmCustomizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AllowLocalAdminPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AllowLocalAdminPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AutoGeneratePassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGeneratePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) AutoGeneratePasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGeneratePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) ChangeSid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) ChangeSidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) Force() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) ForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) Initscript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initscript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) InitscriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initscriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) InternalValue() *VappVmCustomization {
	var returns *VappVmCustomization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainAccountOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainAccountOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainAccountOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainAccountOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinDomainUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinOrgDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinOrgDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) JoinOrgDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinOrgDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) MustChangePasswordOnFirstLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustChangePasswordOnFirstLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) MustChangePasswordOnFirstLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustChangePasswordOnFirstLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) NumberOfAutoLogons() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfAutoLogons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) NumberOfAutoLogonsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfAutoLogonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappVmCustomizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVappVmCustomizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VappVmCustomizationOutputReference {
	_init_.Initialize()

	if err := validateNewVappVmCustomizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappVmCustomizationOutputReference{}

	_jsii_.Create(
		"vcd.vappVm.VappVmCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVappVmCustomizationOutputReference_Override(v VappVmCustomizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.vappVm.VappVmCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetAllowLocalAdminPassword(val interface{}) {
	if err := j.validateSetAllowLocalAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLocalAdminPassword",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetAutoGeneratePassword(val interface{}) {
	if err := j.validateSetAutoGeneratePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoGeneratePassword",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetChangeSid(val interface{}) {
	if err := j.validateSetChangeSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeSid",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetForce(val interface{}) {
	if err := j.validateSetForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"force",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetInitscript(val *string) {
	if err := j.validateSetInitscriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initscript",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetInternalValue(val *VappVmCustomization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinDomain(val interface{}) {
	if err := j.validateSetJoinDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomain",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinDomainAccountOu(val *string) {
	if err := j.validateSetJoinDomainAccountOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomainAccountOu",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinDomainName(val *string) {
	if err := j.validateSetJoinDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomainName",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinDomainPassword(val *string) {
	if err := j.validateSetJoinDomainPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomainPassword",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinDomainUser(val *string) {
	if err := j.validateSetJoinDomainUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomainUser",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetJoinOrgDomain(val interface{}) {
	if err := j.validateSetJoinOrgDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinOrgDomain",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetMustChangePasswordOnFirstLogin(val interface{}) {
	if err := j.validateSetMustChangePasswordOnFirstLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mustChangePasswordOnFirstLogin",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetNumberOfAutoLogons(val *float64) {
	if err := j.validateSetNumberOfAutoLogonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfAutoLogons",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VappVmCustomizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetAllowLocalAdminPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowLocalAdminPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetAutoGeneratePassword() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoGeneratePassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetChangeSid() {
	_jsii_.InvokeVoid(
		v,
		"resetChangeSid",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetForce() {
	_jsii_.InvokeVoid(
		v,
		"resetForce",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetInitscript() {
	_jsii_.InvokeVoid(
		v,
		"resetInitscript",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinDomain() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomain",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinDomainAccountOu() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomainAccountOu",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinDomainName() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomainName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinDomainPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomainPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinDomainUser() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomainUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetJoinOrgDomain() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinOrgDomain",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetMustChangePasswordOnFirstLogin() {
	_jsii_.InvokeVoid(
		v,
		"resetMustChangePasswordOnFirstLogin",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ResetNumberOfAutoLogons() {
	_jsii_.InvokeVoid(
		v,
		"resetNumberOfAutoLogons",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappVmCustomizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

