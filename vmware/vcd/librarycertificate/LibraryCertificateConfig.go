package librarycertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LibraryCertificateConfig struct {
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
	// Alias of certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#alias LibraryCertificate#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Certificate content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#certificate LibraryCertificate#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Certificate description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#description LibraryCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#id LibraryCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of organization to use, optional if defined at provider level.
	//
	// Useful when connected as sysadmin working across different organizations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#org LibraryCertificate#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Certificate private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#private_key LibraryCertificate#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Certificate private pass phrase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/library_certificate#private_key_passphrase LibraryCertificate#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
}

