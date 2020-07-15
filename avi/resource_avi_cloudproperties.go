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
		"cc_props": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCC_PropertiesSchema(),
		},
		"cc_vtypes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"hyp_props": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHypervisor_PropertiesSchema(),
		},
		"info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloudInfoSchema(),
		},
		"uuid": {
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
		Importer: &schema.ResourceImporter{
			State: ResourceCloudPropertiesImporter,
		},
	}
}

func ResourceCloudPropertiesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCloudPropertiesSchema()
	return ResourceImporter(d, m, "cloudproperties", s)
}

func ResourceAviCloudPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := ApiRead(d, meta, "cloudproperties", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCloudPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudproperties", s)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	return err
}

func resourceAviCloudPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudPropertiesSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "cloudproperties", s)
	if err == nil {
		err = ResourceAviCloudPropertiesRead(d, meta)
	}
	return err
}

func resourceAviCloudPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloudproperties"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCloudPropertiesDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
