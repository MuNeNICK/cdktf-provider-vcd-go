package nsxvfirewallrule

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxvfirewallrule/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvFirewallRuleSourceOutputReference interface {
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
	Exclude() interface{}
	SetExclude(val interface{})
	ExcludeInput() interface{}
	// Experimental.
	Fqn() *string
	GatewayInterfaces() *[]*string
	SetGatewayInterfaces(val *[]*string)
	GatewayInterfacesInput() *[]*string
	InternalValue() *NsxvFirewallRuleSource
	SetInternalValue(val *NsxvFirewallRuleSource)
	IpAddresses() *[]*string
	SetIpAddresses(val *[]*string)
	IpAddressesInput() *[]*string
	IpSets() *[]*string
	SetIpSets(val *[]*string)
	IpSetsInput() *[]*string
	OrgNetworks() *[]*string
	SetOrgNetworks(val *[]*string)
	OrgNetworksInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmIds() *[]*string
	SetVmIds(val *[]*string)
	VmIdsInput() *[]*string
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
	ResetExclude()
	ResetGatewayInterfaces()
	ResetIpAddresses()
	ResetIpSets()
	ResetOrgNetworks()
	ResetVmIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NsxvFirewallRuleSourceOutputReference
type jsiiProxy_NsxvFirewallRuleSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) Exclude() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GatewayInterfaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GatewayInterfacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) InternalValue() *NsxvFirewallRuleSource {
	var returns *NsxvFirewallRuleSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) IpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) IpSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) IpSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) OrgNetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) OrgNetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) VmIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference) VmIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmIdsInput",
		&returns,
	)
	return returns
}


func NewNsxvFirewallRuleSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NsxvFirewallRuleSourceOutputReference {
	_init_.Initialize()

	if err := validateNewNsxvFirewallRuleSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxvFirewallRuleSourceOutputReference{}

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNsxvFirewallRuleSourceOutputReference_Override(n NsxvFirewallRuleSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetExclude(val interface{}) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetGatewayInterfaces(val *[]*string) {
	if err := j.validateSetGatewayInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayInterfaces",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetInternalValue(val *NsxvFirewallRuleSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetIpAddresses(val *[]*string) {
	if err := j.validateSetIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetIpSets(val *[]*string) {
	if err := j.validateSetIpSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSets",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetOrgNetworks(val *[]*string) {
	if err := j.validateSetOrgNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgNetworks",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleSourceOutputReference)SetVmIds(val *[]*string) {
	if err := j.validateSetVmIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmIds",
		val,
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		n,
		"resetExclude",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetGatewayInterfaces() {
	_jsii_.InvokeVoid(
		n,
		"resetGatewayInterfaces",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetIpSets() {
	_jsii_.InvokeVoid(
		n,
		"resetIpSets",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetOrgNetworks() {
	_jsii_.InvokeVoid(
		n,
		"resetOrgNetworks",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ResetVmIds() {
	_jsii_.InvokeVoid(
		n,
		"resetVmIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NsxvFirewallRuleSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

