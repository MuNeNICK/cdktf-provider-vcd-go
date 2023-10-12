package datavcdnsxtalbpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVcdNsxtAlbPoolConfig struct {
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
	// Edge gateway ID in which ALB Pool should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#edge_gateway_id DataVcdNsxtAlbPool#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Name of ALB Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#name DataVcdNsxtAlbPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// IDs of associated virtual services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#associated_virtual_service_ids DataVcdNsxtAlbPool#associated_virtual_service_ids}
	AssociatedVirtualServiceIds *[]*string `field:"optional" json:"associatedVirtualServiceIds" yaml:"associatedVirtualServiceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#id DataVcdNsxtAlbPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#org DataVcdNsxtAlbPool#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Monitors if the traffic is accepted by node (default true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#passive_monitoring_enabled DataVcdNsxtAlbPool#passive_monitoring_enabled}
	PassiveMonitoringEnabled interface{} `field:"optional" json:"passiveMonitoringEnabled" yaml:"passiveMonitoringEnabled"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/data-sources/nsxt_alb_pool#vdc DataVcdNsxtAlbPool#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

