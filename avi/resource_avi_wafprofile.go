// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWafProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceWafConfigSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"files": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafDataFileSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafProfileCreate,
		Read:   ResourceAviWafProfileRead,
		Update: resourceAviWafProfileUpdate,
		Delete: resourceAviWafProfileDelete,
		Schema: ResourceWafProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafProfileImporter,
		},
	}
}

func ResourceWafProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafProfileSchema()
	return ResourceImporter(d, m, "wafprofile", s)
}

func ResourceAviWafProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := APIRead(d, meta, "wafprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := APICreateOrUpdate(d, meta, "wafprofile", s)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	return err
}

func resourceAviWafProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafprofile", s)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	return err
}

func resourceAviWafProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
