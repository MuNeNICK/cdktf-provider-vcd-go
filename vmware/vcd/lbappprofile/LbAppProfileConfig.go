package lbappprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbAppProfileConfig struct {
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
	// Edge gateway name in which the LB Application Profile is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#edge_gateway LbAppProfile#edge_gateway}
	EdgeGateway *string `field:"required" json:"edgeGateway" yaml:"edgeGateway"`
	// Unique LB Application Profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#name LbAppProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#type LbAppProfile#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#cookie_mode LbAppProfile#cookie_mode}
	CookieMode *string `field:"optional" json:"cookieMode" yaml:"cookieMode"`
	// Used to uniquely identify the session the first time a client accesses the site.
	//
	// The load balancer refers to this cookie when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for persistence_mechanism 'cookie'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#cookie_name LbAppProfile#cookie_name}
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#enable_pool_side_ssl LbAppProfile#enable_pool_side_ssl}
	EnablePoolSideSsl interface{} `field:"optional" json:"enablePoolSideSsl" yaml:"enablePoolSideSsl"`
	// Enable SSL authentication to be passed through to the virtual server.
	//
	// Otherwise SSL authentication takes place at the destination address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#enable_ssl_passthrough LbAppProfile#enable_ssl_passthrough}
	EnableSslPassthrough interface{} `field:"optional" json:"enableSslPassthrough" yaml:"enableSslPassthrough"`
	// Length of time in seconds that persistence stays in effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#expiration LbAppProfile#expiration}
	Expiration *float64 `field:"optional" json:"expiration" yaml:"expiration"`
	// The URL to which traffic that arrives at the destination address should be redirected.
	//
	// Only applies for types 'http' and 'https'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#http_redirect_url LbAppProfile#http_redirect_url}
	HttpRedirectUrl *string `field:"optional" json:"httpRedirectUrl" yaml:"httpRedirectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#id LbAppProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server through the load balancer.
	//
	// Only applies for types HTTP and HTTPS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#insert_x_forwarded_http_header LbAppProfile#insert_x_forwarded_http_header}
	InsertXForwardedHttpHeader interface{} `field:"optional" json:"insertXForwardedHttpHeader" yaml:"insertXForwardedHttpHeader"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#org LbAppProfile#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#persistence_mechanism LbAppProfile#persistence_mechanism}
	PersistenceMechanism *string `field:"optional" json:"persistenceMechanism" yaml:"persistenceMechanism"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_app_profile#vdc LbAppProfile#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

