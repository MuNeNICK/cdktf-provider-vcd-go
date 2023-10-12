package catalogmedia


type CatalogMediaMetadataEntry struct {
	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#is_system CatalogMedia#is_system}
	IsSystem interface{} `field:"optional" json:"isSystem" yaml:"isSystem"`
	// Key of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#key CatalogMedia#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#type CatalogMedia#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#user_access CatalogMedia#user_access}
	UserAccess *string `field:"optional" json:"userAccess" yaml:"userAccess"`
	// Value of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_media#value CatalogMedia#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

