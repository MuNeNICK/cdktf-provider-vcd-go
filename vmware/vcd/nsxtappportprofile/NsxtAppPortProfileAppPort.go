package nsxtappportprofile


type NsxtAppPortProfileAppPort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#protocol NsxtAppPortProfile#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Set of ports or ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_app_port_profile#port NsxtAppPortProfile#port}
	Port *[]*string `field:"optional" json:"port" yaml:"port"`
}

