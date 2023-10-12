package datavcdcatalogvapptemplate


type DataVcdCatalogVappTemplateFilter struct {
	// Search by date comparison ({>|>=|<|<=|==} yyyy-mm-dd[ hh[:mm[:ss]]]).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#date DataVcdCatalogVappTemplate#date}
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Retrieves the oldest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#earliest DataVcdCatalogVappTemplate#earliest}
	Earliest interface{} `field:"optional" json:"earliest" yaml:"earliest"`
	// Retrieves the newest item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#latest DataVcdCatalogVappTemplate#latest}
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#metadata DataVcdCatalogVappTemplate#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Search by name with a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#name_regex DataVcdCatalogVappTemplate#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
}

