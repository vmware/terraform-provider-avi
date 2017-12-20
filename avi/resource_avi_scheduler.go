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

func ResourceSchedulerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"backup_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"end_date_time": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"frequency": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"frequency_unit": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"run_mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"run_script_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"scheduler_action": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SCHEDULER_ACTION_BACKUP",
		},
		"start_date_time": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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

func resourceAviScheduler() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSchedulerCreate,
		Read:   ResourceAviSchedulerRead,
		Update: resourceAviSchedulerUpdate,
		Delete: resourceAviSchedulerDelete,
		Schema: ResourceSchedulerSchema(),
	}
}

func ResourceAviSchedulerRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	log.Printf("[INFO] ResourceAviSchedulerRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/scheduler/" + uuid.(string)
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
	log.Printf("ResourceAviSchedulerRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSchedulerRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSchedulerRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSchedulerRead Updated %v\n", d)
	return nil
}

func resourceAviSchedulerCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := ApiCreateOrUpdate(d, meta, "scheduler", s)
	log.Printf("[DEBUG] created object %v: %v", "scheduler", d)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "scheduler", d)
	return err
}

func resourceAviSchedulerUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := ApiCreateOrUpdate(d, meta, "scheduler", s)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "scheduler", d)
	return err
}

func resourceAviSchedulerDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "scheduler"
	log.Println("[INFO] ResourceAviSchedulerRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSchedulerDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
