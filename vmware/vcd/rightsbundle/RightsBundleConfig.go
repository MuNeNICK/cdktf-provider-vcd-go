package rightsbundle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RightsBundleConfig struct {
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
	// Rights bundle description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#description RightsBundle#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of rights bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#name RightsBundle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, publishes the rights bundle to all tenants.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#publish_to_all_tenants RightsBundle#publish_to_all_tenants}
	PublishToAllTenants interface{} `field:"required" json:"publishToAllTenants" yaml:"publishToAllTenants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#id RightsBundle#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of rights assigned to this rights bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#rights RightsBundle#rights}
	Rights *[]*string `field:"optional" json:"rights" yaml:"rights"`
	// Set of tenants to which this rights bundle is published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/rights_bundle#tenants RightsBundle#tenants}
	Tenants *[]*string `field:"optional" json:"tenants" yaml:"tenants"`
}

