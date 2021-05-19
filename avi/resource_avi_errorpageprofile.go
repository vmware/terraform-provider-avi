/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceErrorPageProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"error_pages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceErrorPageSchema(),
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

func resourceAviErrorPageProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviErrorPageProfileCreate,
		Read:   ResourceAviErrorPageProfileRead,
		Update: resourceAviErrorPageProfileUpdate,
		Delete: resourceAviErrorPageProfileDelete,
		Schema: ResourceErrorPageProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceErrorPageProfileImporter,
		},
	}
}

func ResourceErrorPageProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceErrorPageProfileSchema()
	return ResourceImporter(d, m, "errorpageprofile", s)
}

func ResourceAviErrorPageProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := APIRead(d, meta, "errorpageprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviErrorPageProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := APICreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpageprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviErrorPageProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
