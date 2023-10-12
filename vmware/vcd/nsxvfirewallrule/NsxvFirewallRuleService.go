package nsxvfirewallrule


type NsxvFirewallRuleService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#protocol NsxvFirewallRule#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#port NsxvFirewallRule#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxv_firewall_rule#source_port NsxvFirewallRule#source_port}.
	SourcePort *string `field:"optional" json:"sourcePort" yaml:"sourcePort"`
}

