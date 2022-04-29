// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceAuthMappingProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"mapping_rules": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceAuthMappingRuleSchema(),
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

func resourceAviAuthMappingProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAuthMappingProfileCreate,
		Read:   ResourceAviAuthMappingProfileRead,
		Update: resourceAviAuthMappingProfileUpdate,
		Delete: resourceAviAuthMappingProfileDelete,
		Schema: ResourceAuthMappingProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAuthMappingProfileImporter,
		},
	}
}

func ResourceAuthMappingProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAuthMappingProfileSchema()
	return ResourceImporter(d, m, "authmappingprofile", s)
}

func ResourceAviAuthMappingProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthMappingProfileSchema()
	err := APIRead(d, meta, "authmappingprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAuthMappingProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthMappingProfileSchema()
	err := APICreateOrUpdate(d, meta, "authmappingprofile", s)
	if err == nil {
		err = ResourceAviAuthMappingProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthMappingProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthMappingProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "authmappingprofile", s)
	if err == nil {
		err = ResourceAviAuthMappingProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthMappingProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "authmappingprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAuthMappingProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
