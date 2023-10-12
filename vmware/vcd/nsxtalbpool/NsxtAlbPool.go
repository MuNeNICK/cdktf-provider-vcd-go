package nsxtalbpool

import (
	_init_ "github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/munenick/cdktf-provider-vcd-go/vmware/vcd/nsxtalbpool/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool vcd_nsxt_alb_pool}.
type NsxtAlbPool interface {
	cdktf.TerraformResource
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
	AssociatedVirtualServiceIds() *[]*string
	AssociatedVirtualServices() *[]*string
	CaCertificateIds() *[]*string
	SetCaCertificateIds(val *[]*string)
	CaCertificateIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CnCheckEnabled() interface{}
	SetCnCheckEnabled(val interface{})
	CnCheckEnabledInput() interface{}
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
	DefaultPort() *float64
	SetDefaultPort(val *float64)
	DefaultPortInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DomainNames() *[]*string
	SetDomainNames(val *[]*string)
	DomainNamesInput() *[]*string
	EdgeGatewayId() *string
	SetEdgeGatewayId(val *string)
	EdgeGatewayIdInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnabledMemberCount() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GracefulTimeoutPeriod() *float64
	SetGracefulTimeoutPeriod(val *float64)
	GracefulTimeoutPeriodInput() *float64
	HealthMessage() *string
	HealthMonitor() NsxtAlbPoolHealthMonitorList
	HealthMonitorInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Member() NsxtAlbPoolMemberList
	MemberCount() *float64
	MemberGroupId() *string
	SetMemberGroupId(val *string)
	MemberGroupIdInput() *string
	MemberInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	PassiveMonitoringEnabled() interface{}
	SetPassiveMonitoringEnabled(val interface{})
	PassiveMonitoringEnabledInput() interface{}
	PersistenceProfile() NsxtAlbPoolPersistenceProfileOutputReference
	PersistenceProfileInput() *NsxtAlbPoolPersistenceProfile
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpMemberCount() *float64
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
	PutHealthMonitor(value interface{})
	PutMember(value interface{})
	PutPersistenceProfile(value *NsxtAlbPoolPersistenceProfile)
	ResetAlgorithm()
	ResetCaCertificateIds()
	ResetCnCheckEnabled()
	ResetDefaultPort()
	ResetDescription()
	ResetDomainNames()
	ResetEnabled()
	ResetGracefulTimeoutPeriod()
	ResetHealthMonitor()
	ResetId()
	ResetMember()
	ResetMemberGroupId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassiveMonitoringEnabled()
	ResetPersistenceProfile()
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

// The jsii proxy struct for NsxtAlbPool
type jsiiProxy_NsxtAlbPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtAlbPool) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) AssociatedVirtualServiceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedVirtualServiceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) AssociatedVirtualServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedVirtualServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) CaCertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"caCertificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) CaCertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"caCertificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) CnCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cnCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) CnCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cnCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DefaultPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) EdgeGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) EdgeGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) EnabledMemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledMemberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) GracefulTimeoutPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracefulTimeoutPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) GracefulTimeoutPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gracefulTimeoutPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) HealthMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) HealthMonitor() NsxtAlbPoolHealthMonitorList {
	var returns NsxtAlbPoolHealthMonitorList
	_jsii_.Get(
		j,
		"healthMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) HealthMonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Member() NsxtAlbPoolMemberList {
	var returns NsxtAlbPoolMemberList
	_jsii_.Get(
		j,
		"member",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) MemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) MemberGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) MemberGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) MemberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) PassiveMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passiveMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) PassiveMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passiveMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) PersistenceProfile() NsxtAlbPoolPersistenceProfileOutputReference {
	var returns NsxtAlbPoolPersistenceProfileOutputReference
	_jsii_.Get(
		j,
		"persistenceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) PersistenceProfileInput() *NsxtAlbPoolPersistenceProfile {
	var returns *NsxtAlbPoolPersistenceProfile
	_jsii_.Get(
		j,
		"persistenceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) UpMemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upMemberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtAlbPool) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool vcd_nsxt_alb_pool} Resource.
func NewNsxtAlbPool(scope constructs.Construct, id *string, config *NsxtAlbPoolConfig) NsxtAlbPool {
	_init_.Initialize()

	if err := validateNewNsxtAlbPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtAlbPool{}

	_jsii_.Create(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool vcd_nsxt_alb_pool} Resource.
func NewNsxtAlbPool_Override(n NsxtAlbPool, scope constructs.Construct, id *string, config *NsxtAlbPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetCaCertificateIds(val *[]*string) {
	if err := j.validateSetCaCertificateIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIds",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetCnCheckEnabled(val interface{}) {
	if err := j.validateSetCnCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetDefaultPort(val *float64) {
	if err := j.validateSetDefaultPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPort",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetEdgeGatewayId(val *string) {
	if err := j.validateSetEdgeGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeGatewayId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetGracefulTimeoutPeriod(val *float64) {
	if err := j.validateSetGracefulTimeoutPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulTimeoutPeriod",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetMemberGroupId(val *string) {
	if err := j.validateSetMemberGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberGroupId",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetPassiveMonitoringEnabled(val interface{}) {
	if err := j.validateSetPassiveMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtAlbPool)SetVdc(val *string) {
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
func NsxtAlbPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtAlbPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtAlbPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtAlbPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtAlbPool.NsxtAlbPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtAlbPool) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtAlbPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtAlbPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtAlbPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtAlbPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtAlbPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtAlbPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtAlbPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtAlbPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtAlbPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtAlbPool) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtAlbPool) PutHealthMonitor(value interface{}) {
	if err := n.validatePutHealthMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHealthMonitor",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtAlbPool) PutMember(value interface{}) {
	if err := n.validatePutMemberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMember",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtAlbPool) PutPersistenceProfile(value *NsxtAlbPoolPersistenceProfile) {
	if err := n.validatePutPersistenceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPersistenceProfile",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		n,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetCaCertificateIds() {
	_jsii_.InvokeVoid(
		n,
		"resetCaCertificateIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetCnCheckEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetCnCheckEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetDefaultPort() {
	_jsii_.InvokeVoid(
		n,
		"resetDefaultPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetDomainNames() {
	_jsii_.InvokeVoid(
		n,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetGracefulTimeoutPeriod() {
	_jsii_.InvokeVoid(
		n,
		"resetGracefulTimeoutPeriod",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetHealthMonitor() {
	_jsii_.InvokeVoid(
		n,
		"resetHealthMonitor",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetMember() {
	_jsii_.InvokeVoid(
		n,
		"resetMember",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetMemberGroupId() {
	_jsii_.InvokeVoid(
		n,
		"resetMemberGroupId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetPassiveMonitoringEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetPassiveMonitoringEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetPersistenceProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetPersistenceProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtAlbPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtAlbPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

