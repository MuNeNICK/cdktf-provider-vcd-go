package datavcdnsxvapplicationfinder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxvApplicationFinderConfig struct {
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
	// Regular expression used to search applications or groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxv_application_finder#search_expression DataVcdNsxvApplicationFinder#search_expression}
	SearchExpression *string `field:"required" json:"searchExpression" yaml:"searchExpression"`
	// Type of object. One of 'application', 'application_group'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxv_application_finder#type DataVcdNsxvApplicationFinder#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxv_application_finder#vdc_id DataVcdNsxvApplicationFinder#vdc_id}
	VdcId *string `field:"required" json:"vdcId" yaml:"vdcId"`
	// Convert the search to case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxv_application_finder#case_sensitive DataVcdNsxvApplicationFinder#case_sensitive}
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxv_application_finder#id DataVcdNsxvApplicationFinder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

