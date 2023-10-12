package lbvirtualserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbVirtualServerConfig struct {
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
	// Edge gateway name in which the LB Virtual Server is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#edge_gateway LbVirtualServer#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// IP address that the load balancer listens on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#ip_address LbVirtualServer#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Unique Virtual Server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#name LbVirtualServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port number that the load balancer listens on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#port LbVirtualServer#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Protocol that the virtual server accepts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#protocol LbVirtualServer#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Application profile ID to be associated with the virtual server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#app_profile_id LbVirtualServer#app_profile_id}
	AppProfileId *string `field:"optional" json:"appProfileId" yaml:"appProfileId"`
	// List of attached application rule IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#app_rule_ids LbVirtualServer#app_rule_ids}
	AppRuleIds *[]*string `field:"optional" json:"appRuleIds" yaml:"appRuleIds"`
	// Maximum concurrent connections that the virtual server can process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#connection_limit LbVirtualServer#connection_limit}
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Maximum incoming new connection requests per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#connection_rate_limit LbVirtualServer#connection_rate_limit}
	ConnectionRateLimit *float64 `field:"optional" json:"connectionRateLimit" yaml:"connectionRateLimit"`
	// Virtual Server description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#description LbVirtualServer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable virtual server acceleration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#enable_acceleration LbVirtualServer#enable_acceleration}
	EnableAcceleration interface{} `field:"optional" json:"enableAcceleration" yaml:"enableAcceleration"`
	// Defines if the virtual server is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#enabled LbVirtualServer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#id LbVirtualServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#org LbVirtualServer#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The server pool that the load balancer will use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#server_pool_id LbVirtualServer#server_pool_id}
	ServerPoolId *string `field:"optional" json:"serverPoolId" yaml:"serverPoolId"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_virtual_server#vdc LbVirtualServer#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

