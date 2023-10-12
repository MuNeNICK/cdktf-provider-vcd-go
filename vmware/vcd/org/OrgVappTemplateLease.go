package org


type OrgVappTemplateLease struct {
	// If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires.
	//
	// If false, the storage is flagged for deletion, but not deleted
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#delete_on_storage_lease_expiration Org#delete_on_storage_lease_expiration}
	DeleteOnStorageLeaseExpiration interface{} `field:"required" json:"deleteOnStorageLeaseExpiration" yaml:"deleteOnStorageLeaseExpiration"`
	// How long vApp templates are available before being automatically cleaned up (in seconds). 0 means never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#maximum_storage_lease_in_sec Org#maximum_storage_lease_in_sec}
	MaximumStorageLeaseInSec *float64 `field:"required" json:"maximumStorageLeaseInSec" yaml:"maximumStorageLeaseInSec"`
}

