package edgegateway


type EdgegatewayExternalNetwork struct {
	// External network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#name Edgegateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Enable rate limiting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#enable_rate_limit Edgegateway#enable_rate_limit}
	EnableRateLimit interface{} `field:"optional" json:"enableRateLimit" yaml:"enableRateLimit"`
	// Incoming rate limit (Mbps).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#incoming_rate_limit Edgegateway#incoming_rate_limit}
	IncomingRateLimit *float64 `field:"optional" json:"incomingRateLimit" yaml:"incomingRateLimit"`
	// Outgoing rate limit (Mbps).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#outgoing_rate_limit Edgegateway#outgoing_rate_limit}
	OutgoingRateLimit *float64 `field:"optional" json:"outgoingRateLimit" yaml:"outgoingRateLimit"`
	// subnet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/edgegateway#subnet Edgegateway#subnet}
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
}

