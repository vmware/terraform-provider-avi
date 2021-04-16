/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceLabelGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleMatchOperationMatchLabelSchema(),
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

func resourceAviLabelGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLabelGroupCreate,
		Read:   ResourceAviLabelGroupRead,
		Update: resourceAviLabelGroupUpdate,
		Delete: resourceAviLabelGroupDelete,
		Schema: ResourceLabelGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLabelGroupImporter,
		},
	}
}

func ResourceLabelGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLabelGroupSchema()
	return ResourceImporter(d, m, "labelgroup", s)
}

func ResourceAviLabelGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	err := APIRead(d, meta, "labelgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLabelGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	err := APICreateOrUpdate(d, meta, "labelgroup", s)
	if err == nil {
		err = ResourceAviLabelGroupRead(d, meta)
	}
	return err
}

func resourceAviLabelGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "labelgroup", s)
	if err == nil {
		err = ResourceAviLabelGroupRead(d, meta)
	}
	return err
}

func resourceAviLabelGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "labelgroup"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviLabelGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
