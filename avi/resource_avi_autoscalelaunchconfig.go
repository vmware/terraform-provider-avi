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

func ResourceAutoScaleLaunchConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"image_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"mesos": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAutoScaleMesosSettingsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"openstack": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAutoScaleOpenStackSettingsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_external_asg": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAutoScaleLaunchConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAutoScaleLaunchConfigCreate,
		Read:   ResourceAviAutoScaleLaunchConfigRead,
		Update: resourceAviAutoScaleLaunchConfigUpdate,
		Delete: resourceAviAutoScaleLaunchConfigDelete,
		Schema: ResourceAutoScaleLaunchConfigSchema(),
	}
}

func ResourceAviAutoScaleLaunchConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/autoscalelaunchconfig/" + uuid.(string)
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

func resourceAviAutoScaleLaunchConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "autoscalelaunchconfig", s)
	if err == nil {
		err = ResourceAviAutoScaleLaunchConfigRead(d, meta)
	}
	return err
}

func resourceAviAutoScaleLaunchConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "autoscalelaunchconfig", s)
	if err == nil {
		err = ResourceAviAutoScaleLaunchConfigRead(d, meta)
	}
	return err
}

func resourceAviAutoScaleLaunchConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "autoscalelaunchconfig"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAutoScaleLaunchConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
