package datavcdvappvm

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdvappvm/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdVappVmCustomizationOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	AllowLocalAdminPassword() cdktf.IResolvable
	AutoGeneratePassword() cdktf.IResolvable
	ChangeSid() cdktf.IResolvable
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
	Enabled() cdktf.IResolvable
	Force() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	Initscript() *string
	InternalValue() *DataVcdVappVmCustomization
	SetInternalValue(val *DataVcdVappVmCustomization)
	JoinDomain() cdktf.IResolvable
	JoinDomainAccountOu() *string
	JoinDomainName() *string
	JoinDomainPassword() *string
	JoinDomainUser() *string
	JoinOrgDomain() cdktf.IResolvable
	MustChangePasswordOnFirstLogin() cdktf.IResolvable
	NumberOfAutoLogons() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataVcdVappVmCustomizationOutputReference
type jsiiProxy_DataVcdVappVmCustomizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) AllowLocalAdminPassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowLocalAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) AutoGeneratePassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoGeneratePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) ChangeSid() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"changeSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) Force() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) Initscript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initscript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) InternalValue() *DataVcdVappVmCustomization {
	var returns *DataVcdVappVmCustomization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinDomain() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"joinDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinDomainAccountOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainAccountOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinDomainPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinDomainUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) JoinOrgDomain() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"joinOrgDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) MustChangePasswordOnFirstLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mustChangePasswordOnFirstLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) NumberOfAutoLogons() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfAutoLogons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataVcdVappVmCustomizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataVcdVappVmCustomizationOutputReference {
	_init_.Initialize()

	if err := validateNewDataVcdVappVmCustomizationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdVappVmCustomizationOutputReference{}

	_jsii_.Create(
		"vcd.dataVcdVappVm.DataVcdVappVmCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataVcdVappVmCustomizationOutputReference_Override(d DataVcdVappVmCustomizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdVappVm.DataVcdVappVmCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference)SetInternalValue(val *DataVcdVappVmCustomization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVcdVappVmCustomizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdVappVmCustomizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

