// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudConnectorUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_serviceprincipal": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
		},
		"azure_userpass": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAzureUserPassCredentialsSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"gcp_credentials": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGCPCredentialsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nsxt_credentials": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceNsxtCredentialsSchema(),
		},
		"oci_credentials": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOCICredentialsSchema(),
		},
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"private_key": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"public_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tencent_credentials": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTencentCredentialsSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vcenter_credentials": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVCenterCredentialsSchema(),
		},
	}
}

func resourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudConnectorUserCreate,
		Read:   ResourceAviCloudConnectorUserRead,
		Update: resourceAviCloudConnectorUserUpdate,
		Delete: resourceAviCloudConnectorUserDelete,
		Schema: ResourceCloudConnectorUserSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCloudConnectorUserImporter,
		},
	}
}

func ResourceCloudConnectorUserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCloudConnectorUserSchema()
	return ResourceImporter(d, m, "cloudconnectoruser", s)
}

func ResourceAviCloudConnectorUserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := APIRead(d, meta, "cloudconnectoruser", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCloudConnectorUserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := APICreateOrUpdate(d, meta, "cloudconnectoruser", s)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	return err
}

func resourceAviCloudConnectorUserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "cloudconnectoruser", s)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	return err
}

func resourceAviCloudConnectorUserDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "cloudconnectoruser")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
