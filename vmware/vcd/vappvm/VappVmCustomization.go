package vappvm


type VappVmCustomization struct {
	// Manually specify admin password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#admin_password VappVm#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Allow local administrator password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#allow_local_admin_password VappVm#allow_local_admin_password}
	AllowLocalAdminPassword interface{} `field:"optional" json:"allowLocalAdminPassword" yaml:"allowLocalAdminPassword"`
	// Auto generate password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#auto_generate_password VappVm#auto_generate_password}
	AutoGeneratePassword interface{} `field:"optional" json:"autoGeneratePassword" yaml:"autoGeneratePassword"`
	// 'true' value will change SID. Applicable only for Windows VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#change_sid VappVm#change_sid}
	ChangeSid interface{} `field:"optional" json:"changeSid" yaml:"changeSid"`
	// 'true' value will enable guest customization. It may occur on first boot or when 'force' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#enabled VappVm#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// 'true' value will cause the VM to reboot on every 'apply' operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#force VappVm#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Script to run on initial boot or with customization.force=true set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#initscript VappVm#initscript}
	Initscript *string `field:"optional" json:"initscript" yaml:"initscript"`
	// Enable this VM to join a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_domain VappVm#join_domain}
	JoinDomain interface{} `field:"optional" json:"joinDomain" yaml:"joinDomain"`
	// Account organizational unit for domain name join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_domain_account_ou VappVm#join_domain_account_ou}
	JoinDomainAccountOu *string `field:"optional" json:"joinDomainAccountOu" yaml:"joinDomainAccountOu"`
	// Custom domain name for join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_domain_name VappVm#join_domain_name}
	JoinDomainName *string `field:"optional" json:"joinDomainName" yaml:"joinDomainName"`
	// Password for custom domain name join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_domain_password VappVm#join_domain_password}
	JoinDomainPassword *string `field:"optional" json:"joinDomainPassword" yaml:"joinDomainPassword"`
	// Username for custom domain name join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_domain_user VappVm#join_domain_user}
	JoinDomainUser *string `field:"optional" json:"joinDomainUser" yaml:"joinDomainUser"`
	// Use organization's domain for joining.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#join_org_domain VappVm#join_org_domain}
	JoinOrgDomain interface{} `field:"optional" json:"joinOrgDomain" yaml:"joinOrgDomain"`
	// Require Administrator to change password on first login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#must_change_password_on_first_login VappVm#must_change_password_on_first_login}
	MustChangePasswordOnFirstLogin interface{} `field:"optional" json:"mustChangePasswordOnFirstLogin" yaml:"mustChangePasswordOnFirstLogin"`
	// Number of times to log on automatically. '0' - disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/vapp_vm#number_of_auto_logons VappVm#number_of_auto_logons}
	NumberOfAutoLogons *float64 `field:"optional" json:"numberOfAutoLogons" yaml:"numberOfAutoLogons"`
}

