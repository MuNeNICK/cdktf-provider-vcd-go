package datavcdrdetype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdRdeTypeConfig struct {
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
	// A unique namespace associated with the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type#nss DataVcdRdeType#nss}
	Nss *string `field:"required" json:"nss" yaml:"nss"`
	// The vendor name of the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type#vendor DataVcdRdeType#vendor}
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
	// The version of the Runtime Defined Entity Type.
	//
	// The version string follows semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type#version DataVcdRdeType#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/rde_type#id DataVcdRdeType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

