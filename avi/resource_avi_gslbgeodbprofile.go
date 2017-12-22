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

func ResourceGslbGeoDbProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"entries": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbGeoDbEntrySchema(),
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviGslbGeoDbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbGeoDbProfileCreate,
		Read:   ResourceAviGslbGeoDbProfileRead,
		Update: resourceAviGslbGeoDbProfileUpdate,
		Delete: resourceAviGslbGeoDbProfileDelete,
		Schema: ResourceGslbGeoDbProfileSchema(),
	}
}

func ResourceAviGslbGeoDbProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/gslbgeodbprofile/" + uuid.(string)
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

func resourceAviGslbGeoDbProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "gslbgeodbprofile", s)
	if err == nil {
		err = ResourceAviGslbGeoDbProfileRead(d, meta)
	}
	return err
}

func resourceAviGslbGeoDbProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "gslbgeodbprofile", s)
	if err == nil {
		err = ResourceAviGslbGeoDbProfileRead(d, meta)
	}
	return err
}

func resourceAviGslbGeoDbProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslbgeodbprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviGslbGeoDbProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
