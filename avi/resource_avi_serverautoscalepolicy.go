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

func ResourceServerAutoScalePolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"intelligent_autoscale": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"intelligent_scalein_margin": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  40,
		},
		"intelligent_scaleout_margin": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"max_scalein_adjustment_step": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_scaleout_adjustment_step": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_size": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"min_size": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"scalein_alertconfig_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"scalein_cooldown": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"scaleout_alertconfig_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"scaleout_cooldown": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_predicted_load": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviServerAutoScalePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServerAutoScalePolicyCreate,
		Read:   ResourceAviServerAutoScalePolicyRead,
		Update: resourceAviServerAutoScalePolicyUpdate,
		Delete: resourceAviServerAutoScalePolicyDelete,
		Schema: ResourceServerAutoScalePolicySchema(),
	}
}

func ResourceAviServerAutoScalePolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	log.Printf("[INFO] ResourceAviServerAutoScalePolicyRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/serverautoscalepolicy/" + uuid.(string)
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
	log.Printf("ResourceAviServerAutoScalePolicyRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviServerAutoScalePolicyRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviServerAutoScalePolicyRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviServerAutoScalePolicyRead Updated %v\n", d)
	return nil
}

func resourceAviServerAutoScalePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := ApiCreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	log.Printf("[DEBUG] created object %v: %v", "serverautoscalepolicy", d)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "serverautoscalepolicy", d)
	return err
}

func resourceAviServerAutoScalePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := ApiCreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "serverautoscalepolicy", d)
	return err
}

func resourceAviServerAutoScalePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serverautoscalepolicy"
	log.Println("[INFO] ResourceAviServerAutoScalePolicyRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviServerAutoScalePolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
