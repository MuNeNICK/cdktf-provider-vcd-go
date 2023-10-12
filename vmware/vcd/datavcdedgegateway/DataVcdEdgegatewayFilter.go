package datavcdedgegateway


type DataVcdEdgegatewayFilter struct {
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/edgegateway#name_regex DataVcdEdgegateway#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

