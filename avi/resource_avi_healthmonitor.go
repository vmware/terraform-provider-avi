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

func ResourceHealthMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"disable_quickstart": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"dns_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorDNSSchema(),
		},
		"external_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorExternalSchema(),
		},
		"failed_checks": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"http_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"https_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"monitor_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"radius_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorRadiusSchema(),
		},
		"receive_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"send_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"sip_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSIPSchema(),
		},
		"successful_checks": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"tcp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorTcpSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"udp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorUdpSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHealthMonitorCreate,
		Read:   ResourceAviHealthMonitorRead,
		Update: resourceAviHealthMonitorUpdate,
		Delete: resourceAviHealthMonitorDelete,
		Schema: ResourceHealthMonitorSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHealthMonitorImporter,
		},
	}
}

func ResourceHealthMonitorImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHealthMonitorSchema()
	return ResourceImporter(d, m, "healthmonitor", s)
}

func ResourceAviHealthMonitorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiRead(d, meta, "healthmonitor", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHealthMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "healthmonitor"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviHealthMonitorDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
