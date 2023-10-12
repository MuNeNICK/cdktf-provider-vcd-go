package datavcdcatalogmedia


type DataVcdCatalogMediaFilter struct {
	// Search by date comparison ({>|>=|<|<=|==} yyyy-mm-dd[ hh[:mm[:ss]]]).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#date DataVcdCatalogMedia#date}
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Retrieves the oldest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#earliest DataVcdCatalogMedia#earliest}
	Earliest interface{} `field:"optional" json:"earliest" yaml:"earliest"`
	// Retrieves the newest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#latest DataVcdCatalogMedia#latest}
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#metadata DataVcdCatalogMedia#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_media#name_regex DataVcdCatalogMedia#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

