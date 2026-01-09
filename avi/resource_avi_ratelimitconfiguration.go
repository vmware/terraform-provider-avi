// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceRateLimitConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"burst": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
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
		"http_methods": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"token_refill_rate": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceTokenRefillRateSchema(),
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "RATE_LIMITER_API_CATEGORY",
			ForceNew: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviRateLimitConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRateLimitConfigurationCreate,
		Read:   ResourceAviRateLimitConfigurationRead,
		Update: resourceAviRateLimitConfigurationUpdate,
		Delete: resourceAviRateLimitConfigurationDelete,
		Schema: ResourceRateLimitConfigurationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRateLimitConfigurationImporter,
		},
	}
}

func ResourceRateLimitConfigurationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRateLimitConfigurationSchema()
	return ResourceImporter(d, m, "ratelimitconfiguration", s)
}

func ResourceAviRateLimitConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRateLimitConfigurationSchema()
	err := APIRead(d, meta, "ratelimitconfiguration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRateLimitConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRateLimitConfigurationSchema()
	err := APICreate(d, meta, "ratelimitconfiguration", s)
	if err == nil {
		err = ResourceAviRateLimitConfigurationRead(d, meta)
	}
	return err
}

func resourceAviRateLimitConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRateLimitConfigurationSchema()
	var err error
	err = APIUpdate(d, meta, "ratelimitconfiguration", s)
	if err == nil {
		err = ResourceAviRateLimitConfigurationRead(d, meta)
	}
	return err
}

func resourceAviRateLimitConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "ratelimitconfiguration")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
