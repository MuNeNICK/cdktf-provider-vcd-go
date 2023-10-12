package orgvdc


type OrgVdcStorageProfile struct {
	// True if this is default storage profile for this VDC.
	//
	// The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#default OrgVdc#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
	// Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#limit OrgVdc#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Name of Provider VDC storage profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#name OrgVdc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// True if this storage profile is enabled for use in the VDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vcd/3.10.0/docs/resources/org_vdc#enabled OrgVdc#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

