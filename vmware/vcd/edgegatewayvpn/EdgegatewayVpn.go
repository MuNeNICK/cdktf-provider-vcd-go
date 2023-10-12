package edgegatewayvpn

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/edgegatewayvpn/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn vcd_edgegateway_vpn}.
type EdgegatewayVpn interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeGateway() *string
	SetEdgeGateway(val *string)
	EdgeGatewayInput() *string
	EncryptionProtocol() *string
	SetEncryptionProtocol(val *string)
	EncryptionProtocolInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalId() *string
	SetLocalId(val *string)
	LocalIdInput() *string
	LocalIpAddress() *string
	SetLocalIpAddress(val *string)
	LocalIpAddressInput() *string
	LocalSubnets() EdgegatewayVpnLocalSubnetsList
	LocalSubnetsInput() interface{}
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	PeerId() *string
	SetPeerId(val *string)
	PeerIdInput() *string
	PeerIpAddress() *string
	SetPeerIpAddress(val *string)
	PeerIpAddressInput() *string
	PeerSubnets() EdgegatewayVpnPeerSubnetsList
	PeerSubnetsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SharedSecret() *string
	SetSharedSecret(val *string)
	SharedSecretInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Vdc() *string
	SetVdc(val *string)
	VdcInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLocalSubnets(value interface{})
	PutPeerSubnets(value interface{})
	ResetDescription()
	ResetId()
	ResetLocalSubnets()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerSubnets()
	ResetVdc()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EdgegatewayVpn
type jsiiProxy_EdgegatewayVpn struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EdgegatewayVpn) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) EdgeGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) EdgeGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) EncryptionProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) EncryptionProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalSubnets() EdgegatewayVpnLocalSubnetsList {
	var returns EdgegatewayVpnLocalSubnetsList
	_jsii_.Get(
		j,
		"localSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) LocalSubnetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerSubnets() EdgegatewayVpnPeerSubnetsList {
	var returns EdgegatewayVpnPeerSubnetsList
	_jsii_.Get(
		j,
		"peerSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) PeerSubnetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"peerSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) SharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) SharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgegatewayVpn) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn vcd_edgegateway_vpn} Resource.
func NewEdgegatewayVpn(scope constructs.Construct, id *string, config *EdgegatewayVpnConfig) EdgegatewayVpn {
	_init_.Initialize()

	if err := validateNewEdgegatewayVpnParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgegatewayVpn{}

	_jsii_.Create(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway_vpn vcd_edgegateway_vpn} Resource.
func NewEdgegatewayVpn_Override(e EdgegatewayVpn, scope constructs.Construct, id *string, config *EdgegatewayVpnConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetEdgeGateway(val *string) {
	if err := j.validateSetEdgeGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGateway",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetEncryptionProtocol(val *string) {
	if err := j.validateSetEncryptionProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionProtocol",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetLocalId(val *string) {
	if err := j.validateSetLocalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localId",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetLocalIpAddress(val *string) {
	if err := j.validateSetLocalIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpAddress",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetPeerId(val *string) {
	if err := j.validateSetPeerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerId",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetPeerIpAddress(val *string) {
	if err := j.validateSetPeerIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddress",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetSharedSecret(val *string) {
	if err := j.validateSetSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedSecret",
		val,
	)
}

func (j *jsiiProxy_EdgegatewayVpn)SetVdc(val *string) {
	if err := j.validateSetVdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdc",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EdgegatewayVpn_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewayVpn_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgegatewayVpn_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewayVpn_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgegatewayVpn_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgegatewayVpn_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EdgegatewayVpn_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.edgegatewayVpn.EdgegatewayVpn",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EdgegatewayVpn) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EdgegatewayVpn) PutLocalSubnets(value interface{}) {
	if err := e.validatePutLocalSubnetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLocalSubnets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgegatewayVpn) PutPeerSubnets(value interface{}) {
	if err := e.validatePutPeerSubnetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPeerSubnets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetLocalSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetOrg() {
	_jsii_.InvokeVoid(
		e,
		"resetOrg",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetPeerSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetPeerSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) ResetVdc() {
	_jsii_.InvokeVoid(
		e,
		"resetVdc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgegatewayVpn) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgegatewayVpn) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

