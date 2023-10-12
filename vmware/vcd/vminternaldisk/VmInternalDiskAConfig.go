package vminternaldisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmInternalDiskAConfig struct {
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
	// The number of the SCSI or IDE controller itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#bus_number VmInternalDiskA#bus_number}
	BusNumber *float64 `field:"required" json:"busNumber" yaml:"busNumber"`
	// The type of disk controller.
	//
	// Possible values: ide, parallel( LSI Logic Parallel SCSI), sas(LSI Logic SAS (SCSI)), paravirtual(Paravirtual (SCSI)), sata, nvme
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#bus_type VmInternalDiskA#bus_type}
	BusType *string `field:"required" json:"busType" yaml:"busType"`
	// The size of the disk in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#size_in_mb VmInternalDiskA#size_in_mb}
	SizeInMb *float64 `field:"required" json:"sizeInMb" yaml:"sizeInMb"`
	// The device number on the SCSI or IDE controller of the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#unit_number VmInternalDiskA#unit_number}
	UnitNumber *float64 `field:"required" json:"unitNumber" yaml:"unitNumber"`
	// The vApp this VM internal disk belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#vapp_name VmInternalDiskA#vapp_name}
	VappName *string `field:"required" json:"vappName" yaml:"vappName"`
	// VM in vApp in which internal disk is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#vm_name VmInternalDiskA#vm_name}
	VmName *string `field:"required" json:"vmName" yaml:"vmName"`
	// Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on.
	//
	// Without this setting enabled, such changes on a powered-on VM would fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#allow_vm_reboot VmInternalDiskA#allow_vm_reboot}
	AllowVmReboot interface{} `field:"optional" json:"allowVmReboot" yaml:"allowVmReboot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#id VmInternalDiskA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the IOPS for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#iops VmInternalDiskA#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#org VmInternalDiskA#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Storage profile to override the VM default one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#storage_profile VmInternalDiskA#storage_profile}
	StorageProfile *string `field:"optional" json:"storageProfile" yaml:"storageProfile"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm_internal_disk#vdc VmInternalDiskA#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

