// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceAlertSyslogConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"sender_hostname": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"syslog_servers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceAlertSyslogServerSchema(),
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

func resourceAviAlertSyslogConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertSyslogConfigCreate,
		Read:   ResourceAviAlertSyslogConfigRead,
		Update: resourceAviAlertSyslogConfigUpdate,
		Delete: resourceAviAlertSyslogConfigDelete,
		Schema: ResourceAlertSyslogConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAlertSyslogConfigImporter,
		},
	}
}

func ResourceAlertSyslogConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAlertSyslogConfigSchema()
	return ResourceImporter(d, m, "alertsyslogconfig", s)
}

func ResourceAviAlertSyslogConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	err := APIRead(d, meta, "alertsyslogconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAlertSyslogConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	err := APICreate(d, meta, "alertsyslogconfig", s)
	if err == nil {
		err = ResourceAviAlertSyslogConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertSyslogConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	var err error
	err = APIUpdate(d, meta, "alertsyslogconfig", s)
	if err == nil {
		err = ResourceAviAlertSyslogConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertSyslogConfigDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "alertsyslogconfig")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
