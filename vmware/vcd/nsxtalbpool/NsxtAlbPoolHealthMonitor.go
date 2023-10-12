package nsxtalbpool


type NsxtAlbPoolHealthMonitor struct {
	// Type of health monitor. One of `HTTP`, `HTTPS`, `TCP`, `UDP`, `PING`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#type NsxtAlbPool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

