package nsxtedgegateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtEdgegatewayConfig struct {
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
	// External network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#external_network_id NsxtEdgegateway#external_network_id}
	ExternalNetworkId *string `field:"required" json:"externalNetworkId" yaml:"externalNetworkId"`
	// Edge Gateway name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#name NsxtEdgegateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#dedicate_external_network NsxtEdgegateway#dedicate_external_network}
	DedicateExternalNetwork interface{} `field:"optional" json:"dedicateExternalNetwork" yaml:"dedicateExternalNetwork"`
	// Edge Gateway description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#description NsxtEdgegateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#edge_cluster_id NsxtEdgegateway#edge_cluster_id}
	EdgeClusterId *string `field:"optional" json:"edgeClusterId" yaml:"edgeClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#id NsxtEdgegateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#org NsxtEdgegateway#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// ID of VDC or VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#owner_id NsxtEdgegateway#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Optional ID of starting VDC if the 'owner_id' is a VDC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#starting_vdc_id NsxtEdgegateway#starting_vdc_id}
	StartingVdcId *string `field:"optional" json:"startingVdcId" yaml:"startingVdcId"`
	// subnet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#subnet NsxtEdgegateway#subnet}
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// subnet_with_ip_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#subnet_with_ip_count NsxtEdgegateway#subnet_with_ip_count}
	SubnetWithIpCount interface{} `field:"optional" json:"subnetWithIpCount" yaml:"subnetWithIpCount"`
	// subnet_with_total_ip_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#subnet_with_total_ip_count NsxtEdgegateway#subnet_with_total_ip_count}
	SubnetWithTotalIpCount interface{} `field:"optional" json:"subnetWithTotalIpCount" yaml:"subnetWithTotalIpCount"`
	// Total number of IP addresses allocated for this gateway. Can be set with 'subnet_with_total_ip_count' definitions only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#total_allocated_ip_count NsxtEdgegateway#total_allocated_ip_count}
	TotalAllocatedIpCount *float64 `field:"optional" json:"totalAllocatedIpCount" yaml:"totalAllocatedIpCount"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_edgegateway#vdc NsxtEdgegateway#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

