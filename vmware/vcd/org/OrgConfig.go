package org

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgConfig struct {
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
	// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains, regardless of their state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#delete_force Org#delete_force}
	DeleteForce interface{} `field:"required" json:"deleteForce" yaml:"deleteForce"`
	// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that normally allows removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#delete_recursive Org#delete_recursive}
	DeleteRecursive interface{} `field:"required" json:"deleteRecursive" yaml:"deleteRecursive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#full_name Org#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#name Org#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// True if this organization is allowed to share catalogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#can_publish_catalogs Org#can_publish_catalogs}
	CanPublishCatalogs interface{} `field:"optional" json:"canPublishCatalogs" yaml:"canPublishCatalogs"`
	// True if this organization is allowed to publish external catalogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#can_publish_external_catalogs Org#can_publish_external_catalogs}
	CanPublishExternalCatalogs interface{} `field:"optional" json:"canPublishExternalCatalogs" yaml:"canPublishExternalCatalogs"`
	// True if this organization is allowed to subscribe to external catalogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#can_subscribe_external_catalogs Org#can_subscribe_external_catalogs}
	CanSubscribeExternalCatalogs interface{} `field:"optional" json:"canSubscribeExternalCatalogs" yaml:"canSubscribeExternalCatalogs"`
	// Specifies this organization's default for virtual machine boot delay after power on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#delay_after_power_on_seconds Org#delay_after_power_on_seconds}
	DelayAfterPowerOnSeconds *float64 `field:"optional" json:"delayAfterPowerOnSeconds" yaml:"delayAfterPowerOnSeconds"`
	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#deployed_vm_quota Org#deployed_vm_quota}
	DeployedVmQuota *float64 `field:"optional" json:"deployedVmQuota" yaml:"deployedVmQuota"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#description Org#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#id Org#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// True if this organization is enabled (allows login and all other operations).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#is_enabled Org#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Key value map of metadata to assign to this organization. Key and value can be any string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#metadata Org#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// metadata_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#metadata_entry Org#metadata_entry}
	MetadataEntry interface{} `field:"optional" json:"metadataEntry" yaml:"metadataEntry"`
	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.
	//
	// (0 = unlimited)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#stored_vm_quota Org#stored_vm_quota}
	StoredVmQuota *float64 `field:"optional" json:"storedVmQuota" yaml:"storedVmQuota"`
	// vapp_lease block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#vapp_lease Org#vapp_lease}
	VappLease *OrgVappLease `field:"optional" json:"vappLease" yaml:"vappLease"`
	// vapp_template_lease block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org#vapp_template_lease Org#vapp_template_lease}
	VappTemplateLease *OrgVappTemplateLease `field:"optional" json:"vappTemplateLease" yaml:"vappTemplateLease"`
}

