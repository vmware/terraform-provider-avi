// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceControllerSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "443",
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
	}
}

func resourceAviControllerSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviControllerSiteCreate,
		Read:   ResourceAviControllerSiteRead,
		Update: resourceAviControllerSiteUpdate,
		Delete: resourceAviControllerSiteDelete,
		Schema: ResourceControllerSiteSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceControllerSiteImporter,
		},
	}
}

func ResourceControllerSiteImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceControllerSiteSchema()
	return ResourceImporter(d, m, "controllersite", s)
}

func ResourceAviControllerSiteRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	err := APIRead(d, meta, "controllersite", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviControllerSiteCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	err := APICreateOrUpdate(d, meta, "controllersite", s)
	if err == nil {
		err = ResourceAviControllerSiteRead(d, meta)
	}
	return err
}

func resourceAviControllerSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "controllersite", s)
	if err == nil {
		err = ResourceAviControllerSiteRead(d, meta)
	}
	return err
}

func resourceAviControllerSiteDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "controllersite")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
