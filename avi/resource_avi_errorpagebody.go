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
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/errorpagebody/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
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
