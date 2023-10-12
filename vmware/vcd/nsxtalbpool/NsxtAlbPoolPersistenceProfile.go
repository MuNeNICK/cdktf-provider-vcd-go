package nsxtalbpool


type NsxtAlbPoolPersistenceProfile struct {
	// Type of persistence strategy. One of `CLIENT_IP`, `HTTP_COOKIE`, `CUSTOM_HTTP_HEADER`, `APP_COOKIE`, `TLS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#type NsxtAlbPool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of attribute based on persistence type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#value NsxtAlbPool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

