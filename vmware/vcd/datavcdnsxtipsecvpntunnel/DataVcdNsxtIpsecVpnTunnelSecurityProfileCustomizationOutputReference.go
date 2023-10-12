package datavcdnsxtipsecvpntunnel

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/datavcdnsxtipsecvpntunnel/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference interface {
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
	DpdProbeInternal() *float64
	// Experimental.
	Fqn() *string
	IkeDhGroups() *[]*string
	IkeDigestAlgorithms() *[]*string
	IkeEncryptionAlgorithms() *[]*string
	IkeSaLifetime() *float64
	IkeVersion() *string
	InternalValue() *DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomization
	SetInternalValue(val *DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomization)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelDfPolicy() *string
	TunnelDhGroups() *[]*string
	TunnelDigestAlgorithms() *[]*string
	TunnelEncryptionAlgorithms() *[]*string
	TunnelPfsEnabled() cdktf.IResolvable
	TunnelSaLifetime() *float64
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

// The jsii proxy struct for DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference
type jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) DpdProbeInternal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdProbeInternal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDhGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDhGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDigestAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDigestAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeEncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeSaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeSaLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InternalValue() *DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomization {
	var returns *DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDfPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelDfPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDhGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDhGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDigestAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDigestAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelEncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelEncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelPfsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tunnelPfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelSaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnelSaLifetime",
		&returns,
	)
	return returns
}


func NewDataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference {
	_init_.Initialize()

	if err := validateNewDataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference{}

	_jsii_.Create(
		"vcd.dataVcdNsxtIpsecVpnTunnel.DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference_Override(d DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.dataVcdNsxtIpsecVpnTunnel.DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetInternalValue(val *DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataVcdNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

