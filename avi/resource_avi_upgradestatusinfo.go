/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func ResourceUpgradeStatusInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
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
		"enqueue_time": {
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
		"params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUpgradeOpsParamSchema(),
		},
		"patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_patch_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_upgrade_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSeUpgradeEventsSchema(),
		},
		"seg_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeGroupStatusSchema(),
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
		"system": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
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
		"upgrade_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEventMapSchema(),
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

func resourceAviUpgradeStatusInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUpgradeStatusInfoCreate,
		Read:   ResourceAviUpgradeStatusInfoRead,
		Update: resourceAviUpgradeStatusInfoUpdate,
		Delete: resourceAviUpgradeStatusInfoDelete,
		Schema: ResourceUpgradeStatusInfoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUpgradeStatusInfoImporter,
		},
	}
}

func ResourceUpgradeStatusInfoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUpgradeStatusInfoSchema()
	return ResourceImporter(d, m, "upgradestatusinfo", s)
}

func ResourceAviUpgradeStatusInfoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	err := ApiRead(d, meta, "upgradestatusinfo", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUpgradeStatusInfoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	err := ApiCreateOrUpdate(d, meta, "upgradestatusinfo", s)
	if err == nil {
		err = ResourceAviUpgradeStatusInfoRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "upgradestatusinfo", s)
	if err == nil {
		err = ResourceAviUpgradeStatusInfoRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusInfoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "upgradestatusinfo"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviUpgradeStatusInfoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
