package globalrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalRoleConfig struct {
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
	// Global role description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#description GlobalRole#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of global role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#name GlobalRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, publishes the global role to all tenants.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#publish_to_all_tenants GlobalRole#publish_to_all_tenants}
	PublishToAllTenants interface{} `field:"required" json:"publishToAllTenants" yaml:"publishToAllTenants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#id GlobalRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// list of rights assigned to this global role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#rights GlobalRole#rights}
	Rights *[]*string `field:"optional" json:"rights" yaml:"rights"`
	// list of tenants to which this global role is published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/global_role#tenants GlobalRole#tenants}
	Tenants *[]*string `field:"optional" json:"tenants" yaml:"tenants"`
}

