package nsxtdistributedfirewallrule

import (
	_init_ "app/internal/cdktf/generated/vmware/vcd/jsii"

	"app/internal/cdktf/generated/vmware/vcd/nsxtdistributedfirewallrule/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule vcd_nsxt_distributed_firewall_rule}.
type NsxtDistributedFirewallRuleA interface {
	cdktf.TerraformResource
	AboveRuleId() *string
	SetAboveRuleId(val *string)
	AboveRuleIdInput() *string
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	AppPortProfileIds() *[]*string
	SetAppPortProfileIds(val *[]*string)
	AppPortProfileIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	DestinationGroupsExcluded() interface{}
	SetDestinationGroupsExcluded(val interface{})
	DestinationGroupsExcludedInput() interface{}
	DestinationIds() *[]*string
	SetDestinationIds(val *[]*string)
	DestinationIdsInput() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
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
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() interface{}
	SetLogging(val interface{})
	LoggingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkContextProfileIds() *[]*string
	SetNetworkContextProfileIds(val *[]*string)
	NetworkContextProfileIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
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
	SourceGroupsExcluded() interface{}
	SetSourceGroupsExcluded(val interface{})
	SourceGroupsExcludedInput() interface{}
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	SourceIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VdcGroupId() *string
	SetVdcGroupId(val *string)
	VdcGroupIdInput() *string
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
	ResetAboveRuleId()
	ResetAppPortProfileIds()
	ResetComment()
	ResetDescription()
	ResetDestinationGroupsExcluded()
	ResetDestinationIds()
	ResetDirection()
	ResetEnabled()
	ResetId()
	ResetIpProtocol()
	ResetLogging()
	ResetNetworkContextProfileIds()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceGroupsExcluded()
	ResetSourceIds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NsxtDistributedFirewallRuleA
type jsiiProxy_NsxtDistributedFirewallRuleA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) AboveRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboveRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) AboveRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboveRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) AppPortProfileIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appPortProfileIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) AppPortProfileIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appPortProfileIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DestinationGroupsExcluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationGroupsExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DestinationGroupsExcludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationGroupsExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DestinationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DestinationIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) NetworkContextProfileIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkContextProfileIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) NetworkContextProfileIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkContextProfileIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) SourceGroupsExcluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceGroupsExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) SourceGroupsExcludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceGroupsExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) SourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) VdcGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA) VdcGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcGroupIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule vcd_nsxt_distributed_firewall_rule} Resource.
func NewNsxtDistributedFirewallRuleA(scope constructs.Construct, id *string, config *NsxtDistributedFirewallRuleAConfig) NsxtDistributedFirewallRuleA {
	_init_.Initialize()

	if err := validateNewNsxtDistributedFirewallRuleAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtDistributedFirewallRuleA{}

	_jsii_.Create(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_distributed_firewall_rule vcd_nsxt_distributed_firewall_rule} Resource.
func NewNsxtDistributedFirewallRuleA_Override(n NsxtDistributedFirewallRuleA, scope constructs.Construct, id *string, config *NsxtDistributedFirewallRuleAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetAboveRuleId(val *string) {
	if err := j.validateSetAboveRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aboveRuleId",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetAppPortProfileIds(val *[]*string) {
	if err := j.validateSetAppPortProfileIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appPortProfileIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetDestinationGroupsExcluded(val interface{}) {
	if err := j.validateSetDestinationGroupsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationGroupsExcluded",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetDestinationIds(val *[]*string) {
	if err := j.validateSetDestinationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetLogging(val interface{}) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetNetworkContextProfileIds(val *[]*string) {
	if err := j.validateSetNetworkContextProfileIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkContextProfileIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetSourceGroupsExcluded(val interface{}) {
	if err := j.validateSetSourceGroupsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceGroupsExcluded",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetSourceIds(val *[]*string) {
	if err := j.validateSetSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_NsxtDistributedFirewallRuleA)SetVdcGroupId(val *string) {
	if err := j.validateSetVdcGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdcGroupId",
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
func NsxtDistributedFirewallRuleA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtDistributedFirewallRuleA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtDistributedFirewallRuleA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtDistributedFirewallRuleA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtDistributedFirewallRuleA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtDistributedFirewallRuleA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtDistributedFirewallRuleA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtDistributedFirewallRule.NsxtDistributedFirewallRuleA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetAboveRuleId() {
	_jsii_.InvokeVoid(
		n,
		"resetAboveRuleId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetAppPortProfileIds() {
	_jsii_.InvokeVoid(
		n,
		"resetAppPortProfileIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetComment() {
	_jsii_.InvokeVoid(
		n,
		"resetComment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetDestinationGroupsExcluded() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationGroupsExcluded",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetDestinationIds() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetDirection() {
	_jsii_.InvokeVoid(
		n,
		"resetDirection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetLogging() {
	_jsii_.InvokeVoid(
		n,
		"resetLogging",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetNetworkContextProfileIds() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkContextProfileIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetSourceGroupsExcluded() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceGroupsExcluded",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ResetSourceIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtDistributedFirewallRuleA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

