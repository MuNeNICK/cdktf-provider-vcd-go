package lbservicemonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbServiceMonitorConfig struct {
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
	// Edge gateway name in which the LB Service Monitor is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#edge_gateway LbServiceMonitor#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Unique LB Service Monitor name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#name LbServiceMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Way in which you want to send the health check request to the server.
	//
	// One of http, https, tcp, icmp, or udp
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#type LbServiceMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#expected LbServiceMonitor#expected}
	Expected *string `field:"optional" json:"expected" yaml:"expected"`
	// Advanced monitor parameters as key=value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#extension LbServiceMonitor#extension}
	Extension *map[string]*string `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#id LbServiceMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Interval in seconds at which a server is to be monitored (defaults to 10).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#interval LbServiceMonitor#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Number of times the specified monitoring Method must fail sequentially before the server is declared down  (defaults to 3).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#max_retries LbServiceMonitor#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#method LbServiceMonitor#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#org LbServiceMonitor#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// String to be matched in the response content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#receive LbServiceMonitor#receive}
	Receive *string `field:"optional" json:"receive" yaml:"receive"`
	// Data to be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#send LbServiceMonitor#send}
	Send *string `field:"optional" json:"send" yaml:"send"`
	// Maximum time in seconds within which a response from the server must be received  (defaults to 15).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#timeout LbServiceMonitor#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// URL to be used in the server status request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#url LbServiceMonitor#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_service_monitor#vdc LbServiceMonitor#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

