package nsxtalbcontroller

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbControllerConfig struct {
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
	// NSX-T ALB Controller name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#name NsxtAlbController#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// NSX-T ALB Controller Password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#password NsxtAlbController#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// NSX-T ALB Controller URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#url NsxtAlbController#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// NSX-T ALB Controller Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#username NsxtAlbController#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// NSX-T ALB Controller description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#description NsxtAlbController#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#id NsxtAlbController#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// NSX-T ALB License type. One of 'BASIC', 'ENTERPRISE'. Must not be used from VCD 10.4.0 onwards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_controller#license_type NsxtAlbController#license_type}
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
}

