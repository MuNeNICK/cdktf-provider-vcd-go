package datavcdcatalogvapptemplate


type DataVcdCatalogVappTemplateFilterMetadata struct {
	// Metadata key (field name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#key DataVcdCatalogVappTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Metadata value (can be a regular expression if "use_api_search" is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#value DataVcdCatalogVappTemplate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// True if is a metadata@SYSTEM key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#is_system DataVcdCatalogVappTemplate#is_system}
	IsSystem interface{} `field:"optional" json:"isSystem" yaml:"isSystem"`
	// Type of metadata value (needed only if "use_api_search" is true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#type DataVcdCatalogVappTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// If true, will search the vCD using native metadata query (without regular expressions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/catalog_vapp_template#use_api_search DataVcdCatalogVappTemplate#use_api_search}
	UseApiSearch interface{} `field:"optional" json:"useApiSearch" yaml:"useApiSearch"`
}

