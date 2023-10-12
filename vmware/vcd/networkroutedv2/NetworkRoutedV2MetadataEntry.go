package networkroutedv2


type NetworkRoutedV2MetadataEntry struct {
	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#is_system NetworkRoutedV2#is_system}
	IsSystem interface{} `field:"optional" json:"isSystem" yaml:"isSystem"`
	// Key of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#key NetworkRoutedV2#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#type NetworkRoutedV2#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#user_access NetworkRoutedV2#user_access}
	UserAccess *string `field:"optional" json:"userAccess" yaml:"userAccess"`
	// Value of this metadata entry. Required if the metadata entry is not empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/network_routed_v2#value NetworkRoutedV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

