package datavcdsubscribedcatalog

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdsubscribedcatalog/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdSubscribedCatalogFilterOutputReference interface {
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
	Date() *string
	SetDate(val *string)
	DateInput() *string
	Earliest() interface{}
	SetEarliest(val interface{})
	EarliestInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataVcdSubscribedCatalogFilter
	SetInternalValue(val *DataVcdSubscribedCatalogFilter)
	Latest() interface{}
	SetLatest(val interface{})
	LatestInput() interface{}
	Metadata() DataVcdSubscribedCatalogFilterMetadataList
	MetadataInput() interface{}
	NameRegex() *string
	SetNameRegex(val *string)
	NameRegexInput() *string
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
	PutMetadata(value interface{})
	ResetDate()
	ResetEarliest()
	ResetLatest()
	ResetMetadata()
	ResetNameRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataVcdSubscribedCatalogFilterOutputReference
type jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Date() *string {
	var returns *string
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) DateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Earliest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earliest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) EarliestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earliestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) InternalValue() *DataVcdSubscribedCatalogFilter {
	var returns *DataVcdSubscribedCatalogFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) LatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Metadata() DataVcdSubscribedCatalogFilterMetadataList {
	var returns DataVcdSubscribedCatalogFilterMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) NameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) NameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataVcdSubscribedCatalogFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataVcdSubscribedCatalogFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataVcdSubscribedCatalogFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference{}

	_jsii_.Create(
		"vcd.dataVcdSubscribedCatalog.DataVcdSubscribedCatalogFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataVcdSubscribedCatalogFilterOutputReference_Override(d DataVcdSubscribedCatalogFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdSubscribedCatalog.DataVcdSubscribedCatalogFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetDate(val *string) {
	if err := j.validateSetDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetEarliest(val interface{}) {
	if err := j.validateSetEarliestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earliest",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetInternalValue(val *DataVcdSubscribedCatalogFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetLatest(val interface{}) {
	if err := j.validateSetLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latest",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetNameRegex(val *string) {
	if err := j.validateSetNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameRegex",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) PutMetadata(value interface{}) {
	if err := d.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetadata",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ResetDate() {
	_jsii_.InvokeVoid(
		d,
		"resetDate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ResetEarliest() {
	_jsii_.InvokeVoid(
		d,
		"resetEarliest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ResetLatest() {
	_jsii_.InvokeVoid(
		d,
		"resetLatest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ResetNameRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetNameRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataVcdSubscribedCatalogFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

