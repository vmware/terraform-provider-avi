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

func ResourceVSDataScriptSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"datascript": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVSDataScriptSchema(),
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ipgroup_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"pool_group_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pool_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"string_group_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviVSDataScriptSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVSDataScriptSetCreate,
		Read:   ResourceAviVSDataScriptSetRead,
		Update: resourceAviVSDataScriptSetUpdate,
		Delete: resourceAviVSDataScriptSetDelete,
		Schema: ResourceVSDataScriptSetSchema(),
	}
}

func ResourceAviVSDataScriptSetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/vsdatascriptset/" + uuid.(string)
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

func resourceAviVSDataScriptSetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	err := ApiCreateOrUpdate(d, meta, "vsdatascriptset", s)
	if err == nil {
		err = ResourceAviVSDataScriptSetRead(d, meta)
	}
	return err
}

func resourceAviVSDataScriptSetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	err := ApiCreateOrUpdate(d, meta, "vsdatascriptset", s)
	if err == nil {
		err = ResourceAviVSDataScriptSetRead(d, meta)
	}
	return err
}

func resourceAviVSDataScriptSetDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vsdatascriptset"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviVSDataScriptSetDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
