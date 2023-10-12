package vm


type VmDisk struct {
	// Bus number on which to place the disk controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#bus_number Vm#bus_number}
	BusNumber *string `field:"required" json:"busNumber" yaml:"busNumber"`
	// Independent disk name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#name Vm#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unit number (slot) on the bus specified by BusNumber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#unit_number Vm#unit_number}
	UnitNumber *string `field:"required" json:"unitNumber" yaml:"unitNumber"`
}

