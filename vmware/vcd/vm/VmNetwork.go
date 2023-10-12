package vm


type VmNetwork struct {
	// Network type to use: 'vapp', 'org' or 'none'.
	//
	// Use 'vapp' for vApp network, 'org' to attach Org VDC network. 'none' for empty NIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#type Vm#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Network card adapter type. (e.g. 'E1000', 'E1000E', 'SRIOVETHERNETCARD', 'VMXNET3', 'PCNet32').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#adapter_type Vm#adapter_type}
	AdapterType *string `field:"optional" json:"adapterType" yaml:"adapterType"`
	// It defines if NIC is connected or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#connected Vm#connected}
	Connected interface{} `field:"optional" json:"connected" yaml:"connected"`
	// IP of the VM. Settings depend on `ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#ip Vm#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// IP address allocation mode. One of POOL, DHCP, MANUAL, NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#ip_allocation_mode Vm#ip_allocation_mode}
	IpAllocationMode *string `field:"optional" json:"ipAllocationMode" yaml:"ipAllocationMode"`
	// Set to true if network interface should be primary.
	//
	// First network card in the list will be primary by default
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#is_primary Vm#is_primary}
	IsPrimary interface{} `field:"optional" json:"isPrimary" yaml:"isPrimary"`
	// Mac address of network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#mac Vm#mac}
	Mac *string `field:"optional" json:"mac" yaml:"mac"`
	// Name of the network this VM should connect to. Always required except for `type` `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vm#name Vm#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

