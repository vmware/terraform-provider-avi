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

func ResourceUpgradeStatusSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_patch_rollback": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_rollback": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"end_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"node_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"start_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUpgradeOpsStateSchema(),
		},
		"tasks_completed": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"total_tasks": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"upgrade_ops": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviUpgradeStatusSummary() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUpgradeStatusSummaryCreate,
		Read:   ResourceAviUpgradeStatusSummaryRead,
		Update: resourceAviUpgradeStatusSummaryUpdate,
		Delete: resourceAviUpgradeStatusSummaryDelete,
		Schema: ResourceUpgradeStatusSummarySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUpgradeStatusSummaryImporter,
		},
	}
}

func ResourceUpgradeStatusSummaryImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUpgradeStatusSummarySchema()
	return ResourceImporter(d, m, "upgradestatussummary", s)
}

func ResourceAviUpgradeStatusSummaryRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	err := ApiRead(d, meta, "upgradestatussummary", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUpgradeStatusSummaryCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	err := ApiCreateOrUpdate(d, meta, "upgradestatussummary", s)
	if err == nil {
		err = ResourceAviUpgradeStatusSummaryRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusSummaryUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "upgradestatussummary", s)
	if err == nil {
		err = ResourceAviUpgradeStatusSummaryRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusSummaryDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "upgradestatussummary"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviUpgradeStatusSummaryDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
