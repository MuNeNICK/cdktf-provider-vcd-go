package networkisolated


type NetworkIsolatedMetadataEntry struct {
	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#is_system NetworkIsolated#is_system}
	IsSystem interface{} `field:"optional" json:"isSystem" yaml:"isSystem"`
	// Key of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#key NetworkIsolated#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#type NetworkIsolated#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#user_access NetworkIsolated#user_access}
	UserAccess *string `field:"optional" json:"userAccess" yaml:"userAccess"`
	// Value of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_isolated#value NetworkIsolated#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

