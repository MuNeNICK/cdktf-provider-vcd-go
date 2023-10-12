package nsxtipsecvpntunnel

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtipsecvpntunnel/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference interface {
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
	SetDpdProbeInternal(val *float64)
	DpdProbeInternalInput() *float64
	// Experimental.
	Fqn() *string
	IkeDhGroups() *[]*string
	SetIkeDhGroups(val *[]*string)
	IkeDhGroupsInput() *[]*string
	IkeDigestAlgorithms() *[]*string
	SetIkeDigestAlgorithms(val *[]*string)
	IkeDigestAlgorithmsInput() *[]*string
	IkeEncryptionAlgorithms() *[]*string
	SetIkeEncryptionAlgorithms(val *[]*string)
	IkeEncryptionAlgorithmsInput() *[]*string
	IkeSaLifetime() *float64
	SetIkeSaLifetime(val *float64)
	IkeSaLifetimeInput() *float64
	IkeVersion() *string
	SetIkeVersion(val *string)
	IkeVersionInput() *string
	InternalValue() *NsxtIpsecVpnTunnelSecurityProfileCustomization
	SetInternalValue(val *NsxtIpsecVpnTunnelSecurityProfileCustomization)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TunnelDfPolicy() *string
	SetTunnelDfPolicy(val *string)
	TunnelDfPolicyInput() *string
	TunnelDhGroups() *[]*string
	SetTunnelDhGroups(val *[]*string)
	TunnelDhGroupsInput() *[]*string
	TunnelDigestAlgorithms() *[]*string
	SetTunnelDigestAlgorithms(val *[]*string)
	TunnelDigestAlgorithmsInput() *[]*string
	TunnelEncryptionAlgorithms() *[]*string
	SetTunnelEncryptionAlgorithms(val *[]*string)
	TunnelEncryptionAlgorithmsInput() *[]*string
	TunnelPfsEnabled() interface{}
	SetTunnelPfsEnabled(val interface{})
	TunnelPfsEnabledInput() interface{}
	TunnelSaLifetime() *float64
	SetTunnelSaLifetime(val *float64)
	TunnelSaLifetimeInput() *float64
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
	ResetDpdProbeInternal()
	ResetIkeDigestAlgorithms()
	ResetIkeSaLifetime()
	ResetTunnelDfPolicy()
	ResetTunnelDigestAlgorithms()
	ResetTunnelPfsEnabled()
	ResetTunnelSaLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference
type jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) DpdProbeInternal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdProbeInternal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) DpdProbeInternalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdProbeInternalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDhGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDhGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDhGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDhGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDigestAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDigestAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeDigestAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeDigestAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeEncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeEncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeSaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeSaLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeSaLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeSaLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) IkeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InternalValue() *NsxtIpsecVpnTunnelSecurityProfileCustomization {
	var returns *NsxtIpsecVpnTunnelSecurityProfileCustomization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDfPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelDfPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDfPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelDfPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDhGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDhGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDhGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDhGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDigestAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDigestAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelDigestAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelDigestAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelEncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelEncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelEncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnelEncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelPfsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnelPfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelPfsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnelPfsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelSaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnelSaLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) TunnelSaLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnelSaLifetimeInput",
		&returns,
	)
	return returns
}


func NewNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference {
	_init_.Initialize()

	if err := validateNewNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference{}

	_jsii_.Create(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference_Override(n NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetDpdProbeInternal(val *float64) {
	if err := j.validateSetDpdProbeInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdProbeInternal",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetIkeDhGroups(val *[]*string) {
	if err := j.validateSetIkeDhGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeDhGroups",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetIkeDigestAlgorithms(val *[]*string) {
	if err := j.validateSetIkeDigestAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeDigestAlgorithms",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetIkeEncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetIkeEncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeEncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetIkeSaLifetime(val *float64) {
	if err := j.validateSetIkeSaLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeSaLifetime",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetIkeVersion(val *string) {
	if err := j.validateSetIkeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeVersion",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetInternalValue(val *NsxtIpsecVpnTunnelSecurityProfileCustomization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelDfPolicy(val *string) {
	if err := j.validateSetTunnelDfPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelDfPolicy",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelDhGroups(val *[]*string) {
	if err := j.validateSetTunnelDhGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelDhGroups",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelDigestAlgorithms(val *[]*string) {
	if err := j.validateSetTunnelDigestAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelDigestAlgorithms",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelEncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnelEncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelEncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelPfsEnabled(val interface{}) {
	if err := j.validateSetTunnelPfsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelPfsEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference)SetTunnelSaLifetime(val *float64) {
	if err := j.validateSetTunnelSaLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelSaLifetime",
		val,
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetDpdProbeInternal() {
	_jsii_.InvokeVoid(
		n,
		"resetDpdProbeInternal",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetIkeDigestAlgorithms() {
	_jsii_.InvokeVoid(
		n,
		"resetIkeDigestAlgorithms",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetIkeSaLifetime() {
	_jsii_.InvokeVoid(
		n,
		"resetIkeSaLifetime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetTunnelDfPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelDfPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetTunnelDigestAlgorithms() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelDigestAlgorithms",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetTunnelPfsEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelPfsEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ResetTunnelSaLifetime() {
	_jsii_.InvokeVoid(
		n,
		"resetTunnelSaLifetime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

