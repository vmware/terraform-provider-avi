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
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"pool_group_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pool_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"protocol_parser_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"string_group_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		Importer: &schema.ResourceImporter{
			State: ResourceVSDataScriptSetImporter,
		},
	}
}

func ResourceVSDataScriptSetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVSDataScriptSetSchema()
	return ResourceImporter(d, m, "vsdatascriptset", s)
}

func ResourceAviVSDataScriptSetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVSDataScriptSetSchema()
	err := ApiRead(d, meta, "vsdatascriptset", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
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
	var err error
	err = ApiCreateOrUpdate(d, meta, "vsdatascriptset", s)
	if err == nil {
		err = ResourceAviVSDataScriptSetRead(d, meta)
	}
	return err
}

func resourceAviVSDataScriptSetDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vsdatascriptset"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviVSDataScriptSetDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
