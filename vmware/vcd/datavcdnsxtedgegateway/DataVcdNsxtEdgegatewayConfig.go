package datavcdnsxtedgegateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtEdgegatewayConfig struct {
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
	// Edge Gateway name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#name DataVcdNsxtEdgegateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// NSX-T Edge Cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#edge_cluster_id DataVcdNsxtEdgegateway#edge_cluster_id}
	EdgeClusterId *string `field:"optional" json:"edgeClusterId" yaml:"edgeClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#id DataVcdNsxtEdgegateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#org DataVcdNsxtEdgegateway#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// ID of VDC or VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#owner_id DataVcdNsxtEdgegateway#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_edgegateway#vdc DataVcdNsxtEdgegateway#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

