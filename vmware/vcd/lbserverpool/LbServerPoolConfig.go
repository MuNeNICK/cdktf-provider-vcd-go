package lbserverpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbServerPoolConfig struct {
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
	// Balancing method for the service. One of 'ip-hash', 'round-robin', 'uri', 'leastconn', 'url', or 'httpheader'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#algorithm LbServerPool#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Edge gateway name in which the LB Server Pool is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#edge_gateway LbServerPool#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Unique LB Server Pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#name LbServerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional options for load balancing algorithm for httpheader or url algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#algorithm_parameters LbServerPool#algorithm_parameters}
	AlgorithmParameters *string `field:"optional" json:"algorithmParameters" yaml:"algorithmParameters"`
	// Server pool description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#description LbServerPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Makes client IP addresses visible to the backend servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#enable_transparency LbServerPool#enable_transparency}
	EnableTransparency interface{} `field:"optional" json:"enableTransparency" yaml:"enableTransparency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#id LbServerPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#member LbServerPool#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
	// Load Balancer Service Monitor ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#monitor_id LbServerPool#monitor_id}
	MonitorId *string `field:"optional" json:"monitorId" yaml:"monitorId"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#org LbServerPool#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#vdc LbServerPool#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

