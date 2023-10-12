package nsxtedgegatewaybgpneighbor

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtedgegatewaybgpneighbor/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor vcd_nsxt_edgegateway_bgp_neighbor}.
type NsxtEdgegatewayBgpNeighbor interface {
	cdktf.TerraformResource
	AllowAsIn() interface{}
	SetAllowAsIn(val interface{})
	AllowAsInInput() interface{}
	BfdDeadMultiple() *float64
	SetBfdDeadMultiple(val *float64)
	BfdDeadMultipleInput() *float64
	BfdEnabled() interface{}
	SetBfdEnabled(val interface{})
	BfdEnabledInput() interface{}
	BfdInterval() *float64
	SetBfdInterval(val *float64)
	BfdIntervalInput() *float64
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
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GracefulRestartMode() *string
	SetGracefulRestartMode(val *string)
	GracefulRestartModeInput() *string
	HoldDownTimer() *float64
	SetHoldDownTimer(val *float64)
	HoldDownTimerInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	InFilterIpPrefixListId() *string
	SetInFilterIpPrefixListId(val *string)
	InFilterIpPrefixListIdInput() *string
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	KeepAliveTimer() *float64
	SetKeepAliveTimer(val *float64)
	KeepAliveTimerInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OutFilterIpPrefixListId() *string
	SetOutFilterIpPrefixListId(val *string)
	OutFilterIpPrefixListIdInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	RemoteAsNumber() *string
	SetRemoteAsNumber(val *string)
	RemoteAsNumberInput() *string
	RouteFiltering() *string
	SetRouteFiltering(val *string)
	RouteFilteringInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAllowAsIn()
	ResetBfdDeadMultiple()
	ResetBfdEnabled()
	ResetBfdInterval()
	ResetGracefulRestartMode()
	ResetHoldDownTimer()
	ResetId()
	ResetInFilterIpPrefixListId()
	ResetKeepAliveTimer()
	ResetOrg()
	ResetOutFilterIpPrefixListId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetRouteFiltering()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxtEdgegatewayBgpNeighbor
type jsiiProxy_NsxtEdgegatewayBgpNeighbor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) AllowAsIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAsIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) AllowAsInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAsInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdDeadMultiple() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bfdDeadMultiple",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdDeadMultipleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bfdDeadMultipleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bfdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bfdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bfdInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) BfdIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bfdIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GracefulRestartMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulRestartMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GracefulRestartModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulRestartModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) HoldDownTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdDownTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) HoldDownTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdDownTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) InFilterIpPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inFilterIpPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) InFilterIpPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inFilterIpPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) KeepAliveTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) KeepAliveTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) OutFilterIpPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outFilterIpPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) OutFilterIpPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outFilterIpPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) RemoteAsNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAsNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) RemoteAsNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAsNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) RouteFiltering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFiltering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) RouteFilteringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeFilteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor vcd_nsxt_edgegateway_bgp_neighbor} Resource.
func NewNsxtEdgegatewayBgpNeighbor(scope constructs.Construct, id *string, config *NsxtEdgegatewayBgpNeighborConfig) NsxtEdgegatewayBgpNeighbor {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewayBgpNeighborParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegatewayBgpNeighbor{}

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway_bgp_neighbor vcd_nsxt_edgegateway_bgp_neighbor} Resource.
func NewNsxtEdgegatewayBgpNeighbor_Override(n NsxtEdgegatewayBgpNeighbor, scope constructs.Construct, id *string, config *NsxtEdgegatewayBgpNeighborConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetAllowAsIn(val interface{}) {
	if err := j.validateSetAllowAsInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAsIn",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetBfdDeadMultiple(val *float64) {
	if err := j.validateSetBfdDeadMultipleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bfdDeadMultiple",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetBfdEnabled(val interface{}) {
	if err := j.validateSetBfdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bfdEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetBfdInterval(val *float64) {
	if err := j.validateSetBfdIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bfdInterval",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetGracefulRestartMode(val *string) {
	if err := j.validateSetGracefulRestartModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulRestartMode",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetHoldDownTimer(val *float64) {
	if err := j.validateSetHoldDownTimerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"holdDownTimer",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetInFilterIpPrefixListId(val *string) {
	if err := j.validateSetInFilterIpPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inFilterIpPrefixListId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetKeepAliveTimer(val *float64) {
	if err := j.validateSetKeepAliveTimerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveTimer",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetOutFilterIpPrefixListId(val *string) {
	if err := j.validateSetOutFilterIpPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outFilterIpPrefixListId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetRemoteAsNumber(val *string) {
	if err := j.validateSetRemoteAsNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteAsNumber",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegatewayBgpNeighbor)SetRouteFiltering(val *string) {
	if err := j.validateSetRouteFilteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeFiltering",
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
func NsxtEdgegatewayBgpNeighbor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpNeighbor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegatewayBgpNeighbor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpNeighbor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegatewayBgpNeighbor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegatewayBgpNeighbor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtEdgegatewayBgpNeighbor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtEdgegatewayBgpNeighbor.NsxtEdgegatewayBgpNeighbor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetAllowAsIn() {
	_jsii_.InvokeVoid(
		n,
		"resetAllowAsIn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetBfdDeadMultiple() {
	_jsii_.InvokeVoid(
		n,
		"resetBfdDeadMultiple",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetBfdEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetBfdEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetBfdInterval() {
	_jsii_.InvokeVoid(
		n,
		"resetBfdInterval",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetGracefulRestartMode() {
	_jsii_.InvokeVoid(
		n,
		"resetGracefulRestartMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetHoldDownTimer() {
	_jsii_.InvokeVoid(
		n,
		"resetHoldDownTimer",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetInFilterIpPrefixListId() {
	_jsii_.InvokeVoid(
		n,
		"resetInFilterIpPrefixListId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetKeepAliveTimer() {
	_jsii_.InvokeVoid(
		n,
		"resetKeepAliveTimer",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetOutFilterIpPrefixListId() {
	_jsii_.InvokeVoid(
		n,
		"resetOutFilterIpPrefixListId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetPassword() {
	_jsii_.InvokeVoid(
		n,
		"resetPassword",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ResetRouteFiltering() {
	_jsii_.InvokeVoid(
		n,
		"resetRouteFiltering",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegatewayBgpNeighbor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

