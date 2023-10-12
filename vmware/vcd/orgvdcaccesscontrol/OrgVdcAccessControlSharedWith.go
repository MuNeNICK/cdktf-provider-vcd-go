package orgvdcaccesscontrol


type OrgVdcAccessControlSharedWith struct {
	// The access level for the user or group to which we are sharing. (Only ReadOnly is available).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#access_level OrgVdcAccessControl#access_level}
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// ID of the group to which we are sharing. Required if user_id is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#group_id OrgVdcAccessControl#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// ID of the user to which we are sharing. Required if group_id is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc_access_control#user_id OrgVdcAccessControl#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

