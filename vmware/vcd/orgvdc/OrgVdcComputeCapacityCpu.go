package orgvdc


type OrgVdcComputeCapacityCpu struct {
	// Capacity that is committed to be available.
	//
	// Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#allocated OrgVdc#allocated}
	Allocated *float64 `field:"optional" json:"allocated" yaml:"allocated"`
	// Capacity limit relative to the value specified for Allocation.
	//
	// It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#limit OrgVdc#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

