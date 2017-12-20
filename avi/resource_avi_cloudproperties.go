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

func ResourceCloudPropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_props": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceCC_PropertiesSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"cc_vtypes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"hyp_props": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHypervisor_PropertiesSchema(),
		},
		"info": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloudInfoSchema(),
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviCloudProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudPropertiesCreate,
		Read:   ResourceAviCloudPropertiesRead,
		Update: resourceAviCloudPropertiesUpdate,
		Delete: resourceAviCloudPropertiesDelete,
		Schema: ResourceCloudPropertiesSchema(),
	}
}

func ResourceAviCloudPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	log.Printf("[INFO] ResourceAviCloudPropertiesRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/cloudproperties/" + uuid.(string)
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
	log.Printf("ResourceAviCloudPropertiesRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviCloudPropertiesRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviCloudPropertiesRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviCloudPropertiesRead Updated %v\n", d)
	return nil
}

func resourceAviCloudPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudproperties", s)
	log.Printf("[DEBUG] created object %v: %v", "cloudproperties", d)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "cloudproperties", d)
	return err
}

func resourceAviCloudPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudproperties", s)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "cloudproperties", d)
	return err
}

func resourceAviCloudPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloudproperties"
	log.Println("[INFO] ResourceAviCloudPropertiesRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviCloudPropertiesDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
