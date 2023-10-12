package nsxtalbserviceenginegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbServiceEngineGroupConfig struct {
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
	// NSX-T ALB backing Cloud ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#alb_cloud_id NsxtAlbServiceEngineGroup#alb_cloud_id}
	AlbCloudId *string `field:"required" json:"albCloudId" yaml:"albCloudId"`
	// NSX-T ALB Importable Service Engine Group Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#importable_service_engine_group_name NsxtAlbServiceEngineGroup#importable_service_engine_group_name}
	ImportableServiceEngineGroupName *string `field:"required" json:"importableServiceEngineGroupName" yaml:"importableServiceEngineGroupName"`
	// NSX-T ALB Service Engine Group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#name NsxtAlbServiceEngineGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// NSX-T ALB Service Engine Group reservation model. One of 'DEDICATED', 'SHARED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#reservation_model NsxtAlbServiceEngineGroup#reservation_model}
	ReservationModel *string `field:"required" json:"reservationModel" yaml:"reservationModel"`
	// NSX-T ALB Service Engine Group description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#description NsxtAlbServiceEngineGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#id NsxtAlbServiceEngineGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Boolean value that shows if virtual services are overallocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#overallocated NsxtAlbServiceEngineGroup#overallocated}
	Overallocated interface{} `field:"optional" json:"overallocated" yaml:"overallocated"`
	// Feature set for this ALB Service Engine Group. One of 'STANDARD', 'PREMIUM'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#supported_feature_set NsxtAlbServiceEngineGroup#supported_feature_set}
	SupportedFeatureSet *string `field:"optional" json:"supportedFeatureSet" yaml:"supportedFeatureSet"`
	// Boolean value that shows if sync should be performed on every refresh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_service_engine_group#sync_on_refresh NsxtAlbServiceEngineGroup#sync_on_refresh}
	SyncOnRefresh interface{} `field:"optional" json:"syncOnRefresh" yaml:"syncOnRefresh"`
}

