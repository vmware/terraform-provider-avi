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

func ResourceErrorPageBodySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error_page_body": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ERROR_PAGE_FORMAT_HTML",
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

func resourceAviErrorPageBody() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviErrorPageBodyCreate,
		Read:   ResourceAviErrorPageBodyRead,
		Update: resourceAviErrorPageBodyUpdate,
		Delete: resourceAviErrorPageBodyDelete,
		Schema: ResourceErrorPageBodySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceErrorPageBodyImporter,
		},
	}
}

func ResourceErrorPageBodyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceErrorPageBodySchema()
	return ResourceImporter(d, m, "errorpagebody", s)
}

func ResourceAviErrorPageBodyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := APIRead(d, meta, "errorpagebody", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviErrorPageBodyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := APICreateOrUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpagebody"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviErrorPageBodyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
