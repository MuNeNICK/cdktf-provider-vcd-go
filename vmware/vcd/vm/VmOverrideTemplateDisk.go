package vm


type VmOverrideTemplateDisk struct {
	// The number of the SCSI or IDE controller itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#bus_number Vm#bus_number}
	BusNumber *float64 `field:"required" json:"busNumber" yaml:"busNumber"`
	// The type of disk controller.
	//
	// Possible values: ide, parallel( LSI Logic Parallel SCSI), sas(LSI Logic SAS (SCSI)), paravirtual(Paravirtual (SCSI)), sata, nvme
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#bus_type Vm#bus_type}
	BusType *string `field:"required" json:"busType" yaml:"busType"`
	// The size of the disk in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#size_in_mb Vm#size_in_mb}
	SizeInMb *float64 `field:"required" json:"sizeInMb" yaml:"sizeInMb"`
	// The device number on the SCSI or IDE controller of the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#unit_number Vm#unit_number}
	UnitNumber *float64 `field:"required" json:"unitNumber" yaml:"unitNumber"`
	// Specifies the IOPS for the disk. Default is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#iops Vm#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Storage profile to override the VM default one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#storage_profile Vm#storage_profile}
	StorageProfile *string `field:"optional" json:"storageProfile" yaml:"storageProfile"`
}

