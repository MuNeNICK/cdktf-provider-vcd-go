package lbserverpool


type LbServerPoolMember struct {
	// Defines member state. One of enabled, drain, disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#condition LbServerPool#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// IP address of member in server pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#ip_address LbServerPool#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Port at which the member is to receive health monitor requests. Can be the same as port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#monitor_port LbServerPool#monitor_port}
	MonitorPort *float64 `field:"required" json:"monitorPort" yaml:"monitorPort"`
	// Name of pool member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#name LbServerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port at which the member is to receive traffic from the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#port LbServerPool#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Proportion of traffic this member is to handle. Must be an integer in the range 1-256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#weight LbServerPool#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// The maximum number of concurrent connections the member can handle.
	//
	// If exceeded requests are queued and the load balancer waits for a connection to be released
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#max_connections LbServerPool#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// Minimum number of concurrent connections a member must always accept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/lb_server_pool#min_connections LbServerPool#min_connections}
	MinConnections *float64 `field:"optional" json:"minConnections" yaml:"minConnections"`
}

