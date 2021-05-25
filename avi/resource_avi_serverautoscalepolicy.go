/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceServerAutoScalePolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delay_for_server_garbage_collection": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"intelligent_autoscale": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"intelligent_scalein_margin": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  40,
		},
		"intelligent_scaleout_margin": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"max_scalein_adjustment_step": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_scaleout_adjustment_step": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"min_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"scalein_alertconfig_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"scalein_cooldown": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"scaleout_alertconfig_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"scaleout_cooldown": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_predicted_load": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": {
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
		Importer: &schema.ResourceImporter{
			State: ResourceServerAutoScalePolicyImporter,
		},
	}
}

func ResourceServerAutoScalePolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServerAutoScalePolicySchema()
	return ResourceImporter(d, m, "serverautoscalepolicy", s)
}

func ResourceAviServerAutoScalePolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := APIRead(d, meta, "serverautoscalepolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServerAutoScalePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	err := APICreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	return err
}

func resourceAviServerAutoScalePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServerAutoScalePolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "serverautoscalepolicy", s)
	if err == nil {
		err = ResourceAviServerAutoScalePolicyRead(d, meta)
	}
	return err
}

func resourceAviServerAutoScalePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serverautoscalepolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviServerAutoScalePolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
