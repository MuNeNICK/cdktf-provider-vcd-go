package providervdc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProviderVdcConfig struct {
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
	// The highest virtual hardware version supported by this Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#highest_supported_hardware_version ProviderVdc#highest_supported_hardware_version}
	HighestSupportedHardwareVersion *string `field:"required" json:"highestSupportedHardwareVersion" yaml:"highestSupportedHardwareVersion"`
	// Name of the Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#name ProviderVdc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#nsxt_manager_id ProviderVdc#nsxt_manager_id}
	NsxtManagerId *string `field:"required" json:"nsxtManagerId" yaml:"nsxtManagerId"`
	// Set of IDs of the resource pools backing this provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#resource_pool_ids ProviderVdc#resource_pool_ids}
	ResourcePoolIds *[]*string `field:"required" json:"resourcePoolIds" yaml:"resourcePoolIds"`
	// Set of storage profile names used to create this Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#storage_profile_names ProviderVdc#storage_profile_names}
	StorageProfileNames *[]*string `field:"required" json:"storageProfileNames" yaml:"storageProfileNames"`
	// ID of the vCenter server that provides the resource pools and datastores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#vcenter_id ProviderVdc#vcenter_id}
	VcenterId *string `field:"required" json:"vcenterId" yaml:"vcenterId"`
	// Optional description of the Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#description ProviderVdc#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#id ProviderVdc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// True if this Provider VDC is enabled and can provide resources to organization VDCs.
	//
	// A Provider VDC is always enabled on creation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#is_enabled ProviderVdc#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Set IDs of the network pools used by this Provider VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/provider_vdc#network_pool_ids ProviderVdc#network_pool_ids}
	NetworkPoolIds *[]*string `field:"optional" json:"networkPoolIds" yaml:"networkPoolIds"`
}

