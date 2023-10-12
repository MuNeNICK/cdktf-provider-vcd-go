package serviceaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceAccountConfig struct {
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
	// Name of service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#name ServiceAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role ID of service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#role_id ServiceAccount#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Any valid UUID, depends on the user, e.g: 12345678-1234-5678-90ab-1234567890ab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#software_id ServiceAccount#software_id}
	SoftwareId *string `field:"required" json:"softwareId" yaml:"softwareId"`
	// Status of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#active ServiceAccount#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Set this to true if you understand the security risks of using API token files and would like to suppress the warnings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#allow_token_file ServiceAccount#allow_token_file}
	AllowTokenFile interface{} `field:"optional" json:"allowTokenFile" yaml:"allowTokenFile"`
	// Name of the file that the API token will be saved to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#file_name ServiceAccount#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#id ServiceAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#org ServiceAccount#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Version of software using the service account, can be freely defined by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#software_version ServiceAccount#software_version}
	SoftwareVersion *string `field:"optional" json:"softwareVersion" yaml:"softwareVersion"`
	// URI of the client using the service account, can be freely defined by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/service_account#uri ServiceAccount#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

