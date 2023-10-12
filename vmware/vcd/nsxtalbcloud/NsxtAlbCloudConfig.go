package nsxtalbcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NsxtAlbCloudConfig struct {
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
	// NSX-T ALB Controller ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#controller_id NsxtAlbCloud#controller_id}
	ControllerId *string `field:"required" json:"controllerId" yaml:"controllerId"`
	// NSX-T ALB Importable Cloud ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#importable_cloud_id NsxtAlbCloud#importable_cloud_id}
	ImportableCloudId *string `field:"required" json:"importableCloudId" yaml:"importableCloudId"`
	// NSX-T ALB Cloud name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#name NsxtAlbCloud#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network pool ID for NSX-T ALB Importable Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#network_pool_id NsxtAlbCloud#network_pool_id}
	NetworkPoolId *string `field:"required" json:"networkPoolId" yaml:"networkPoolId"`
	// NSX-T ALB Cloud description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#description NsxtAlbCloud#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/nsxt_alb_cloud#id NsxtAlbCloud#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

