package nsxvfirewallrule

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/nsxvfirewallrule/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxvFirewallRuleDestinationOutputReference interface {
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
	InternalValue() *NsxvFirewallRuleDestination
	SetInternalValue(val *NsxvFirewallRuleDestination)
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

// The jsii proxy struct for NsxvFirewallRuleDestinationOutputReference
type jsiiProxy_NsxvFirewallRuleDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) Exclude() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GatewayInterfaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GatewayInterfacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) InternalValue() *NsxvFirewallRuleDestination {
	var returns *NsxvFirewallRuleDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) IpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) IpSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) IpSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) OrgNetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) OrgNetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"orgNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) VmIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) VmIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmIdsInput",
		&returns,
	)
	return returns
}


func NewNsxvFirewallRuleDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NsxvFirewallRuleDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewNsxvFirewallRuleDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxvFirewallRuleDestinationOutputReference{}

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRuleDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNsxvFirewallRuleDestinationOutputReference_Override(n NsxvFirewallRuleDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxvFirewallRule.NsxvFirewallRuleDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetExclude(val interface{}) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetGatewayInterfaces(val *[]*string) {
	if err := j.validateSetGatewayInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayInterfaces",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetInternalValue(val *NsxvFirewallRuleDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetIpAddresses(val *[]*string) {
	if err := j.validateSetIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetIpSets(val *[]*string) {
	if err := j.validateSetIpSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSets",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetOrgNetworks(val *[]*string) {
	if err := j.validateSetOrgNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgNetworks",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NsxvFirewallRuleDestinationOutputReference)SetVmIds(val *[]*string) {
	if err := j.validateSetVmIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmIds",
		val,
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		n,
		"resetExclude",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetGatewayInterfaces() {
	_jsii_.InvokeVoid(
		n,
		"resetGatewayInterfaces",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetIpSets() {
	_jsii_.InvokeVoid(
		n,
		"resetIpSets",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetOrgNetworks() {
	_jsii_.InvokeVoid(
		n,
		"resetOrgNetworks",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ResetVmIds() {
	_jsii_.InvokeVoid(
		n,
		"resetVmIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NsxvFirewallRuleDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

