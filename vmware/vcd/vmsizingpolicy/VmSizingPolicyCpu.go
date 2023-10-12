package vmsizingpolicy


type VmSizingPolicyCpu struct {
	// The number of cores per socket for a VM.
	//
	// This is a VM hardware configuration. The number of vCPUs that is defined in the VM sizing policy must be divisible by the number of cores per socket. If the number of vCPUs is not divisible by the number of cores per socket, the number of cores per socket becomes invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#cores_per_socket VmSizingPolicy#cores_per_socket}
	CoresPerSocket *string `field:"optional" json:"coresPerSocket" yaml:"coresPerSocket"`
	// Defines the number of vCPUs configured for a VM.
	//
	// This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, this count becomes the configured number of vCPUs for the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#count VmSizingPolicy#count}
	Count *string `field:"optional" json:"count" yaml:"count"`
	// Defines the CPU limit in MHz for a VM.
	//
	// If not defined in the VDC compute policy, CPU limit is equal to the vCPU speed multiplied by the number of vCPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#limit_in_mhz VmSizingPolicy#limit_in_mhz}
	LimitInMhz *string `field:"optional" json:"limitInMhz" yaml:"limitInMhz"`
	// Defines how much of the CPU resources of a VM are reserved.
	//
	// The allocated CPU for a VM equals the number of vCPUs times the vCPU speed in MHz. The value of the attribute ranges between 0 and one. Value of 0 CPU reservation guarantee defines no CPU reservation. Value of 1 defines 100% of CPU reserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#reservation_guarantee VmSizingPolicy#reservation_guarantee}
	ReservationGuarantee *string `field:"optional" json:"reservationGuarantee" yaml:"reservationGuarantee"`
	// Defines the number of CPU shares for a VM.
	//
	// Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of CPU as another VM, it is entitled to consume twice as much CPU when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#shares VmSizingPolicy#shares}
	Shares *string `field:"optional" json:"shares" yaml:"shares"`
	// Defines the vCPU speed of a core in MHz.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_sizing_policy#speed_in_mhz VmSizingPolicy#speed_in_mhz}
	SpeedInMhz *string `field:"optional" json:"speedInMhz" yaml:"speedInMhz"`
}

