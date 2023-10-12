package catalogaccesscontrol


type CatalogAccessControlSharedWith struct {
	// The access level for the org, user, or group to which we are sharing.
	//
	// One of [ReadOnly, Change, FullControl] for users and groups, but just ReadOnly for Organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#access_level CatalogAccessControl#access_level}
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// ID of the group to which we are sharing. Required if user_id or org_id is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#group_id CatalogAccessControl#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// ID of the Org to which we are sharing. Required if user_id or group_id is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#org_id CatalogAccessControl#org_id}
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// ID of the user to which we are sharing. Required if group_id or org_id is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/catalog_access_control#user_id CatalogAccessControl#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

