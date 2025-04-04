// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

//nolint
func ResourceIpamDnsProviderProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocate_ip_in_vrf": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"aws_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsAwsProfileSchema(),
		},
		"azure_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsAzureProfileSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"custom_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsCustomProfileSchema(),
		},
		"gcp_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsGCPProfileSchema(),
		},
		"infoblox_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsInfobloxProfileSchema(),
		},
		"internal_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsInternalProfileSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"oci_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsOCIProfileSchema(),
		},
		"openstack_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsOpenstackProfileSchema(),
		},
		"proxy_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tencent_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsTencentProfileSchema(),
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

//nolint
func resourceAviIpamDnsProviderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIpamDnsProviderProfileCreate,
		Read:   ResourceAviIpamDnsProviderProfileRead,
		Update: resourceAviIpamDnsProviderProfileUpdate,
		Delete: resourceAviIpamDnsProviderProfileDelete,
		Schema: ResourceIpamDnsProviderProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIpamDnsProviderProfileImporter,
		},
	}
}

//nolint
func ResourceIpamDnsProviderProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIpamDnsProviderProfileSchema()
	return ResourceImporter(d, m, "ipamdnsproviderprofile", s)
}

//nolint
func ResourceAviIpamDnsProviderProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	err := APIRead(d, meta, "ipamdnsproviderprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

//nolint
func resourceAviIpamDnsProviderProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	err := APICreate(d, meta, "ipamdnsproviderprofile", s)
	if err == nil {
		err = ResourceAviIpamDnsProviderProfileRead(d, meta)
	}
	return err
}

//nolint
func resourceAviIpamDnsProviderProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	var err error
	err = APIUpdate(d, meta, "ipamdnsproviderprofile", s)
	if err == nil {
		err = ResourceAviIpamDnsProviderProfileRead(d, meta)
	}
	return err
}

//nolint
func resourceAviIpamDnsProviderProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "ipamdnsproviderprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
