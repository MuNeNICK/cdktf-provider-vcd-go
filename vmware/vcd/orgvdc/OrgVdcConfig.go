package orgvdc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgVdcConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#allocation_model OrgVdc#allocation_model}
	AllocationModel *string `field:"required" json:"allocationModel" yaml:"allocationModel"`
	// compute_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#compute_capacity OrgVdc#compute_capacity}
	ComputeCapacity *OrgVdcComputeCapacity `field:"required" json:"computeCapacity" yaml:"computeCapacity"`
	// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#delete_force OrgVdc#delete_force}
	DeleteForce interface{} `field:"required" json:"deleteForce" yaml:"deleteForce"`
	// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that normally allows removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#delete_recursive OrgVdc#delete_recursive}
	DeleteRecursive interface{} `field:"required" json:"deleteRecursive" yaml:"deleteRecursive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#name OrgVdc#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to the Provider VDC from which this organization VDC is provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#provider_vdc_name OrgVdc#provider_vdc_name}
	ProviderVdcName *string `field:"required" json:"providerVdcName" yaml:"providerVdcName"`
	// storage_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#storage_profile OrgVdc#storage_profile}
	StorageProfile interface{} `field:"required" json:"storageProfile" yaml:"storageProfile"`
	// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply.
	//
	// Default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#allow_over_commit OrgVdc#allow_over_commit}
	AllowOverCommit interface{} `field:"optional" json:"allowOverCommit" yaml:"allowOverCommit"`
	// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC.
	//
	// For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the element is empty, vCD sets a value
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#cpu_guaranteed OrgVdc#cpu_guaranteed}
	CpuGuaranteed *float64 `field:"optional" json:"cpuGuaranteed" yaml:"cpuGuaranteed"`
	// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM.
	//
	// A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#cpu_speed OrgVdc#cpu_speed}
	CpuSpeed *float64 `field:"optional" json:"cpuSpeed" yaml:"cpuSpeed"`
	// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#default_compute_policy_id OrgVdc#default_compute_policy_id}
	DefaultComputePolicyId *string `field:"optional" json:"defaultComputePolicyId" yaml:"defaultComputePolicyId"`
	// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#default_vm_sizing_policy_id OrgVdc#default_vm_sizing_policy_id}
	DefaultVmSizingPolicyId *string `field:"optional" json:"defaultVmSizingPolicyId" yaml:"defaultVmSizingPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#description OrgVdc#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#edge_cluster_id OrgVdc#edge_cluster_id}
	EdgeClusterId *string `field:"optional" json:"edgeClusterId" yaml:"edgeClusterId"`
	// Set to true to indicate if the Flex VDC is to be elastic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#elasticity OrgVdc#elasticity}
	Elasticity interface{} `field:"optional" json:"elasticity" yaml:"elasticity"`
	// True if this VDC is enabled for use by the organization VDCs. Default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enabled OrgVdc#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Request for fast provisioning.
	//
	// Request will be honored only if the underlying datas tore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enable_fast_provisioning OrgVdc#enable_fast_provisioning}
	EnableFastProvisioning interface{} `field:"optional" json:"enableFastProvisioning" yaml:"enableFastProvisioning"`
	// Set to true to enable distributed firewall - Only applies to NSX-V VDCs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enable_nsxv_distributed_firewall OrgVdc#enable_nsxv_distributed_firewall}
	EnableNsxvDistributedFirewall interface{} `field:"optional" json:"enableNsxvDistributedFirewall" yaml:"enableNsxvDistributedFirewall"`
	// Boolean to request thin provisioning.
	//
	// Request will be honored only if the underlying datastore supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enable_thin_provisioning OrgVdc#enable_thin_provisioning}
	EnableThinProvisioning interface{} `field:"optional" json:"enableThinProvisioning" yaml:"enableThinProvisioning"`
	// True if discovery of vCenter VMs is enabled for resource pools backing this VDC.
	//
	// If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enable_vm_discovery OrgVdc#enable_vm_discovery}
	EnableVmDiscovery interface{} `field:"optional" json:"enableVmDiscovery" yaml:"enableVmDiscovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#id OrgVdc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#include_vm_memory_overhead OrgVdc#include_vm_memory_overhead}
	IncludeVmMemoryOverhead interface{} `field:"optional" json:"includeVmMemoryOverhead" yaml:"includeVmMemoryOverhead"`
	// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC.
	//
	// For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#memory_guaranteed OrgVdc#memory_guaranteed}
	MemoryGuaranteed *float64 `field:"optional" json:"memoryGuaranteed" yaml:"memoryGuaranteed"`
	// Key and value pairs for Org VDC metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#metadata OrgVdc#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#metadata_entry OrgVdc#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// The name of a network pool in the Provider VDC.
	//
	// Required if this VDC will contain routed or isolated networks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#network_pool_name OrgVdc#network_pool_name}
	NetworkPoolName *string `field:"optional" json:"networkPoolName" yaml:"networkPoolName"`
	// Maximum number of network objects that can be deployed in this VDC.
	//
	// Defaults to 0, which means no networks can be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#network_quota OrgVdc#network_quota}
	NetworkQuota *float64 `field:"optional" json:"networkQuota" yaml:"networkQuota"`
	// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#nic_quota OrgVdc#nic_quota}
	NicQuota *float64 `field:"optional" json:"nicQuota" yaml:"nicQuota"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#org OrgVdc#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Set of VM Placement Policy IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#vm_placement_policy_ids OrgVdc#vm_placement_policy_ids}
	VmPlacementPolicyIds *[]*string `field:"optional" json:"vmPlacementPolicyIds" yaml:"vmPlacementPolicyIds"`
	// The maximum number of VMs that can be created in this VDC.
	//
	// Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#vm_quota OrgVdc#vm_quota}
	VmQuota *float64 `field:"optional" json:"vmQuota" yaml:"vmQuota"`
	// Set of VM Sizing Policy IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#vm_sizing_policy_ids OrgVdc#vm_sizing_policy_ids}
	VmSizingPolicyIds *[]*string `field:"optional" json:"vmSizingPolicyIds" yaml:"vmSizingPolicyIds"`
}

