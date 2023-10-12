package provider


type VcdProviderConfig struct {
	// The VCD Org for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#org VcdProvider#org}
	Org *string `field:"required" json:"org" yaml:"org"`
	// The VCD url for VCD API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#url VcdProvider#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#alias VcdProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Set this to true if you understand the security risks of using API token files and would like to suppress the warnings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#allow_api_token_file VcdProvider#allow_api_token_file}
	AllowApiTokenFile interface{} `field:"optional" json:"allowApiTokenFile" yaml:"allowApiTokenFile"`
	// Set this to true if you understand the security risks of using Service Account token files and would like to suppress the warnings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#allow_service_account_token_file VcdProvider#allow_service_account_token_file}
	AllowServiceAccountTokenFile interface{} `field:"optional" json:"allowServiceAccountTokenFile" yaml:"allowServiceAccountTokenFile"`
	// If set, VCDClient will permit unverifiable SSL certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#allow_unverified_ssl VcdProvider#allow_unverified_ssl}
	AllowUnverifiedSsl interface{} `field:"optional" json:"allowUnverifiedSsl" yaml:"allowUnverifiedSsl"`
	// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#api_token VcdProvider#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#api_token_file VcdProvider#api_token_file}
	ApiTokenFile *string `field:"optional" json:"apiTokenFile" yaml:"apiTokenFile"`
	// 'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported. 'integrated' is default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#auth_type VcdProvider#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// ignore_metadata_changes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#ignore_metadata_changes VcdProvider#ignore_metadata_changes}
	IgnoreMetadataChanges interface{} `field:"optional" json:"ignoreMetadataChanges" yaml:"ignoreMetadataChanges"`
	// Defines the import separation string to be used with 'terraform import'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#import_separator VcdProvider#import_separator}
	ImportSeparator *string `field:"optional" json:"importSeparator" yaml:"importSeparator"`
	// If set, it will enable logging of API requests and responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#logging VcdProvider#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// Defines the full name of the logging file for API calls (requires 'logging').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#logging_file VcdProvider#logging_file}
	LoggingFile *string `field:"optional" json:"loggingFile" yaml:"loggingFile"`
	// Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#max_retry_timeout VcdProvider#max_retry_timeout}
	MaxRetryTimeout *float64 `field:"optional" json:"maxRetryTimeout" yaml:"maxRetryTimeout"`
	// The user password for VCD API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#password VcdProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#saml_adfs_rpt_id VcdProvider#saml_adfs_rpt_id}
	SamlAdfsRptId *string `field:"optional" json:"samlAdfsRptId" yaml:"samlAdfsRptId"`
	// The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#service_account_token_file VcdProvider#service_account_token_file}
	ServiceAccountTokenFile *string `field:"optional" json:"serviceAccountTokenFile" yaml:"serviceAccountTokenFile"`
	// The VCD Org for user authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#sysorg VcdProvider#sysorg}
	Sysorg *string `field:"optional" json:"sysorg" yaml:"sysorg"`
	// The token used instead of username/password for VCD API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#token VcdProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The user name for VCD API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#user VcdProvider#user}
	User *string `field:"optional" json:"user" yaml:"user"`
	// The VDC for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs#vdc VcdProvider#vdc}
	Vdc *string `field:"optional" json:"vdc" yaml:"vdc"`
}

