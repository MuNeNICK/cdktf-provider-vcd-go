package uiplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UiPluginConfig struct {
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
	// true to make the UI Plugin enabled. 'false' to make it disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#enabled UiPlugin#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#id UiPlugin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Absolute or relative path to the ZIP file containing the UI Plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#plugin_path UiPlugin#plugin_path}
	PluginPath *string `field:"optional" json:"pluginPath" yaml:"pluginPath"`
	// This value is calculated automatically on create by reading the UI Plugin ZIP file contents.
	//
	// You can updateit to `true` to make it provider scoped or `false` otherwise
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#provider_scoped UiPlugin#provider_scoped}
	ProviderScoped interface{} `field:"optional" json:"providerScoped" yaml:"providerScoped"`
	// Set of organization IDs to which this UI Plugin must be published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#tenant_ids UiPlugin#tenant_ids}
	TenantIds *[]*string `field:"optional" json:"tenantIds" yaml:"tenantIds"`
	// This value is calculated automatically on create by reading the UI Plugin ZIP file contents.
	//
	// You can updateit to `true` to make it tenant scoped or `false` otherwise
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/ui_plugin#tenant_scoped UiPlugin#tenant_scoped}
	TenantScoped interface{} `field:"optional" json:"tenantScoped" yaml:"tenantScoped"`
}

