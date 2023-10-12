package orgldap


type OrgLdapCustomSettings struct {
	// authentication method: one of SIMPLE, MD5DIGEST, NTLM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#authentication_method OrgLdap#authentication_method}
	AuthenticationMethod *string `field:"required" json:"authenticationMethod" yaml:"authenticationMethod"`
	// type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#connector_type OrgLdap#connector_type}
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// group_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#group_attributes OrgLdap#group_attributes}
	GroupAttributes *OrgLdapCustomSettingsGroupAttributes `field:"required" json:"groupAttributes" yaml:"groupAttributes"`
	// Port number for LDAP service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#port OrgLdap#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// host name or IP of the LDAP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#server OrgLdap#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// user_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#user_attributes OrgLdap#user_attributes}
	UserAttributes *OrgLdapCustomSettingsUserAttributes `field:"required" json:"userAttributes" yaml:"userAttributes"`
	// LDAP search base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#base_distinguished_name OrgLdap#base_distinguished_name}
	BaseDistinguishedName *string `field:"optional" json:"baseDistinguishedName" yaml:"baseDistinguishedName"`
	// True if the LDAP service requires an SSL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#is_ssl OrgLdap#is_ssl}
	IsSsl interface{} `field:"optional" json:"isSsl" yaml:"isSsl"`
	// Password for the user identified by UserName.
	//
	// This value is never returned by GET. It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#password OrgLdap#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs (for example: cn="ldap-admin", c="example", dc="com").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_ldap#username OrgLdap#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

