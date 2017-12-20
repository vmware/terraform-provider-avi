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
	log.Printf("[INFO] ResourceAviErrorPageBodyRead Avi Client %v\n", d)
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
	// no need to set the ID
	log.Printf("ResourceAviErrorPageBodyRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviErrorPageBodyRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviErrorPageBodyRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviErrorPageBodyRead Updated %v\n", d)
	return nil
}

func resourceAviErrorPageBodyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := ApiCreateOrUpdate(d, meta, "errorpagebody", s)
	log.Printf("[DEBUG] created object %v: %v", "errorpagebody", d)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "errorpagebody", d)
	return err
}

func resourceAviErrorPageBodyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageBodySchema()
	err := ApiCreateOrUpdate(d, meta, "errorpagebody", s)
	if err == nil {
		err = ResourceAviErrorPageBodyRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "errorpagebody", d)
	return err
}

func resourceAviErrorPageBodyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpagebody"
	log.Println("[INFO] ResourceAviErrorPageBodyRead Avi Client")
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
