package datavcdcatalog


type DataVcdCatalogFilter struct {
	// Search by date comparison ({>|>=|<|<=|==} yyyy-mm-dd[ hh[:mm[:ss]]]).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog#date DataVcdCatalog#date}
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Retrieves the oldest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog#earliest DataVcdCatalog#earliest}
	Earliest interface{} `field:"optional" json:"earliest" yaml:"earliest"`
	// Retrieves the newest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog#latest DataVcdCatalog#latest}
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog#metadata DataVcdCatalog#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog#name_regex DataVcdCatalog#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

