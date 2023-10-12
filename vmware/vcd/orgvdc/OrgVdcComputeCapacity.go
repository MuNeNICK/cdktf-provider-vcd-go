package orgvdc


type OrgVdcComputeCapacity struct {
	// cpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#cpu OrgVdc#cpu}
	Cpu *OrgVdcComputeCapacityCpu `field:"required" json:"cpu" yaml:"cpu"`
	// memory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#memory OrgVdc#memory}
	Memory *OrgVdcComputeCapacityMemory `field:"required" json:"memory" yaml:"memory"`
}

