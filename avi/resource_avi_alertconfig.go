// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceAlertConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"alert_rule": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceAlertRuleSchema(),
		},
		"autoscale_alert": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"category": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"expiry_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "86400",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"obj_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"object_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"recommendation": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rolling_window": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"source": {
			Type:     schema.TypeString,
			Required: true,
		},
		"summary": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"throttle": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "600",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAlertConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertConfigCreate,
		Read:   ResourceAviAlertConfigRead,
		Update: resourceAviAlertConfigUpdate,
		Delete: resourceAviAlertConfigDelete,
		Schema: ResourceAlertConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAlertConfigImporter,
		},
	}
}

func ResourceAlertConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAlertConfigSchema()
	return ResourceImporter(d, m, "alertconfig", s)
}

func ResourceAviAlertConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	err := APIRead(d, meta, "alertconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAlertConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	err := APICreateOrUpdate(d, meta, "alertconfig", s)
	if err == nil {
		err = ResourceAviAlertConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "alertconfig", s)
	if err == nil {
		err = ResourceAviAlertConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertconfig"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAlertConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
