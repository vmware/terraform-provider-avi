// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceActionGroupConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_script_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"autoscale_trigger_notification": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"email_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"external_only": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"level": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"snmp_trap_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"syslog_config_ref": {
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

func resourceAviActionGroupConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviActionGroupConfigCreate,
		Read:   ResourceAviActionGroupConfigRead,
		Update: resourceAviActionGroupConfigUpdate,
		Delete: resourceAviActionGroupConfigDelete,
		Schema: ResourceActionGroupConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceActionGroupConfigImporter,
		},
	}
}

func ResourceActionGroupConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceActionGroupConfigSchema()
	return ResourceImporter(d, m, "actiongroupconfig", s)
}

func ResourceAviActionGroupConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	err := APIRead(d, meta, "actiongroupconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviActionGroupConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	err := APICreateOrUpdate(d, meta, "actiongroupconfig", s)
	if err == nil {
		err = ResourceAviActionGroupConfigRead(d, meta)
	}
	return err
}

func resourceAviActionGroupConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "actiongroupconfig", s)
	if err == nil {
		err = ResourceAviActionGroupConfigRead(d, meta)
	}
	return err
}

func resourceAviActionGroupConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "actiongroupconfig"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviActionGroupConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
