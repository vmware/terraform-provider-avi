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
	}
}

func ResourceAviPriorityLabelsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	log.Printf("[INFO] ResourceAviPriorityLabelsRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/prioritylabels/" + uuid.(string)
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
	log.Printf("ResourceAviPriorityLabelsRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviPriorityLabelsRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviPriorityLabelsRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviPriorityLabelsRead Updated %v\n", d)
	return nil
}

func resourceAviPriorityLabelsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := ApiCreateOrUpdate(d, meta, "prioritylabels", s)
	log.Printf("[DEBUG] created object %v: %v", "prioritylabels", d)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "prioritylabels", d)
	return err
}

func resourceAviPriorityLabelsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePriorityLabelsSchema()
	err := ApiCreateOrUpdate(d, meta, "prioritylabels", s)
	if err == nil {
		err = ResourceAviPriorityLabelsRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "prioritylabels", d)
	return err
}

func resourceAviPriorityLabelsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "prioritylabels"
	log.Println("[INFO] ResourceAviPriorityLabelsRead Avi Client")
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
