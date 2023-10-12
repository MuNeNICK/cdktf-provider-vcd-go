package orguser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgUserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// User's name. Only lowercase letters allowed. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#name OrgUser#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#role OrgUser#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#deployed_vm_quota OrgUser#deployed_vm_quota}
	DeployedVmQuota *float64 `field:"optional" json:"deployedVmQuota" yaml:"deployedVmQuota"`
	// The user's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#description OrgUser#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The user's email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#email_address OrgUser#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// True if the user is enabled and can log in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#enabled OrgUser#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The user's full name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#full_name OrgUser#full_name}
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#id OrgUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The user's telephone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#instant_messaging OrgUser#instant_messaging}
	InstantMessaging *string `field:"optional" json:"instantMessaging" yaml:"instantMessaging"`
	// True if this user is imported from an external resource, like an LDAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#is_external OrgUser#is_external}
	IsExternal interface{} `field:"optional" json:"isExternal" yaml:"isExternal"`
	// True if this user has a group role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#is_group_role OrgUser#is_group_role}
	IsGroupRole interface{} `field:"optional" json:"isGroupRole" yaml:"isGroupRole"`
	// If the user account has been locked due to too many invalid login attempts, the value will change to true (only the system can lock the user).
	//
	// To unlock the user re-set this flag to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#is_locked OrgUser#is_locked}
	IsLocked interface{} `field:"optional" json:"isLocked" yaml:"isLocked"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#org OrgUser#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The user's password.
	//
	// This value is never returned on read. Either "password" or "password_file" must be included on creation unless is_external is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#password OrgUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Name of a file containing the user's password.
	//
	// Either "password_file" or "password" must be included on creation unless is_external is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#password_file OrgUser#password_file}
	PasswordFile *string `field:"optional" json:"passwordFile" yaml:"passwordFile"`
	// Identity provider type for this this user.
	//
	// One of: 'INTEGRATED', 'SAML', 'OAUTH'. When empty, the default value 'INTEGRATED' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#provider_type OrgUser#provider_type}
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#stored_vm_quota OrgUser#stored_vm_quota}
	StoredVmQuota *float64 `field:"optional" json:"storedVmQuota" yaml:"storedVmQuota"`
	// Take ownership of user's objects on deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#take_ownership OrgUser#take_ownership}
	TakeOwnership interface{} `field:"optional" json:"takeOwnership" yaml:"takeOwnership"`
	// The user's telephone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_user#telephone OrgUser#telephone}
	Telephone *string `field:"optional" json:"telephone" yaml:"telephone"`
}

