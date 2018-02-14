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

func ResourcePriorityLabelsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud",
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"equivalent_labels": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEquivalentLabelsSchema(),
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

func resourceAviPriorityLabels() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPriorityLabelsCreate,
		Read:   ResourceAviPriorityLabelsRead,
		Update: resourceAviPriorityLabelsUpdate,
		Delete: resourceAviPriorityLabelsDelete,
		Schema: ResourcePriorityLabelsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePriorityLabelsImporter,
		},
	}
}

func ResourcePriorityLabelsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePriorityLabelsSchema()
	return ResourceImporter(d, m, "prioritylabels", s)
}

func ResourceAviPriorityLabelsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := ApiRead(d, meta, "prioritylabels", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPriorityLabelsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := ApiCreateOrUpdate(d, meta, "prioritylabels", s)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	return err
}

func resourceAviPriorityLabelsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	var err error

	err = ApiCreateOrUpdate(d, meta, "prioritylabels", s)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	return err
}

func resourceAviPriorityLabelsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "prioritylabels"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviPriorityLabelsDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
