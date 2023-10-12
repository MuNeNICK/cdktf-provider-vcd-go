package nsxtipsecvpntunnel

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/nsxtipsecvpntunnel/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel vcd_nsxt_ipsec_vpn_tunnel}.
type NsxtIpsecVpnTunnel interface {
	cdktf.TerraformResource
	AuthenticationMode() *string
	SetAuthenticationMode(val *string)
	AuthenticationModeInput() *string
	CaCertificateId() *string
	SetCaCertificateId(val *string)
	CaCertificateIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
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
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	IkeFailReason() *string
	IkeServiceStatus() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalIpAddress() *string
	SetLocalIpAddress(val *string)
	LocalIpAddressInput() *string
	LocalNetworks() *[]*string
	SetLocalNetworks(val *[]*string)
	LocalNetworksInput() *[]*string
	Logging() interface{}
	SetLogging(val interface{})
	LoggingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	PreSharedKey() *string
	SetPreSharedKey(val *string)
	PreSharedKeyInput() *string
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
	RemoteId() *string
	SetRemoteId(val *string)
	RemoteIdInput() *string
	RemoteIpAddress() *string
	SetRemoteIpAddress(val *string)
	RemoteIpAddressInput() *string
	RemoteNetworks() *[]*string
	SetRemoteNetworks(val *[]*string)
	RemoteNetworksInput() *[]*string
	SecurityProfile() *string
	SecurityProfileCustomization() NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference
	SecurityProfileCustomizationInput() *NsxtIpsecVpnTunnelSecurityProfileCustomization
	Status() *string
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
	PutSecurityProfileCustomization(value *NsxtIpsecVpnTunnelSecurityProfileCustomization)
	ResetAuthenticationMode()
	ResetCaCertificateId()
	ResetCertificateId()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetLogging()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoteId()
	ResetRemoteNetworks()
	ResetSecurityProfileCustomization()
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

// The jsii proxy struct for NsxtIpsecVpnTunnel
type jsiiProxy_NsxtIpsecVpnTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) AuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) AuthenticationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) CaCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) CaCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) IkeFailReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeFailReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) IkeServiceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeServiceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) LocalIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) LocalIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) LocalNetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) LocalNetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) PreSharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) PreSharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteNetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) RemoteNetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) SecurityProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) SecurityProfileCustomization() NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference {
	var returns NsxtIpsecVpnTunnelSecurityProfileCustomizationOutputReference
	_jsii_.Get(
		j,
		"securityProfileCustomization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) SecurityProfileCustomizationInput() *NsxtIpsecVpnTunnelSecurityProfileCustomization {
	var returns *NsxtIpsecVpnTunnelSecurityProfileCustomization
	_jsii_.Get(
		j,
		"securityProfileCustomizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel vcd_nsxt_ipsec_vpn_tunnel} Resource.
func NewNsxtIpsecVpnTunnel(scope constructs.Construct, id *string, config *NsxtIpsecVpnTunnelConfig) NsxtIpsecVpnTunnel {
	_init_.Initialize()

	if err := validateNewNsxtIpsecVpnTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtIpsecVpnTunnel{}

	_jsii_.Create(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_ipsec_vpn_tunnel vcd_nsxt_ipsec_vpn_tunnel} Resource.
func NewNsxtIpsecVpnTunnel_Override(n NsxtIpsecVpnTunnel, scope constructs.Construct, id *string, config *NsxtIpsecVpnTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetAuthenticationMode(val *string) {
	if err := j.validateSetAuthenticationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMode",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetCaCertificateId(val *string) {
	if err := j.validateSetCaCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateId",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetLocalIpAddress(val *string) {
	if err := j.validateSetLocalIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetLocalNetworks(val *[]*string) {
	if err := j.validateSetLocalNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localNetworks",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetLogging(val interface{}) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetPreSharedKey(val *string) {
	if err := j.validateSetPreSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKey",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetRemoteId(val *string) {
	if err := j.validateSetRemoteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteId",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetRemoteIpAddress(val *string) {
	if err := j.validateSetRemoteIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetRemoteNetworks(val *[]*string) {
	if err := j.validateSetRemoteNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteNetworks",
		val,
	)
}

func (j *jsiiProxy_NsxtIpsecVpnTunnel)SetVdc(val *string) {
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
func NsxtIpsecVpnTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtIpsecVpnTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtIpsecVpnTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtIpsecVpnTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtIpsecVpnTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtIpsecVpnTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtIpsecVpnTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtIpsecVpnTunnel.NsxtIpsecVpnTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtIpsecVpnTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) PutSecurityProfileCustomization(value *NsxtIpsecVpnTunnelSecurityProfileCustomization) {
	if err := n.validatePutSecurityProfileCustomizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSecurityProfileCustomization",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetAuthenticationMode() {
	_jsii_.InvokeVoid(
		n,
		"resetAuthenticationMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetCaCertificateId() {
	_jsii_.InvokeVoid(
		n,
		"resetCaCertificateId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetCertificateId() {
	_jsii_.InvokeVoid(
		n,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetLogging() {
	_jsii_.InvokeVoid(
		n,
		"resetLogging",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetRemoteId() {
	_jsii_.InvokeVoid(
		n,
		"resetRemoteId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetRemoteNetworks() {
	_jsii_.InvokeVoid(
		n,
		"resetRemoteNetworks",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetSecurityProfileCustomization() {
	_jsii_.InvokeVoid(
		n,
		"resetSecurityProfileCustomization",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtIpsecVpnTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

