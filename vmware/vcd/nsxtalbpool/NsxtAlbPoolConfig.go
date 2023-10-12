package nsxtalbpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#edge_gateway_id NsxtAlbPool#edge_gateway_id}
	EdgeGatewayId *string `field:"required" json:"edgeGatewayId" yaml:"edgeGatewayId"`
	// Name of ALB Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#name NsxtAlbPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`, `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#algorithm NsxtAlbPool#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// A set of root certificate IDs to use when validating certificates presented by pool members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#ca_certificate_ids NsxtAlbPool#ca_certificate_ids}
	CaCertificateIds *[]*string `field:"optional" json:"caCertificateIds" yaml:"caCertificateIds"`
	// Boolean flag if common name check of the certificate should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#cn_check_enabled NsxtAlbPool#cn_check_enabled}
	CnCheckEnabled interface{} `field:"optional" json:"cnCheckEnabled" yaml:"cnCheckEnabled"`
	// Default Port defines destination server port used by the traffic sent to the member (default 80).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#default_port NsxtAlbPool#default_port}
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// Description of ALB Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#description NsxtAlbPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of domain names which will be used to verify common names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#domain_names NsxtAlbPool#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Boolean value if ALB Pool is enabled or not (default true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#enabled NsxtAlbPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Maximum time in minutes to gracefully disable pool member (default 1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#graceful_timeout_period NsxtAlbPool#graceful_timeout_period}
	GracefulTimeoutPeriod *float64 `field:"optional" json:"gracefulTimeoutPeriod" yaml:"gracefulTimeoutPeriod"`
	// health_monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#health_monitor NsxtAlbPool#health_monitor}
	HealthMonitor interface{} `field:"optional" json:"healthMonitor" yaml:"healthMonitor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#id NsxtAlbPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#member NsxtAlbPool#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
	// ID of Firewall Group to use for Pool Membership (VCD 10.4.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#member_group_id NsxtAlbPool#member_group_id}
	MemberGroupId *string `field:"optional" json:"memberGroupId" yaml:"memberGroupId"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#org NsxtAlbPool#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Monitors if the traffic is accepted by node (default true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#passive_monitoring_enabled NsxtAlbPool#passive_monitoring_enabled}
	PassiveMonitoringEnabled interface{} `field:"optional" json:"passiveMonitoringEnabled" yaml:"passiveMonitoringEnabled"`
	// persistence_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#persistence_profile NsxtAlbPool#persistence_profile}
	PersistenceProfile *NsxtAlbPoolPersistenceProfile `field:"optional" json:"persistenceProfile" yaml:"persistenceProfile"`
	// The name of VDC to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_pool#vdc NsxtAlbPool#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

