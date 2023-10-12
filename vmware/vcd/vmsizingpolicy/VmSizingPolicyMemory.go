package vmsizingpolicy


type VmSizingPolicyMemory struct {
	// Defines the memory limit in MB for a VM.
	//
	// If not defined in the VM sizing policy, memory limit is equal to the allocated memory for the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#limit_in_mb VmSizingPolicy#limit_in_mb}
	LimitInMb *string `field:"optional" json:"limitInMb" yaml:"limitInMb"`
	// Defines the reserved amount of memory that is configured for a VM.
	//
	// The value of the attribute ranges between 0 and one. Value of 0 memory reservation guarantee defines no memory reservation. Value of 1 defines 100% of memory reserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#reservation_guarantee VmSizingPolicy#reservation_guarantee}
	ReservationGuarantee *string `field:"optional" json:"reservationGuarantee" yaml:"reservationGuarantee"`
	// Defines the number of memory shares for a VM.
	//
	// Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of memory as another VM, it is entitled to consume twice as much memory when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#shares VmSizingPolicy#shares}
	Shares *string `field:"optional" json:"shares" yaml:"shares"`
	// Defines the memory configured for a VM in MB.
	//
	// This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, the VM receives the amount of memory defined by this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#size_in_mb VmSizingPolicy#size_in_mb}
	SizeInMb *string `field:"optional" json:"sizeInMb" yaml:"sizeInMb"`
}

