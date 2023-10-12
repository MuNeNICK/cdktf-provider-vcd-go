package orgldap


type OrgLdapCustomSettingsUserAttributes struct {
	// LDAP attribute to use for the user's full name. For example, displayName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#display_name OrgLdap#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// LDAP attribute to use for the user's email address. For example, mail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#email OrgLdap#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// LDAP attribute to use for the user's given name. For example, givenName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#given_name OrgLdap#given_name}
	GivenName *string `field:"required" json:"givenName" yaml:"givenName"`
	// LDAP attribute that identifies a user as a member of a group. For example, dn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#group_membership_identifier OrgLdap#group_membership_identifier}
	GroupMembershipIdentifier *string `field:"required" json:"groupMembershipIdentifier" yaml:"groupMembershipIdentifier"`
	// LDAP objectClass of which imported users are members. For example, user or person.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#object_class OrgLdap#object_class}
	ObjectClass *string `field:"required" json:"objectClass" yaml:"objectClass"`
	// LDAP attribute to use for the user's surname. For example, sn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#surname OrgLdap#surname}
	Surname *string `field:"required" json:"surname" yaml:"surname"`
	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#telephone OrgLdap#telephone}
	Telephone *string `field:"required" json:"telephone" yaml:"telephone"`
	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#unique_identifier OrgLdap#unique_identifier}
	UniqueIdentifier *string `field:"required" json:"uniqueIdentifier" yaml:"uniqueIdentifier"`
	// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#username OrgLdap#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// LDAP attribute that returns the identifiers of all the groups of which the user is a member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#group_back_link_identifier OrgLdap#group_back_link_identifier}
	GroupBackLinkIdentifier *string `field:"optional" json:"groupBackLinkIdentifier" yaml:"groupBackLinkIdentifier"`
}

