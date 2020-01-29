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
		"backup_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"end_date_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"frequency_unit": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"run_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"run_script_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"scheduler_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SCHEDULER_ACTION_BACKUP",
		},
		"start_date_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
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
		Importer: &schema.ResourceImporter{
			State: ResourceSchedulerImporter,
		},
	}
}

func ResourceSchedulerImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSchedulerSchema()
	return ResourceImporter(d, m, "scheduler", s)
}

func ResourceAviSchedulerRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := ApiRead(d, meta, "scheduler", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSchedulerCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := ApiCreateOrUpdate(d, meta, "scheduler", s)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	return err
}

func resourceAviSchedulerUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "scheduler", s)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	return err
}

func resourceAviSchedulerDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "scheduler"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSchedulerDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
