package vm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmConfig struct {
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
	// A name for the VM, unique within the vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#name Vm#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Automatically accept EULA if OVA has it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#accept_all_eulas Vm#accept_all_eulas}
	AcceptAllEulas interface{} `field:"optional" json:"acceptAllEulas" yaml:"acceptAllEulas"`
	// Media name to add as boot image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#boot_image Vm#boot_image}
	BootImage *string `field:"optional" json:"bootImage" yaml:"bootImage"`
	// The URN of the media to use as boot image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#boot_image_id Vm#boot_image_id}
	BootImageId *string `field:"optional" json:"bootImageId" yaml:"bootImageId"`
	// The catalog name in which to find the given vApp Template or media for boot_image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#catalog_name Vm#catalog_name}
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Computer name to assign to this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#computer_name Vm#computer_name}
	ComputerName *string `field:"optional" json:"computerName" yaml:"computerName"`
	// The number of cores per socket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_cores Vm#cpu_cores}
	CpuCores *float64 `field:"optional" json:"cpuCores" yaml:"cpuCores"`
	// True if the virtual machine supports addition of virtual CPUs while powered on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_hot_add_enabled Vm#cpu_hot_add_enabled}
	CpuHotAddEnabled interface{} `field:"optional" json:"cpuHotAddEnabled" yaml:"cpuHotAddEnabled"`
	// The limit for how much of CPU can be consumed on the underlying virtualization infrastructure.
	//
	// This is only valid when the resource allocation is not unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_limit Vm#cpu_limit}
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_priority Vm#cpu_priority}
	CpuPriority *string `field:"optional" json:"cpuPriority" yaml:"cpuPriority"`
	// The amount of MHz reservation on the underlying virtualization infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_reservation Vm#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The number of virtual CPUs to allocate to the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpus Vm#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// Custom priority for the resource. This is a read-only, unless the `cpu_priority` is CUSTOM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#cpu_shares Vm#cpu_shares}
	CpuShares *float64 `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// customization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#customization Vm#customization}
	Customization *VmCustomization `field:"optional" json:"customization" yaml:"customization"`
	// The VM description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#description Vm#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#disk Vm#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Expose hardware-assisted CPU virtualization to guest OS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#expose_hardware_virtualization Vm#expose_hardware_virtualization}
	ExposeHardwareVirtualization interface{} `field:"optional" json:"exposeHardwareVirtualization" yaml:"exposeHardwareVirtualization"`
	// Key/value settings for guest properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#guest_properties Vm#guest_properties}
	GuestProperties *map[string]*string `field:"optional" json:"guestProperties" yaml:"guestProperties"`
	// Virtual Hardware Version (e.g.`vmx-14`, `vmx-13`, `vmx-12`, etc.).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#hardware_version Vm#hardware_version}
	HardwareVersion *string `field:"optional" json:"hardwareVersion" yaml:"hardwareVersion"`
	// VM Hyper Reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#href Vm#href}
	Href *string `field:"optional" json:"href" yaml:"href"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#id Vm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The amount of RAM (in MB) to allocate to the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory Vm#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// True if the virtual machine supports addition of memory while powered on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory_hot_add_enabled Vm#memory_hot_add_enabled}
	MemoryHotAddEnabled interface{} `field:"optional" json:"memoryHotAddEnabled" yaml:"memoryHotAddEnabled"`
	// The limit for how much of memory can be consumed on the underlying virtualization infrastructure.
	//
	// This is only valid when the resource allocation is not unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory_limit Vm#memory_limit}
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory_priority Vm#memory_priority}
	MemoryPriority *string `field:"optional" json:"memoryPriority" yaml:"memoryPriority"`
	// The amount of RAM (in MB) reservation on the underlying virtualization infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory_reservation Vm#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// Custom priority for the resource. This is a read-only, unless the `memory_priority` is CUSTOM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#memory_shares Vm#memory_shares}
	MemoryShares *float64 `field:"optional" json:"memoryShares" yaml:"memoryShares"`
	// Key value map of metadata to assign to this VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#metadata Vm#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#metadata_entry Vm#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#network Vm#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
	// Optional number of seconds to try and wait for DHCP IP (valid for 'network' block only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#network_dhcp_wait_seconds Vm#network_dhcp_wait_seconds}
	NetworkDhcpWaitSeconds *float64 `field:"optional" json:"networkDhcpWaitSeconds" yaml:"networkDhcpWaitSeconds"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#org Vm#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Operating System type. Possible values can be found in documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#os_type Vm#os_type}
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// override_template_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#override_template_disk Vm#override_template_disk}
	OverrideTemplateDisk interface{} `field:"optional" json:"overrideTemplateDisk" yaml:"overrideTemplateDisk"`
	// VM placement policy ID. Has to be assigned to Org VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#placement_policy_id Vm#placement_policy_id}
	PlacementPolicyId *string `field:"optional" json:"placementPolicyId" yaml:"placementPolicyId"`
	// A boolean value stating if this VM should be powered on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#power_on Vm#power_on}
	PowerOn interface{} `field:"optional" json:"powerOn" yaml:"powerOn"`
	// True if the update of resource should fail when virtual machine power off needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#prevent_update_power_off Vm#prevent_update_power_off}
	PreventUpdatePowerOff interface{} `field:"optional" json:"preventUpdatePowerOff" yaml:"preventUpdatePowerOff"`
	// Security tags to assign to this VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#security_tags Vm#security_tags}
	SecurityTags *[]*string `field:"optional" json:"securityTags" yaml:"securityTags"`
	// VM sizing policy ID. Has to be assigned to Org VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#sizing_policy_id Vm#sizing_policy_id}
	SizingPolicyId *string `field:"optional" json:"sizingPolicyId" yaml:"sizingPolicyId"`
	// Storage profile to override the default one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#storage_profile Vm#storage_profile}
	StorageProfile *string `field:"optional" json:"storageProfile" yaml:"storageProfile"`
	// The name of the vApp Template to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#template_name Vm#template_name}
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The vApp this VM belongs to - Required, unless it is a standalone VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#vapp_name Vm#vapp_name}
	VappName *string `field:"optional" json:"vappName" yaml:"vappName"`
	// The URN of the vApp Template to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#vapp_template_id Vm#vapp_template_id}
	VappTemplateId *string `field:"optional" json:"vappTemplateId" yaml:"vappTemplateId"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#vdc Vm#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
	// The name of the VM in vApp Template to use.
	//
	// In cases when vApp template has more than one VM
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#vm_name_in_template Vm#vm_name_in_template}
	VmNameInTemplate *string `field:"optional" json:"vmNameInTemplate" yaml:"vmNameInTemplate"`
}

