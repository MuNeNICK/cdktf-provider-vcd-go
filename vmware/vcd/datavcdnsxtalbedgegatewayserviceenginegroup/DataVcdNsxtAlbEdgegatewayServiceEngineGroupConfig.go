package datavcdnsxtalbedgegatewayserviceenginegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtAlbEdgegatewayServiceEngineGroupConfig struct {
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
	// Edge Gateway ID in which ALB Service Engine Group should be located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#edge_gateway_id DataVcdNsxtAlbEdgegatewayServiceEngineGroup#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#id DataVcdNsxtAlbEdgegatewayServiceEngineGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#org DataVcdNsxtAlbEdgegatewayServiceEngineGroup#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Service Engine Group ID which is attached to NSX-T Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#service_engine_group_id DataVcdNsxtAlbEdgegatewayServiceEngineGroup#service_engine_group_id}
	ServiceEngineGroupId *string `field:"optional" json:"serviceEngineGroupId" yaml:"serviceEngineGroupId"`
	// Service Engine Group Name which is attached to NSX-T Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#service_engine_group_name DataVcdNsxtAlbEdgegatewayServiceEngineGroup#service_engine_group_name}
	ServiceEngineGroupName *string `field:"optional" json:"serviceEngineGroupName" yaml:"serviceEngineGroupName"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_edgegateway_service_engine_group#vdc DataVcdNsxtAlbEdgegatewayServiceEngineGroup#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

