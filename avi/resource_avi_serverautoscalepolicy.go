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
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviServerAutoScalePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := ApiCreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	return err
}

func resourceAviServerAutoScalePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := ApiCreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	return err
}

func resourceAviServerAutoScalePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serverautoscalepolicy"
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
