// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceIcapProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_204": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"buffer_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "51200",
			ValidateFunc: validateInteger,
		},
		"buffer_size_exceed_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_FAIL_OPEN",
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_preview": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"fail_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_FAIL_OPEN",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nsx_defender_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIcapNsxDefenderConfigSchema(),
		},
		"pool_group_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preview_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5000",
			ValidateFunc: validateInteger,
		},
		"response_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60000",
			ValidateFunc: validateInteger,
		},
		"service_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"slow_response_warning_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10000",
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vendor": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_VENDOR_OPSWAT",
		},
	}
}

func resourceAviIcapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIcapProfileCreate,
		Read:   ResourceAviIcapProfileRead,
		Update: resourceAviIcapProfileUpdate,
		Delete: resourceAviIcapProfileDelete,
		Schema: ResourceIcapProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIcapProfileImporter,
		},
	}
}

func ResourceIcapProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIcapProfileSchema()
	return ResourceImporter(d, m, "icapprofile", s)
}

func ResourceAviIcapProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	err := APIRead(d, meta, "icapprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviIcapProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	err := APICreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "icapprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
