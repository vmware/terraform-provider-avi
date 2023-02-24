// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceTenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"config_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTenantConfigurationSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enforce_label_group": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"label_group_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"local": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"name": {
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

func resourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantCreate,
		Read:   ResourceAviTenantRead,
		Update: resourceAviTenantUpdate,
		Delete: resourceAviTenantDelete,
		Schema: ResourceTenantSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTenantImporter,
		},
	}
}

func ResourceTenantImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTenantSchema()
	return ResourceImporter(d, m, "tenant", s)
}

func ResourceAviTenantRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := APIRead(d, meta, "tenant", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTenantCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := APICreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "tenant"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTenantDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
