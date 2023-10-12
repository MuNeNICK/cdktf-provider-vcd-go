package nsxtalbedgegatewayserviceenginegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbEdgegatewayServiceEngineGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#edge_gateway_id NsxtAlbEdgegatewayServiceEngineGroup#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Service Engine Group ID to attach to this NSX-T Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#service_engine_group_id NsxtAlbEdgegatewayServiceEngineGroup#service_engine_group_id}
	ServiceEngineGroupId *string `field:"required" json:"serviceEngineGroupId" yaml:"serviceEngineGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#id NsxtAlbEdgegatewayServiceEngineGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum number of virtual services to be used in this Service Engine Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#max_virtual_services NsxtAlbEdgegatewayServiceEngineGroup#max_virtual_services}
	MaxVirtualServices *float64 `field:"optional" json:"maxVirtualServices" yaml:"maxVirtualServices"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#org NsxtAlbEdgegatewayServiceEngineGroup#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Number of reserved virtual services for this Service Engine Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#reserved_virtual_services NsxtAlbEdgegatewayServiceEngineGroup#reserved_virtual_services}
	ReservedVirtualServices *string `field:"optional" json:"reservedVirtualServices" yaml:"reservedVirtualServices"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_edgegateway_service_engine_group#vdc NsxtAlbEdgegatewayServiceEngineGroup#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

