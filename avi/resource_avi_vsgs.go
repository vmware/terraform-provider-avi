// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceVsGsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"geodb_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gslb_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviVsGs() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVsGsCreate,
		Read:   ResourceAviVsGsRead,
		Update: resourceAviVsGsUpdate,
		Delete: resourceAviVsGsDelete,
		Schema: ResourceVsGsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVsGsImporter,
		},
	}
}

func ResourceVsGsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVsGsSchema()
	return ResourceImporter(d, m, "vsgs", s)
}

func ResourceAviVsGsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	err := APIRead(d, meta, "vsgs", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVsGsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	err := APICreateOrUpdate(d, meta, "vsgs", s)
	if err == nil {
		err = ResourceAviVsGsRead(d, meta)
	}
	return err
}

func resourceAviVsGsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "vsgs", s)
	if err == nil {
		err = ResourceAviVsGsRead(d, meta)
	}
	return err
}

func resourceAviVsGsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vsgs"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviVsGsDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
