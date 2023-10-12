package vapp


type VappLease struct {
	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended.
	//
	// 0 means never expires
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#runtime_lease_in_sec Vapp#runtime_lease_in_sec}
	RuntimeLeaseInSec *float64 `field:"required" json:"runtimeLeaseInSec" yaml:"runtimeLeaseInSec"`
	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp#storage_lease_in_sec Vapp#storage_lease_in_sec}
	StorageLeaseInSec *float64 `field:"required" json:"storageLeaseInSec" yaml:"storageLeaseInSec"`
}

