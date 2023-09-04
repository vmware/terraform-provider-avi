// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSiteVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datetime": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"prev_target_version": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"replication_state": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"target_timeline": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"target_version": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"timeline": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"version_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSiteVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSiteVersionCreate,
		Read:   ResourceAviSiteVersionRead,
		Update: resourceAviSiteVersionUpdate,
		Delete: resourceAviSiteVersionDelete,
		Schema: ResourceSiteVersionSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSiteVersionImporter,
		},
	}
}

func ResourceSiteVersionImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSiteVersionSchema()
	return ResourceImporter(d, m, "siteversion", s)
}

func ResourceAviSiteVersionRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	err := APIRead(d, meta, "siteversion", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSiteVersionCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	err := APICreateOrUpdate(d, meta, "siteversion", s)
	if err == nil {
		err = ResourceAviSiteVersionRead(d, meta)
	}
	return err
}

func resourceAviSiteVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSiteVersionSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "siteversion", s)
	if err == nil {
		err = ResourceAviSiteVersionRead(d, meta)
	}
	return err
}

func resourceAviSiteVersionDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "siteversion")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
