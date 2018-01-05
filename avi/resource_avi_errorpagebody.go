/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceErrorPageBodySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error_page_body": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
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
	}
}

func ResourceAviErrorPageBodyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := ApiRead(d, meta, "errorpagebody", s)
	return err
}

func resourceAviErrorPageBodyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := ApiCreateOrUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := ApiCreateOrUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	return err
}

func resourceAviErrorPageBodyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpagebody"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviErrorPageBodyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
