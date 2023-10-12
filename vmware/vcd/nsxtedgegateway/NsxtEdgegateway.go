package nsxtedgegateway

import (
	_init_ "github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/jsii"

	"github.com/MuNeNICK/cdktf-provider-vcd-go/vmware/vcd/nsxtedgegateway/internal"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway vcd_nsxt_edgegateway}.
type NsxtEdgegateway interface {
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
	DedicateExternalNetwork() interface{}
	SetDedicateExternalNetwork(val interface{})
	DedicateExternalNetworkInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeClusterId() *string
	SetEdgeClusterId(val *string)
	EdgeClusterIdInput() *string
	ExternalNetworkId() *string
	SetExternalNetworkId(val *string)
	ExternalNetworkIdInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OwnerId() *string
	SetOwnerId(val *string)
	OwnerIdInput() *string
	PrimaryIp() *string
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
	StartingVdcId() *string
	SetStartingVdcId(val *string)
	StartingVdcIdInput() *string
	Subnet() NsxtEdgegatewaySubnetList
	SubnetInput() interface{}
	SubnetWithIpCount() NsxtEdgegatewaySubnetWithIpCountList
	SubnetWithIpCountInput() interface{}
	SubnetWithTotalIpCount() NsxtEdgegatewaySubnetWithTotalIpCountList
	SubnetWithTotalIpCountInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TotalAllocatedIpCount() *float64
	SetTotalAllocatedIpCount(val *float64)
	TotalAllocatedIpCountInput() *float64
	UnusedIpCount() *float64
	UsedIpCount() *float64
	UseIpSpaces() cdktf.IResolvable
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
	PutSubnet(value interface{})
	PutSubnetWithIpCount(value interface{})
	PutSubnetWithTotalIpCount(value interface{})
	ResetDedicateExternalNetwork()
	ResetDescription()
	ResetEdgeClusterId()
	ResetId()
	ResetOrg()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerId()
	ResetStartingVdcId()
	ResetSubnet()
	ResetSubnetWithIpCount()
	ResetSubnetWithTotalIpCount()
	ResetTotalAllocatedIpCount()
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

// The jsii proxy struct for NsxtEdgegateway
type jsiiProxy_NsxtEdgegateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NsxtEdgegateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) DedicateExternalNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicateExternalNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) DedicateExternalNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicateExternalNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) EdgeClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) EdgeClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) ExternalNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) ExternalNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) OwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) PrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) StartingVdcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingVdcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) StartingVdcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingVdcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Subnet() NsxtEdgegatewaySubnetList {
	var returns NsxtEdgegatewaySubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) SubnetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) SubnetWithIpCount() NsxtEdgegatewaySubnetWithIpCountList {
	var returns NsxtEdgegatewaySubnetWithIpCountList
	_jsii_.Get(
		j,
		"subnetWithIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) SubnetWithIpCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetWithIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) SubnetWithTotalIpCount() NsxtEdgegatewaySubnetWithTotalIpCountList {
	var returns NsxtEdgegatewaySubnetWithTotalIpCountList
	_jsii_.Get(
		j,
		"subnetWithTotalIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) SubnetWithTotalIpCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetWithTotalIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) TotalAllocatedIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAllocatedIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) TotalAllocatedIpCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAllocatedIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) UnusedIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unusedIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) UsedIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) UseIpSpaces() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useIpSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) Vdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NsxtEdgegateway) VdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway vcd_nsxt_edgegateway} Resource.
func NewNsxtEdgegateway(scope constructs.Construct, id *string, config *NsxtEdgegatewayConfig) NsxtEdgegateway {
	_init_.Initialize()

	if err := validateNewNsxtEdgegatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NsxtEdgegateway{}

	_jsii_.Create(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway vcd_nsxt_edgegateway} Resource.
func NewNsxtEdgegateway_Override(n NsxtEdgegateway, scope constructs.Construct, id *string, config *NsxtEdgegatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetDedicateExternalNetwork(val interface{}) {
	if err := j.validateSetDedicateExternalNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicateExternalNetwork",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetEdgeClusterId(val *string) {
	if err := j.validateSetEdgeClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeClusterId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetExternalNetworkId(val *string) {
	if err := j.validateSetExternalNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalNetworkId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetOwnerId(val *string) {
	if err := j.validateSetOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetStartingVdcId(val *string) {
	if err := j.validateSetStartingVdcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingVdcId",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetTotalAllocatedIpCount(val *float64) {
	if err := j.validateSetTotalAllocatedIpCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalAllocatedIpCount",
		val,
	)
}

func (j *jsiiProxy_NsxtEdgegateway)SetVdc(val *string) {
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
func NsxtEdgegateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NsxtEdgegateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNsxtEdgegateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NsxtEdgegateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"vcd.nsxtEdgegateway.NsxtEdgegateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NsxtEdgegateway) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NsxtEdgegateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NsxtEdgegateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NsxtEdgegateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NsxtEdgegateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NsxtEdgegateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NsxtEdgegateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NsxtEdgegateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NsxtEdgegateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NsxtEdgegateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NsxtEdgegateway) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NsxtEdgegateway) PutSubnet(value interface{}) {
	if err := n.validatePutSubnetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSubnet",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtEdgegateway) PutSubnetWithIpCount(value interface{}) {
	if err := n.validatePutSubnetWithIpCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSubnetWithIpCount",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtEdgegateway) PutSubnetWithTotalIpCount(value interface{}) {
	if err := n.validatePutSubnetWithTotalIpCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSubnetWithTotalIpCount",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetDedicateExternalNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetDedicateExternalNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetEdgeClusterId() {
	_jsii_.InvokeVoid(
		n,
		"resetEdgeClusterId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetOrg() {
	_jsii_.InvokeVoid(
		n,
		"resetOrg",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetOwnerId() {
	_jsii_.InvokeVoid(
		n,
		"resetOwnerId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetStartingVdcId() {
	_jsii_.InvokeVoid(
		n,
		"resetStartingVdcId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetSubnet() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetSubnetWithIpCount() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnetWithIpCount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetSubnetWithTotalIpCount() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnetWithTotalIpCount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetTotalAllocatedIpCount() {
	_jsii_.InvokeVoid(
		n,
		"resetTotalAllocatedIpCount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) ResetVdc() {
	_jsii_.InvokeVoid(
		n,
		"resetVdc",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NsxtEdgegateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NsxtEdgegateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

