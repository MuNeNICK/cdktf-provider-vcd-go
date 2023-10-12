package orgldap


type OrgLdapCustomSettingsGroupAttributes struct {
	// LDAP attribute that identifies a group as a member of another group. For example, dn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#group_membership_identifier OrgLdap#group_membership_identifier}
	GroupMembershipIdentifier *string `field:"required" json:"groupMembershipIdentifier" yaml:"groupMembershipIdentifier"`
	// LDAP attribute to use when getting the members of a group. For example, member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#membership OrgLdap#membership}
	Membership *string `field:"required" json:"membership" yaml:"membership"`
	// LDAP attribute to use for the group name. For example, cn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#name OrgLdap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// LDAP objectClass of which imported groups are members. For example, group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#object_class OrgLdap#object_class}
	ObjectClass *string `field:"required" json:"objectClass" yaml:"objectClass"`
	// LDAP attribute to use as the unique identifier for a group. For example, objectGuid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#unique_identifier OrgLdap#unique_identifier}
	UniqueIdentifier *string `field:"required" json:"uniqueIdentifier" yaml:"uniqueIdentifier"`
	// LDAP group attribute used to identify a group member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#group_back_link_identifier OrgLdap#group_back_link_identifier}
	GroupBackLinkIdentifier *string `field:"optional" json:"groupBackLinkIdentifier" yaml:"groupBackLinkIdentifier"`
}

