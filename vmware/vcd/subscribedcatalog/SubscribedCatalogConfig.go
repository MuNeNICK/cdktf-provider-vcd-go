package subscribedcatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscribedCatalogConfig struct {
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
	// The name of the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#name SubscribedCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL to subscribe to the external catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#subscription_url SubscribedCatalog#subscription_url}
	SubscriptionUrl *string `field:"required" json:"subscriptionUrl" yaml:"subscriptionUrl"`
	// When true, the subscribed catalog will attempt canceling failed tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#cancel_failed_tasks SubscribedCatalog#cancel_failed_tasks}
	CancelFailedTasks interface{} `field:"optional" json:"cancelFailedTasks" yaml:"cancelFailedTasks"`
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains, regardless of their state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#delete_force SubscribedCatalog#delete_force}
	DeleteForce interface{} `field:"optional" json:"deleteForce" yaml:"deleteForce"`
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#delete_recursive SubscribedCatalog#delete_recursive}
	DeleteRecursive interface{} `field:"optional" json:"deleteRecursive" yaml:"deleteRecursive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#id SubscribedCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, subscription to a catalog creates a local copy of all items.
	//
	// Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#make_local_copy SubscribedCatalog#make_local_copy}
	MakeLocalCopy interface{} `field:"optional" json:"makeLocalCopy" yaml:"makeLocalCopy"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#org SubscribedCatalog#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Optional storage profile ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#storage_profile_id SubscribedCatalog#storage_profile_id}
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// If true, saves list of tasks to file for later update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#store_tasks SubscribedCatalog#store_tasks}
	StoreTasks interface{} `field:"optional" json:"storeTasks" yaml:"storeTasks"`
	// An optional password to access the catalog.
	//
	// Only ASCII characters are allowed in a valid password. Passing in six asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#subscription_password SubscribedCatalog#subscription_password}
	SubscriptionPassword *string `field:"optional" json:"subscriptionPassword" yaml:"subscriptionPassword"`
	// If true, synchronise this catalog and all items.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_all SubscribedCatalog#sync_all}
	SyncAll interface{} `field:"optional" json:"syncAll" yaml:"syncAll"`
	// If true, synchronises all media items.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_all_media_items SubscribedCatalog#sync_all_media_items}
	SyncAllMediaItems interface{} `field:"optional" json:"syncAllMediaItems" yaml:"syncAllMediaItems"`
	// If true, synchronises all vApp templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_all_vapp_templates SubscribedCatalog#sync_all_vapp_templates}
	SyncAllVappTemplates interface{} `field:"optional" json:"syncAllVappTemplates" yaml:"syncAllVappTemplates"`
	// If true, synchronise this catalog.
	//
	// This operation fetches the list of items. If `make_local_copy` is set, it also fetches the items data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_catalog SubscribedCatalog#sync_catalog}
	SyncCatalog interface{} `field:"optional" json:"syncCatalog" yaml:"syncCatalog"`
	// Synchronises media items from this list of names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_media_items SubscribedCatalog#sync_media_items}
	SyncMediaItems *[]*string `field:"optional" json:"syncMediaItems" yaml:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_on_refresh SubscribedCatalog#sync_on_refresh}
	SyncOnRefresh interface{} `field:"optional" json:"syncOnRefresh" yaml:"syncOnRefresh"`
	// Synchronises vApp templates from this list of names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/subscribed_catalog#sync_vapp_templates SubscribedCatalog#sync_vapp_templates}
	SyncVappTemplates *[]*string `field:"optional" json:"syncVappTemplates" yaml:"syncVappTemplates"`
}

