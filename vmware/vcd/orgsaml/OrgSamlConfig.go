package orgsaml

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgSamlConfig struct {
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
	// Enable SAML authentication. When this option is set, authentication is deferred to the SAML identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#enabled OrgSaml#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Organization ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#org_id OrgSaml#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Optional email attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#email OrgSaml#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Your service provider entity ID. Once you set this field, it cannot be changed back to empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#entity_id OrgSaml#entity_id}
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// Optional first name attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#first_name OrgSaml#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Optional full name attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#full_name OrgSaml#full_name}
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Optional group attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#group OrgSaml#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#id OrgSaml#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the file containing the metadata from the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#identity_provider_metadata_file OrgSaml#identity_provider_metadata_file}
	IdentityProviderMetadataFile *string `field:"optional" json:"identityProviderMetadataFile" yaml:"identityProviderMetadataFile"`
	// The text of the metadata from the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#identity_provider_metadata_text OrgSaml#identity_provider_metadata_text}
	IdentityProviderMetadataText *string `field:"optional" json:"identityProviderMetadataText" yaml:"identityProviderMetadataText"`
	// Optional role attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#role OrgSaml#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Optional surname attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#surname OrgSaml#surname}
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// Optional username attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_saml#user_name OrgSaml#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

