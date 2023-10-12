package org


type OrgVappLease struct {
	// If true, storage for a vApp is deleted when the vApp's lease expires.
	//
	// If false, the storage is flagged for deletion, but not deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#delete_on_storage_lease_expiration Org#delete_on_storage_lease_expiration}
	DeleteOnStorageLeaseExpiration interface{} `field:"required" json:"deleteOnStorageLeaseExpiration" yaml:"deleteOnStorageLeaseExpiration"`
	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#maximum_runtime_lease_in_sec Org#maximum_runtime_lease_in_sec}
	MaximumRuntimeLeaseInSec *float64 `field:"required" json:"maximumRuntimeLeaseInSec" yaml:"maximumRuntimeLeaseInSec"`
	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#maximum_storage_lease_in_sec Org#maximum_storage_lease_in_sec}
	MaximumStorageLeaseInSec *float64 `field:"required" json:"maximumStorageLeaseInSec" yaml:"maximumStorageLeaseInSec"`
	// When true, vApps are powered off when the runtime lease expires.
	//
	// When false, vApps are suspended when the runtime lease expires
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#power_off_on_runtime_lease_expiration Org#power_off_on_runtime_lease_expiration}
	PowerOffOnRuntimeLeaseExpiration interface{} `field:"required" json:"powerOffOnRuntimeLeaseExpiration" yaml:"powerOffOnRuntimeLeaseExpiration"`
}

